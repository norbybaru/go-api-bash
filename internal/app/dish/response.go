package dish

type DishResourceResponse struct {
	Dish
	Ratings []RatingResource `json:"ratings,omitempty"`
}

type RatingResource struct {
	Rating int `json:"rating"`
	UserId int `json:"user_id"`
}
