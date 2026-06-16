package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 区块管理 ====================

type GetBlockListReq struct {
	g.Meta   `path:"/blockchain/blocks" method:"get" tags:"区块链管理" summary:"区块列表"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}
type GetBlockListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
	Total  int         `json:"total"`
}

type GetBlockByHashReq struct {
	g.Meta `path:"/blockchain/block/:hash" method:"get" tags:"区块链管理" summary:"按哈希查询区块"`
	Hash   string `json:"hash"`
}
type GetBlockByHashRes struct {
	g.Meta `mime:"application/json"`
	Block  interface{} `json:"block"`
}

// ==================== 交易管理 ====================

type GetTransactionListReq struct {
	g.Meta   `path:"/blockchain/transactions" method:"get" tags:"区块链管理" summary:"交易列表"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}
type GetTransactionListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
	Total  int         `json:"total"`
}

type GetTransactionDetailReq struct {
	g.Meta `path:"/blockchain/transaction/:txId" method:"get" tags:"区块链管理" summary:"交易详情"`
	TxId   string `json:"txId"`
}
type GetTransactionDetailRes struct {
	g.Meta `mime:"application/json"`
	Tx     interface{} `json:"tx"`
}

// ==================== 合约管理 ====================

type DeployContractReq struct {
	g.Meta `path:"/blockchain/contract" method:"post" tags:"区块链管理" summary:"部署合约"`
	Data   map[string]interface{} `json:"data"`
}
type DeployContractRes struct {
	g.Meta  `mime:"application/json"`
	Contract interface{} `json:"contract"`
}

type GetContractInfoReq struct {
	g.Meta  `path:"/blockchain/contract/:address" method:"get" tags:"区块链管理" summary:"合约详情"`
	Address string `json:"address"`
}
type GetContractInfoRes struct {
	g.Meta   `mime:"application/json"`
	Contract interface{} `json:"contract"`
}

type GetContractListReq struct {
	g.Meta `path:"/blockchain/contract/list" method:"get" tags:"区块链管理" summary:"合约列表"`
}
type GetContractListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
}

// ==================== 数据上链 ====================

type UploadToChainReq struct {
	g.Meta `path:"/blockchain/data/upload" method:"post" tags:"区块链管理" summary:"数据上链"`
	Data   map[string]interface{} `json:"data"`
}
type UploadToChainRes struct {
	g.Meta `mime:"application/json"`
	Result interface{} `json:"result"`
}

type VerifyDataReq struct {
	g.Meta `path:"/blockchain/data/verify" method:"get" tags:"区块链管理" summary:"数据验真"`
	Hash   string `json:"hash"`
	TxId   string `json:"txId"`
}
type VerifyDataRes struct {
	g.Meta  `mime:"application/json"`
	Result  interface{} `json:"result"`
}
