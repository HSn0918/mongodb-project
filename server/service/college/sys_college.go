package college

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/college"
	collegeReq "github.com/flipped-aurora/gin-vue-admin/server/model/college/request"
	"go.mongodb.org/mongo-driver/bson"
)

type SysCollegeService struct {
}

// CreateSysCollege 创建学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) CreateSysCollege(sysCollege *college.SysCollege) (err error) {
	err = global.GVA_DB.Create(sysCollege).Error
	return err
}

// DeleteSysCollege 删除学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) DeleteSysCollege(ID string) (err error) {
	err = global.GVA_DB.Delete(&college.SysCollege{}, "id = ?", ID).Error
	return err
}

// DeleteSysCollegeByIds 批量删除学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) DeleteSysCollegeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]college.SysCollege{}, "id in ?", IDs).Error
	return err
}

// UpdateSysCollege 更新学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) UpdateSysCollege(sysCollege college.SysCollege) (err error) {
	err = global.GVA_DB.Model(&college.SysCollege{}).Where("id = ?", sysCollege.ID).Updates(&sysCollege).Error
	return err
}

// GetSysCollege 根据ID获取学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) GetSysCollege(ID string) (sysCollege college.SysCollege, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysCollege).Error
	return
}

// GetSysCollegeInfoList 分页获取学院记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysCollegeService *SysCollegeService) GetSysCollegeInfoList(info collegeReq.SysCollegeSearch) (list []college.SysCollege, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&college.SysCollege{})
	var sysColleges []college.SysCollege
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysColleges).Error
	return sysColleges, total, err
}

func (sysCollegeService *SysCollegeService) FindAllCollegeName() ([]string, error) {
	var collegesName []string
	collection := global.GVA_MONGO.Database.Collection("college")

	var college []college.College
	err := collection.Find(context.TODO(), bson.M{}).All(&college)
	if err != nil {
		return nil, err
	}
	// 遍历college切片，提取学院名称并添加到collegesName切片中
	for _, v := range college {
		collegesName = append(collegesName, v.Name)
	}

	return collegesName, nil
}

func (sysCollegeService *SysCollegeService) FindDean(name string) (dean *college.Dean, err error) {
	var College *college.College

	collection := global.GVA_MONGO.Database.Collection("college")
	_ = collection.Find(context.TODO(), bson.M{"name": name}).One(&College)
	if College == nil {
		return nil, errors.New("未找到该学院")
	}

	return &college.Dean{
		ID:      College.Dean.ID,
		Name:    College.Dean.Name,
		Contact: College.Dean.Contact,
	}, nil
}
func (sysCollegeService *SysCollegeService) FindAllTeacher(name string) ([]*college.Teacher, error) {
	var College *college.College
	collection := global.GVA_MONGO.Database.Collection("college")
	err := collection.Find(context.TODO(), bson.M{"name": name}).One(&College)
	if err != nil {
		return nil, err
	}
	teacherIDs := College.Teachers
	collection = global.GVA_MONGO.Database.Collection("teacher")
	var teachers []*college.Teacher
	err = collection.Find(context.TODO(), bson.M{"id": bson.M{"$in": teacherIDs}}).All(&teachers)
	return teachers, nil
}
