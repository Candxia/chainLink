package trace

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	traceV1 "cl_system/api/trace/v1"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) CreateProduct(r *ghttp.Request) {
	var req traceV1.CreateProductReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Trace().CreateProduct(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateProduct(r *ghttp.Request) {
	var req traceV1.UpdateProductReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Trace().UpdateProduct(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteProduct(r *ghttp.Request) {
	var req traceV1.DeleteProductReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Trace().DeleteProduct(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetProduct(r *ghttp.Request) {
	var req traceV1.GetProductReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	product, err := service.Trace().GetProduct(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": product})
}

func (c *Controller) GetProductList(r *ghttp.Request) {
	var req traceV1.GetProductListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Trace().GetProductList(r.Context(), req.Page, req.PageSize, req.Name, req.Category)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) CreateBatch(r *ghttp.Request) {
	var req traceV1.CreateBatchReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Trace().CreateBatch(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) GetBatch(r *ghttp.Request) {
	var req traceV1.GetBatchReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	batch, err := service.Trace().GetBatch(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": batch})
}

func (c *Controller) GetBatchList(r *ghttp.Request) {
	var req traceV1.GetBatchListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Trace().GetBatchList(r.Context(), req.Page, req.PageSize, req.ProductId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateTraceRecord(r *ghttp.Request) {
	var req traceV1.CreateTraceRecordReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	record, err := service.Trace().CreateTraceRecord(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "记录创建成功", "data": record})
}

func (c *Controller) GetTraceRecord(r *ghttp.Request) {
	var req traceV1.GetTraceRecordReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	record, err := service.Trace().GetTraceRecord(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": record})
}

func (c *Controller) GetTraceRecordList(r *ghttp.Request) {
	var req traceV1.GetTraceRecordListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	records, err := service.Trace().GetTraceRecordList(r.Context(), req.ProductId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": records})
}

func (c *Controller) GetTraceChain(r *ghttp.Request) {
	var req traceV1.GetTraceChainReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	chain, err := service.Trace().GetTraceChain(r.Context(), req.ProductId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": chain})
}

func (c *Controller) GenerateQRCode(r *ghttp.Request) {
	var req traceV1.GenerateQRCodeReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	qrcode, err := service.Trace().GenerateQRCode(r.Context(), req.ProductId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"qrcode": qrcode}})
}

func (c *Controller) PublicQuery(r *ghttp.Request) {
	var req traceV1.PublicQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.Trace().PublicQuery(r.Context(), req.Q)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": result})
}
