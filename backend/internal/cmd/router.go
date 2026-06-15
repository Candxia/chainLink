package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/controller/user"
	"cl_system/internal/controller/trace"
	"cl_system/internal/controller/supply"
	"cl_system/internal/controller/blockchain"
	"cl_system/internal/controller/system"
)

func Router(group *ghttp.RouterGroup) {
	// ==================== 用户与权限 ====================
	userGroup := group.Group("/api/user")
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
	traceGroup := group.Group("/api/trace")
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
	supplyGroup := group.Group("/api/supply")
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
	blockGroup := group.Group("/api/blockchain")
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
	sysGroup := group.Group("/api/system")
	sysGroup.Middleware(AuthMiddleware)
	sysGroup.GET("/dashboard", system.Ctl.GetDashboard)
	sysGroup.GET("/config", system.Ctl.GetConfig)
	sysGroup.PUT("/config", system.Ctl.UpdateConfig)
	sysGroup.GET("/logs", system.Ctl.GetLogList)
	sysGroup.DELETE("/logs/:id", system.Ctl.DeleteLog)
	sysGroup.GET("/notifications", system.Ctl.GetNotificationList)
	sysGroup.POST("/notification/read/:id", system.Ctl.MarkNotificationRead)

	// ==================== 文件上传 ====================
	group.Group("/api/upload").Middleware(AuthMiddleware).POST("/", system.Ctl.UploadFile)
}

func AuthMiddleware(r *ghttp.Request) {
	token := r.GetHeader("Authorization")
	if token == "" {
		r.Response.WriteJson(g.Map{"code": 401, "message": "未登录"})
		r.Exit()
	}
	r.Middleware.Next()
}
