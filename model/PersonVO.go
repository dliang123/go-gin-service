package model

// Post请求实体
type PersonVO struct {
	Name string `json:"name" binding:"required"`
}

