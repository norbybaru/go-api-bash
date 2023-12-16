package dish

type CreateDishRequest struct {
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required,lte=300,url"`
	Price       uint64 `json:"price" validate:"required,number"`
}

type UpdateDishRequest struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required,http_url"`
	Price       uint64 `json:"price" validate:"required,numeric"`
}
