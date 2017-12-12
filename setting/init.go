package setting

import (
	"github.com/apiadmin/models"
)

// 初始化框架
func Init() (err error) {
	/*// 初始化配置
	err = InitConfig()
	if err != nil {
		return err
	}*/

	// 初始化model
	err = models.Init()
	if err != nil {
		return err
	}

	return nil
}
