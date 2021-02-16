package model

import (
	"time"
)

type Staffs struct {
	StaffID          uint32 `gorm:"type:int(11);column:staff_id;primary_key"`
	StaffName        string
	Account          string
	Password         string
	DeptID           uint32
	StaffsStatusID   int
	CreateTime       uint64
	DeleteTime       uint64
	WorkShiftID      int
	EntryTime        uint64
	RoleID           int
	DeleteTimeFormat time.Time
	EntryTimeFormat  time.Time
}

//获取所有
func FindStaffs(where interface{}) Staffs {
	var staffs Staffs
	if where == nil {
		db.Find(&staffs)
	} else {
		db.Where(&where).Find(&staffs)
	}
	return staffs
}
