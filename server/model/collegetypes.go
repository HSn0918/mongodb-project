package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type College struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Dean     Dean               `bson:"dean,omitempty" json:"dean,omitempty"`
	Teachers []string           `bson:"teachers,omitempty" json:"teachers,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type Dean struct {
	ID      string `bson:"id,omitempty" json:"id,omitempty"`
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Contact string `bson:"contact,omitempty" json:"contact,omitempty"`
}
