package blockchain

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type sBlockchain struct{}
func init() {
	service.RegisterBlockchain(&sBlockchain{})
}

func randStr(n int) string {
	const letters = "0123456789abcdef"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func nowUnix() uint64 {
	return uint64(time.Now().Unix())
}

func (s *sBlockchain) GetBlockList(ctx context.Context, page, pageSize int) ([]*model.BlockInfo, int, error) {
	total := 100
	var blocks []*model.BlockInfo
	for i := 0; i < pageSize && i < 10; i++ {
		blocks = append(blocks, &model.BlockInfo{
			Number:     uint64(total - i),
			Hash:       "0x" + randStr(64),
			ParentHash: "0x" + randStr(64),
			Timestamp:  nowUnix(),
			TxCount:    rand.Intn(19) + 1,
		})
	}
	return blocks, total, nil
}

func (s *sBlockchain) GetBlockByHash(ctx context.Context, hash string) (*model.BlockInfo, error) {
	return &model.BlockInfo{
		Number:     uint64(rand.Intn(10000) + 1),
		Hash:       hash,
		ParentHash: "0x" + randStr(64),
		Timestamp:  nowUnix(),
		TxCount:    rand.Intn(19) + 1,
	}, nil
}

func (s *sBlockchain) GetTransactionList(ctx context.Context, page, pageSize int) ([]*model.TransactionInfo, int, error) {
	total := 500
	var txs []*model.TransactionInfo
	for i := 0; i < pageSize; i++ {
		txs = append(txs, &model.TransactionInfo{
			TxHash:    "0x" + randStr(64),
			From:      "0x" + randStr(40),
			To:        "0x" + randStr(40),
			Value:     fmt.Sprintf("%d", rand.Intn(10000)+1),
			BlockNum:  uint64(rand.Intn(1000) + 1),
			Status:    1,
			Data:      "0x",
			Timestamp: nowUnix(),
		})
	}
	return txs, total, nil
}

func (s *sBlockchain) GetTransactionDetail(ctx context.Context, txId string) (*model.TransactionInfo, error) {
	return &model.TransactionInfo{
		TxHash:    txId,
		From:      "0x" + randStr(40),
		To:        "0x" + randStr(40),
		Value:     "1000",
		BlockNum:  uint64(rand.Intn(10000) + 1),
		Status:    1,
		Data:      "0x54726163655265636f7264",
		Timestamp: nowUnix(),
	}, nil
}

func (s *sBlockchain) DeployContract(ctx context.Context, data map[string]interface{}) (*model.ContractInfo, error) {
	name := ""
	if v, ok := data["name"]; ok {
		name = fmt.Sprintf("%v", v)
	}
	contract := &model.ContractInfo{
		Address:     "0x" + randStr(40),
		Name:        name,
		ABI:         "[]",
		DeployTx:    "0x" + randStr(64),
		DeployTime:  time.Now().Format("2006-01-02 15:04:05"),
		Owner:       "0x" + randStr(40),
		Status:      "active",
	}
	_, err := g.DB().Model("contract").Insert(g.Map{
		"address":     contract.Address,
		"name":        contract.Name,
		"abi":         contract.ABI,
		"deploy_tx":   contract.DeployTx,
		"deploy_time": contract.DeployTime,
		"owner":       contract.Owner,
		"status":      contract.Status,
	})
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (s *sBlockchain) GetContractInfo(ctx context.Context, address string) (*model.ContractInfo, error) {
	contract, err := g.DB().Model("contract").Where("address", address).One()
	if err != nil {
		return nil, err
	}
	if contract == nil {
		return nil, fmt.Errorf("合约不存在")
	}
	return &model.ContractInfo{
		Address:    contract["address"].String(),
		Name:       contract["name"].String(),
		ABI:        contract["abi"].String(),
		DeployTx:   contract["deploy_tx"].String(),
		DeployTime: contract["deploy_time"].String(),
		Owner:      contract["owner"].String(),
		Status:     contract["status"].String(),
	}, nil
}

func (s *sBlockchain) GetContractList(ctx context.Context) ([]*model.ContractInfo, error) {
	contracts, err := g.DB().Model("contract").All()
	if err != nil {
		return nil, err
	}
	var list []*model.ContractInfo
	for _, c := range contracts {
		list = append(list, &model.ContractInfo{
			Address: c["address"].String(),
			Name:    c["name"].String(),
			Status:  c["status"].String(),
		})
	}
	return list, nil
}

func (s *sBlockchain) UploadToChain(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return g.Map{
		"txHash":      "0x" + randStr(64),
		"blockNumber": rand.Intn(90000) + 10000,
		"dataHash":    "0x" + randStr(64),
		"status":      "confirmed",
	}, nil
}

func (s *sBlockchain) VerifyData(ctx context.Context, hash, txId string) (map[string]interface{}, error) {
	return g.Map{
		"verified":    true,
		"dataHash":    hash,
		"txHash":      txId,
		"blockNumber": rand.Intn(90000) + 10000,
		"message":     "数据验证通过，已上链存证",
	}, nil
}
