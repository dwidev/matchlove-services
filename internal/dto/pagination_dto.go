package dto

type (
	PaginationParams struct {
		Page     int `query:"page" validate:"required,min=1"`
		PageSize int `query:"pageSize" validate:"required,min=1,max=100"`
	}

	PaginationResultDTO struct {
		CurrentPage int         `json:"currentPage"`
		TotalPage   int         `json:"totalPage"`
		Data        interface{} `json:"data"`
	}
)
