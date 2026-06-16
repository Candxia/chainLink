package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/model/entity"
	"cl_system/internal/service"
)

type sApi struct{}

func init() {
	service.RegisterApi(NewApi())
}

func NewApi() service.IApi {
	return &sApi{}
}

func (s *sApi) exist(ctx context.Context, in mdlSys.SysApiAdd) bool {
	where := g.Map{
		"path":   in.Path,
		"action": in.Action,
	}
	res, _ := g.DB().Model("sys_api").Where(where).Fields("api_id").One()
	return !res.IsEmpty()
}

func (s *sApi) Add(ctx context.Context, in mdlSys.SysApiAdd) (err error) {
	if s.exist(ctx, in) {
		return fmt.Errorf("API已存在")
	}
	_, err = g.DB().Model("sys_api").Insert(in.Data())
	return err
}

func (s *sApi) Del(ctx context.Context, id int64) (err error) {
	_, err = g.DB().Model("sys_api").Where("api_id", id).Delete()
	return err
}

func (s *sApi) Edit(ctx context.Context, in mdlSys.SysApiEdit) (err error) {
	_, err = g.DB().Model("sys_api").Where("api_id", in.ApiId).Data(in.Data()).Update()
	return err
}

func (s *sApi) List(ctx context.Context, in mdlSys.SysApiSearch) (list []entity.SysApi, total int, err error) {
	list = make([]entity.SysApi, 0, in.Limit)
	where := in.Condition()
	m := g.DB().Model("sys_api").Where(where)
	total, err = m.Count()
	if err != nil {
		return list, total, err
	}
	if total > 0 {
		limit, offset := in.Paginate()
		err = m.Limit(offset, limit).OrderDesc("api_id").Scan(&list)
	}
	return list, total, err
}

func (s *sApi) Drop(ctx context.Context) (list []g.Map, err error) {
	err = g.DB().Model("sys_api").Fields("title AS label,api_id AS value").
		OrderDesc("api_id").Scan(&list)
	return list, err
}

func (s *sApi) Path(ctx context.Context) (list []g.Map, err error) {
	err = g.DB().Model("sys_api").Fields("title AS label,path AS value").
		Where("logger=1").OrderDesc("api_id").Scan(&list)
	return list, err
}
