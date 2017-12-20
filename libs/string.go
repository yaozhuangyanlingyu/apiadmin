package libs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

/**
 * 生成随机字符串
 * @param lens int 	// 要生成随机字符串的长度
 * @return string
 */
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := []byte{}
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * 生成密码
 * @param lens int 	// 要生成随机字符串的长度
 * @return string
 */
func GetPasswdString(passwd string) (string, string) {
	if passwd == "" {
		passwd = strings.TrimSpace(beego.AppConfig.String("site.default_passwd"))
	}
	salt := GetRandomString(4)
	md5Sum := md5.New()
	md5Sum.Write([]byte(passwd + salt))
	pwd := hex.EncodeToString(md5Sum.Sum(nil))
	fmt.Println(passwd + salt)
	return pwd, salt
}
