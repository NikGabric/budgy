package helpers

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count"`
}

func CreatePaginationResponse(data interface{}, totalCount int) PaginationResponse {
	return PaginationResponse{
		Data:       data,
		TotalCount: totalCount,
	}
}
