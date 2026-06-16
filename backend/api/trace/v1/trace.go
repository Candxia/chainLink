package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 产品管理 ====================

type CreateProductReq struct {
	g.Meta `path:"/trace/product" method:"post" tags:"产品溯源" summary:"创建产品"`
	Data   map[string]interface{} `json:"data"`
}
type CreateProductRes struct {
	g.Meta `mime:"application/json"`
	Id     uint `json:"id"`
}

type UpdateProductReq struct {
	g.Meta `path:"/trace/product" method:"put" tags:"产品溯源" summary:"更新产品"`
	Data   map[string]interface{} `json:"data"`
}
type UpdateProductRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteProductReq struct {
	g.Meta `path:"/trace/product/:id" method:"delete" tags:"产品溯源" summary:"删除产品"`
	Id     uint `json:"id"`
}
type DeleteProductRes struct {
	g.Meta `mime:"application/json"`
}

type GetProductReq struct {
	g.Meta `path:"/trace/product/:id" method:"get" tags:"产品溯源" summary:"产品详情"`
	Id     uint `json:"id"`
}
type GetProductRes struct {
	g.Meta   `mime:"application/json"`
	Product  interface{} `json:"product"`
}

type GetProductListReq struct {
	g.Meta   `path:"/trace/product/list" method:"get" tags:"产品溯源" summary:"产品列表"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
type GetProductListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// ==================== 批次管理 ====================

type CreateBatchReq struct {
	g.Meta `path:"/trace/batch" method:"post" tags:"批次管理" summary:"创建批次"`
	Data   map[string]interface{} `json:"data"`
}
type CreateBatchRes struct {
	g.Meta `mime:"application/json"`
	Id     uint `json:"id"`
}

type GetBatchReq struct {
	g.Meta `path:"/trace/batch/:id" method:"get" tags:"批次管理" summary:"批次详情"`
	Id     uint `json:"id"`
}
type GetBatchRes struct {
	g.Meta `mime:"application/json"`
	Batch  interface{} `json:"batch"`
}

type GetBatchListReq struct {
	g.Meta    `path:"/trace/batch/list" method:"get" tags:"批次管理" summary:"批次列表"`
	Page      int  `json:"page" d:"1"`
	PageSize  int  `json:"pageSize" d:"10"`
	ProductId uint `json:"productId"`
}
type GetBatchListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

// ==================== 溯源记录 ====================

type CreateTraceRecordReq struct {
	g.Meta `path:"/trace/record" method:"post" tags:"溯源记录" summary:"创建溯源记录"`
	Data   map[string]interface{} `json:"data"`
}
type CreateTraceRecordRes struct {
	g.Meta  `mime:"application/json"`
	Record  interface{} `json:"record"`
}

type GetTraceRecordReq struct {
	g.Meta `path:"/trace/record/:id" method:"get" tags:"溯源记录" summary:"溯源记录详情"`
	Id     uint `json:"id"`
}
type GetTraceRecordRes struct {
	g.Meta  `mime:"application/json"`
	Record  interface{} `json:"record"`
}

type GetTraceRecordListReq struct {
	g.Meta    `path:"/trace/record/list/:productId" method:"get" tags:"溯源记录" summary:"溯源记录列表"`
	ProductId uint `json:"productId"`
}
type GetTraceRecordListRes struct {
	g.Meta  `mime:"application/json"`
	Records interface{} `json:"records"`
}

// ==================== 溯源链 ====================

type GetTraceChainReq struct {
	g.Meta    `path:"/trace/chain/:productId" method:"get" tags:"溯源链" summary:"溯源链查询"`
	ProductId uint `json:"productId"`
}
type GetTraceChainRes struct {
	g.Meta `mime:"application/json"`
	Chain  interface{} `json:"chain"`
}

type GenerateQRCodeReq struct {
	g.Meta    `path:"/trace/qrcode/:productId" method:"get" tags:"溯源链" summary:"生成二维码"`
	ProductId uint `json:"productId"`
}
type GenerateQRCodeRes struct {
	g.Meta  `mime:"application/json"`
	QRCode  string `json:"qrcode"`
}

// ==================== 公开查询 ====================

type PublicQueryReq struct {
	g.Meta `path:"/trace/public/query" method:"get" tags:"公开查询" summary:"公开溯源查询"`
	Q      string `json:"q"`
}
type PublicQueryRes struct {
	g.Meta  `mime:"application/json"`
	Result  interface{} `json:"result"`
}
