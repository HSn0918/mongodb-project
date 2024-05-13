package model

import (
	"context"
	"github.com/hsn0918/mongodb_project/resp"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ CollegeModel = (*customCollegeModel)(nil)

type (
	// CollegeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollegeModel.
	CollegeModel interface {
		collegeModel
		FindDean(college string) (*resp.DeanResp, error)
		FindAllCollegeName() (*[]string, error)
		FindAllCollegeTeacherID(college string) (*[]resp.TeacherResp, error)
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
func (m *customCollegeModel) FindAllCollegeTeacherID(college string) (*[]resp.TeacherResp, error) {
	filter := bson.M{"name": college}
	res, err := m.conn.Collection.FindOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var c College
	err = res.Decode(&c)
	if err != nil {
		return nil, err
	}
	err = res.Decode(&c)
	if err != nil {
		return nil, err
	}
	var teachers []resp.TeacherResp
	for _, v := range c.Teachers {
		teachers = append(teachers, resp.TeacherResp{
			TeacherID: v,
		})
	}
	return &teachers, nil
}

func (m *customCollegeModel) FindAllCollegeName() (*[]string, error) {
	filter := bson.M{}
	res, err := m.conn.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var c []College
	err = res.All(context.TODO(), &c)
	if err != nil {
		return nil, err
	}
	var collegeName []string
	for _, v := range c {
		collegeName = append(collegeName, v.Name)
	}
	return &collegeName, nil
}
func (m *customCollegeModel) FindDean(college string) (*resp.DeanResp, error) {
	filter := bson.M{"name": college}
	res, err := m.conn.Collection.FindOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var c College
	err = res.Decode(&c)
	if err != nil {
		return nil, err
	}
	dean := &resp.DeanResp{
		ID:      c.Dean.ID,
		Name:    c.Dean.Name,
		Contact: c.Dean.Contact,
	}
	return dean, nil
}
