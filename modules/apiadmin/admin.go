package apiadmin

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/apiadmin/models"
)

/**
 * 根据管理员ID获取数据
 * @param userId int
 * @return *Admin
 */
func GetUserInfoById(userId int) (*models.Admin, error) {
	adminObj := models.Admin{}
	adminUser, err := adminObj.AdminGetByUserId(userId)
	return adminUser, err
}

/**
 * 获取管理员列表
 * @param page  int 	// 分页页数
 * @param limit int		// 分页大小
 * return []*admin
 */
func GetAdminList(page, pageSize int, filters map[string]interface{}) (list []*models.Admin, count int64, err error) {
	adminObj := models.Admin{}
	list, count, err = adminObj.AdminGetList(page, pageSize, filters)
	return list, count, err
}

/**
 * 根据用户名获取用户数据
 * @param page  int 	// 分页页数
 * @param limit int		// 分页大小
 * return []*admin
 */
func GetUserInfoByName(username string) (*models.Admin, error) {
	adminObj := models.Admin{}
	return adminObj.AdminGetByName(username)
}

/**
 * 根据用户名获取用户数据
 * @param page  int 	// 分页页数
 * @param limit int		// 分页大小
 * return []*admin
 */
func AdminCreate(a *models.Admin) (int64, error) {
	adminObj := models.Admin{}
	return adminObj.AdminCreate(a)
}

/**
 * 管理员登录处理
 * @param username string // 用户名
 * @param password string // 登录密码
 * @return error
 */
func Login(req *beego.Controller, username, password string) (err error) {
	// 查询用户数据
	adminObj := models.Admin{}
	userInfo, err := adminObj.AdminGetByName(username)
	if err != nil {
		return err
	}
	if userInfo.Id < 0 {
		return errors.New("用户名或密码不正确")
	}

	// 对比用户密码
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password + userInfo.Salt))
	md5sum := hex.EncodeToString(md5Ctx.Sum(nil))
	if md5sum != userInfo.Password {
		return errors.New("用户名或密码不正确")
	}

	// 记录登录信息
	lastIp := strings.Split(req.Ctx.Request.RemoteAddr, ":")
	userInfo.LastIp = lastIp[0]
	userInfo.LastLogin = time.Now().Unix()
	userInfo.UpdateTime = time.Now().Unix()
	userInfo.Update()

	// 设置cookie
	md5Ctx = md5.New()
	md5Ctx.Write([]byte(md5sum + userInfo.LastIp + userInfo.Salt))
	authKey := hex.EncodeToString(md5Ctx.Sum(nil))
	req.Ctx.SetCookie("auth", strconv.Itoa(userInfo.Id)+"|"+authKey, 7*86400)

	return nil
}
