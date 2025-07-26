package api

import (
	"backend/api/test_api"
)

type ApiGroup struct {
	TestApi test_api.TestApi
}

var ApiGroupApp = new(ApiGroup)
