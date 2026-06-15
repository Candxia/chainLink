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
