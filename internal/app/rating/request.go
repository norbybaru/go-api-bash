package rating

type CreateRatingRequest struct {
	Rate   uint8 `json:"rate" validate:"required,gte=1,lte=5"`
	DishId int   `json:"dish_id" validate:"required"`
	UserId int   `json:"user_id"`
}
