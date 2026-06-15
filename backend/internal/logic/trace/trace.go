package trace

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type sTrace struct{}
func init() {
	service.RegisterTrace(&sTrace{})
}

func (s *sTrace) CreateProduct(ctx context.Context, data map[string]interface{}) (uint, error) {
	if data["name"] == nil || data["name"] == "" {
		return 0, fmt.Errorf("产品名称不能为空")
	}
	r, err := g.DB().Model("product").Insert(data)
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()
	return uint(id), nil
}

func (s *sTrace) UpdateProduct(ctx context.Context, data map[string]interface{}) error {
	id := data["id"]
	if id == nil {
		return fmt.Errorf("缺少产品ID")
	}
	delete(data, "id")
	_, err := g.DB().Model("product").Where("id", id).Update(data)
	return err
}

func (s *sTrace) DeleteProduct(ctx context.Context, id uint) error {
	_, err := g.DB().Model("product").Where("id", id).Delete()
	return err
}

func (s *sTrace) GetProduct(ctx context.Context, id uint) (*model.ProductInfo, error) {
	product, err := g.DB().Model("product").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("产品不存在")
	}
	return &model.ProductInfo{
		Id:           product["id"].Uint(),
		Name:         product["name"].String(),
		Category:     product["category"].String(),
		Spec:         product["spec"].String(),
		Unit:         product["unit"].String(),
		Description:  product["description"].String(),
		EnterpriseId: product["enterprise_id"].Uint(),
		Status:       product["status"].String(),
	}, nil
}

func (s *sTrace) GetProductList(ctx context.Context, page, pageSize int, name, category string) ([]*model.ProductInfo, int, error) {
	m := g.DB().Model("product")
	if name != "" {
		m = m.WhereLike("name", "%"+name+"%")
	}
	if category != "" {
		m = m.Where("category", category)
	}
	total, _ := m.Count()
	products, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.ProductInfo
	for _, p := range products {
		list = append(list, &model.ProductInfo{
			Id:          p["id"].Uint(),
			Name:        p["name"].String(),
			Category:    p["category"].String(),
			Spec:        p["spec"].String(),
			Unit:        p["unit"].String(),
			Description: p["description"].String(),
			Status:      p["status"].String(),
		})
	}
	return list, total, nil
}

func (s *sTrace) CreateBatch(ctx context.Context, data map[string]interface{}) (uint, error) {
	r, err := g.DB().Model("batch").Insert(data)
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()
	return uint(id), nil
}

func (s *sTrace) GetBatch(ctx context.Context, id uint) (*model.BatchInfo, error) {
	batch, err := g.DB().Model("batch").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if batch == nil {
		return nil, fmt.Errorf("批次不存在")
	}
	return &model.BatchInfo{
		Id:          batch["id"].Uint(),
		BatchNo:     batch["batch_no"].String(),
		ProductId:   batch["product_id"].Uint(),
		Quantity:    batch["quantity"].Int(),
		ProduceDate: batch["produce_date"].String(),
		ExpireDate:  batch["expire_date"].String(),
		Status:      batch["status"].String(),
	}, nil
}

func (s *sTrace) GetBatchList(ctx context.Context, page, pageSize int, productId uint) ([]*model.BatchInfo, int, error) {
	m := g.DB().Model("batch")
	if productId > 0 {
		m = m.Where("product_id", productId)
	}
	total, _ := m.Count()
	batches, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.BatchInfo
	for _, b := range batches {
		list = append(list, &model.BatchInfo{
			Id:          b["id"].Uint(),
			BatchNo:     b["batch_no"].String(),
			ProductId:   b["product_id"].Uint(),
			Quantity:    b["quantity"].Int(),
			ProduceDate: b["produce_date"].String(),
			ExpireDate:  b["expire_date"].String(),
			Status:      b["status"].String(),
		})
	}
	return list, total, nil
}

func (s *sTrace) CreateTraceRecord(ctx context.Context, data map[string]interface{}) (*model.TraceRecordInfo, error) {
	// 生成模拟的交易哈希（实际项目中会调用区块链接口）
	data["tx_hash"] = "0x" + fmt.Sprintf("%x", gtime.TimestampNano())
	r, err := g.DB().Model("trace_record").Insert(data)
	if err != nil {
		return nil, err
	}
	id, _ := r.LastInsertId()
	record, _ := s.GetTraceRecord(ctx, uint(id))
	return record, nil
}

func (s *sTrace) GetTraceRecord(ctx context.Context, id uint) (*model.TraceRecordInfo, error) {
	record, err := g.DB().Model("trace_record").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, fmt.Errorf("记录不存在")
	}
	return &model.TraceRecordInfo{
		Id:          record["id"].Uint(),
		ProductId:   record["product_id"].Uint(),
		BatchId:     record["batch_id"].Uint(),
		RecordType:  record["record_type"].String(),
		Content:     record["content"].String(),
		Operator:    record["operator"].String(),
		Location:    record["location"].String(),
		TxHash:      record["tx_hash"].String(),
		BlockNumber: record["block_number"].Uint64(),
	}, nil
}

func (s *sTrace) GetTraceRecordList(ctx context.Context, productId uint) ([]*model.TraceRecordInfo, error) {
	records, err := g.DB().Model("trace_record").Where("product_id", productId).Order("id asc").All()
	if err != nil {
		return nil, err
	}
	var list []*model.TraceRecordInfo
	for _, r := range records {
		list = append(list, &model.TraceRecordInfo{
			Id:          r["id"].Uint(),
			ProductId:   r["product_id"].Uint(),
			BatchId:     r["batch_id"].Uint(),
			RecordType:  r["record_type"].String(),
			Content:     r["content"].String(),
			Operator:    r["operator"].String(),
			Location:    r["location"].String(),
			TxHash:      r["tx_hash"].String(),
			BlockNumber: r["block_number"].Uint64(),
		})
	}
	return list, nil
}

func (s *sTrace) GetTraceChain(ctx context.Context, productId uint) ([]*model.TraceRecordInfo, error) {
	return s.GetTraceRecordList(ctx, productId)
}

func (s *sTrace) GenerateQRCode(ctx context.Context, productId uint) (string, error) {
	product, err := s.GetProduct(ctx, productId)
	if err != nil {
		return "", err
	}
	// 生成二维码内容（实际项目会调用二维码库）
	qrcodeContent := fmt.Sprintf("链环溯源 - 产品: %s, ID: %d", product.Name, productId)
	return qrcodeContent, nil
}

func (s *sTrace) PublicQuery(ctx context.Context, query string) (*model.TraceRecordInfo, error) {
	record, err := g.DB().Model("trace_record").Where("tx_hash", query).One()
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, fmt.Errorf("未找到相关溯源记录")
	}
	return &model.TraceRecordInfo{
		Id:          record["id"].Uint(),
		ProductId:   record["product_id"].Uint(),
		RecordType:  record["record_type"].String(),
		Content:     record["content"].String(),
		TxHash:      record["tx_hash"].String(),
		BlockNumber: record["block_number"].Uint64(),
	}, nil
}
