package rating

type CreateRatingRequest struct {
	Rate   uint8 `json:"rate" validate:"required,gte=1,lte=10"`
	DishId int   `json:"dish_id" validate:"required"`
	UserId int   `json:"-"`
}
