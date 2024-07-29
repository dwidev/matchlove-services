package dto

type (
	PaginationParams struct {
		Page    int `query:"page" validate:"required,min=1"`
		PerPage int `query:"per_page" validate:"required,min=1,max=100"`
	}

	PaginationResultDTO struct {
		CurrentPage int         `json:"current_page"`
		TotalPage   int         `json:"total_page"`
		TotalData   int         `json:"total_data"`
		Data        interface{} `json:"data"`
	}
)
