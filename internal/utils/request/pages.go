// Package request defines request DTOs.
package request

type PageList struct {
	Page      int64  `json:"page" form:"page"`
	PageSize  int64  `json:"pageSize" form:"pageSize"`
	SearchKey string `json:"key" form:"key"`
}
