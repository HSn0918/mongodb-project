// 自动生成模板SysBook
package book

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 教材 结构体  SysBook
type SysBook struct {
	global.GVA_MODEL
	ISBN string `json:"iSBN" form:"iSBN" gorm:"column:ISBN;comment:;"` //ISBN
}

// TableName 教材 SysBook自定义表名 sys_book
func (SysBook) TableName() string {
	return "sys_book"
}

type FindBookReq struct {
	Name string `json:"name"`
}
type FindBookResp struct {
	Name      string   `bson:"name,omitempty" json:"name,omitempty"`
	College   string   `bson:"college,omitempty" json:"college,omitempty"`
	Publisher string   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string   `bson:"year,omitempty" json:"year,omitempty"`
	ISBN      string   `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Authors   []Author `bson:"authors,omitempty" json:"authors,omitempty"`
}
type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	College   string             `bson:"college,omitempty" json:"college,omitempty"`
	Publisher string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string             `bson:"year,omitempty" json:"year,omitempty"`
	ISBN      string             `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Authors   []Author           `bson:"authors,omitempty" json:"authors,omitempty"`
}
type Author struct {
	ID   string `bson:"id,omitempty" json:"id,omitempty"`
	Rank int    `bson:"rank,omitempty" json:"rank,omitempty"`
}
type UpdateBookReq struct {
	Name      string   `bson:"name,omitempty" json:"name,omitempty"`
	College   string   `bson:"college,omitempty" json:"college,omitempty"`
	Publisher string   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string   `bson:"year,omitempty" json:"year,omitempty"`
	ISBN      string   `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Authors   []Author `bson:"authors,omitempty" json:"authors,omitempty"`
}
type UpdateBookResp struct {
	Flag bool `json:"flag"`
}
type InsertBookReq struct {
	Name      string   `bson:"name,omitempty" json:"name,omitempty"`
	College   string   `bson:"college,omitempty" json:"college,omitempty"`
	Publisher string   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string   `bson:"year,omitempty" json:"year,omitempty"`
	ISBN      string   `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Authors   []Author `bson:"authors,omitempty" json:"authors,omitempty"`
}
type InsertBookResp struct {
	Flag bool `json:"flag"`
}
type EChart struct {
	CollegesEChart []College `bson:"college,omitempty" json:"college,omitempty"`
	PublishEChart  []Publish `bson:"publish,omitempty" json:"publish,omitempty"`
	YearEChart     []Year    `bson:"year,omitempty" json:"year,omitempty"`
}
type College struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Publish struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Year struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
