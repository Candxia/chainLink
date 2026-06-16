package system

import (
	"context"
	"fmt"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/model/entity"
	"cl_system/internal/service"
)

type sDept struct{}

func init() {
	service.RegisterDept(NewDept())
}

func NewDept() service.IDept {
	return &sDept{}
}

func (s *sDept) Add(ctx context.Context, in mdlSys.SysDeptAdd) (err error) {
	now := time.Now().Unix()
	_, err = g.DB().Model("sys_dept").Data(g.Map{
		"parent_id":  in.ParentId,
		"name":       in.Name,
		"order_num":  in.OrderNum,
		"leader":     in.Leader,
		"phone":      in.Phone,
		"email":      in.Email,
		"status":     in.Status,
		"is_del":     2,
		"created_at": now,
		"updated_at": now,
	}).Insert()
	return err
}

func (s *sDept) Del(ctx context.Context, id int64) (err error) {
	// 检查是否有子部门
	count, err := g.DB().Model("sys_dept").Where("parent_id", id).Where("is_del", 2).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该部门存在子部门，无法删除")
	}
	// 检查是否有管理员关联
	adminCount, err := g.DB().Model("sys_admin").Where("dept_id", id).Count()
	if err != nil {
		return err
	}
	if adminCount > 0 {
		return fmt.Errorf("该部门存在关联管理员，无法删除")
	}
	_, err = g.DB().Model("sys_dept").Where("id", id).Update(g.Map{"is_del": 1, "updated_at": time.Now().Unix()})
	return err
}

func (s *sDept) Exist(ctx context.Context, in mdlSys.SysDeptExist) (has bool) {
	where := in.Condition()
	count, err := g.DB().Model("sys_dept").Where(where).Where("is_del", 2).Count()
	if err != nil {
		return false
	}
	return count > 0
}

func (s *sDept) Edit(ctx context.Context, in mdlSys.SysDeptEdit) (err error) {
	data := in.Update()
	data["updated_at"] = time.Now().Unix()
	_, err = g.DB().Model("sys_dept").Where("id", in.Id).Update(data)
	return err
}

func (s *sDept) Info(ctx context.Context, id int64) (result entity.SysDept, err error) {
	res, err := g.DB().Model("sys_dept").Where("id", id).Where("is_del", 2).One()
	if err != nil {
		return result, err
	}
	if res == nil {
		return result, fmt.Errorf("部门不存在或已删除")
	}
	err = res.Struct(&result)
	return result, err
}

func (s *sDept) List(ctx context.Context, in mdlSys.SysDeptSearch) (result []mdlSys.SysDeptInfo, err error) {
	result = make([]mdlSys.SysDeptInfo, 0)
	where := in.Condition()
	where["is_del"] = 2

	m := g.DB().Model("sys_dept").Where(where)
	all, err := m.Order("order_num asc, id asc").All()
	if err != nil {
		return result, err
	}
	// 转换为树形结构
	var list []mdlSys.SysDeptInfo
	for _, v := range all {
		list = append(list, mdlSys.SysDeptInfo{
			Id:        v["id"].Int64(),
			ParentId:  v["parent_id"].Int64(),
			OrderNum:  v["order_num"].Int(),
			Name:      v["name"].String(),
			Leader:    v["leader"].String(),
			Phone:     v["phone"].String(),
			Email:     v["email"].String(),
			Status:    v["status"].Int(),
			CreatedAt: v["created_at"].Int64(),
		})
	}
	result = buildDeptTree(list, 0)
	return result, err
}

func (s *sDept) Drop(ctx context.Context) (result []mdlSys.SysDeptInfo, err error) {
	result = make([]mdlSys.SysDeptInfo, 0)
	all, err := g.DB().Model("sys_dept").Where("status", 1).Where("is_del", 2).Order("order_num asc, id asc").All()
	if err != nil {
		return result, err
	}
	var list []mdlSys.SysDeptInfo
	for _, v := range all {
		list = append(list, mdlSys.SysDeptInfo{
			Id:       v["id"].Int64(),
			ParentId: v["parent_id"].Int64(),
			Name:     v["name"].String(),
		})
	}
	result = buildDeptTree(list, 0)
	return result, err
}

// buildDeptTree 递归构建部门树
func buildDeptTree(list []mdlSys.SysDeptInfo, parentId int64) []mdlSys.SysDeptInfo {
	var tree []mdlSys.SysDeptInfo
	for _, v := range list {
		if v.ParentId == parentId {
			children := buildDeptTree(list, v.Id)
			if len(children) > 0 {
				v.Children = children
			}
			tree = append(tree, v)
		}
	}
	return tree
}
