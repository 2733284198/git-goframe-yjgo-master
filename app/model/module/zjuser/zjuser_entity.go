// ==========================================================================
// 云捷GO自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2020-12-07 01:38:53
// 生成路径: app/model/module/zjuser/zjuser_entity.go
// 生成人：yunjie
// ==========================================================================

package zjuser

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table t_zjuser.
type Entity struct { 
	 Id       int64         `orm:"id,primary" json:"id"`    // 主键  
	 Uname    string         `orm:"uname" json:"uname"`    // 手机号  
	 Puid    int64         `orm:"puid" json:"puid"`    // 推荐人ID  
	 Puname    string         `orm:"puname" json:"puname"`    // 推荐人手机号  
	 RealName    string         `orm:"real_name" json:"real_name"`    // 姓名  
	 Idno    string         `orm:"idno" json:"idno"`    // 身份证号  
	 Avatar    string         `orm:"avatar" json:"avatar"`    // 头像  
	 Job    string         `orm:"job" json:"job"`    // 职业  
	 Utype    int         `orm:"utype" json:"utype"`    // 用户类别  
	 Upwd    string         `orm:"upwd" json:"upwd"`    // 密码  
	 Salt    string         `orm:"salt" json:"salt"`    // 密码盐  
	 Idpic1    string         `orm:"idpic1" json:"idpic1"`    // 身份证正面  
	 Idpic2    string         `orm:"idpic2" json:"idpic2"`    // 身份证反面  
	 Pimg1    string         `orm:"pimg1" json:"pimg1"`    // 职业资格认证1  
	 Pimg2    string         `orm:"pimg2" json:"pimg2"`    // 职业资格认证2  
	 Pimg3    string         `orm:"pimg3" json:"pimg3"`    // 职业资格认证3  
	 Status    int         `orm:"status" json:"status"`    // 状态  
	 ImMoney    float64         `orm:"im_money" json:"im_money"`    // 在线咨询费用  
	 TelMoney    float64         `orm:"tel_money" json:"tel_money"`    // 电话咨询费用  
	 CreateTime    *gtime.Time         `orm:"create_time" json:"create_time"`    // 创建时间  
	 UpdateTime    *gtime.Time         `orm:"update_time" json:"update_time"`    // 更新时间  
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
