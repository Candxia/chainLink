package warehouse

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type Controller struct{}

var Ctl = &Controller{}

// ==================== 仓库信息 ====================

func (c *Controller) GetWarehouseList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	name := r.GetQuery("name").String()
	list, total, err := service.Warehouse().GetWarehouseList(r.Context(), page, pageSize, name)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) CreateWarehouse(r *ghttp.Request) {
	var req model.WarehouseCreateInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateWarehouse(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) UpdateWarehouse(r *ghttp.Request) {
	var req model.WarehouseUpdateInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().UpdateWarehouse(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteWarehouse(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Warehouse().DeleteWarehouse(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) UpdateWarehouseStatus(r *ghttp.Request) {
	id := r.Get("id").Uint()
	status := r.Get("status").Int()
	err := service.Warehouse().UpdateWarehouseStatus(r.Context(), id, status)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "状态更新成功"})
}

// ==================== 区域 ====================

func (c *Controller) GetAreaList(r *ghttp.Request) {
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, err := service.Warehouse().GetAreaList(r.Context(), warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) CreateArea(r *ghttp.Request) {
	var req model.WarehouseAreaCreateInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateArea(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新增区域成功"})
}

func (c *Controller) DeleteArea(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Warehouse().DeleteArea(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除区域成功"})
}

// ==================== 货架 ====================

func (c *Controller) GetShelfList(r *ghttp.Request) {
	areaId := r.GetQuery("areaId", 0).Uint()
	list, err := service.Warehouse().GetShelfList(r.Context(), areaId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) CreateShelf(r *ghttp.Request) {
	var req model.WarehouseShelfCreateInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateShelf(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新增货架成功"})
}

func (c *Controller) DeleteShelf(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Warehouse().DeleteShelf(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除货架成功"})
}

// ==================== 盘存 ====================

func (c *Controller) GetStocktakeList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetStocktakeList(r.Context(), page, pageSize, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) CreateStocktake(r *ghttp.Request) {
	var req model.StocktakeCreateInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().CreateStocktake(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "新建盘存成功"})
}

// ==================== 整车库存 ====================

func (c *Controller) GetInventoryCarList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	brand := r.GetQuery("brand").String()
	modelName := r.GetQuery("model").String()
	vin := r.GetQuery("vin").String()
	list, total, err := service.Warehouse().GetInventoryCarList(r.Context(), page, pageSize, warehouseId, brand, modelName, vin)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

// ==================== 原材料库存 ====================

func (c *Controller) GetRawMaterialList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetRawMaterialList(r.Context(), page, pageSize, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) RawMaterialIn(r *ghttp.Request) {
	var req model.RawMaterialInInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().RawMaterialIn(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "原材料入库成功"})
}

func (c *Controller) RawMaterialOut(r *ghttp.Request) {
	var req model.RawMaterialOutInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().RawMaterialOut(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "原材料出库成功"})
}

// ==================== 危固废库存 ====================

func (c *Controller) GetWasteList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetWasteList(r.Context(), page, pageSize, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) WasteIn(r *ghttp.Request) {
	var req model.WasteInInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().WasteIn(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "危废入库成功"})
}

// ==================== 溯源件 ====================

func (c *Controller) GetPartTraceableList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	partName := r.GetQuery("partName").String()
	list, total, err := service.Warehouse().GetPartTraceableList(r.Context(), page, pageSize, warehouseId, partName)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) PartTraceableOut(r *ghttp.Request) {
	var req model.PartTraceableOutInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartTraceableOut(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "溯源件出库成功"})
}

// ==================== 非溯源件 ====================

func (c *Controller) GetPartUntraceableList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetPartUntraceableList(r.Context(), page, pageSize, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) PartUntraceableIn(r *ghttp.Request) {
	var req model.PartUntraceableInInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartUntraceableIn(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "非溯源件入库成功"})
}

func (c *Controller) PartUntraceableOut(r *ghttp.Request) {
	var req model.PartUntraceableOutInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Warehouse().PartUntraceableOut(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "非溯源件出库成功"})
}

// ==================== 操作记录 ====================

func (c *Controller) GetInboundRecordList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	inboundType := r.GetQuery("inboundType").String()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetInboundRecordList(r.Context(), page, pageSize, inboundType, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) GetInboundRecordDetail(r *ghttp.Request) {
	id := r.Get("id").Uint()
	detail, err := service.Warehouse().GetInboundRecordDetail(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}

func (c *Controller) GetOutboundRecordList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	outboundType := r.GetQuery("outboundType").String()
	warehouseId := r.GetQuery("warehouseId", 0).Uint()
	list, total, err := service.Warehouse().GetOutboundRecordList(r.Context(), page, pageSize, outboundType, warehouseId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) GetOutboundRecordDetail(r *ghttp.Request) {
	id := r.Get("id").Uint()
	detail, err := service.Warehouse().GetOutboundRecordDetail(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}

func (c *Controller) GetTransferRecordList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	transferType := r.GetQuery("transferType").String()
	list, total, err := service.Warehouse().GetTransferRecordList(r.Context(), page, pageSize, transferType)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) GetTransferRecordDetail(r *ghttp.Request) {
	id := r.Get("id").Uint()
	detail, err := service.Warehouse().GetTransferRecordDetail(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": detail})
}
