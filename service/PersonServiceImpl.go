package service

import (
	"go-gin-demo/model"
	"go-gin-demo/config"
	"go-gin-demo/entity"
)

// 实现类
type PersonServiceImpl struct{}

func (personServiceImpl *PersonServiceImpl) GetPersonInfo(name string) (personBo model.PersonBO) {

	DbEngine := config.Init()
	crmUser := entity.CrmUser{}
	DbEngine.Where("name=?", name).Get(&crmUser)
	return model.PersonBO{Name: crmUser.Name, Age: crmUser.Age, LastName: crmUser.LastName}
}

type PersonServiceImpl2 struct {
	PersonServiceImpl
	name string
}

func (personServiceImpl *PersonServiceImpl2) print() {

}
