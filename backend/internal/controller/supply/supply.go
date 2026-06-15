package supply

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) CreateSupplier(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Supply().CreateSupplier(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateSupplier(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().UpdateSupplier(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetSupplier(r *ghttp.Request) {
	id := r.Get("id").Uint()
	supplier, err := service.Supply().GetSupplier(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": supplier})
}

func (c *Controller) GetSupplierList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	name := r.GetQuery("name").String()
	list, total, err := service.Supply().GetSupplierList(r.Context(), page, pageSize, name)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteSupplier(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Supply().DeleteSupplier(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) CreateOrder(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Supply().CreateOrder(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateOrder(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().UpdateOrder(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetOrder(r *ghttp.Request) {
	id := r.Get("id").Uint()
	order, err := service.Supply().GetOrder(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": order})
}

func (c *Controller) GetOrderList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	status := r.GetQuery("status").String()
	list, total, err := service.Supply().GetOrderList(r.Context(), page, pageSize, status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteOrder(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Supply().DeleteOrder(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) WarehouseIn(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().WarehouseOp(r.Context(), data, "in")
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "入库成功"})
}

func (c *Controller) WarehouseOut(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().WarehouseOp(r.Context(), data, "out")
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "出库成功"})
}

func (c *Controller) GetStockList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	productId := r.GetQuery("productId").Uint()
	list, total, err := service.Supply().GetStockList(r.Context(), page, pageSize, productId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateLogisticsRecord(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().CreateLogisticsRecord(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "物流记录创建成功"})
}

func (c *Controller) GetLogisticsList(r *ghttp.Request) {
	orderId := r.Get("orderId").Uint()
	list, err := service.Supply().GetLogisticsList(r.Context(), orderId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}
