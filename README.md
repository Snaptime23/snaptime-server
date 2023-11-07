# 服务端

## 快速开始

### 需要提前设置的环境变量
```
DB_PASSWORD
HOST=0.0.0.0
DB_USERNAME
QINIU_ACCESS_KEY
QINIU_SECRET_KEY
REDIS_ADDR
REDIS_PASSWORD
DB_ADDR
```

### 下载依赖

在项目的根目录下

```bash
go mod tidy
```

运行

```bash
go run main.go
```

### 如果想 修改/生成 gRPC 代码

```bash
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto
go get google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/gogo/protobuf/gogoproto
```

然后在 proto 的文件夹下执行：

但是需要确保 `github.com/gogo/protobuf/gogoproto/gogo.proto` 在 `$GOPATH/pkg/mod` 目录下

最后执行

```bash
protoc -I=. -I=$GOPATH/pkg/mod --go-grpc_out=. --gogo_out=. *.proto
```

## 技术选型与开发设计
### 技术选型

![img.png](./docs/images/img.png)

### 开发设计

- $JWT$ 鉴权： 基于 $golang$ - $jwt$ 包实现了 token 的颁发与解析
- 使用原生 $gRPC$ 框架， $proto$ 自动生成 $pb.go$ 文件。
- 用户头像、背景的随机获取：图片与视频的存储采用了 $OSS$ 云对象存储，相较于一般本地服务器，具有更加安全、持久、方便、快捷等特性，由于客户端没有设置头像的设置，这里采用的是随机生成，在用户注册后会自动得获取一组随机得头像与背景。

### 架构设计
使用七牛云的云对象存储 $kodo$，先将视频上传到指定的 $bucket$ ，上传完成后出发工作流，回调项目中的接口，从而完成视频链接的存储。

![img_7.png](./docs/images/img_7.png)

### 数据库设计

![img_2.png](./docs/images/img_2.png) ![img_3.png](./docs/images/img_3.png) ![img_4.png](./docs/images/img_4.png)
![img_5.png](./docs/images/img_5.png) ![img_6.png](./docs/images/img_6.png) ![img_8.png](./docs/images/img_8.png)
![img_9.png](./docs/images/img_9.png) ![img_10.png](./docs/images/img_10.png) ![img_11.png](./docs/images/img_11.png)

- 关注数、点赞数等计数操作使用额外字段记录，避免使用 $COUNT()$ 计数。
- 每张表都添加了 $deleted_at$ 字段用于软删除。
- 索引：
  - 用户表：对( $username$ , $deleted_at$ )做唯一索引；
  - 评论表：$vid$ 、$uid$ 做索引；
  - 点赞表：$uid$ 、$vid$ 做索引，（ $uid$, $vid$ , $deleted$\_$at$ ）做唯一约束；

缓存层，使用 $Redis$ 的 $String$ 的数据结构，使用 $key$ - $value$，存储一个视频的点赞数量。

采取使用定时任务的方式，间隔一段时间将缓存中的数据与数据库中的数据进行同步，一定程度上解决数据一致性的问题。