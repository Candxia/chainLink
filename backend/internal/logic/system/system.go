package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type sSystem struct{}
func init() {
	service.RegisterSystem(&sSystem{})
}

func (s *sSystem) GetDashboard(ctx context.Context) (*model.DashboardData, error) {
	totalUsers, _ := g.DB().Model("sys_user").Count()
	totalProducts, _ := g.DB().Model("product").Count()
	totalOrders, _ := g.DB().Model("order_info").Count()
	totalSuppliers, _ := g.DB().Model("supplier").Count()
	pendingOrders, _ := g.DB().Model("order_info").Where("status", "pending").Count()
	// 今日溯源记录
	todayStr := gtime.Now().Format("Y-m-d")
	todayRecords, _ := g.DB().Model("trace_record").WhereGTE("created_at", todayStr).Count()

	return &model.DashboardData{
		TotalUsers:        totalUsers,
		TotalProducts:     totalProducts,
		TotalOrders:       totalOrders,
		TotalSuppliers:    totalSuppliers,
		TodayTraceRecords: todayRecords,
		PendingOrders:     pendingOrders,
		TotalBlockCount:   100,
		ActiveContracts:   5,
	}, nil
}

func (s *sSystem) GetConfigList(ctx context.Context) ([]*model.SystemConfig, error) {
	configs, err := g.DB().Model("sys_config").All()
	if err != nil {
		return nil, err
	}
	var list []*model.SystemConfig
	for _, c := range configs {
		list = append(list, &model.SystemConfig{
			Id:    c["id"].Uint(),
			Key:   c["key"].String(),
			Value: c["value"].String(),
			Desc:  c["desc"].String(),
		})
	}
	return list, nil
}

func (s *sSystem) UpdateConfig(ctx context.Context, data map[string]interface{}) error {
	id := data["id"]
	if id == nil {
		return fmt.Errorf("缺少配置ID")
	}
	delete(data, "id")
	_, err := g.DB().Model("sys_config").Where("id", id).Update(data)
	return err
}

func (s *sSystem) GetLogList(ctx context.Context, page, pageSize int, logType string) ([]map[string]interface{}, int, error) {
	m := g.DB().Model("sys_log")
	if logType != "" {
		m = m.Where("type", logType)
	}
	total, _ := m.Count()
	logs, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	return logs.List(), total, nil
}

func (s *sSystem) DeleteLog(ctx context.Context, id uint) error {
	_, err := g.DB().Model("sys_log").Where("id", id).Delete()
	return err
}

func (s *sSystem) GetNotificationList(ctx context.Context, userId uint) ([]*model.NotificationInfo, error) {
	notifications, err := g.DB().Model("notification").Where("user_id", userId).Order("id desc").Limit(50).All()
	if err != nil {
		return nil, err
	}
	var list []*model.NotificationInfo
	for _, n := range notifications {
		list = append(list, &model.NotificationInfo{
			Id:      n["id"].Uint(),
			UserId:  n["user_id"].Uint(),
			Title:   n["title"].String(),
			Content: n["content"].String(),
			Type:    n["type"].String(),
			IsRead:  n["is_read"].Int(),
		})
	}
	return list, nil
}

func (s *sSystem) MarkNotificationRead(ctx context.Context, id uint) error {
	_, err := g.DB().Model("notification").Where("id", id).Update(g.Map{"is_read": 1})
	return err
}

func (s *sSystem) UploadFile(ctx context.Context, file *ghttp.UploadFile) (string, error) {
	url, err := file.Save("resource/public/upload")
	if err != nil {
		return "", err
	}
	return "/upload/" + url, nil
}
