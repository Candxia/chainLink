package consts

const (
	// 用户角色
	RoleSuperAdmin  = "super_admin"
	RoleAdmin       = "admin"
	RoleInspector   = "inspector"
	RoleUser        = "user"

	// 溯源状态
	TraceStatusPending  = "pending"
	TraceStatusActive   = "active"
	TraceStatusArchived = "archived"

	// 订单状态
	OrderStatusPending   = "pending"
	OrderStatusConfirmed = "confirmed"
	OrderStatusShipped   = "shipped"
	OrderStatusDelivered = "delivered"
	OrderStatusCancelled = "cancelled"

	// 物流状态
	LogisticsStatusCreated    = "created"
	LogisticsStatusInTransit  = "in_transit"
	LogisticsStatusArrived    = "arrived"
	LogisticsStatusCompleted  = "completed"

	// 仓库操作类型
	WarehouseOpIn  = "in"
	WarehouseOpOut = "out"

	// 供应商状态
	SupplierStatusPending   = "pending"
	SupplierStatusActive    = "active"
	SupplierStatusSuspended = "suspended"

	// JWT
	JWTSecret     = "cl_system_jwt_secret_key_2024"
	JWTExpire     = 86400

	// 上下文
	CtxKeyUserId   = "userId"
	CtxKeyUserRole = "userRole"
	CtxKeyUserName = "userName"
)
