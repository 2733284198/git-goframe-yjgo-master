// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package answer

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table t_answer.
type Entity struct {
	Id         int64       `orm:"id,primary"  json:"id"`          // 主键
	Pid        int64       `orm:"pid"         json:"pid"`         // 问题ID
	Atype      int         `orm:"atype"       json:"atype"`       // 回复人类别
	UserId     int64       `orm:"user_id"     json:"user_id"`     // 回复人ID
	NickName   string      `orm:"nick_name"   json:"nick_name"`   // 回复人名称
	Avatar     string      `orm:"avatar"      json:"avatar"`      // 回复人头像
	Remark     string      `orm:"remark"      json:"remark"`      // 回复内容
	Img1       string      `orm:"img1"        json:"img_1"`       // 回复图片1
	Img2       string      `orm:"img2"        json:"img_2"`       // 回复图片2
	Img3       string      `orm:"img3"        json:"img_3"`       // 回复图片3
	Status     int         `orm:"status"      json:"status"`      // 状态
	CreateTime *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime *gtime.Time `orm:"update_time" json:"update_time"` // 更新时间
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
