package flag

import (
	"backend/global"
	"backend/models"
	"github.com/sirupsen/logrus"
)

func Makemigrations() {
	var err error
	//生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.TestModel{},
		)
	if err != nil {
		logrus.Error("初始化数据库失败", err)
		return
	}
	logrus.Info("初始化数据库成功")
}
