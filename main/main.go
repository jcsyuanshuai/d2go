package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type BloodSugar struct {
	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PatientId  string `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"` // 患者id
	RecordTime string `protobuf:"bytes,3,opt,name=record_time,json=recordTime,proto3" json:"record_time,omitempty"`
	DiningType string `protobuf:"bytes,4,opt,name=dining_type,json=diningType,proto3" json:"dining_type,omitempty"` // 用餐类型
	Value      string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Creator    string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`                         // 创建人
	CreateTime string `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间
	Updater    string `protobuf:"bytes,8,opt,name=updater,proto3" json:"updater,omitempty"`
	UpdateTime string `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Remark     string `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (BloodSugar) TableName() string {
	return "blood_sugar"
}

func main() {
	d := "2022-01-03 15:00:01"
	t1, err := time.Parse("2006-01-02 15:04:05", d)
	if err != nil {
		return
	}

	t2 := t1.AddDate(0, 0, 1)

	s1 := t1.Format("2006-01-02")
	s2 := t2.Format("2006-01-02")

	println(s1)
	println(s2)

	dsn := "root:zjyNJBrmerOIWgtai9wtjxLRh8djdJVa@tcp(smart.hub.biomind.com.cn:3306)/icdss?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}

	get := &BloodSugar{}

	db.Debug().Where(
		"patient_id = ? and dining_type = ? and record_time > ? and record_time < ?",
		"111", "111",
		s1,
		s2,
	).Find(&get)

	fmt.Printf("%v\n", get)

	ret := &BloodSugar{
		Id:         get.Id,
		PatientId:  get.PatientId,
		RecordTime: get.RecordTime,
		DiningType: get.DiningType,
		Value:      get.Value + "2",
		Creator:    get.Creator,
		CreateTime: get.CreateTime,
		Updater:    get.Updater,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		Remark:     get.Remark,
	}

	db.Debug().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"patient_id", "record_time", "dining_type", "value", "updater", "update_time",
		}),
	}).Create(ret)

	//db.Debug().Omit("id").Clauses(clause.OnConflict{
	//    DoUpdates: clause.AssignmentColumns([]string{
	//        "patient_id", "record_time", "dining_type", "value", "updater", "update_time",
	//    }),
	//}).Create(ret)

	fmt.Printf("%v\n", ret)
}
