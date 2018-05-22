package models

// QueryParams 查询参数
type QueryParams struct {
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form:"age"`
	Weight int    `json:"weight" form:"weight"`
}
