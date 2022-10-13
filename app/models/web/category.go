package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
