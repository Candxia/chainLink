package model

import "github.com/gogf/gf/v2/os/gtime"

type UserLoginInput struct {
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}

type UserLoginOutput struct {
	Token    string `json:"token"`
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

type UserRegisterInput struct {
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码|length:6,32"`
	Email    string `json:"email" v:"required#请输入邮箱"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   int    `json:"status"`
}

type UserInfo struct {
	Id        uint        `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Role      string      `json:"role"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type RoleInfo struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Permissions string      `json:"permissions"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type EnterpriseInfo struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Contact     string      `json:"contact"`
	Phone       string      `json:"phone"`
	Address     string      `json:"address"`
	Status      int         `json:"status"`
	UserId      uint        `json:"userId"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type ProductInfo struct {
	Id           uint        `json:"id"`
	Name         string      `json:"name"`
	Category     string      `json:"category"`
	Spec         string      `json:"spec"`
	Unit         string      `json:"unit"`
	Description  string      `json:"description"`
	EnterpriseId uint        `json:"enterpriseId"`
	Status       string      `json:"status"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

type BatchInfo struct {
	Id          uint        `json:"id"`
	BatchNo     string      `json:"batchNo"`
	ProductId   uint        `json:"productId"`
	Quantity    int         `json:"quantity"`
	ProduceDate string      `json:"produceDate"`
	ExpireDate  string      `json:"expireDate"`
	Status      string      `json:"status"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type TraceRecordInfo struct {
	Id          uint        `json:"id"`
	ProductId   uint        `json:"productId"`
	BatchId     uint        `json:"batchId"`
	RecordType  string      `json:"recordType"`
	Content     string      `json:"content"`
	Operator    string      `json:"operator"`
	Location    string      `json:"location"`
	TxHash      string      `json:"txHash"`
	BlockNumber uint64      `json:"blockNumber"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type SupplierInfo struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Contact     string      `json:"contact"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	Category    string      `json:"category"`
	QualLevel   string      `json:"qualLevel"`
	Status      string      `json:"status"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type OrderInfo struct {
	Id           uint        `json:"id"`
	OrderNo      string      `json:"orderNo"`
	ProductId    uint        `json:"productId"`
	BatchId      uint        `json:"batchId"`
	SupplierId   uint        `json:"supplierId"`
	Quantity     int         `json:"quantity"`
	TotalAmount  float64     `json:"totalAmount"`
	Status       string      `json:"status"`
	Remark       string      `json:"remark"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

type WarehouseRecord struct {
	Id         uint        `json:"id"`
	OrderId    uint        `json:"orderId"`
	ProductId  uint        `json:"productId"`
	BatchId    uint        `json:"batchId"`
	OpType     string      `json:"opType"`
	Quantity   int         `json:"quantity"`
	Operator   string      `json:"operator"`
	Remark     string      `json:"remark"`
	CreatedAt  *gtime.Time `json:"createdAt"`
}

type LogisticsRecord struct {
	Id        uint        `json:"id"`
	OrderId   uint        `json:"orderId"`
	NodeName  string      `json:"nodeName"`
	Location  string      `json:"location"`
	Status    string      `json:"status"`
	Operator  string      `json:"operator"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type BlockInfo struct {
	Number       uint64   `json:"number"`
	Hash         string   `json:"hash"`
	ParentHash   string   `json:"parentHash"`
	Timestamp    uint64   `json:"timestamp"`
	TxCount      int      `json:"txCount"`
	Transactions []string `json:"transactions"`
}

type TransactionInfo struct {
	TxHash    string `json:"txHash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	BlockNum  uint64 `json:"blockNum"`
	Status    uint64 `json:"status"`
	Data      string `json:"data"`
	Timestamp uint64 `json:"timestamp"`
}

type ContractInfo struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	ABI         string `json:"abi"`
	DeployTx    string `json:"deployTx"`
	DeployTime  string `json:"deployTime"`
	Owner       string `json:"owner"`
	Status      string `json:"status"`
}

type SystemConfig struct {
	Id        uint        `json:"id"`
	Key       string      `json:"key"`
	Value     string      `json:"value"`
	Desc      string      `json:"desc"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type NotificationInfo struct {
	Id         uint        `json:"id"`
	UserId     uint        `json:"userId"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Type       string      `json:"type"`
	IsRead     int         `json:"isRead"`
	CreatedAt  *gtime.Time `json:"createdAt"`
}

type DashboardData struct {
	TotalUsers         int `json:"totalUsers"`
	TotalProducts      int `json:"totalProducts"`
	TotalOrders        int `json:"totalOrders"`
	TotalSuppliers     int `json:"totalSuppliers"`
	TodayTraceRecords  int `json:"todayTraceRecords"`
	PendingOrders      int `json:"pendingOrders"`
	TotalBlockCount    int `json:"totalBlockCount"`
	ActiveContracts    int `json:"activeContracts"`
}

// ==================== 分页信息 ====================

type PageInfo struct {
	Page   int   `json:"page"  default:"1"  v:"required|min:1"`
	Limit  int   `json:"limit" default:"20" v:"required|min:10|max:1000"`
	Cursor int64 `json:"cursor,omitempty" default:"0"`
}

func (p *PageInfo) Paginate() (limit, offset int) {
	limit = p.Limit
	offset = p.Limit * (p.Page - 1)
	return limit, offset
}

func (p *PageInfo) MaxPage(total int) (page int) {
	page = total / p.Limit
	if total%p.Limit != 0 {
		page += 1
	}
	return page
}

// ==================== 仓库管理 数据模型 ====================

type WarehouseInfo struct {
	Id        uint        `json:"id"`
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	Company   string      `json:"company"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type WarehouseCreateInput struct {
	Code    string `json:"code" v:"required#请输入仓库编码"`
	Name    string `json:"name" v:"required#请输入仓库名称"`
	Address string `json:"address"`
	Company string `json:"company"`
}

type WarehouseUpdateInput struct {
	Id      uint   `json:"id" v:"required#请输入仓库ID"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Company string `json:"company"`
	Status  int    `json:"status"`
}

type WarehouseAreaInfo struct {
	Id          uint        `json:"id"`
	WarehouseId uint        `json:"warehouseId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type WarehouseAreaCreateInput struct {
	WarehouseId uint   `json:"warehouseId" v:"required#请选择所属仓库"`
	Name        string `json:"name" v:"required#请输入区域名称"`
	Description string `json:"description"`
}

type WarehouseShelfInfo struct {
	Id          uint        `json:"id"`
	AreaId      uint        `json:"areaId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedAt   *gtime.Time `json:"createdAt"`
}

type WarehouseShelfCreateInput struct {
	AreaId      uint   `json:"areaId" v:"required#请选择所属区域"`
	Name        string `json:"name" v:"required#请输入货架名称"`
	Description string `json:"description"`
}

type StocktakeInfo struct {
	Id           uint        `json:"id"`
	WarehouseId  uint        `json:"warehouseId"`
	StocktakeNo  string      `json:"stocktakeNo"`
	Type         int         `json:"type"`
	Status       int         `json:"status"`
	Operator     string      `json:"operator"`
	StocktakeTime *gtime.Time `json:"stocktakeTime"`
	Remark       string      `json:"remark"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

type StocktakeCreateInput struct {
	WarehouseId uint   `json:"warehouseId" v:"required#请选择仓库"`
	Type        int    `json:"type" v:"required#请选择盘点类型"`
	Operator    string `json:"operator" v:"required#请输入盘点人员"`
	Remark      string `json:"remark"`
}

type InventoryCarInfo struct {
	Id          uint        `json:"id"`
	WarehouseId uint        `json:"warehouseId"`
	Vin         string      `json:"vin"`
	Brand       string      `json:"brand"`
	Model       string      `json:"model"`
	PlateNo     string      `json:"plateNo"`
	Color       string      `json:"color"`
	Year        int         `json:"year"`
	Status      int         `json:"status"`
	ShelfId     uint        `json:"shelfId"`
	EntryDate   *gtime.Time `json:"entryDate"`
	Remark      string      `json:"remark"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}

type InventoryRawMaterialInfo struct {
	Id           uint        `json:"id"`
	WarehouseId  uint        `json:"warehouseId"`
	MaterialNo   string      `json:"materialNo"`
	MaterialType string      `json:"materialType"`
	Quantity     float64     `json:"quantity"`
	ShelfId      uint        `json:"shelfId"`
	Remark       string      `json:"remark"`
	CreatedAt    *gtime.Time `json:"createdAt"`
	UpdatedAt    *gtime.Time `json:"updatedAt"`
}

type RawMaterialInInput struct {
	WarehouseId  uint    `json:"warehouseId" v:"required#请选择仓库"`
	MaterialNo   string  `json:"materialNo" v:"required#请输入原材料编号"`
	MaterialType string  `json:"materialType" v:"required#请输入原材料类型"`
	Quantity     float64 `json:"quantity" v:"required#请输入数量"`
	ShelfId      uint    `json:"shelfId"`
	Operator     string  `json:"operator"`
	Remark       string  `json:"remark"`
}

type RawMaterialOutInput struct {
	Id       uint    `json:"id" v:"required#请选择原材料"`
	Quantity float64 `json:"quantity" v:"required#请输入出库数量"`
	Operator string  `json:"operator"`
	Remark   string  `json:"remark"`
}

type InventoryWasteInfo struct {
	Id          uint        `json:"id"`
	WarehouseId uint        `json:"warehouseId"`
	WasteNo     string      `json:"wasteNo"`
	WasteType   string      `json:"wasteType"`
	Quantity    float64     `json:"quantity"`
	ShelfId     uint        `json:"shelfId"`
	Remark      string      `json:"remark"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}

type WasteInInput struct {
	WarehouseId uint    `json:"warehouseId" v:"required#请选择仓库"`
	WasteNo     string  `json:"wasteNo" v:"required#请输入废料编号"`
	WasteType   string  `json:"wasteType" v:"required#请选择废料类型"`
	Quantity    float64 `json:"quantity" v:"required#请输入数量"`
	ShelfId     uint    `json:"shelfId"`
	Operator    string  `json:"operator"`
	Remark      string  `json:"remark"`
}

type InventoryPartTraceableInfo struct {
	Id          uint        `json:"id"`
	WarehouseId uint        `json:"warehouseId"`
	PartName    string      `json:"partName"`
	PartType    string      `json:"partType"`
	CarModel    string      `json:"carModel"`
	CarSeries   string      `json:"carSeries"`
	Description string      `json:"description"`
	Vin         string      `json:"vin"`
	ShelfId     uint        `json:"shelfId"`
	Quantity    int         `json:"quantity"`
	Remark      string      `json:"remark"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}

type PartTraceableOutInput struct {
	Id       uint   `json:"id" v:"required#请选择配件"`
	Quantity int    `json:"quantity" v:"required#请输入出库数量"`
	Operator string `json:"operator"`
	Remark   string `json:"remark"`
}

type InventoryPartUntraceableInfo struct {
	Id          uint        `json:"id"`
	WarehouseId uint        `json:"warehouseId"`
	PartName    string      `json:"partName"`
	PartType    string      `json:"partType"`
	Quantity    int         `json:"quantity"`
	Remark      string      `json:"remark"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}

type PartUntraceableInInput struct {
	WarehouseId uint   `json:"warehouseId" v:"required#请选择仓库"`
	PartName    string `json:"partName" v:"required#请输入配件名称"`
	PartType    string `json:"partType"`
	Quantity    int    `json:"quantity" v:"required#请输入数量"`
	Operator    string `json:"operator"`
	Remark      string `json:"remark"`
}

type PartUntraceableOutInput struct {
	Id       uint   `json:"id" v:"required#请选择配件"`
	Quantity int    `json:"quantity" v:"required#请输入出库数量"`
	Operator string `json:"operator"`
	Remark   string `json:"remark"`
}

type InboundRecordInfo struct {
	Id            uint        `json:"id"`
	RecordNo      string      `json:"recordNo"`
	InboundType   string      `json:"inboundType"`
	InboundMethod string      `json:"inboundMethod"`
	WarehouseId   uint        `json:"warehouseId"`
	Operator      string      `json:"operator"`
	TotalQuantity int         `json:"totalQuantity"`
	Status        int         `json:"status"`
	Remark        string      `json:"remark"`
	InboundTime   *gtime.Time `json:"inboundTime"`
	CreatedAt     *gtime.Time `json:"createdAt"`
}

type OutboundRecordInfo struct {
	Id             uint        `json:"id"`
	RecordNo       string      `json:"recordNo"`
	OutboundType   string      `json:"outboundType"`
	OutboundMethod string      `json:"outboundMethod"`
	WarehouseId    uint        `json:"warehouseId"`
	Operator       string      `json:"operator"`
	TotalQuantity  int         `json:"totalQuantity"`
	Status         int         `json:"status"`
	Remark         string      `json:"remark"`
	OutboundTime   *gtime.Time `json:"outboundTime"`
	CreatedAt      *gtime.Time `json:"createdAt"`
}

type TransferRecordInfo struct {
	Id              uint        `json:"id"`
	RecordNo        string      `json:"recordNo"`
	TransferType    string      `json:"transferType"`
	FromWarehouseId uint        `json:"fromWarehouseId"`
	ToWarehouseId   uint        `json:"toWarehouseId"`
	ItemTypeDesc    string      `json:"itemTypeDesc"`
	TotalCategories int         `json:"totalCategories"`
	TotalQuantity   int         `json:"totalQuantity"`
	Status          int         `json:"status"`
	Operator        string      `json:"operator"`
	Remark          string      `json:"remark"`
	TransferTime    *gtime.Time `json:"transferTime"`
	CreatedAt       *gtime.Time `json:"createdAt"`
}
