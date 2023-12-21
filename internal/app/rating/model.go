package rating

import "time"

func (d *Rating) TableName() string {
	return "ratings"
}

type Rating struct {
	UserId    int       `json:"user_id" db:"user_id"`
	DishId    int       `json:"dish_id" db:"dish_id"`
	Rate      uint8     `json:"rate" db:"rate"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func NewRating(userId int, dishId int, rate uint8) *Rating {
	return &Rating{
		UserId: userId,
		DishId: dishId,
		Rate:   rate,
	}
}
