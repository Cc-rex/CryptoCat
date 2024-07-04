package api

import "myServer/api/system_api"

type ApiGroup struct {
	SystemsApi system_api.SystemApi
}

var MyApigroup = new(ApiGroup)
