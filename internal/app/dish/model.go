package dish

import (
	"time"

	"github.com/gosimple/slug"
)

func (d *Dish) TableName() string {
	return "dishes"
}

type Dish struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Slug        string    `json:"slug" db:"slug"`
	Price       uint64    `json:"price" db:"price"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	UserId      int       `json:"-" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewDish(name string, description string, imageUrl string, price uint64, userId int) *Dish {
	now := time.Now().UTC()
	return &Dish{
		Name:        name,
		Description: description,
		Slug:        slug.Make(name),
		Price:       price,
		ImageUrl:    imageUrl,
		UserId:      userId,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
