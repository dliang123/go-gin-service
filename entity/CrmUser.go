package entity

type CrmUser struct {
	Id       int
	Name     string
	Age      int
	LastName string `xorm:"varchar(255)  'last_name'"`
}

// 设置表名
func (*CrmUser) TableName() string {
	return "crm_user"
}
