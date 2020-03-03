package api

import "math"

type Result struct {
	Code    string      `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

type PagingResult struct {
	TotalCount     int         `json:"totalCount" xml:"totalCount"`
	TotalPageCount int         `json:"totalPageCount" xml:"totalPageCount"`
	Page           int         `json:"page" xml:"page"`
	PageSize       int         `json:"pageSize" xml:"pageSize"`
	Lists          interface{} `json:"lists" xml:"lists"`
}

func NewResult(code string, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
	}
}

func NewPagingResult(lists interface{}, totalCount, page int, pageSize int) *PagingResult {
	totalPageCount := 0
	if pageSize > 0 {
		totalPageCount = int(math.Ceil(float64(totalCount / pageSize)))
	}
	return &PagingResult{
		TotalCount:     totalCount,
		TotalPageCount: totalPageCount,
		Page:           page,
		PageSize:       pageSize,
		Lists:          lists,
	}
}

func (r *Result) Simple(data interface{}) *Result {
	r.Data = data
	return r
}

func (r *Result) Paging(lists interface{}, totalCount, page int, pageSize int) *Result {
	r.Data = NewPagingResult(lists, totalCount, page, pageSize)
	return r
}
