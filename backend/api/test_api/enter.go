package test_api

import (
	"backend/models"
	"backend/models/res"
	"backend/service/common"
	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (TestApi) TestView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.TestModel{},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	if err != nil {
		res.FailWithMessage("获取失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
