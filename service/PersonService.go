package service

import "go-gin-demo/model"

type PersonService interface {
	GetPersonInfo(name string)(personBo model.PersonBO)
	
}

