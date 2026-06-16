package consts

const (
	// ==================== 系统管理常量 ====================
	SuperAdminId = 1  // 超级管理员id
	RoleSysId    = 1  // 默认系统角色id
	TopMenu      = 0  // 顶级菜单parent_id
	TopRole      = 0  // 顶级角色parent_id
	RoleMaxLv    = 10 // 角色最大层级
	RoleSys      = "role_sys"

	// 开关状态
	OnEnable    = 1
	OnDisabled  = 2

	// 在线状态
	IsOnline  = 1
	IsOffline = 2

	// 存在/删除
	EnumIs  = 1 // 是
	EnumNot = 2 // 否

	// 菜单类型
	MenuDir  = "M" // 路由
	MenuOwn  = "C" // 菜单
	MenuTag  = "T" // 页签
	MenuButton = "F" // 按钮

	// API类型
	ApiBus = "BUS" // 业务接口
	ApiSys = "SYS" // 系统接口
	ApiDef = "DEF" // 下拉/选项

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
