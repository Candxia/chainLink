package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 供应商管理 ====================

type CreateSupplierReq struct {
	g.Meta `path:"/supply/supplier" method:"post" tags:"供应链管理" summary:"创建供应商"`
	Data   map[string]interface{} `json:"data"`
}
type CreateSupplierRes struct {
	g.Meta `mime:"application/json"`
	Id     uint `json:"id"`
}

type UpdateSupplierReq struct {
	g.Meta `path:"/supply/supplier" method:"put" tags:"供应链管理" summary:"更新供应商"`
	Data   map[string]interface{} `json:"data"`
}
type UpdateSupplierRes struct {
	g.Meta `mime:"application/json"`
}

type GetSupplierReq struct {
	g.Meta `path:"/supply/supplier/:id" method:"get" tags:"供应链管理" summary:"供应商详情"`
	Id     uint `json:"id"`
}
type GetSupplierRes struct {
	g.Meta   `mime:"application/json"`
	Supplier interface{} `json:"supplier"`
}

type GetSupplierListReq struct {
	g.Meta   `path:"/supply/supplier/list" method:"get" tags:"供应链管理" summary:"供应商列表"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Name     string `json:"name"`
}
type GetSupplierListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type DeleteSupplierReq struct {
	g.Meta `path:"/supply/supplier/:id" method:"delete" tags:"供应链管理" summary:"删除供应商"`
	Id     uint `json:"id"`
}
type DeleteSupplierRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 订单管理 ====================

type CreateOrderReq struct {
	g.Meta `path:"/supply/order" method:"post" tags:"供应链管理" summary:"创建订单"`
	Data   map[string]interface{} `json:"data"`
}
type CreateOrderRes struct {
	g.Meta `mime:"application/json"`
	Id     uint `json:"id"`
}

type UpdateOrderReq struct {
	g.Meta `path:"/supply/order" method:"put" tags:"供应链管理" summary:"更新订单"`
	Data   map[string]interface{} `json:"data"`
}
type UpdateOrderRes struct {
	g.Meta `mime:"application/json"`
}

type GetOrderReq struct {
	g.Meta `path:"/supply/order/:id" method:"get" tags:"供应链管理" summary:"订单详情"`
	Id     uint `json:"id"`
}
type GetOrderRes struct {
	g.Meta `mime:"application/json"`
	Order  interface{} `json:"order"`
}

type GetOrderListReq struct {
	g.Meta   `path:"/supply/order/list" method:"get" tags:"供应链管理" summary:"订单列表"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Status   string `json:"status"`
}
type GetOrderListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type DeleteOrderReq struct {
	g.Meta `path:"/supply/order/:id" method:"delete" tags:"供应链管理" summary:"删除订单"`
	Id     uint `json:"id"`
}
type DeleteOrderRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 仓库出入库 ====================

type WarehouseInReq struct {
	g.Meta `path:"/supply/warehouse/in" method:"post" tags:"供应链管理" summary:"入库"`
	Data   map[string]interface{} `json:"data"`
}
type WarehouseInRes struct {
	g.Meta `mime:"application/json"`
}

type WarehouseOutReq struct {
	g.Meta `path:"/supply/warehouse/out" method:"post" tags:"供应链管理" summary:"出库"`
	Data   map[string]interface{} `json:"data"`
}
type WarehouseOutRes struct {
	g.Meta `mime:"application/json"`
}

type GetStockListReq struct {
	g.Meta    `path:"/supply/warehouse/stock" method:"get" tags:"供应链管理" summary:"库存列表"`
	Page      int  `json:"page" d:"1"`
	PageSize  int  `json:"pageSize" d:"10"`
	ProductId uint `json:"productId"`
}
type GetStockListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

// ==================== 物流管理 ====================

type CreateLogisticsRecordReq struct {
	g.Meta `path:"/supply/logistics" method:"post" tags:"供应链管理" summary:"创建物流记录"`
	Data   map[string]interface{} `json:"data"`
}
type CreateLogisticsRecordRes struct {
	g.Meta `mime:"application/json"`
}

type GetLogisticsListReq struct {
	g.Meta  `path:"/supply/logistics/:orderId" method:"get" tags:"供应链管理" summary:"物流记录列表"`
	OrderId uint `json:"orderId"`
}
type GetLogisticsListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
}
