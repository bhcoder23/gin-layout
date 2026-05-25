// Package response defines API response DTOs and helpers.
package response

import "math"

type Pages struct {
	Count     int64 `json:"count"`
	CurPage   int64 `json:"page"`
	TotalPage int64 `json:"totalPage"`
	PageSize  int64 `json:"pageSize"`
}

type PageList struct {
	Pages Pages `json:"pages"`
	List  any   `json:"data"`
}
type PageListT[T any] struct {
	Pages Pages `json:"pages"`
	List  []T   `json:"data"`
}

func MakePages(count int64, curPage int64, pageSize int64) (pages Pages) {
	pages.Count = count
	pages.PageSize = pageSize
	pages.CurPage = curPage
	var totalPageFloat float64 = 0
	if pageSize != 0 {
		totalPageFloat = float64(count) / float64(pageSize)
	}
	pages.TotalPage = int64(math.Ceil(totalPageFloat))
	return
}
