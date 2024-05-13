package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ TeacherModel = (*customTeacherModel)(nil)

type (
	// TeacherModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeacherModel.
	TeacherModel interface {
		teacherModel
	}

	customTeacherModel struct {
		*defaultTeacherModel
	}
)

// NewTeacherModel returns a model for the mongo.
func NewTeacherModel(url, db, collection string) TeacherModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customTeacherModel{
		defaultTeacherModel: newDefaultTeacherModel(conn),
	}
}
