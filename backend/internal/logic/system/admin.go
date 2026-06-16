package system

import (
	"context"
	"fmt"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
	"cl_system/utility"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(NewAdmin())
}

func NewAdmin() service.IAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Add(ctx context.Context, in mdlSys.SysAdminAdd) (err error) {
	pwd, _ := utility.Password.Encrypt(in.Password)
	now := time.Now().Unix()
	_, err = g.DB().Model("sys_admin").Data(g.Map{
		"username":   in.Username,
		"password":   pwd,
		"nickname":   in.Nickname,
		"role_id":    in.RoleId,
		"dept_id":    in.DeptId,
		"status":     2,
		"online":     2,
		"created_at": now,
	}).Insert()
	return err
}

func (s *sAdmin) Del(ctx context.Context, id int64) (err error) {
	if id <= 1 {
		return fmt.Errorf("不能删除超级管理员")
	}
	_, err = g.DB().Model("sys_admin").Where("id", id).Delete()
	return err
}

func (s *sAdmin) Edit(ctx context.Context, in mdlSys.SysAdminEdit) (err error) {
	if in.Id <= 1 {
		return fmt.Errorf("不能编辑超级管理员")
	}
	_, err = g.DB().Model("sys_admin").Update(in.Update(), g.Map{"id": in.Id, "username": in.Username})
	return err
}

func (s *sAdmin) Info(ctx context.Context, id int64) (out mdlSys.SysAdminInfo, err error) {
	if id <= 1 {
		return out, fmt.Errorf("无效的ID")
	}
	res, err := g.DB().Model("sys_admin").Where("id=?", id).One()
	if err != nil {
		return out, err
	}
	err = res.Struct(&out)
	return out, err
}

func (s *sAdmin) List(ctx context.Context, in mdlSys.SysAdminSearch) (list []mdlSys.SysAdminInfo, total int, err error) {
	list = make([]mdlSys.SysAdminInfo, 0, in.Limit)
	where := in.Condition()
	m := g.DB().Model("sys_admin").Where(where)
	total, err = m.Count()
	if err != nil {
		return list, total, err
	}
	if total > 0 {
		limit, offset := in.Paginate()
		err = m.Limit(limit).Offset(offset).LeftJoin("sys_role", "sys_role.id=sys_admin.role_id").
			Fields("sys_admin.*,sys_role.name as role_name").OrderDesc("sys_admin.created_at").Scan(&list)
	}
	return list, total, err
}

func (s *sAdmin) Password(ctx context.Context, id int64, pwd string) (err error) {
	if id <= 1 {
		return fmt.Errorf("不能修改超级管理员密码")
	}
	password, _ := utility.Password.Encrypt(pwd)
	_, err = g.DB().Model("sys_admin").Where("id", id).Data(g.Map{"password": password}).Update()
	return err
}

func (s *sAdmin) Status(ctx context.Context, username string, status int) (err error) {
	data := g.Map{"status": status}
	if status == 2 {
		data["online"] = 2
	}
	_, err = g.DB().Model("sys_admin").Where("username", username).Where("id > 1").Data(data).Update()
	return err
}

func (s *sAdmin) ClickOut(ctx context.Context, username string) (err error) {
	_, err = g.DB().Model("sys_admin").Where("username", username).Where("id > 1").
		Data(g.Map{"online": 2, "remark": ""}).Update()
	return err
}

func (s *sAdmin) SetOnline(ctx context.Context, info mdlSys.SysAdminOnline) (err error) {
	data := g.Map{"online": 1, "remark": info.AuthKey}
	_, err = g.DB().Model("sys_admin").Update(data, info.Condition())
	return err
}
