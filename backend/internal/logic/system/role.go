package system

import (
	"context"
	"fmt"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/consts"
	"cl_system/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(NewRole())
}

func NewRole() service.IRole {
	return &sRole{}
}

func (s *sRole) exist(ctx context.Context, id int64) bool {
	res, _ := g.DB().Model("sys_role").Where("id", id).Where("is_del", consts.EnumNot).Fields("id").One()
	return !res.IsEmpty()
}

func (s *sRole) existUser(ctx context.Context, id int64) bool {
	res, _ := g.DB().Model("sys_admin").Where("role_id", id).Fields("id").One()
	return !res.IsEmpty()
}

func (s *sRole) Status(ctx context.Context, in mdlSys.SysRoleStatus) (err error) {
	info, errs := s.Info(ctx, in.Id)
	if errs != nil {
		return errs
	}
	if info.Status == in.Status {
		return nil
	}
	now := time.Now().Unix()
	_, err = g.DB().Model("sys_role").Update(g.Map{
		"status":     in.Status,
		"updated_at": now,
	}, g.Map{"id": in.Id})
	return err
}

func (s *sRole) Add(ctx context.Context, in mdlSys.SysRoleAdd) (err error) {
	if !in.Valid() {
		return fmt.Errorf("权限参数无效")
	}
	now := time.Now().Unix()
	_, err = g.DB().Model("sys_role").Insert(g.Map{
		"parent_id":      0,
		"code":           fmt.Sprintf("role_%d", now),
		"name":           in.Name,
		"status":         in.Status,
		"remark":         in.Remark,
		"menu_all":       in.MenuAll,
		"is_show_mobile": in.IsShowMobile,
		"is_del":         consts.EnumNot,
		"level":          1,
		"created_at":     now,
	})
	return err
}

func (s *sRole) Del(ctx context.Context, id int64) (err error) {
	if id <= consts.RoleSysId {
		return fmt.Errorf("不能删除系统角色")
	}
	if s.existUser(ctx, id) {
		return fmt.Errorf("该角色下存在管理员，无法删除")
	}
	res, _ := g.DB().Model("sys_role").Where(g.Map{"is_del": consts.EnumNot, "parent_id": id}).Fields("id").One()
	if !res.IsEmpty() {
		return fmt.Errorf("该角色下存在子角色，无法删除")
	}
	_, err = g.DB().Model("sys_role").Update(g.Map{"is_del": consts.EnumIs}, g.Map{"id": id})
	return err
}

func (s *sRole) Edit(ctx context.Context, in mdlSys.SysRoleEdit) (err error) {
	if !in.Valid() {
		return fmt.Errorf("权限参数无效")
	}
	if !s.exist(ctx, in.Id) {
		return fmt.Errorf("角色不存在")
	}
	now := time.Now().Unix()
	_, err = g.DB().Model("sys_role").Update(g.Map{
		"name":           in.Name,
		"status":         in.Status,
		"remark":         in.Remark,
		"menu_all":       in.MenuAll,
		"is_show_mobile": in.IsShowMobile,
		"updated_at":     now,
		"updated_by":     in.UpdatedBy,
	}, g.Map{"id": in.Id})
	return err
}

func (s *sRole) Info(ctx context.Context, id int64) (res mdlSys.SysRoleInfo, err error) {
	if id < consts.RoleSysId {
		return res, fmt.Errorf("角色不存在")
	}
	record, err := g.DB().Model("sys_role").Where("id", id).One()
	if err != nil {
		return res, err
	}
	if record.IsEmpty() {
		return res, fmt.Errorf("角色不存在")
	}
	err = record.Struct(&res)
	if err != nil {
		return res, err
	}
	// 加载关联菜单ID
	if res.MenuAll != consts.EnumIs {
		menuRecords, _ := g.DB().Model("sys_role_menu").Fields("menu_id").Where("role_id", id).All()
		for _, m := range menuRecords {
			res.MenuIds = append(res.MenuIds, m["menu_id"].Int64())
		}
	}
	return res, err
}

func (s *sRole) List(ctx context.Context, in mdlSys.SysRoleSearch) (list []mdlSys.SysRoleTree, total int, err error) {
	list = make([]mdlSys.SysRoleTree, 0)
	where := in.Condition()
	m := g.DB().Model("sys_role").Where(where)
	total, err = m.Count()
	if err != nil {
		return list, total, err
	}
	if total > 0 {
		var all []mdlSys.SysRoleList
		allRecords, err2 := m.Fields("id,parent_id,code,name,status,is_del,is_show_mobile,level,remark,menu_all,created_at,created_by,updated_at,updated_by").OrderAsc("parent_id").All()
		if err2 != nil {
			return list, total, err2
		}
		for _, r := range allRecords {
			var item mdlSys.SysRoleList
			_ = r.Struct(&item)
			all = append(all, item)
		}
		if len(all) > 0 {
			list = treeList(all, 0)
		}
	}
	return list, total, nil
}

func treeList(all []mdlSys.SysRoleList, parent int64) (list []mdlSys.SysRoleTree) {
	list = make([]mdlSys.SysRoleTree, 0)
	for _, v := range all {
		if v.ParentId == parent {
			t := mdlSys.SysRoleTree{
				SysRoleList: v,
				Children:    treeList(all, v.Id),
			}
			list = append(list, t)
		}
	}
	return list
}

func (s *sRole) Drop(ctx context.Context) (list []g.Map, err error) {
	list = make([]g.Map, 0)
	err = g.DB().Model("sys_role").Where("is_del", consts.EnumNot).
		Fields("id as value,name as label").
		OrderDesc("created_at").Scan(&list)
	return list, err
}

func (s *sRole) IsMobile(ctx context.Context, params mdlSys.SysRoleIsMobile) (err error) {
	_, err = g.DB().Model("sys_role").Update(g.Map{"is_show_mobile": params.IsShowMobile}, g.Map{"id": params.Id})
	return err
}
