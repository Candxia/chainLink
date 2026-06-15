package trace

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) CreateProduct(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Trace().CreateProduct(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) UpdateProduct(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Trace().UpdateProduct(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteProduct(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.Trace().DeleteProduct(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetProduct(r *ghttp.Request) {
	id := r.Get("id").Uint()
	product, err := service.Trace().GetProduct(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": product})
}

func (c *Controller) GetProductList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	name := r.GetQuery("name").String()
	category := r.GetQuery("category").String()
	list, total, err := service.Trace().GetProductList(r.Context(), page, pageSize, name, category)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) CreateBatch(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	id, err := service.Trace().CreateBatch(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功", "data": g.Map{"id": id}})
}

func (c *Controller) GetBatch(r *ghttp.Request) {
	id := r.Get("id").Uint()
	batch, err := service.Trace().GetBatch(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": batch})
}

func (c *Controller) GetBatchList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	productId := r.GetQuery("productId").Uint()
	list, total, err := service.Trace().GetBatchList(r.Context(), page, pageSize, productId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateTraceRecord(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	record, err := service.Trace().CreateTraceRecord(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "记录创建成功", "data": record})
}

func (c *Controller) GetTraceRecord(r *ghttp.Request) {
	id := r.Get("id").Uint()
	record, err := service.Trace().GetTraceRecord(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": record})
}

func (c *Controller) GetTraceRecordList(r *ghttp.Request) {
	productId := r.Get("productId").Uint()
	records, err := service.Trace().GetTraceRecordList(r.Context(), productId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": records})
}

func (c *Controller) GetTraceChain(r *ghttp.Request) {
	productId := r.Get("productId").Uint()
	chain, err := service.Trace().GetTraceChain(r.Context(), productId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": chain})
}

func (c *Controller) GenerateQRCode(r *ghttp.Request) {
	productId := r.Get("productId").Uint()
	qrcode, err := service.Trace().GenerateQRCode(r.Context(), productId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"qrcode": qrcode}})
}

func (c *Controller) PublicQuery(r *ghttp.Request) {
	query := r.GetQuery("q").String()
	result, err := service.Trace().PublicQuery(r.Context(), query)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": result})
}
