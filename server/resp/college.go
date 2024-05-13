package resp

type DeanResp struct {
	ID      string `bson:"id,omitempty" json:"id,omitempty"`
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Contact string `bson:"contact,omitempty" json:"contact,omitempty"`
}
type TeacherResp struct {
	TeacherID string `bson:"teacherID,omitempty" json:"id,omitempty"`
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Contact   string `bson:"contact,omitempty" json:"contact,omitempty"`
}
