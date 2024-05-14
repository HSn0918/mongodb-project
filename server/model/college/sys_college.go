// 自动生成模板SysCollege
package college

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 学院 结构体  SysCollege
type SysCollege struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"` //学院名
}

// TableName 学院 SysCollege自定义表名 college
func (SysCollege) TableName() string {
	return "college"
}

type College struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Dean     Dean               `bson:"dean,omitempty" json:"dean,omitempty"`
	Teachers []string           `bson:"teachers,omitempty" json:"teachers,omitempty"`
}
type Dean struct {
	ID      string `bson:"id,omitempty" json:"id,omitempty"`
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Contact string `bson:"contact,omitempty" json:"contact,omitempty"`
}
type Teacher struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TeacherID string             `bson:"id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Contact   string             `bson:"contact,omitempty" json:"contact,omitempty"`
}
type DeanReq struct {
	Name string `json:"name"`
}
type TeachersReq struct {
	Name string `json:"name"`
}
