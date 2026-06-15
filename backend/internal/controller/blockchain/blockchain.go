package blockchain

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) GetBlockList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	blocks, total, err := service.Blockchain().GetBlockList(r.Context(), page, pageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": blocks, "total": total}})
}

func (c *Controller) GetBlockByHash(r *ghttp.Request) {
	hash := r.Get("hash").String()
	block, err := service.Blockchain().GetBlockByHash(r.Context(), hash)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": block})
}

func (c *Controller) GetTransactionList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	txs, total, err := service.Blockchain().GetTransactionList(r.Context(), page, pageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": txs, "total": total}})
}

func (c *Controller) GetTransactionDetail(r *ghttp.Request) {
	txId := r.Get("txId").String()
	tx, err := service.Blockchain().GetTransactionDetail(r.Context(), txId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": tx})
}

func (c *Controller) DeployContract(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	contract, err := service.Blockchain().DeployContract(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "部署成功", "data": contract})
}

func (c *Controller) GetContractInfo(r *ghttp.Request) {
	address := r.Get("address").String()
	contract, err := service.Blockchain().GetContractInfo(r.Context(), address)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": contract})
}

func (c *Controller) GetContractList(r *ghttp.Request) {
	list, err := service.Blockchain().GetContractList(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) UploadToChain(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.Blockchain().UploadToChain(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "上链成功", "data": result})
}

func (c *Controller) VerifyData(r *ghttp.Request) {
	hash := r.GetQuery("hash").String()
	txId := r.GetQuery("txId").String()
	result, err := service.Blockchain().VerifyData(r.Context(), hash, txId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": result})
}
