package dish

import (
	"dancing-pony/internal/common/utils"
	"time"

	"github.com/gosimple/slug"
	"github.com/segmentio/ksuid"
)

func (d *Dish) TableName() string {
	return "dishes"
}

type Dish struct {
	Id          ksuid.KSUID `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Description string      `json:"description" db:"description"`
	Slug        string      `json:"slug" db:"slug"`
	Price       uint64      `json:"price" db:"price"`
	ImageUrl    string      `json:"image_url" db:"image_url"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}

func (d *Dish) GetId() string {
	return d.Id.String()
}

func NewDish(name string, description string, imageUrl string, price uint64) *Dish {
	now := time.Now().UTC()
	return &Dish{
		Id:          utils.GenUUID(),
		Name:        name,
		Description: description,
		Slug:        slug.Make(name),
		Price:       price,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
