package supply

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	supplyV1 "cl_system/api/supply/v1"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) CreateSupplier(r *ghttp.Request) {
	var req supplyV1.CreateSupplierReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Supply().CreateSupplier(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateSupplier(r *ghttp.Request) {
	var req supplyV1.UpdateSupplierReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().UpdateSupplier(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetSupplier(r *ghttp.Request) {
	var req supplyV1.GetSupplierReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	supplier, err := service.Supply().GetSupplier(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": supplier})
}

func (c *Controller) GetSupplierList(r *ghttp.Request) {
	var req supplyV1.GetSupplierListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Supply().GetSupplierList(r.Context(), req.Page, req.PageSize, req.Name)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteSupplier(r *ghttp.Request) {
	var req supplyV1.DeleteSupplierReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().DeleteSupplier(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) CreateOrder(r *ghttp.Request) {
	var req supplyV1.CreateOrderReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Supply().CreateOrder(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateOrder(r *ghttp.Request) {
	var req supplyV1.UpdateOrderReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().UpdateOrder(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetOrder(r *ghttp.Request) {
	var req supplyV1.GetOrderReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	order, err := service.Supply().GetOrder(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": order})
}

func (c *Controller) GetOrderList(r *ghttp.Request) {
	var req supplyV1.GetOrderListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Supply().GetOrderList(r.Context(), req.Page, req.PageSize, req.Status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteOrder(r *ghttp.Request) {
	var req supplyV1.DeleteOrderReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().DeleteOrder(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) WarehouseIn(r *ghttp.Request) {
	var req supplyV1.WarehouseInReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().WarehouseOp(r.Context(), req.Data, "in")
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "入库成功"})
}

func (c *Controller) WarehouseOut(r *ghttp.Request) {
	var req supplyV1.WarehouseOutReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().WarehouseOp(r.Context(), req.Data, "out")
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "出库成功"})
}

func (c *Controller) GetStockList(r *ghttp.Request) {
	var req supplyV1.GetStockListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Supply().GetStockList(r.Context(), req.Page, req.PageSize, req.ProductId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateLogisticsRecord(r *ghttp.Request) {
	var req supplyV1.CreateLogisticsRecordReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Supply().CreateLogisticsRecord(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "物流记录创建成功"})
}

func (c *Controller) GetLogisticsList(r *ghttp.Request) {
	var req supplyV1.GetLogisticsListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, err := service.Supply().GetLogisticsList(r.Context(), req.OrderId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}
