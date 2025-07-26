package common

import (
	"backend/global"
	"backend/models"
	"fmt"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool     //是否开启debug模式,是否查看日志
	Likes []string //模糊匹配的字段
}

// ComList 通用列表分页查询服务
func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	//排序
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认排序创建顺序desc从晚到早.asc从早到晚
	}

	//嵌套一层查询数据库中的这个结构体
	DB = DB.Where(model)

	// 新增模糊查询
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	count = DB.Find(&list).RowsAffected
	//复位
	query := DB.Where(model)
	//偏移量
	offset := (option.Page - 1) * option.Limit
	//如果偏移量小于0，则从0开始
	if offset < 0 {
		offset = 0
	}
	//如果limit为0，则查询所有
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
