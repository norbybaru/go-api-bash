package dish

type CreateDishRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
}

type UpdateDishRequest struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
}
