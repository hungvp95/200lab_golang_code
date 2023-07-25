package common

type SQLModel struct {
	Id     int  `json:"-" gorm:"column:id;"`
	FakeId *UID `json:"id" gorm:"-"`
	Status int  `json:"status" gorm:"column:status;default:1;"`
	//CreateAt *time.Time `json:"createAt,omitempty" gorm:"column:created_at;"`
	//UpdateAt *time.Time `json:"updateAt,omitempty" gorm:"column:updated_at;"`
}

func (model *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(model.Id), dbType, 1)
	model.FakeId = &uid
}
