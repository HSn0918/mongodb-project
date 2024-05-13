package svc

import "github.com/hsn0918/mongodb_project/model"

type ServiceContext struct {
	BookModel    model.BookModel
	TeacherModel model.TeacherModel
	CollegeModel model.CollegeModel
}

func NexServiceContext() *ServiceContext {
	url := "mongodb://admin:admin@localhost:27017"
	db := "manger"
	return &ServiceContext{
		BookModel:    model.NewBookModel(url, db, "book"),
		TeacherModel: model.NewTeacherModel(url, db, "teacher"),
		CollegeModel: model.NewCollegeModel(url, db, "college"),
	}
}
