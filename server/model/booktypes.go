package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	College   string             `bson:"college,omitempty" json:"college,omitempty"`
	Publisher string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string             `bson:"year,omitempty" json:"year,omitempty"`
	ISBN      string             `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Authors   []Author           `bson:"authors,omitempty" json:"authors,omitempty"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
type Author struct {
	ID   string `bson:"id,omitempty" json:"id,omitempty"`
	Rank int    `bson:"rank,omitempty" json:"rank,omitempty"`
}
