package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ TeacherModel = (*customTeacherModel)(nil)

type (
	// TeacherModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeacherModel.
	TeacherModel interface {
		teacherModel
		FindOneByTeacherIds(teacherIds []string) (*[]Teacher, error)
	}

	customTeacherModel struct {
		*defaultTeacherModel
	}
)

func (m *customTeacherModel) FindOneByTeacherIds(teacherIds []string) (*[]Teacher, error) {
	var resp []Teacher
	query := bson.M{"teacherID": bson.M{"$in": teacherIds}}
	cursor, err := m.conn.Collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var teacher Teacher
		err := cursor.Decode(&teacher)
		if err != nil {
			return nil, err
		}
		resp = append(resp, teacher)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewTeacherModel returns a model for the mongo.
func NewTeacherModel(url, db, collection string) TeacherModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customTeacherModel{
		defaultTeacherModel: newDefaultTeacherModel(conn),
	}
}
