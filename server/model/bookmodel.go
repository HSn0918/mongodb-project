package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ BookModel = (*customBookModel)(nil)

type (
	// BookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookModel.
	BookModel interface {
		bookModel
	}

	customBookModel struct {
		*defaultBookModel
	}
)

// NewBookModel returns a model for the mongo.
func NewBookModel(url, db, collection string) BookModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customBookModel{
		defaultBookModel: newDefaultBookModel(conn),
	}
}
