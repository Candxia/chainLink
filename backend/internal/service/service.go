package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/model"
)

type IUser interface {
	Login(ctx context.Context, req model.UserLoginInput) (*model.UserLoginOutput, error)
	Register(ctx context.Context, req model.UserRegisterInput) error
	GetUserInfo(ctx context.Context, userId uint) (*model.UserInfo, error)
	GetUserList(ctx context.Context, page, pageSize int) ([]*model.UserInfo, int, error)
	UpdateUser(ctx context.Context, userId uint, data map[string]interface{}) error
	DeleteUser(ctx context.Context, id uint) error
	GetRoleList(ctx context.Context) ([]*model.RoleInfo, error)
	CreateRole(ctx context.Context, role model.RoleInfo) error
	UpdateRole(ctx context.Context, role model.RoleInfo) error
	DeleteRole(ctx context.Context, id uint) error
	GetEnterpriseList(ctx context.Context, page, pageSize int) ([]*model.EnterpriseInfo, int, error)
	CreateEnterprise(ctx context.Context, ent model.EnterpriseInfo) error
	UpdateEnterprise(ctx context.Context, ent model.EnterpriseInfo) error
}

type ITrace interface {
	CreateProduct(ctx context.Context, data map[string]interface{}) (uint, error)
	UpdateProduct(ctx context.Context, data map[string]interface{}) error
	DeleteProduct(ctx context.Context, id uint) error
	GetProduct(ctx context.Context, id uint) (*model.ProductInfo, error)
	GetProductList(ctx context.Context, page, pageSize int, name, category string) ([]*model.ProductInfo, int, error)
	CreateBatch(ctx context.Context, data map[string]interface{}) (uint, error)
	GetBatch(ctx context.Context, id uint) (*model.BatchInfo, error)
	GetBatchList(ctx context.Context, page, pageSize int, productId uint) ([]*model.BatchInfo, int, error)
	CreateTraceRecord(ctx context.Context, data map[string]interface{}) (*model.TraceRecordInfo, error)
	GetTraceRecord(ctx context.Context, id uint) (*model.TraceRecordInfo, error)
	GetTraceRecordList(ctx context.Context, productId uint) ([]*model.TraceRecordInfo, error)
	GetTraceChain(ctx context.Context, productId uint) ([]*model.TraceRecordInfo, error)
	GenerateQRCode(ctx context.Context, productId uint) (string, error)
	PublicQuery(ctx context.Context, query string) (*model.TraceRecordInfo, error)
}

type ISupply interface {
	CreateSupplier(ctx context.Context, data map[string]interface{}) (uint, error)
	UpdateSupplier(ctx context.Context, data map[string]interface{}) error
	GetSupplier(ctx context.Context, id uint) (*model.SupplierInfo, error)
	GetSupplierList(ctx context.Context, page, pageSize int, name string) ([]*model.SupplierInfo, int, error)
	DeleteSupplier(ctx context.Context, id uint) error
	CreateOrder(ctx context.Context, data map[string]interface{}) (uint, error)
	UpdateOrder(ctx context.Context, data map[string]interface{}) error
	GetOrder(ctx context.Context, id uint) (*model.OrderInfo, error)
	GetOrderList(ctx context.Context, page, pageSize int, status string) ([]*model.OrderInfo, int, error)
	DeleteOrder(ctx context.Context, id uint) error
	WarehouseOp(ctx context.Context, data map[string]interface{}, opType string) error
	GetStockList(ctx context.Context, page, pageSize int, productId uint) ([]*model.WarehouseRecord, int, error)
	CreateLogisticsRecord(ctx context.Context, data map[string]interface{}) error
	GetLogisticsList(ctx context.Context, orderId uint) ([]*model.LogisticsRecord, error)
}

type IBlockchain interface {
	GetBlockList(ctx context.Context, page, pageSize int) ([]*model.BlockInfo, int, error)
	GetBlockByHash(ctx context.Context, hash string) (*model.BlockInfo, error)
	GetTransactionList(ctx context.Context, page, pageSize int) ([]*model.TransactionInfo, int, error)
	GetTransactionDetail(ctx context.Context, txId string) (*model.TransactionInfo, error)
	DeployContract(ctx context.Context, data map[string]interface{}) (*model.ContractInfo, error)
	GetContractInfo(ctx context.Context, address string) (*model.ContractInfo, error)
	GetContractList(ctx context.Context) ([]*model.ContractInfo, error)
	UploadToChain(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error)
	VerifyData(ctx context.Context, hash, txId string) (map[string]interface{}, error)
}

type ISystem interface {
	GetDashboard(ctx context.Context) (*model.DashboardData, error)
	GetConfigList(ctx context.Context) ([]*model.SystemConfig, error)
	UpdateConfig(ctx context.Context, data map[string]interface{}) error
	GetLogList(ctx context.Context, page, pageSize int, logType string) ([]map[string]interface{}, int, error)
	DeleteLog(ctx context.Context, id uint) error
	GetNotificationList(ctx context.Context, userId uint) ([]*model.NotificationInfo, error)
	MarkNotificationRead(ctx context.Context, id uint) error
	UploadFile(ctx context.Context, file *ghttp.UploadFile) (string, error)
}

var (
	localUser       IUser
	localTrace      ITrace
	localSupply     ISupply
	localBlockchain IBlockchain
	localSystem     ISystem
)

func User() IUser {
	if localUser == nil {
		panic("User service not initialized")
	}
	return localUser
}

func Trace() ITrace {
	if localTrace == nil {
		panic("Trace service not initialized")
	}
	return localTrace
}

func Supply() ISupply {
	if localSupply == nil {
		panic("Supply service not initialized")
	}
	return localSupply
}

func Blockchain() IBlockchain {
	if localBlockchain == nil {
		panic("Blockchain service not initialized")
	}
	return localBlockchain
}

func System() ISystem {
	if localSystem == nil {
		panic("System service not initialized")
	}
	return localSystem
}

func RegisterUser(s IUser)       { localUser = s }
func RegisterTrace(s ITrace)     { localTrace = s }
func RegisterSupply(s ISupply)   { localSupply = s }
func RegisterBlockchain(s IBlockchain) { localBlockchain = s }
func RegisterSystem(s ISystem)       { localSystem = s }
