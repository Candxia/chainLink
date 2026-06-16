package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/consts"
	"cl_system/internal/model/entity"
	"cl_system/internal/service"
)

type sMenu struct{}

func init() {
	service.RegisterMenu(NewMenu())
}

func NewMenu() service.IMenu {
	return &sMenu{}
}

func (s *sMenu) validate(ctx context.Context, in mdlSys.SysMenuAdd) (err error) {
	if in.MenuType != consts.MenuDir && in.ParentId == 0 {
		return fmt.Errorf("非目录类型菜单必须指定上级菜单")
	}
	if in.ParentId > 0 {
		parent, errs := s.Info(ctx, in.ParentId)
		if errs != nil {
			return errs
		}
		switch in.MenuType {
		case consts.MenuButton:
			if !(parent.MenuType == consts.MenuOwn || parent.MenuType == consts.MenuTag) {
				return fmt.Errorf("按钮必须挂在菜单或页签下")
			}
		case consts.MenuTag:
			if parent.MenuType != consts.MenuOwn {
				return fmt.Errorf("页签必须挂在菜单下")
			}
		case consts.MenuOwn:
			if parent.MenuType != consts.MenuDir {
				return fmt.Errorf("菜单必须挂在路由下")
			}
		case consts.MenuDir:
			if parent.MenuType != consts.MenuDir {
				return fmt.Errorf("路由必须挂在路由下")
			}
		}
	}
	return nil
}

func (s *sMenu) Add(ctx context.Context, in mdlSys.SysMenuAdd) (err error) {
	if err = s.validate(ctx, in); err != nil {
		return err
	}
	if in.MenuType == consts.MenuDir {
		in.ApiId = []int64{}
	}
	_, err = g.DB().Model("sys_menu").Insert(in)
	return err
}

func (s *sMenu) Del(ctx context.Context, id int64) (err error) {
	// 判断是否存在子菜单
	res, _ := g.DB().Model("sys_menu").Where("parent_id", id).Fields("menu_id").One()
	if !res.IsEmpty() {
		return fmt.Errorf("存在子菜单，无法删除")
	}
	_, err = g.DB().Model("sys_menu").Where("menu_id", id).Delete()
	return err
}

func (s *sMenu) Edit(ctx context.Context, in mdlSys.SysMenuEdit) (err error) {
	if err = s.validate(ctx, in.SysMenuAdd); err != nil {
		return err
	}
	if in.MenuType == consts.MenuDir {
		in.ApiId = []int64{}
	}
	_, err = g.DB().Model("sys_menu").Update(in, g.Map{"menu_id": in.MenuId})
	return err
}

func (s *sMenu) Info(ctx context.Context, id int64) (res entity.SysMenu, err error) {
	record, err := g.DB().Model("sys_menu").Where("menu_id", id).One()
	if err != nil {
		return res, err
	}
	if record.IsEmpty() {
		return res, fmt.Errorf("菜单不存在")
	}
	err = record.Struct(&res)
	return res, err
}

func (s *sMenu) List(ctx context.Context) (list []mdlSys.SysMenuTree, err error) {
	var all []entity.SysMenu
	list = make([]mdlSys.SysMenuTree, 0)
	err = g.DB().Model("sys_menu").OrderAsc("order_num").Scan(&all)
	if err != nil || len(all) == 0 {
		return list, err
	}
	list = menuListTree(all, consts.TopMenu)
	return
}

func menuListTree(all []entity.SysMenu, parentId int64) (list []mdlSys.SysMenuTree) {
	list = make([]mdlSys.SysMenuTree, 0)
	for _, v := range all {
		if v.ParentId == parentId {
			one := mdlSys.SysMenuTree{
				Title:     v.Title,
				Icon:      v.Icon,
				Component: v.Component,
				OrderNum:  v.OrderNum,
				Visible:   v.Visible,
				MenuId:    v.MenuId,
				MenuType:  v.MenuType,
				Path:      v.Path,
				CreatedAt: v.CreatedAt,
			}
			if c := menuListTree(all, v.MenuId); len(c) > 0 {
				one.Children = c
			}
			list = append(list, one)
		}
	}
	return list
}

func (s *sMenu) Role(ctx context.Context) (menus []mdlSys.SysMenuRole, err error) {
	menus = make([]mdlSys.SysMenuRole, 0)
	var all []entity.SysMenu
	err = g.DB().Model("sys_menu").Where("status = ?", consts.OnEnable).
		OrderAsc("order_num").Scan(&all)
	if err != nil || len(all) == 0 {
		return menus, err
	}
	menus = menuRoleTree(all, 1, consts.TopMenu)
	return
}

func menuRoleTree(all []entity.SysMenu, lv int, parentId int64) (list []mdlSys.SysMenuRole) {
	list = make([]mdlSys.SysMenuRole, 0)
	for _, v := range all {
		if lv >= consts.RoleMaxLv && v.MenuName == "System" {
			continue
		}
		if v.ParentId == parentId && v.MenuType != consts.MenuButton {
			one := mdlSys.SysMenuRole{
				Path:      v.Path,
				Name:      v.MenuName,
				Component: v.Component,
				Meta: mdlSys.SysMenuMeta{
					Title:   v.Title,
					Icon:    v.Icon,
					NoCache: v.IsCache == consts.EnumNot,
					Hidden:  v.Visible == consts.EnumNot,
				},
				Permission: make([]string, 0),
				Tags:       make([]string, 0),
				Children:   make([]mdlSys.SysMenuRole, 0),
			}
			if v.MenuType == consts.MenuOwn {
				one.Permission = getButton(all, v.MenuId)
				t, b := getTags(all, v.MenuId)
				one.Tags = t
				one.Permission = append(one.Permission, b...)
			} else {
				one.Children = menuRoleTree(all, lv, v.MenuId)
			}
			list = append(list, one)
		}
	}
	return list
}

func getButton(all []entity.SysMenu, parentId int64) (b []string) {
	b = make([]string, 0)
	for _, v := range all {
		if v.ParentId == parentId && v.MenuType == consts.MenuButton {
			b = append(b, v.MenuName)
		}
	}
	return b
}

func getTags(all []entity.SysMenu, parentId int64) (t []string, b []string) {
	t = make([]string, 0)
	b = make([]string, 0)
	for _, v := range all {
		if v.ParentId == parentId && v.MenuType == consts.MenuTag {
			t = append(t, v.MenuName)
			b = append(b, getButton(all, v.MenuId)...)
		}
	}
	return t, b
}

func (s *sMenu) Drop(ctx context.Context, exMenutype []string) (list []g.Map, total int, err error) {
	var all []entity.SysMenu
	err = g.DB().Model("sys_menu").Where("status = ?", consts.OnEnable).
		Fields("menu_id,title,menu_name,parent_id,menu_type").
		OrderAsc("order_num").Scan(&all)
	if err != nil || len(all) == 0 {
		return list, 0, err
	}
	ex := make(map[string]bool)
	for _, v := range exMenutype {
		ex[v] = true
	}
	var filtered []entity.SysMenu
	for _, v := range all {
		if ex[v.MenuType] {
			continue
		}
		filtered = append(filtered, v)
	}
	treeList := menuDropTree(filtered, consts.TopMenu)
	return treeList, 1, nil
}

func menuDropTree(all []entity.SysMenu, parent int64) []g.Map {
	list := make([]g.Map, 0)
	for _, v := range all {
		if v.ParentId == parent {
			item := g.Map{
				"id":    v.MenuId,
				"label": v.Title,
			}
			if c := menuDropTree(all, v.MenuId); len(c) > 0 {
				item["children"] = c
			}
			list = append(list, item)
		}
	}
	return list
}
