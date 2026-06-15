package supply

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type sSupply struct{}
func init() {
	service.RegisterSupply(&sSupply{})
}

func (s *sSupply) CreateSupplier(ctx context.Context, data map[string]interface{}) (uint, error) {
	if data["name"] == nil {
		return 0, fmt.Errorf("供应商名称不能为空")
	}
	r, err := g.DB().Model("supplier").Insert(data)
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()
	return uint(id), nil
}

func (s *sSupply) UpdateSupplier(ctx context.Context, data map[string]interface{}) error {
	id := data["id"]
	if id == nil {
		return fmt.Errorf("缺少供应商ID")
	}
	delete(data, "id")
	_, err := g.DB().Model("supplier").Where("id", id).Update(data)
	return err
}

func (s *sSupply) GetSupplier(ctx context.Context, id uint) (*model.SupplierInfo, error) {
	supplier, err := g.DB().Model("supplier").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if supplier == nil {
		return nil, fmt.Errorf("供应商不存在")
	}
	return &model.SupplierInfo{
		Id:        supplier["id"].Uint(),
		Name:      supplier["name"].String(),
		Code:      supplier["code"].String(),
		Contact:   supplier["contact"].String(),
		Phone:     supplier["phone"].String(),
		Email:     supplier["email"].String(),
		Address:   supplier["address"].String(),
		Category:  supplier["category"].String(),
		QualLevel: supplier["qual_level"].String(),
		Status:    supplier["status"].String(),
	}, nil
}

func (s *sSupply) GetSupplierList(ctx context.Context, page, pageSize int, name string) ([]*model.SupplierInfo, int, error) {
	m := g.DB().Model("supplier")
	if name != "" {
		m = m.WhereLike("name", "%"+name+"%")
	}
	total, _ := m.Count()
	suppliers, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.SupplierInfo
	for _, sp := range suppliers {
		list = append(list, &model.SupplierInfo{
			Id:        sp["id"].Uint(),
			Name:      sp["name"].String(),
			Code:      sp["code"].String(),
			Contact:   sp["contact"].String(),
			Phone:     sp["phone"].String(),
			Status:    sp["status"].String(),
		})
	}
	return list, total, nil
}

func (s *sSupply) DeleteSupplier(ctx context.Context, id uint) error {
	_, err := g.DB().Model("supplier").Where("id", id).Delete()
	return err
}

func (s *sSupply) CreateOrder(ctx context.Context, data map[string]interface{}) (uint, error) {
	r, err := g.DB().Model("order_info").Insert(data)
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()
	return uint(id), nil
}

func (s *sSupply) UpdateOrder(ctx context.Context, data map[string]interface{}) error {
	id := data["id"]
	if id == nil {
		return fmt.Errorf("缺少订单ID")
	}
	delete(data, "id")
	_, err := g.DB().Model("order_info").Where("id", id).Update(data)
	return err
}

func (s *sSupply) GetOrder(ctx context.Context, id uint) (*model.OrderInfo, error) {
	order, err := g.DB().Model("order_info").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, fmt.Errorf("订单不存在")
	}
	return &model.OrderInfo{
		Id:          order["id"].Uint(),
		OrderNo:     order["order_no"].String(),
		ProductId:   order["product_id"].Uint(),
		SupplierId:  order["supplier_id"].Uint(),
		Quantity:    order["quantity"].Int(),
		TotalAmount: order["total_amount"].Float64(),
		Status:      order["status"].String(),
		Remark:      order["remark"].String(),
	}, nil
}

func (s *sSupply) GetOrderList(ctx context.Context, page, pageSize int, status string) ([]*model.OrderInfo, int, error) {
	m := g.DB().Model("order_info")
	if status != "" {
		m = m.Where("status", status)
	}
	total, _ := m.Count()
	orders, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.OrderInfo
	for _, o := range orders {
		list = append(list, &model.OrderInfo{
			Id:          o["id"].Uint(),
			OrderNo:     o["order_no"].String(),
			ProductId:   o["product_id"].Uint(),
			SupplierId:  o["supplier_id"].Uint(),
			Quantity:    o["quantity"].Int(),
			TotalAmount: o["total_amount"].Float64(),
			Status:      o["status"].String(),
		})
	}
	return list, total, nil
}

func (s *sSupply) DeleteOrder(ctx context.Context, id uint) error {
	_, err := g.DB().Model("order_info").Where("id", id).Delete()
	return err
}

func (s *sSupply) WarehouseOp(ctx context.Context, data map[string]interface{}, opType string) error {
	data["op_type"] = opType
	_, err := g.DB().Model("warehouse_record").Insert(data)
	return err
}

func (s *sSupply) GetStockList(ctx context.Context, page, pageSize int, productId uint) ([]*model.WarehouseRecord, int, error) {
	m := g.DB().Model("warehouse_record")
	if productId > 0 {
		m = m.Where("product_id", productId)
	}
	total, _ := m.Count()
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.WarehouseRecord
	for _, r := range records {
		list = append(list, &model.WarehouseRecord{
			Id:        r["id"].Uint(),
			OrderId:   r["order_id"].Uint(),
			ProductId: r["product_id"].Uint(),
			OpType:    r["op_type"].String(),
			Quantity:  r["quantity"].Int(),
			Operator:  r["operator"].String(),
			Remark:    r["remark"].String(),
		})
	}
	return list, total, nil
}

func (s *sSupply) CreateLogisticsRecord(ctx context.Context, data map[string]interface{}) error {
	_, err := g.DB().Model("logistics_record").Insert(data)
	return err
}

func (s *sSupply) GetLogisticsList(ctx context.Context, orderId uint) ([]*model.LogisticsRecord, error) {
	records, err := g.DB().Model("logistics_record").Where("order_id", orderId).Order("id asc").All()
	if err != nil {
		return nil, err
	}
	var list []*model.LogisticsRecord
	for _, r := range records {
		list = append(list, &model.LogisticsRecord{
			Id:       r["id"].Uint(),
			OrderId:  r["order_id"].Uint(),
			NodeName: r["node_name"].String(),
			Location: r["location"].String(),
			Status:   r["status"].String(),
			Operator: r["operator"].String(),
			Remark:   r["remark"].String(),
		})
	}
	return list, nil
}
