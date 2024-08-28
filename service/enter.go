package service

import "myServer/service/user_service"

type ServiceGroup struct {
	UserService user_service.UserService
}

var MyService ServiceGroup
