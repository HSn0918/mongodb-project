package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ CollegeModel = (*customCollegeModel)(nil)

type (
	// CollegeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollegeModel.
	CollegeModel interface {
		collegeModel
	}

	customCollegeModel struct {
		*defaultCollegeModel
	}
)

// NewCollegeModel returns a model for the mongo.
func NewCollegeModel(url, db, collection string) CollegeModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCollegeModel{
		defaultCollegeModel: newDefaultCollegeModel(conn),
	}
}
