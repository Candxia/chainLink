package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/model"
)

// ==================== 仓库信息 ====================

type GetWarehouseListReq struct {
	g.Meta   `path:"/warehouse/list" method:"get" tags:"仓库管理" summary:"仓库列表"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Name     string `json:"name"`
}
type GetWarehouseListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.WarehouseInfo `json:"list"`
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
}

type CreateWarehouseReq struct {
	g.Meta `path:"/warehouse" method:"post" tags:"仓库管理" summary:"创建仓库"`
	*model.WarehouseCreateInput
}
type CreateWarehouseRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateWarehouseReq struct {
	g.Meta `path:"/warehouse" method:"put" tags:"仓库管理" summary:"更新仓库"`
	*model.WarehouseUpdateInput
}
type UpdateWarehouseRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteWarehouseReq struct {
	g.Meta `path:"/warehouse/:id" method:"delete" tags:"仓库管理" summary:"删除仓库"`
	Id     uint `json:"id"`
}
type DeleteWarehouseRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateWarehouseStatusReq struct {
	g.Meta `path:"/warehouse/:id/status" method:"put" tags:"仓库管理" summary:"更新仓库状态"`
	Id     uint `json:"id"`
	Status int  `json:"status"`
}
type UpdateWarehouseStatusRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 区域 ====================

type GetAreaListReq struct {
	g.Meta      `path:"/warehouse/area/list" method:"get" tags:"仓库管理" summary:"区域列表"`
	WarehouseId uint `json:"warehouseId"`
}
type GetAreaListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.WarehouseAreaInfo `json:"list"`
}

type CreateAreaReq struct {
	g.Meta `path:"/warehouse/area" method:"post" tags:"仓库管理" summary:"创建区域"`
	*model.WarehouseAreaCreateInput
}
type CreateAreaRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteAreaReq struct {
	g.Meta `path:"/warehouse/area/:id" method:"delete" tags:"仓库管理" summary:"删除区域"`
	Id     uint `json:"id"`
}
type DeleteAreaRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 货架 ====================

type GetShelfListReq struct {
	g.Meta `path:"/warehouse/shelf/list" method:"get" tags:"仓库管理" summary:"货架列表"`
	AreaId uint `json:"areaId"`
}
type GetShelfListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.WarehouseShelfInfo `json:"list"`
}

type CreateShelfReq struct {
	g.Meta `path:"/warehouse/shelf" method:"post" tags:"仓库管理" summary:"创建货架"`
	*model.WarehouseShelfCreateInput
}
type CreateShelfRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteShelfReq struct {
	g.Meta `path:"/warehouse/shelf/:id" method:"delete" tags:"仓库管理" summary:"删除货架"`
	Id     uint `json:"id"`
}
type DeleteShelfRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 盘存 ====================

type GetStocktakeListReq struct {
	g.Meta      `path:"/warehouse/stocktake/list" method:"get" tags:"仓库管理" summary:"盘点列表"`
	Page        int  `json:"page" d:"1"`
	PageSize    int  `json:"pageSize" d:"10"`
	WarehouseId uint `json:"warehouseId"`
}
type GetStocktakeListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.StocktakeInfo `json:"list"`
	Total    int                    `json:"total"`
}

type CreateStocktakeReq struct {
	g.Meta `path:"/warehouse/stocktake" method:"post" tags:"仓库管理" summary:"创建盘点"`
	*model.StocktakeCreateInput
}
type CreateStocktakeRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 整车库存 ====================

type GetInventoryCarListReq struct {
	g.Meta      `path:"/inventory/car/list" method:"get" tags:"库存管理" summary:"整车库存列表"`
	Page        int    `json:"page" d:"1"`
	PageSize    int    `json:"pageSize" d:"10"`
	WarehouseId uint   `json:"warehouseId"`
	Brand       string `json:"brand"`
	ModelName   string `json:"modelName"`
	Vin         string `json:"vin"`
}
type GetInventoryCarListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.InventoryCarInfo `json:"list"`
	Total    int                       `json:"total"`
}

// ==================== 原材料 ====================

type GetRawMaterialListReq struct {
	g.Meta      `path:"/inventory/raw-material/list" method:"get" tags:"库存管理" summary:"原材料列表"`
	Page        int  `json:"page" d:"1"`
	PageSize    int  `json:"pageSize" d:"10"`
	WarehouseId uint `json:"warehouseId"`
}
type GetRawMaterialListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.InventoryRawMaterialInfo `json:"list"`
	Total    int                               `json:"total"`
}

type RawMaterialInReq struct {
	g.Meta `path:"/inventory/raw-material/in" method:"post" tags:"库存管理" summary:"原材料入库"`
	*model.RawMaterialInInput
}
type RawMaterialInRes struct {
	g.Meta `mime:"application/json"`
}

type RawMaterialOutReq struct {
	g.Meta `path:"/inventory/raw-material/out" method:"post" tags:"库存管理" summary:"原材料出库"`
	*model.RawMaterialOutInput
}
type RawMaterialOutRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 危固废 ====================

type GetWasteListReq struct {
	g.Meta      `path:"/inventory/waste/list" method:"get" tags:"库存管理" summary:"危固废列表"`
	Page        int  `json:"page" d:"1"`
	PageSize    int  `json:"pageSize" d:"10"`
	WarehouseId uint `json:"warehouseId"`
}
type GetWasteListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.InventoryWasteInfo `json:"list"`
	Total    int                         `json:"total"`
}

type WasteInReq struct {
	g.Meta `path:"/inventory/waste/in" method:"post" tags:"库存管理" summary:"危固废入库"`
	*model.WasteInInput
}
type WasteInRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 溯源件 ====================

type GetPartTraceableListReq struct {
	g.Meta      `path:"/inventory/part/traceable/list" method:"get" tags:"库存管理" summary:"溯源件列表"`
	Page        int    `json:"page" d:"1"`
	PageSize    int    `json:"pageSize" d:"10"`
	WarehouseId uint   `json:"warehouseId"`
	PartName    string `json:"partName"`
}
type GetPartTraceableListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.InventoryPartTraceableInfo `json:"list"`
	Total    int                                 `json:"total"`
}

type PartTraceableOutReq struct {
	g.Meta `path:"/inventory/part/traceable/out" method:"post" tags:"库存管理" summary:"溯源件出库"`
	*model.PartTraceableOutInput
}
type PartTraceableOutRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 非溯源件 ====================

type GetPartUntraceableListReq struct {
	g.Meta      `path:"/inventory/part/untraceable/list" method:"get" tags:"库存管理" summary:"非溯源件列表"`
	Page        int    `json:"page" d:"1"`
	PageSize    int    `json:"pageSize" d:"10"`
	WarehouseId uint   `json:"warehouseId"`
}
type GetPartUntraceableListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.InventoryPartUntraceableInfo `json:"list"`
	Total    int                                   `json:"total"`
}

type PartUntraceableInReq struct {
	g.Meta `path:"/inventory/part/untraceable/in" method:"post" tags:"库存管理" summary:"非溯源件入库"`
	*model.PartUntraceableInInput
}
type PartUntraceableInRes struct {
	g.Meta `mime:"application/json"`
}

type PartUntraceableOutReq struct {
	g.Meta `path:"/inventory/part/untraceable/out" method:"post" tags:"库存管理" summary:"非溯源件出库"`
	*model.PartUntraceableOutInput
}
type PartUntraceableOutRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 操作记录 ====================

type GetInboundRecordListReq struct {
	g.Meta      `path:"/records/inbound" method:"get" tags:"操作记录" summary:"入库记录列表"`
	Page        int    `json:"page" d:"1"`
	PageSize    int    `json:"pageSize" d:"10"`
	InboundType string `json:"inboundType"`
	WarehouseId uint   `json:"warehouseId"`
}
type GetInboundRecordListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type GetInboundRecordDetailReq struct {
	g.Meta `path:"/records/inbound/:id" method:"get" tags:"操作记录" summary:"入库记录详情"`
	Id     uint `json:"id"`
}
type GetInboundRecordDetailRes struct {
	g.Meta  `mime:"application/json"`
	Record  interface{} `json:"record"`
}

type GetOutboundRecordListReq struct {
	g.Meta       `path:"/records/outbound" method:"get" tags:"操作记录" summary:"出库记录列表"`
	Page         int    `json:"page" d:"1"`
	PageSize     int    `json:"pageSize" d:"10"`
	OutboundType string `json:"outboundType"`
	WarehouseId  uint   `json:"warehouseId"`
}
type GetOutboundRecordListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type GetOutboundRecordDetailReq struct {
	g.Meta `path:"/records/outbound/:id" method:"get" tags:"操作记录" summary:"出库记录详情"`
	Id     uint `json:"id"`
}
type GetOutboundRecordDetailRes struct {
	g.Meta  `mime:"application/json"`
	Record  interface{} `json:"record"`
}

type GetTransferRecordListReq struct {
	g.Meta       `path:"/records/transfer" method:"get" tags:"操作记录" summary:"调拨记录列表"`
	Page         int    `json:"page" d:"1"`
	PageSize     int    `json:"pageSize" d:"10"`
	TransferType string `json:"transferType"`
}
type GetTransferRecordListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type GetTransferRecordDetailReq struct {
	g.Meta `path:"/records/transfer/:id" method:"get" tags:"操作记录" summary:"调拨记录详情"`
	Id     uint `json:"id"`
}
type GetTransferRecordDetailRes struct {
	g.Meta  `mime:"application/json"`
	Record  interface{} `json:"record"`
}

// ==================== 文件上传 ====================

type UploadFileReq struct {
	g.Meta `path:"/upload" method:"post" tags:"系统管理" summary:"上传文件"`
}
type UploadFileRes struct {
	g.Meta  `mime:"application/json"`
	Url     string `json:"url"`
}
