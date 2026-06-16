package blockchain

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	blockchainV1 "cl_system/api/blockchain/v1"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) GetBlockList(r *ghttp.Request) {
	var req blockchainV1.GetBlockListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	blocks, total, err := service.Blockchain().GetBlockList(r.Context(), req.Page, req.PageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": blocks, "total": total}})
}

func (c *Controller) GetBlockByHash(r *ghttp.Request) {
	var req blockchainV1.GetBlockByHashReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	block, err := service.Blockchain().GetBlockByHash(r.Context(), req.Hash)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": block})
}

func (c *Controller) GetTransactionList(r *ghttp.Request) {
	var req blockchainV1.GetTransactionListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	txs, total, err := service.Blockchain().GetTransactionList(r.Context(), req.Page, req.PageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": txs, "total": total}})
}

func (c *Controller) GetTransactionDetail(r *ghttp.Request) {
	var req blockchainV1.GetTransactionDetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	tx, err := service.Blockchain().GetTransactionDetail(r.Context(), req.TxId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": tx})
}

func (c *Controller) DeployContract(r *ghttp.Request) {
	var req blockchainV1.DeployContractReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	contract, err := service.Blockchain().DeployContract(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "部署成功", "data": contract})
}

func (c *Controller) GetContractInfo(r *ghttp.Request) {
	var req blockchainV1.GetContractInfoReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	contract, err := service.Blockchain().GetContractInfo(r.Context(), req.Address)
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
	var req blockchainV1.UploadToChainReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.Blockchain().UploadToChain(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "上链成功", "data": result})
}

func (c *Controller) VerifyData(r *ghttp.Request) {
	var req blockchainV1.VerifyDataReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.Blockchain().VerifyData(r.Context(), req.Hash, req.TxId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": result})
}
