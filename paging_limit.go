package api

type PagingLimit struct {
	Page     uint64 `json:"page" form:"page" valid:"int~分页必须在1-2100000000之间,range(1|2100000000)~分页必须在1-2100000000之间"`
	PageSize uint64 `json:"pageSize" form:"pageSize" valid:"int~分页大小必须在1-2100000000之间,range(1|2100000000)~分页大小必须在1-2100000000之间"`
}

func NewPagingLimit(page uint64, pageSize uint64) *PagingLimit {
	return &PagingLimit{
		Page:     page,
		PageSize: pageSize,
	}
}
