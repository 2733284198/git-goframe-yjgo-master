// ==========================================================================
// 云捷GO自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-12-07 01:38:53
// 生成路径: app/service/module/zjuser/zjuser_service.go
// 生成人：yunjie
// ==========================================================================
package zjuser

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	zjuserModel "yj-app/app/model/module/zjuser"
	userService "yj-app/app/service/system/user"
	"yj-app/app/utils/convert"
	"yj-app/app/utils/excel"
	"yj-app/app/utils/page"
)

//根据主键查询数据
func SelectRecordById(id int64) (*zjuserModel.Entity, error) {
	return zjuserModel.FindOne("id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := zjuserModel.Delete("id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := zjuserModel.Delete("id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *zjuserModel.AddReq, session *ghttp.Session) (int64, error) {
	var entity zjuserModel.Entity
	 
	entity.Id = req.Id  
	entity.Uname = req.Uname  
	entity.Puid = req.Puid  
	entity.Puname = req.Puname  
	entity.RealName = req.RealName  
	entity.Idno = req.Idno  
	entity.Avatar = req.Avatar  
	entity.Job = req.Job  
	entity.Utype = req.Utype  
	entity.Upwd = req.Upwd  
	entity.Salt = req.Salt  
	entity.Idpic1 = req.Idpic1  
	entity.Idpic2 = req.Idpic2  
	entity.Pimg1 = req.Pimg1  
	entity.Pimg2 = req.Pimg2  
	entity.Pimg3 = req.Pimg3  
	entity.Status = req.Status  
	entity.ImMoney = req.ImMoney  
	entity.TelMoney = req.TelMoney     
	entity.CreateTime = gtime.Now()
	entity.CreateBy = ""

	user := userService.GetProfile(session)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	result, err := entity.Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

//修改数据
func EditSave(req *zjuserModel.EditReq, session *ghttp.Session) (int64, error) {

	entity, err := zjuserModel.FindOne("id=?", req.Id)

	if err != nil {
		return 0, err
	}

	if entity == nil {
		return 0, gerror.New("数据不存在")
	}
	   
	entity.Uname = req.Uname  
	entity.Puid = req.Puid  
	entity.Puname = req.Puname  
	entity.RealName = req.RealName  
	entity.Idno = req.Idno  
	entity.Avatar = req.Avatar  
	entity.Job = req.Job  
	entity.Utype = req.Utype  
	entity.Upwd = req.Upwd  
	entity.Salt = req.Salt  
	entity.Idpic1 = req.Idpic1  
	entity.Idpic2 = req.Idpic2  
	entity.Pimg1 = req.Pimg1  
	entity.Pimg2 = req.Pimg2  
	entity.Pimg3 = req.Pimg3  
	entity.Status = req.Status  
	entity.ImMoney = req.ImMoney  
	entity.TelMoney = req.TelMoney     
	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = ""

	user := userService.GetProfile(session)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	result, err := entity.Update()

	if err != nil {
		return 0, err
	}

	rs, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rs, nil
}

//根据条件查询数据
func SelectListAll(params *zjuserModel.SelectPageReq) ([]zjuserModel.Entity, error) {
	return zjuserModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListByPage(params *zjuserModel.SelectPageReq) ([]zjuserModel.Entity, *page.Paging, error) {
	return zjuserModel.SelectListByPage(params)
}

// 导出excel
func Export(param *zjuserModel.SelectPageReq) (string, error) {
	result, err := zjuserModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{  "主键" ,"手机号" ,"推荐人ID" ,"推荐人手机号" ,"姓名" ,"身份证号" ,"头像" ,"职业" ,"用户类别" ,"密码" ,"密码盐" ,"身份证正面" ,"身份证反面" ,"职业资格认证1" ,"职业资格认证2" ,"职业资格认证3" ,"状态" ,"在线咨询费用" ,"电话咨询费用" ,"创建时间" ,"更新时间"}
	key := []string{  "id" ,"uname" ,"puid" ,"puname" ,"real_name" ,"idno" ,"avatar" ,"job" ,"utype" ,"upwd" ,"salt" ,"idpic1" ,"idpic2" ,"pimg1" ,"pimg2" ,"pimg3" ,"status" ,"im_money" ,"tel_money" ,"create_time" ,"update_time"}
	url, err := excel.DownlaodExcel(head,key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}