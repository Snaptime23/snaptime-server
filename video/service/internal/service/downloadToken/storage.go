package downloadToken

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func hmacSha1(data, key string) []byte {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return h.Sum(nil)
}

func urlsafeBase64Encode(data []byte) string {
	encoded := base64.URLEncoding.EncodeToString(data)
	return encoded
}

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	//bucket    = os.Getenv("QINIU_TEST_BUCKET")
)

func GetToken(resourcekey string) string {
	resourcekey = "https://bucket.snaptime.jiyeon.club/" + resourcekey
	// 添加过期时间参数
	downloadURLWithExpire := resourcekey + fmt.Sprintf("?e=%d", time.Now().Unix()+7200)

	// 计算签名
	sign := hmacSha1(downloadURLWithExpire, secretKey)

	// 进行 URL 安全的 Base64 编码
	encodedSign := urlsafeBase64Encode(sign)

	// 拼接访问密钥与签名
	token := accessKey + ":" + encodedSign

	// 拼接最终的下载 URL
	realDownloadURL := downloadURLWithExpire + "&token=" + token
	return realDownloadURL
}
