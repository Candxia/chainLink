package warehouse

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	warehouseV1 "cl_system/api/warehouse/v1"
	"cl_system/internal/service"
)

type Controller struct{}

var Ctl = &Controller{}

// ==================== 仓库信息 ====================

func (c *Controller) GetWarehouseList(r *ghttp.Request) {
	var req warehouseV1.GetWarehouseListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetWarehouseList(r.Context(), req.Page, req.PageSize, req.Name)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) CreateWarehouse(r *ghttp.Request) {
	var req warehouseV1.CreateWarehouseReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateWarehouse(r.Context(), *req.WarehouseCreateInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) UpdateWarehouse(r *ghttp.Request) {
	var req warehouseV1.UpdateWarehouseReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().UpdateWarehouse(r.Context(), *req.WarehouseUpdateInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteWarehouse(r *ghttp.Request) {
	var req warehouseV1.DeleteWarehouseReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().DeleteWarehouse(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) UpdateWarehouseStatus(r *ghttp.Request) {
	var req warehouseV1.UpdateWarehouseStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().UpdateWarehouseStatus(r.Context(), req.Id, req.Status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "状态更新成功"})
}

// ==================== 区域 ====================

func (c *Controller) GetAreaList(r *ghttp.Request) {
	var req warehouseV1.GetAreaListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, err := service.Warehouse().GetAreaList(r.Context(), req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) CreateArea(r *ghttp.Request) {
	var req warehouseV1.CreateAreaReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateArea(r.Context(), *req.WarehouseAreaCreateInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新增区域成功"})
}

func (c *Controller) DeleteArea(r *ghttp.Request) {
	var req warehouseV1.DeleteAreaReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().DeleteArea(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除区域成功"})
}

// ==================== 货架 ====================

func (c *Controller) GetShelfList(r *ghttp.Request) {
	var req warehouseV1.GetShelfListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, err := service.Warehouse().GetShelfList(r.Context(), req.AreaId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) CreateShelf(r *ghttp.Request) {
	var req warehouseV1.CreateShelfReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateShelf(r.Context(), *req.WarehouseShelfCreateInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新增货架成功"})
}

func (c *Controller) DeleteShelf(r *ghttp.Request) {
	var req warehouseV1.DeleteShelfReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().DeleteShelf(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除货架成功"})
}

// ==================== 盘存 ====================

func (c *Controller) GetStocktakeList(r *ghttp.Request) {
	var req warehouseV1.GetStocktakeListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetStocktakeList(r.Context(), req.Page, req.PageSize, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) CreateStocktake(r *ghttp.Request) {
	var req warehouseV1.CreateStocktakeReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateStocktake(r.Context(), *req.StocktakeCreateInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新建盘存成功"})
}

// ==================== 整车库存 ====================

func (c *Controller) GetInventoryCarList(r *ghttp.Request) {
	var req warehouseV1.GetInventoryCarListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetInventoryCarList(r.Context(), req.Page, req.PageSize, req.WarehouseId, req.Brand, req.ModelName, req.Vin)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

// ==================== 原材料库存 ====================

func (c *Controller) GetRawMaterialList(r *ghttp.Request) {
	var req warehouseV1.GetRawMaterialListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetRawMaterialList(r.Context(), req.Page, req.PageSize, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) RawMaterialIn(r *ghttp.Request) {
	var req warehouseV1.RawMaterialInReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().RawMaterialIn(r.Context(), *req.RawMaterialInInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "原材料入库成功"})
}

func (c *Controller) RawMaterialOut(r *ghttp.Request) {
	var req warehouseV1.RawMaterialOutReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().RawMaterialOut(r.Context(), *req.RawMaterialOutInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "原材料出库成功"})
}

// ==================== 危固废库存 ====================

func (c *Controller) GetWasteList(r *ghttp.Request) {
	var req warehouseV1.GetWasteListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetWasteList(r.Context(), req.Page, req.PageSize, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) WasteIn(r *ghttp.Request) {
	var req warehouseV1.WasteInReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().WasteIn(r.Context(), *req.WasteInInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "危废入库成功"})
}

// ==================== 溯源件 ====================

func (c *Controller) GetPartTraceableList(r *ghttp.Request) {
	var req warehouseV1.GetPartTraceableListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetPartTraceableList(r.Context(), req.Page, req.PageSize, req.WarehouseId, req.PartName)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) PartTraceableOut(r *ghttp.Request) {
	var req warehouseV1.PartTraceableOutReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartTraceableOut(r.Context(), *req.PartTraceableOutInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "溯源件出库成功"})
}

// ==================== 非溯源件 ====================

func (c *Controller) GetPartUntraceableList(r *ghttp.Request) {
	var req warehouseV1.GetPartUntraceableListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetPartUntraceableList(r.Context(), req.Page, req.PageSize, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) PartUntraceableIn(r *ghttp.Request) {
	var req warehouseV1.PartUntraceableInReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartUntraceableIn(r.Context(), *req.PartUntraceableInInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "非溯源件入库成功"})
}

func (c *Controller) PartUntraceableOut(r *ghttp.Request) {
	var req warehouseV1.PartUntraceableOutReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartUntraceableOut(r.Context(), *req.PartUntraceableOutInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "非溯源件出库成功"})
}

// ==================== 操作记录 ====================

func (c *Controller) GetInboundRecordList(r *ghttp.Request) {
	var req warehouseV1.GetInboundRecordListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetInboundRecordList(r.Context(), req.Page, req.PageSize, req.InboundType, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) GetInboundRecordDetail(r *ghttp.Request) {
	var req warehouseV1.GetInboundRecordDetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	detail, err := service.Warehouse().GetInboundRecordDetail(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}

func (c *Controller) GetOutboundRecordList(r *ghttp.Request) {
	var req warehouseV1.GetOutboundRecordListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetOutboundRecordList(r.Context(), req.Page, req.PageSize, req.OutboundType, req.WarehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) GetOutboundRecordDetail(r *ghttp.Request) {
	var req warehouseV1.GetOutboundRecordDetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	detail, err := service.Warehouse().GetOutboundRecordDetail(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}

func (c *Controller) GetTransferRecordList(r *ghttp.Request) {
	var req warehouseV1.GetTransferRecordListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Warehouse().GetTransferRecordList(r.Context(), req.Page, req.PageSize, req.TransferType)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) GetTransferRecordDetail(r *ghttp.Request) {
	var req warehouseV1.GetTransferRecordDetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	detail, err := service.Warehouse().GetTransferRecordDetail(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}
