package warehouse

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"cl_system/internal/model"
	"cl_system/internal/service"
	"time"
)

type sWarehouse struct{}

func init() {
	service.RegisterWarehouse(&sWarehouse{})
}

// ==================== 仓库信息 ====================

func (s *sWarehouse) GetWarehouseList(ctx context.Context, page, pageSize int, name string) ([]*model.WarehouseInfo, int, error) {
	m := g.DB().Model("warehouse")
	if name != "" {
		m = m.WhereLike("name", "%"+name+"%")
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.WarehouseInfo
	for _, r := range records {
		list = append(list, &model.WarehouseInfo{
			Id:        r["id"].Uint(),
			Code:      r["code"].String(),
			Name:      r["name"].String(),
			Address:   r["address"].String(),
			Company:   r["company"].String(),
			Status:    r["status"].Int(),
			CreatedAt: r["created_at"].GTime(),
			UpdatedAt: r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) CreateWarehouse(ctx context.Context, req model.WarehouseCreateInput) error {
	_, err := g.DB().Model("warehouse").Insert(g.Map{
		"code":    req.Code,
		"name":    req.Name,
		"address": req.Address,
		"company": req.Company,
		"status":  1,
	})
	return err
}

func (s *sWarehouse) UpdateWarehouse(ctx context.Context, req model.WarehouseUpdateInput) error {
	data := g.Map{}
	if req.Code != "" {
		data["code"] = req.Code
	}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Address != "" {
		data["address"] = req.Address
	}
	if req.Company != "" {
		data["company"] = req.Company
	}
	if req.Status != 0 {
		data["status"] = req.Status
	}
	_, err := g.DB().Model("warehouse").Where("id", req.Id).Update(data)
	return err
}

func (s *sWarehouse) DeleteWarehouse(ctx context.Context, id uint) error {
	// 校验是否有关联库存
	count, err := g.DB().Model("inventory_car").Where("warehouse_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该仓库下有关联整车库存，无法删除")
	}
	count, err = g.DB().Model("inventory_raw_material").Where("warehouse_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该仓库下有关联原材料库存，无法删除")
	}
	count, err = g.DB().Model("inventory_waste").Where("warehouse_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该仓库下有关联危固废库存，无法删除")
	}
	count, err = g.DB().Model("inventory_part_traceable").Where("warehouse_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该仓库下有关联溯源件库存，无法删除")
	}
	count, err = g.DB().Model("inventory_part_untraceable").Where("warehouse_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该仓库下有关联非溯源件库存，无法删除")
	}
	_, err = g.DB().Model("warehouse").Where("id", id).Delete()
	return err
}

func (s *sWarehouse) UpdateWarehouseStatus(ctx context.Context, id uint, status int) error {
	if status != 0 && status != 1 {
		return fmt.Errorf("状态值无效")
	}
	_, err := g.DB().Model("warehouse").Where("id", id).Update(g.Map{"status": status})
	return err
}

// ==================== 区域 ====================

func (s *sWarehouse) GetAreaList(ctx context.Context, warehouseId uint) ([]*model.WarehouseAreaInfo, error) {
	m := g.DB().Model("warehouse_area")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	records, err := m.Order("id asc").All()
	if err != nil {
		return nil, err
	}
	var list []*model.WarehouseAreaInfo
	for _, r := range records {
		list = append(list, &model.WarehouseAreaInfo{
			Id:          r["id"].Uint(),
			WarehouseId: r["warehouse_id"].Uint(),
			Name:        r["name"].String(),
			Description: r["description"].String(),
			CreatedAt:   r["created_at"].GTime(),
		})
	}
	return list, nil
}

func (s *sWarehouse) CreateArea(ctx context.Context, req model.WarehouseAreaCreateInput) error {
	_, err := g.DB().Model("warehouse_area").Insert(g.Map{
		"warehouse_id": req.WarehouseId,
		"name":         req.Name,
		"description":  req.Description,
	})
	return err
}

func (s *sWarehouse) DeleteArea(ctx context.Context, id uint) error {
	// 校验是否有子货架
	count, err := g.DB().Model("warehouse_shelf").Where("area_id", id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该区域下有关联货架，请先删除货架")
	}
	_, err = g.DB().Model("warehouse_area").Where("id", id).Delete()
	return err
}

// ==================== 货架 ====================

func (s *sWarehouse) GetShelfList(ctx context.Context, areaId uint) ([]*model.WarehouseShelfInfo, error) {
	m := g.DB().Model("warehouse_shelf")
	if areaId > 0 {
		m = m.Where("area_id", areaId)
	}
	records, err := m.Order("id asc").All()
	if err != nil {
		return nil, err
	}
	var list []*model.WarehouseShelfInfo
	for _, r := range records {
		list = append(list, &model.WarehouseShelfInfo{
			Id:          r["id"].Uint(),
			AreaId:      r["area_id"].Uint(),
			Name:        r["name"].String(),
			Description: r["description"].String(),
			CreatedAt:   r["created_at"].GTime(),
		})
	}
	return list, nil
}

func (s *sWarehouse) CreateShelf(ctx context.Context, req model.WarehouseShelfCreateInput) error {
	_, err := g.DB().Model("warehouse_shelf").Insert(g.Map{
		"area_id":     req.AreaId,
		"name":        req.Name,
		"description": req.Description,
	})
	return err
}

func (s *sWarehouse) DeleteShelf(ctx context.Context, id uint) error {
	_, err := g.DB().Model("warehouse_shelf").Where("id", id).Delete()
	return err
}

// ==================== 盘存 ====================

func (s *sWarehouse) GetStocktakeList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.StocktakeInfo, int, error) {
	m := g.DB().Model("stocktake")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.StocktakeInfo
	for _, r := range records {
		list = append(list, &model.StocktakeInfo{
			Id:            r["id"].Uint(),
			WarehouseId:   r["warehouse_id"].Uint(),
			StocktakeNo:   r["stocktake_no"].String(),
			Type:          r["type"].Int(),
			Status:        r["status"].Int(),
			Operator:      r["operator"].String(),
			StocktakeTime: r["stocktake_time"].GTime(),
			Remark:        r["remark"].String(),
			CreatedAt:     r["created_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) CreateStocktake(ctx context.Context, req model.StocktakeCreateInput) error {
	stocktakeNo := "ST-" + time.Now().Format("20060102") + grand.S(6)
	_, err := g.DB().Model("stocktake").Insert(g.Map{
		"warehouse_id": req.WarehouseId,
		"stocktake_no": stocktakeNo,
		"type":         req.Type,
		"status":       0,
		"operator":     req.Operator,
		"remark":       req.Remark,
	})
	return err
}

// ==================== 整车库存 ====================

func (s *sWarehouse) GetInventoryCarList(ctx context.Context, page, pageSize int, warehouseId uint, brand, modelName, vin string) ([]*model.InventoryCarInfo, int, error) {
	m := g.DB().Model("inventory_car")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	if brand != "" {
		m = m.WhereLike("brand", "%"+brand+"%")
	}
	if modelName != "" {
		m = m.WhereLike("model", "%"+modelName+"%")
	}
	if vin != "" {
		m = m.WhereLike("vin", "%"+vin+"%")
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InventoryCarInfo
	for _, r := range records {
		list = append(list, &model.InventoryCarInfo{
			Id:          r["id"].Uint(),
			WarehouseId: r["warehouse_id"].Uint(),
			Vin:         r["vin"].String(),
			Brand:       r["brand"].String(),
			Model:       r["model"].String(),
			PlateNo:     r["plate_no"].String(),
			Color:       r["color"].String(),
			Year:        r["year"].Int(),
			Status:      r["status"].Int(),
			ShelfId:     r["shelf_id"].Uint(),
			EntryDate:   r["entry_date"].GTime(),
			Remark:      r["remark"].String(),
			CreatedAt:   r["created_at"].GTime(),
			UpdatedAt:   r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

// ==================== 原材料库存 ====================

func (s *sWarehouse) GetRawMaterialList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryRawMaterialInfo, int, error) {
	m := g.DB().Model("inventory_raw_material")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InventoryRawMaterialInfo
	for _, r := range records {
		list = append(list, &model.InventoryRawMaterialInfo{
			Id:           r["id"].Uint(),
			WarehouseId:  r["warehouse_id"].Uint(),
			MaterialNo:   r["material_no"].String(),
			MaterialType: r["material_type"].String(),
			Quantity:     r["quantity"].Float64(),
			ShelfId:      r["shelf_id"].Uint(),
			Remark:       r["remark"].String(),
			CreatedAt:    r["created_at"].GTime(),
			UpdatedAt:    r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) RawMaterialIn(ctx context.Context, req model.RawMaterialInInput) error {
	_, err := g.DB().Model("inventory_raw_material").Insert(g.Map{
		"warehouse_id":  req.WarehouseId,
		"material_no":   req.MaterialNo,
		"material_type": req.MaterialType,
		"quantity":      req.Quantity,
		"shelf_id":      req.ShelfId,
		"remark":        req.Remark,
	})
	if err != nil {
		return err
	}
	// 记录入库
	inboundNo := "IN-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("inbound_record").Insert(g.Map{
		"record_no":    inboundNo,
		"inbound_type": "raw_material",
		"inbound_method": "dismantle",
		"warehouse_id": req.WarehouseId,
		"operator":     req.Operator,
		"total_quantity": 1,
		"status":       1,
		"inbound_time": time.Now(),
	})
	return err
}

func (s *sWarehouse) RawMaterialOut(ctx context.Context, req model.RawMaterialOutInput) error {
	// 检查库存
	record, err := g.DB().Model("inventory_raw_material").Where("id", req.Id).One()
	if err != nil {
		return err
	}
	if record == nil {
		return fmt.Errorf("原材料不存在")
	}
	currentQty := record["quantity"].Float64()
	if currentQty < req.Quantity {
		return fmt.Errorf("库存不足，当前库存: %.2f, 需要出库: %.2f", currentQty, req.Quantity)
	}
	newQty := currentQty - req.Quantity
	if newQty <= 0 {
		_, err = g.DB().Model("inventory_raw_material").Where("id", req.Id).Delete()
	} else {
		_, err = g.DB().Model("inventory_raw_material").Where("id", req.Id).Update(g.Map{"quantity": newQty})
	}
	if err != nil {
		return err
	}
	// 记录出库
	outboundNo := "OUT-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("outbound_record").Insert(g.Map{
		"record_no":      outboundNo,
		"outbound_type":  "raw_material",
		"outbound_method": "sale",
		"warehouse_id":   record["warehouse_id"].Uint(),
		"operator":       req.Operator,
		"total_quantity": req.Quantity,
		"status":         1,
		"outbound_time":  time.Now(),
		"remark":         req.Remark,
	})
	return err
}

// ==================== 危固废库存 ====================

func (s *sWarehouse) GetWasteList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryWasteInfo, int, error) {
	m := g.DB().Model("inventory_waste")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InventoryWasteInfo
	for _, r := range records {
		list = append(list, &model.InventoryWasteInfo{
			Id:          r["id"].Uint(),
			WarehouseId: r["warehouse_id"].Uint(),
			WasteNo:     r["waste_no"].String(),
			WasteType:   r["waste_type"].String(),
			Quantity:    r["quantity"].Float64(),
			ShelfId:     r["shelf_id"].Uint(),
			Remark:      r["remark"].String(),
			CreatedAt:   r["created_at"].GTime(),
			UpdatedAt:   r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) WasteIn(ctx context.Context, req model.WasteInInput) error {
	_, err := g.DB().Model("inventory_waste").Insert(g.Map{
		"warehouse_id": req.WarehouseId,
		"waste_no":     req.WasteNo,
		"waste_type":   req.WasteType,
		"quantity":     req.Quantity,
		"shelf_id":     req.ShelfId,
		"remark":       req.Remark,
	})
	if err != nil {
		return err
	}
	// 记录入库
	inboundNo := "IN-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("inbound_record").Insert(g.Map{
		"record_no":      inboundNo,
		"inbound_type":   "waste",
		"inbound_method": "other",
		"warehouse_id":   req.WarehouseId,
		"operator":       req.Operator,
		"total_quantity": 1,
		"status":         1,
		"inbound_time":   time.Now(),
	})
	return err
}

// ==================== 溯源件 ====================

func (s *sWarehouse) GetPartTraceableList(ctx context.Context, page, pageSize int, warehouseId uint, partName string) ([]*model.InventoryPartTraceableInfo, int, error) {
	m := g.DB().Model("inventory_part_traceable")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	if partName != "" {
		m = m.WhereLike("part_name", "%"+partName+"%")
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InventoryPartTraceableInfo
	for _, r := range records {
		list = append(list, &model.InventoryPartTraceableInfo{
			Id:          r["id"].Uint(),
			WarehouseId: r["warehouse_id"].Uint(),
			PartName:    r["part_name"].String(),
			PartType:    r["part_type"].String(),
			CarModel:    r["car_model"].String(),
			CarSeries:   r["car_series"].String(),
			Description: r["description"].String(),
			Vin:         r["vin"].String(),
			ShelfId:     r["shelf_id"].Uint(),
			Quantity:    r["quantity"].Int(),
			Remark:      r["remark"].String(),
			CreatedAt:   r["created_at"].GTime(),
			UpdatedAt:   r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) PartTraceableOut(ctx context.Context, req model.PartTraceableOutInput) error {
	record, err := g.DB().Model("inventory_part_traceable").Where("id", req.Id).One()
	if err != nil {
		return err
	}
	if record == nil {
		return fmt.Errorf("溯源件不存在")
	}
	currentQty := record["quantity"].Int()
	if currentQty < req.Quantity {
		return fmt.Errorf("库存不足，当前库存: %d, 需要出库: %d", currentQty, req.Quantity)
	}
	newQty := currentQty - req.Quantity
	if newQty <= 0 {
		_, err = g.DB().Model("inventory_part_traceable").Where("id", req.Id).Delete()
	} else {
		_, err = g.DB().Model("inventory_part_traceable").Where("id", req.Id).Update(g.Map{"quantity": newQty})
	}
	if err != nil {
		return err
	}
	// 记录出库
	outboundNo := "OUT-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("outbound_record").Insert(g.Map{
		"record_no":       outboundNo,
		"outbound_type":   "part_traceable",
		"outbound_method": "sale",
		"warehouse_id":    record["warehouse_id"].Uint(),
		"operator":        req.Operator,
		"total_quantity":  req.Quantity,
		"status":          1,
		"outbound_time":   time.Now(),
		"remark":          req.Remark,
	})
	return err
}

// ==================== 非溯源件 ====================

func (s *sWarehouse) GetPartUntraceableList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryPartUntraceableInfo, int, error) {
	m := g.DB().Model("inventory_part_untraceable")
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InventoryPartUntraceableInfo
	for _, r := range records {
		list = append(list, &model.InventoryPartUntraceableInfo{
			Id:          r["id"].Uint(),
			WarehouseId: r["warehouse_id"].Uint(),
			PartName:    r["part_name"].String(),
			PartType:    r["part_type"].String(),
			Quantity:    r["quantity"].Int(),
			Remark:      r["remark"].String(),
			CreatedAt:   r["created_at"].GTime(),
			UpdatedAt:   r["updated_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) PartUntraceableIn(ctx context.Context, req model.PartUntraceableInInput) error {
	_, err := g.DB().Model("inventory_part_untraceable").Insert(g.Map{
		"warehouse_id": req.WarehouseId,
		"part_name":    req.PartName,
		"part_type":    req.PartType,
		"quantity":     req.Quantity,
		"remark":       req.Remark,
	})
	if err != nil {
		return err
	}
	// 记录入库
	inboundNo := "IN-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("inbound_record").Insert(g.Map{
		"record_no":      inboundNo,
		"inbound_type":   "part_untraceable",
		"inbound_method": "other",
		"warehouse_id":   req.WarehouseId,
		"operator":       req.Operator,
		"total_quantity": req.Quantity,
		"status":         1,
		"inbound_time":   time.Now(),
	})
	return err
}

func (s *sWarehouse) PartUntraceableOut(ctx context.Context, req model.PartUntraceableOutInput) error {
	record, err := g.DB().Model("inventory_part_untraceable").Where("id", req.Id).One()
	if err != nil {
		return err
	}
	if record == nil {
		return fmt.Errorf("非溯源件不存在")
	}
	currentQty := record["quantity"].Int()
	if currentQty < req.Quantity {
		return fmt.Errorf("库存不足，当前库存: %d, 需要出库: %d", currentQty, req.Quantity)
	}
	newQty := currentQty - req.Quantity
	if newQty <= 0 {
		_, err = g.DB().Model("inventory_part_untraceable").Where("id", req.Id).Delete()
	} else {
		_, err = g.DB().Model("inventory_part_untraceable").Where("id", req.Id).Update(g.Map{"quantity": newQty})
	}
	if err != nil {
		return err
	}
	// 记录出库
	outboundNo := "OUT-" + time.Now().Format("20060102") + grand.S(6)
	_, err = g.DB().Model("outbound_record").Insert(g.Map{
		"record_no":       outboundNo,
		"outbound_type":   "part_untraceable",
		"outbound_method": "sale",
		"warehouse_id":    record["warehouse_id"].Uint(),
		"operator":        req.Operator,
		"total_quantity":  req.Quantity,
		"status":          1,
		"outbound_time":   time.Now(),
		"remark":          req.Remark,
	})
	return err
}

// ==================== 操作记录 ====================

func (s *sWarehouse) GetInboundRecordList(ctx context.Context, page, pageSize int, inboundType string, warehouseId uint) ([]*model.InboundRecordInfo, int, error) {
	m := g.DB().Model("inbound_record")
	if inboundType != "" {
		m = m.Where("inbound_type", inboundType)
	}
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.InboundRecordInfo
	for _, r := range records {
		list = append(list, &model.InboundRecordInfo{
			Id:            r["id"].Uint(),
			RecordNo:      r["record_no"].String(),
			InboundType:   r["inbound_type"].String(),
			InboundMethod: r["inbound_method"].String(),
			WarehouseId:   r["warehouse_id"].Uint(),
			Operator:      r["operator"].String(),
			TotalQuantity: r["total_quantity"].Int(),
			Status:        r["status"].Int(),
			Remark:        r["remark"].String(),
			InboundTime:   r["inbound_time"].GTime(),
			CreatedAt:     r["created_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) GetInboundRecordDetail(ctx context.Context, id uint) (*model.InboundRecordInfo, error) {
	r, err := g.DB().Model("inbound_record").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, fmt.Errorf("入库记录不存在")
	}
	return &model.InboundRecordInfo{
		Id:            r["id"].Uint(),
		RecordNo:      r["record_no"].String(),
		InboundType:   r["inbound_type"].String(),
		InboundMethod: r["inbound_method"].String(),
		WarehouseId:   r["warehouse_id"].Uint(),
		Operator:      r["operator"].String(),
		TotalQuantity: r["total_quantity"].Int(),
		Status:        r["status"].Int(),
		Remark:        r["remark"].String(),
		InboundTime:   r["inbound_time"].GTime(),
		CreatedAt:     r["created_at"].GTime(),
	}, nil
}

func (s *sWarehouse) GetOutboundRecordList(ctx context.Context, page, pageSize int, outboundType string, warehouseId uint) ([]*model.OutboundRecordInfo, int, error) {
	m := g.DB().Model("outbound_record")
	if outboundType != "" {
		m = m.Where("outbound_type", outboundType)
	}
	if warehouseId > 0 {
		m = m.Where("warehouse_id", warehouseId)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.OutboundRecordInfo
	for _, r := range records {
		list = append(list, &model.OutboundRecordInfo{
			Id:             r["id"].Uint(),
			RecordNo:       r["record_no"].String(),
			OutboundType:   r["outbound_type"].String(),
			OutboundMethod: r["outbound_method"].String(),
			WarehouseId:    r["warehouse_id"].Uint(),
			Operator:       r["operator"].String(),
			TotalQuantity:  r["total_quantity"].Int(),
			Status:         r["status"].Int(),
			Remark:         r["remark"].String(),
			OutboundTime:   r["outbound_time"].GTime(),
			CreatedAt:      r["created_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) GetOutboundRecordDetail(ctx context.Context, id uint) (*model.OutboundRecordInfo, error) {
	r, err := g.DB().Model("outbound_record").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, fmt.Errorf("出库记录不存在")
	}
	return &model.OutboundRecordInfo{
		Id:             r["id"].Uint(),
		RecordNo:       r["record_no"].String(),
		OutboundType:   r["outbound_type"].String(),
		OutboundMethod: r["outbound_method"].String(),
		WarehouseId:    r["warehouse_id"].Uint(),
		Operator:       r["operator"].String(),
		TotalQuantity:  r["total_quantity"].Int(),
		Status:         r["status"].Int(),
		Remark:         r["remark"].String(),
		OutboundTime:   r["outbound_time"].GTime(),
		CreatedAt:      r["created_at"].GTime(),
	}, nil
}

func (s *sWarehouse) GetTransferRecordList(ctx context.Context, page, pageSize int, transferType string) ([]*model.TransferRecordInfo, int, error) {
	m := g.DB().Model("transfer_record")
	if transferType != "" {
		m = m.Where("transfer_type", transferType)
	}
	total, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	records, err := m.Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.TransferRecordInfo
	for _, r := range records {
		list = append(list, &model.TransferRecordInfo{
			Id:              r["id"].Uint(),
			RecordNo:        r["record_no"].String(),
			TransferType:    r["transfer_type"].String(),
			FromWarehouseId: r["from_warehouse_id"].Uint(),
			ToWarehouseId:   r["to_warehouse_id"].Uint(),
			ItemTypeDesc:    r["item_type_desc"].String(),
			TotalCategories: r["total_categories"].Int(),
			TotalQuantity:   r["total_quantity"].Int(),
			Status:          r["status"].Int(),
			Operator:        r["operator"].String(),
			Remark:          r["remark"].String(),
			TransferTime:    r["transfer_time"].GTime(),
			CreatedAt:       r["created_at"].GTime(),
		})
	}
	return list, total, nil
}

func (s *sWarehouse) GetTransferRecordDetail(ctx context.Context, id uint) (*model.TransferRecordInfo, error) {
	r, err := g.DB().Model("transfer_record").Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, fmt.Errorf("调拨记录不存在")
	}
	return &model.TransferRecordInfo{
		Id:              r["id"].Uint(),
		RecordNo:        r["record_no"].String(),
		TransferType:    r["transfer_type"].String(),
		FromWarehouseId: r["from_warehouse_id"].Uint(),
		ToWarehouseId:   r["to_warehouse_id"].Uint(),
		ItemTypeDesc:    r["item_type_desc"].String(),
		TotalCategories: r["total_categories"].Int(),
		TotalQuantity:   r["total_quantity"].Int(),
		Status:          r["status"].Int(),
		Operator:        r["operator"].String(),
		Remark:          r["remark"].String(),
		TransferTime:    r["transfer_time"].GTime(),
		CreatedAt:       r["created_at"].GTime(),
	}, nil
}
