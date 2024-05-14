package book

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/book"
	bookReq "github.com/flipped-aurora/gin-vue-admin/server/model/book/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/college"
	"go.mongodb.org/mongo-driver/bson"
)

type SysBookService struct {
}

func (sysBookService *SysBookService) FindBookByTeacherName(name string) ([]*book.FindBookResp, error) {
	var teacher college.Teacher
	_ = global.GVA_MONGO.Database.Collection("teacher").Find(context.TODO(), bson.M{"name": name}).One(&teacher)
	var books []*book.FindBookResp
	_ = global.GVA_MONGO.Database.Collection("book").Find(context.TODO(), bson.M{"authors.id": teacher.TeacherID}).All(&books)
	return books, nil
}

// CreateSysBook 创建教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) CreateSysBook(sysBook *book.SysBook) (err error) {
	err = global.GVA_DB.Create(sysBook).Error
	return err
}

// DeleteSysBook 删除教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) DeleteSysBook(ID string) (err error) {
	err = global.GVA_DB.Delete(&book.SysBook{}, "id = ?", ID).Error
	return err
}

// DeleteSysBookByIds 批量删除教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) DeleteSysBookByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]book.SysBook{}, "id in ?", IDs).Error
	return err
}

// UpdateSysBook 更新教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) UpdateSysBook(sysBook book.SysBook) (err error) {
	err = global.GVA_DB.Model(&book.SysBook{}).Where("id = ?", sysBook.ID).Updates(&sysBook).Error
	return err
}

// GetSysBook 根据ID获取教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) GetSysBook(ID string) (sysBook book.SysBook, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysBook).Error
	return
}

// GetSysBookInfoList 分页获取教材记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysBookService *SysBookService) GetSysBookInfoList(info bookReq.SysBookSearch) (list []book.SysBook, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&book.SysBook{})
	var sysBooks []book.SysBook
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

	err = db.Find(&sysBooks).Error
	return sysBooks, total, err
}

func (sysBookService *SysBookService) InsertBook(req book.InsertBookReq) (*book.InsertBookResp, error) {
	coll := global.GVA_MONGO.Database.Collection("book")
	_, err := coll.InsertOne(context.TODO(), req)
	if err != nil {
		return &book.InsertBookResp{Flag: false}, err
	}
	return &book.InsertBookResp{Flag: true}, nil
}

func (sysBookService *SysBookService) EChart() (*book.EChart, error) {
	var books []book.Book
	err := global.GVA_MONGO.Database.Collection("book").Find(context.TODO(), bson.M{}).All(&books)
	if err != nil {
		return &book.EChart{}, err
	}

	collegeMap := make(map[string]int)
	publishMap := make(map[string]int)
	yearMap := make(map[string]int)

	for _, book := range books {
		collegeMap[book.College]++
		publishMap[book.Publisher]++
		yearMap[book.Year]++
	}

	var collegesEChart []book.College
	for name, value := range collegeMap {
		collegesEChart = append(collegesEChart, book.College{Name: name, Value: value})
	}

	var publishEChart []book.Publish
	for name, value := range publishMap {
		publishEChart = append(publishEChart, book.Publish{Name: name, Value: value})
	}

	var yearEChart []book.Year
	for name, value := range yearMap {
		yearEChart = append(yearEChart, book.Year{Name: name, Value: value})
	}

	eChart := &book.EChart{
		CollegesEChart: collegesEChart,
		PublishEChart:  publishEChart,
		YearEChart:     yearEChart,
	}

	return eChart, nil
}
