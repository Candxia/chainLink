package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/controller/user"
	"cl_system/internal/controller/trace"
	"cl_system/internal/controller/supply"
	"cl_system/internal/controller/blockchain"
	"cl_system/internal/controller/system"
	"cl_system/internal/controller/warehouse"
)

func Router(group *ghttp.RouterGroup) {
	// ==================== 用户与权限 ====================
	userGroup := group.Group("/user")
	userGroup.POST("/login", user.Ctl.Login)
	userGroup.POST("/register", user.Ctl.Register)
	userGroup.POST("/logout", user.Ctl.Logout)

	// 需要认证的接口
	userGroup.Middleware(AuthMiddleware)
	userGroup.GET("/info", user.Ctl.GetUserInfo)
	userGroup.GET("/list", user.Ctl.GetUserList)
	userGroup.PUT("/update", user.Ctl.UpdateUser)
	userGroup.DELETE("/delete/:id", user.Ctl.DeleteUser)
	userGroup.GET("/roles", user.Ctl.GetRoleList)
	userGroup.POST("/role", user.Ctl.CreateRole)
	userGroup.PUT("/role", user.Ctl.UpdateRole)
	userGroup.DELETE("/role/:id", user.Ctl.DeleteRole)
	userGroup.GET("/enterprise", user.Ctl.GetEnterpriseList)
	userGroup.POST("/enterprise", user.Ctl.CreateEnterprise)
	userGroup.PUT("/enterprise", user.Ctl.UpdateEnterprise)

	// ==================== 产品溯源 ====================
	traceGroup := group.Group("/trace")
	traceGroup.Middleware(AuthMiddleware)
	traceGroup.POST("/product", trace.Ctl.CreateProduct)
	traceGroup.PUT("/product", trace.Ctl.UpdateProduct)
	traceGroup.DELETE("/product/:id", trace.Ctl.DeleteProduct)
	traceGroup.GET("/product/:id", trace.Ctl.GetProduct)
	traceGroup.GET("/product/list", trace.Ctl.GetProductList)
	traceGroup.POST("/batch", trace.Ctl.CreateBatch)
	traceGroup.GET("/batch/:id", trace.Ctl.GetBatch)
	traceGroup.GET("/batch/list", trace.Ctl.GetBatchList)
	traceGroup.POST("/record", trace.Ctl.CreateTraceRecord)
	traceGroup.GET("/record/:id", trace.Ctl.GetTraceRecord)
	traceGroup.GET("/record/list/:productId", trace.Ctl.GetTraceRecordList)
	traceGroup.GET("/chain/:productId", trace.Ctl.GetTraceChain)
	traceGroup.GET("/qrcode/:productId", trace.Ctl.GenerateQRCode)

	// 公开查询接口
	traceGroup.GET("/public/query", trace.Ctl.PublicQuery)

	// ==================== 供应链管理 ====================
	supplyGroup := group.Group("/supply")
	supplyGroup.Middleware(AuthMiddleware)
	supplyGroup.POST("/supplier", supply.Ctl.CreateSupplier)
	supplyGroup.PUT("/supplier", supply.Ctl.UpdateSupplier)
	supplyGroup.GET("/supplier/:id", supply.Ctl.GetSupplier)
	supplyGroup.GET("/supplier/list", supply.Ctl.GetSupplierList)
	supplyGroup.DELETE("/supplier/:id", supply.Ctl.DeleteSupplier)
	supplyGroup.POST("/order", supply.Ctl.CreateOrder)
	supplyGroup.PUT("/order", supply.Ctl.UpdateOrder)
	supplyGroup.GET("/order/:id", supply.Ctl.GetOrder)
	supplyGroup.GET("/order/list", supply.Ctl.GetOrderList)
	supplyGroup.DELETE("/order/:id", supply.Ctl.DeleteOrder)
	supplyGroup.POST("/warehouse/in", supply.Ctl.WarehouseIn)
	supplyGroup.POST("/warehouse/out", supply.Ctl.WarehouseOut)
	supplyGroup.GET("/warehouse/stock", supply.Ctl.GetStockList)
	supplyGroup.POST("/logistics", supply.Ctl.CreateLogisticsRecord)
	supplyGroup.GET("/logistics/:orderId", supply.Ctl.GetLogisticsList)

	// ==================== 区块链管理 ====================
	blockGroup := group.Group("/blockchain")
	blockGroup.Middleware(AuthMiddleware)
	blockGroup.GET("/blocks", blockchain.Ctl.GetBlockList)
	blockGroup.GET("/block/:hash", blockchain.Ctl.GetBlockByHash)
	blockGroup.GET("/transactions", blockchain.Ctl.GetTransactionList)
	blockGroup.GET("/transaction/:txId", blockchain.Ctl.GetTransactionDetail)
	blockGroup.POST("/contract", blockchain.Ctl.DeployContract)
	blockGroup.GET("/contract/:address", blockchain.Ctl.GetContractInfo)
	blockGroup.GET("/contract/list", blockchain.Ctl.GetContractList)
	blockGroup.POST("/data/upload", blockchain.Ctl.UploadToChain)
	blockGroup.GET("/data/verify", blockchain.Ctl.VerifyData)

	// ==================== 系统管理 ====================
	sysGroup := group.Group("/system")
	sysGroup.Middleware(AuthMiddleware)
	sysGroup.GET("/dashboard", system.Ctl.GetDashboard)
	sysGroup.GET("/config", system.Ctl.GetConfig)
	sysGroup.PUT("/config", system.Ctl.UpdateConfig)
	sysGroup.GET("/logs", system.Ctl.GetLogList)
	sysGroup.DELETE("/logs/:id", system.Ctl.DeleteLog)
	sysGroup.GET("/notifications", system.Ctl.GetNotificationList)
	sysGroup.POST("/notification/read/:id", system.Ctl.MarkNotificationRead)

	// ==================== 账号管理 ====================
	sysGroup.POST("/admin/add", system.Ctl.AdminAdd)
	sysGroup.DELETE("/admin/del", system.Ctl.AdminDel)
	sysGroup.PUT("/admin/edit", system.Ctl.AdminEdit)
	sysGroup.GET("/admin/info", system.Ctl.AdminInfo)
	sysGroup.GET("/admin/list", system.Ctl.AdminList)
	sysGroup.PUT("/admin/password", system.Ctl.AdminPassword)
	sysGroup.PUT("/admin/status", system.Ctl.AdminStatus)
	sysGroup.PUT("/admin/clickout", system.Ctl.AdminClickOut)

	// ==================== API管理 ====================
	sysGroup.POST("/api/add", system.Ctl.ApiAdd)
	sysGroup.DELETE("/api/del", system.Ctl.ApiDel)
	sysGroup.PUT("/api/edit", system.Ctl.ApiEdit)
	sysGroup.GET("/api/list", system.Ctl.ApiList)
	sysGroup.GET("/api/dropdown", system.Ctl.ApiDrop)
	sysGroup.GET("/api/path", system.Ctl.ApiPath)

	// ==================== 角色管理 ====================
	sysGroup.POST("/role/add", system.Ctl.RoleAdd)
	sysGroup.DELETE("/role/del", system.Ctl.RoleDel)
	sysGroup.PUT("/role/edit", system.Ctl.RoleEdit)
	sysGroup.PUT("/role/status", system.Ctl.RoleStatus)
	sysGroup.GET("/role/info", system.Ctl.RoleInfo)
	sysGroup.GET("/role/list", system.Ctl.RoleList)
	sysGroup.GET("/role/dropdown", system.Ctl.RoleDrop)
	sysGroup.PUT("/role/is_mobile", system.Ctl.RoleIsMobile)

	// ==================== 菜单管理 ====================
	sysGroup.POST("/menu/add", system.Ctl.MenuAdd)
	sysGroup.DELETE("/menu/del", system.Ctl.MenuDel)
	sysGroup.PUT("/menu/edit", system.Ctl.MenuEdit)
	sysGroup.GET("/menu/info", system.Ctl.MenuInfo)
	sysGroup.GET("/menu/list", system.Ctl.MenuList)
	sysGroup.GET("/menu/role", system.Ctl.MenuRole)
	sysGroup.GET("/menu/dropdown", system.Ctl.MenuDrop)

	// ==================== 部门管理 ====================
	sysGroup.POST("/dept/add", system.Ctl.DeptAdd)
	sysGroup.DELETE("/dept/del", system.Ctl.DeptDel)
	sysGroup.PUT("/dept/edit", system.Ctl.DeptEdit)
	sysGroup.GET("/dept/info", system.Ctl.DeptInfo)
	sysGroup.GET("/dept/list", system.Ctl.DeptList)
	sysGroup.GET("/dept/dropdown", system.Ctl.DeptDrop)

	// ==================== 仓库管理 ====================
	warehouseGroup := group.Group("/warehouse")
	warehouseGroup.Middleware(AuthMiddleware)
	warehouseGroup.GET("/list", warehouse.Ctl.GetWarehouseList)
	warehouseGroup.POST("/", warehouse.Ctl.CreateWarehouse)
	warehouseGroup.PUT("/", warehouse.Ctl.UpdateWarehouse)
	warehouseGroup.DELETE("/:id", warehouse.Ctl.DeleteWarehouse)
	warehouseGroup.PUT("/:id/status", warehouse.Ctl.UpdateWarehouseStatus)

	// 区域
	warehouseGroup.GET("/area/list", warehouse.Ctl.GetAreaList)
	warehouseGroup.POST("/area", warehouse.Ctl.CreateArea)
	warehouseGroup.DELETE("/area/:id", warehouse.Ctl.DeleteArea)

	// 货架
	warehouseGroup.GET("/shelf/list", warehouse.Ctl.GetShelfList)
	warehouseGroup.POST("/shelf", warehouse.Ctl.CreateShelf)
	warehouseGroup.DELETE("/shelf/:id", warehouse.Ctl.DeleteShelf)

	// 盘存
	warehouseGroup.GET("/stocktake/list", warehouse.Ctl.GetStocktakeList)
	warehouseGroup.POST("/stocktake", warehouse.Ctl.CreateStocktake)

	// ==================== 库存管理 ====================
	inventoryGroup := group.Group("/inventory")
	inventoryGroup.Middleware(AuthMiddleware)

	// 整车库存
	inventoryGroup.GET("/car/list", warehouse.Ctl.GetInventoryCarList)

	// 原材料库存
	inventoryGroup.GET("/raw-material/list", warehouse.Ctl.GetRawMaterialList)
	inventoryGroup.POST("/raw-material/in", warehouse.Ctl.RawMaterialIn)
	inventoryGroup.POST("/raw-material/out", warehouse.Ctl.RawMaterialOut)

	// 危固废
	inventoryGroup.GET("/waste/list", warehouse.Ctl.GetWasteList)
	inventoryGroup.POST("/waste/in", warehouse.Ctl.WasteIn)

	// 溯源件
	inventoryGroup.GET("/part/traceable/list", warehouse.Ctl.GetPartTraceableList)
	inventoryGroup.POST("/part/traceable/out", warehouse.Ctl.PartTraceableOut)

	// 非溯源件
	inventoryGroup.GET("/part/untraceable/list", warehouse.Ctl.GetPartUntraceableList)
	inventoryGroup.POST("/part/untraceable/in", warehouse.Ctl.PartUntraceableIn)
	inventoryGroup.POST("/part/untraceable/out", warehouse.Ctl.PartUntraceableOut)

	// ==================== 操作记录 ====================
	recordGroup := group.Group("/records")
	recordGroup.Middleware(AuthMiddleware)
	recordGroup.GET("/inbound", warehouse.Ctl.GetInboundRecordList)
	recordGroup.GET("/inbound/:id", warehouse.Ctl.GetInboundRecordDetail)
	recordGroup.GET("/outbound", warehouse.Ctl.GetOutboundRecordList)
	recordGroup.GET("/outbound/:id", warehouse.Ctl.GetOutboundRecordDetail)
	recordGroup.GET("/transfer", warehouse.Ctl.GetTransferRecordList)
	recordGroup.GET("/transfer/:id", warehouse.Ctl.GetTransferRecordDetail)

	// ==================== 文件上传 ====================
	group.Group("/upload").Middleware(AuthMiddleware).POST("/", system.Ctl.UploadFile)
}


