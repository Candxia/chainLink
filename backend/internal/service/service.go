package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/model"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/model/entity"
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

// ==================== 系统管理 - 账号管理 ====================

type IAdmin interface {
	Add(ctx context.Context, params mdlSys.SysAdminAdd) error
	Del(ctx context.Context, id int64) error
	Edit(ctx context.Context, params mdlSys.SysAdminEdit) error
	Info(ctx context.Context, id int64) (result mdlSys.SysAdminInfo, err error)
	List(ctx context.Context, params mdlSys.SysAdminSearch) (result []mdlSys.SysAdminInfo, total int, err error)
	Password(ctx context.Context, id int64, pwd string) error
	Status(ctx context.Context, username string, status int) error
	ClickOut(ctx context.Context, username string) error
	SetOnline(ctx context.Context, info mdlSys.SysAdminOnline) error
}

// ==================== 系统管理 - API管理 ====================

type IApi interface {
	Add(ctx context.Context, params mdlSys.SysApiAdd) error
	Del(ctx context.Context, id int64) error
	Edit(ctx context.Context, params mdlSys.SysApiEdit) error
	List(ctx context.Context, params mdlSys.SysApiSearch) (result []entity.SysApi, total int, err error)
	Drop(ctx context.Context) (list []g.Map, err error)
	Path(ctx context.Context) (list []g.Map, err error)
}

// ==================== 系统管理 - 角色管理 ====================

type IRole interface {
	Add(ctx context.Context, params mdlSys.SysRoleAdd) error
	Del(ctx context.Context, id int64) error
	Status(ctx context.Context, params mdlSys.SysRoleStatus) error
	Edit(ctx context.Context, params mdlSys.SysRoleEdit) error
	Info(ctx context.Context, id int64) (result mdlSys.SysRoleInfo, err error)
	Drop(ctx context.Context) (result []g.Map, err error)
	List(ctx context.Context, params mdlSys.SysRoleSearch) (result []mdlSys.SysRoleTree, total int, err error)
	IsMobile(ctx context.Context, params mdlSys.SysRoleIsMobile) error
}

// ==================== 系统管理 - 菜单管理 ====================

type IMenu interface {
	Add(ctx context.Context, params mdlSys.SysMenuAdd) error
	Del(ctx context.Context, id int64) error
	Edit(ctx context.Context, params mdlSys.SysMenuEdit) error
	Info(ctx context.Context, id int64) (result entity.SysMenu, err error)
	List(ctx context.Context) (result []mdlSys.SysMenuTree, err error)
	Role(ctx context.Context) (result []mdlSys.SysMenuRole, err error)
	Drop(ctx context.Context, exMenutype []string) (result []g.Map, total int, err error)
}

// ==================== 系统管理 - 部门管理 ====================

type IDept interface {
	Add(ctx context.Context, params mdlSys.SysDeptAdd) error
	Del(ctx context.Context, id int64) error
	Exist(ctx context.Context, params mdlSys.SysDeptExist) (has bool)
	Edit(ctx context.Context, params mdlSys.SysDeptEdit) error
	Info(ctx context.Context, id int64) (result entity.SysDept, err error)
	List(ctx context.Context, params mdlSys.SysDeptSearch) (result []mdlSys.SysDeptInfo, err error)
	Drop(ctx context.Context) (result []mdlSys.SysDeptInfo, err error)
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

type IWarehouse interface {
	// 仓库信息
	GetWarehouseList(ctx context.Context, page, pageSize int, name string) ([]*model.WarehouseInfo, int, error)
	CreateWarehouse(ctx context.Context, req model.WarehouseCreateInput) error
	UpdateWarehouse(ctx context.Context, req model.WarehouseUpdateInput) error
	DeleteWarehouse(ctx context.Context, id uint) error
	UpdateWarehouseStatus(ctx context.Context, id uint, status int) error
	// 区域
	GetAreaList(ctx context.Context, warehouseId uint) ([]*model.WarehouseAreaInfo, error)
	CreateArea(ctx context.Context, req model.WarehouseAreaCreateInput) error
	DeleteArea(ctx context.Context, id uint) error
	// 货架
	GetShelfList(ctx context.Context, areaId uint) ([]*model.WarehouseShelfInfo, error)
	CreateShelf(ctx context.Context, req model.WarehouseShelfCreateInput) error
	DeleteShelf(ctx context.Context, id uint) error
	// 盘存
	GetStocktakeList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.StocktakeInfo, int, error)
	CreateStocktake(ctx context.Context, req model.StocktakeCreateInput) error
	// 整车库存
	GetInventoryCarList(ctx context.Context, page, pageSize int, warehouseId uint, brand, modelName, vin string) ([]*model.InventoryCarInfo, int, error)
	// 原材料库存
	GetRawMaterialList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryRawMaterialInfo, int, error)
	RawMaterialIn(ctx context.Context, req model.RawMaterialInInput) error
	RawMaterialOut(ctx context.Context, req model.RawMaterialOutInput) error
	// 危固废库存
	GetWasteList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryWasteInfo, int, error)
	WasteIn(ctx context.Context, req model.WasteInInput) error
	// 溯源件
	GetPartTraceableList(ctx context.Context, page, pageSize int, warehouseId uint, partName string) ([]*model.InventoryPartTraceableInfo, int, error)
	PartTraceableOut(ctx context.Context, req model.PartTraceableOutInput) error
	// 非溯源件
	GetPartUntraceableList(ctx context.Context, page, pageSize int, warehouseId uint) ([]*model.InventoryPartUntraceableInfo, int, error)
	PartUntraceableIn(ctx context.Context, req model.PartUntraceableInInput) error
	PartUntraceableOut(ctx context.Context, req model.PartUntraceableOutInput) error
	// 操作记录
	GetInboundRecordList(ctx context.Context, page, pageSize int, inboundType string, warehouseId uint) ([]*model.InboundRecordInfo, int, error)
	GetInboundRecordDetail(ctx context.Context, id uint) (*model.InboundRecordInfo, error)
	GetOutboundRecordList(ctx context.Context, page, pageSize int, outboundType string, warehouseId uint) ([]*model.OutboundRecordInfo, int, error)
	GetOutboundRecordDetail(ctx context.Context, id uint) (*model.OutboundRecordInfo, error)
	GetTransferRecordList(ctx context.Context, page, pageSize int, transferType string) ([]*model.TransferRecordInfo, int, error)
	GetTransferRecordDetail(ctx context.Context, id uint) (*model.TransferRecordInfo, error)
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
	localDept       IDept
	localUser       IUser
	localTrace      ITrace
	localSupply     ISupply
	localBlockchain IBlockchain
	localAdmin      IAdmin
	localApi        IApi
	localRole       IRole
	localMenu       IMenu
	localSystem     ISystem
	localWarehouse  IWarehouse
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

func Dept() IDept {
	if localDept == nil {
		panic("Dept service not initialized")
	}
	return localDept
}

func RegisterDept(s IDept) { localDept = s }

func RegisterUser(s IUser)       { localUser = s }
func RegisterTrace(s ITrace)     { localTrace = s }
func RegisterSupply(s ISupply)   { localSupply = s }
func RegisterBlockchain(s IBlockchain) { localBlockchain = s }
func RegisterAdmin(s IAdmin)       { localAdmin = s }
func RegisterApi(s IApi)           { localApi = s }
func RegisterRole(s IRole)         { localRole = s }
func RegisterMenu(s IMenu)         { localMenu = s }
func RegisterSystem(s ISystem)       { localSystem = s }

func Admin() IAdmin {
	if localAdmin == nil {
		panic("Admin service not initialized")
	}
	return localAdmin
}

func Api() IApi {
	if localApi == nil {
		panic("Api service not initialized")
	}
	return localApi
}

func Role() IRole {
	if localRole == nil {
		panic("Role service not initialized")
	}
	return localRole
}

func Menu() IMenu {
	if localMenu == nil {
		panic("Menu service not initialized")
	}
	return localMenu
}

func Warehouse() IWarehouse {
	if localWarehouse == nil {
		panic("Warehouse service not initialized")
	}
	return localWarehouse
}

func RegisterWarehouse(s IWarehouse) { localWarehouse = s }
