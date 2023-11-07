# 服务端

## 技术选型与开发设计
### 技术选型
![img.png](img.png)

### 开发设计
* $JWT$ 鉴权： 基于 $golang$ - $jwt$ 包实现了token的颁发与解析
* 用户头像、背景的随机获取：图片与视频的存储采用了 $OSS$ 云对象存储，相较于一般本地服务器，具有更加安全、持久、方便、快捷等特性，由于客户端没有设置头像的设置，这里采用的是随机生成，在用户注册后会自动得获取一组随机得头像与背景。

### 架构设计
![img_7.png](img_7.png)

### 数据库设计

![img_2.png](img_2.png) ![img_3.png](img_3.png) ![img_4.png](img_4.png)
![img_5.png](img_5.png) ![img_6.png](img_6.png) ![img_8.png](img_8.png)
![img_9.png](img_9.png) ![img_10.png](img_10.png) ![img_11.png](img_11.png)
* 关注数、点赞数等计数操作使用额外字段记录，避免使用 $COUNT()$ 计数。 
* 每张表都添加了 $deleted_at$ 字段用于软删除。
* 索引：
    - 用户表：对( $username$ , $deleted_at$ )做唯一索引；
    - 评论表：$vid$、$uid$ 做索引；
    - 点赞表：$uid$、$vid$ 做索引，（$uid$, $vid$, $deleted_at$ ）做唯一约束；


缓存层，使用 $Redis$ 的 $String$ 的数据结构，使用 $key$ - $value$，存储一个视频的点赞数量。

采取使用定时任务的方式，间隔一段时间将缓存中的数据与数据库中的数据进行同步，一定程度上解决数据一致性的问题。