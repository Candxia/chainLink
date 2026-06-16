-- ============================================
-- 系统管理模块表结构 + 模拟数据
-- 数据库: cl_system
-- ============================================

-- ----------------------------
-- Table: sys_admin 系统管理员
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_admin` (
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `username`   varchar(32)  DEFAULT NULL COMMENT '用户名',
    `nickname`   varchar(32)  DEFAULT '' COMMENT '昵称',
    `password`   varchar(128) DEFAULT NULL COMMENT '密码',
    `role_id`    bigint(20) NOT NULL DEFAULT 0 COMMENT '角色id',
    `dept_id`    bigint(20) NOT NULL DEFAULT 0 COMMENT '部门id',
    `avatar`     varchar(128) DEFAULT NULL COMMENT '头像',
    `remark`     varchar(128) DEFAULT NULL COMMENT '备注',
    `status`     tinyint(4) DEFAULT 1 COMMENT '账号状态 1=正常 2=冻结 3=注销',
    `online`     tinyint(4) DEFAULT 2 COMMENT '在线状态 2=离线 1=在线',
    `login_at`   bigint(13) DEFAULT NULL COMMENT '最后登录时间',
    `created_by` varchar(32)  DEFAULT NULL COMMENT '创建者',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_admin_username` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统管理员';

-- ----------------------------
-- Table: sys_role 角色表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_role` (
    `id`        bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `parent_id` bigint(20) DEFAULT 0 COMMENT '父级id',
    `name`      varchar(32)  DEFAULT NULL COMMENT '角色名称',
    `code`      varchar(32)  DEFAULT NULL COMMENT '角色代码',
    `remark`    varchar(128) DEFAULT '' COMMENT '备注',
    `status`    tinyint(4) DEFAULT 1 COMMENT '状态 1=启用 2=禁用',
    `is_del`    tinyint(4) DEFAULT 2 COMMENT '删除 1=是 2=否',
    `level`     int(11) DEFAULT 1 COMMENT '层级',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(32) DEFAULT NULL COMMENT '创建者',
    `updated_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    `updated_by` varchar(32) DEFAULT NULL COMMENT '更新者',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Table: sys_menu 菜单表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_menu` (
    `menu_id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单编号',
    `menu_name`  varchar(64)  DEFAULT NULL COMMENT '菜单名称',
    `title`      varchar(64)  DEFAULT NULL COMMENT '显示名称',
    `parent_id`  bigint(20) DEFAULT 0 COMMENT '上级菜单id',
    `path`       varchar(128) DEFAULT NULL COMMENT '视图地址',
    `icon`       varchar(64)  DEFAULT NULL COMMENT '菜单图标',
    `component`  varchar(200) DEFAULT NULL COMMENT '组件路径',
    `permission` varchar(200) DEFAULT NULL COMMENT '权限标识',
    `order_num`  int(11) DEFAULT 0 COMMENT '排序',
    `menu_type`  char(1) DEFAULT 'M' COMMENT '菜单类型 M=路由 C=菜单 T=页签 F=按钮',
    `visible`    tinyint(4) DEFAULT 1 COMMENT '展示状态 1=显示 2=隐藏',
    `status`     tinyint(4) DEFAULT 1 COMMENT '菜单状态 1=正常 2=停用',
    `is_frame`   tinyint(4) DEFAULT 2 COMMENT '是否外链 1=是 2=否',
    `is_cache`   tinyint(4) DEFAULT 1 COMMENT '是否缓存 1=是 2=否',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';

-- ----------------------------
-- Table: sys_api 系统API
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_api` (
    `api_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
    `title`  varchar(128) DEFAULT NULL COMMENT '标题',
    `path`   varchar(128) DEFAULT NULL COMMENT '地址',
    `type`   varchar(8)   DEFAULT 'BUS' COMMENT '接口类型 BUS=业务 SYS=系统 DEF=自定义',
    `action` varchar(8)   DEFAULT 'GET' COMMENT '请求方式 GET/POST/PUT/DELETE',
    `logger` tinyint(4) DEFAULT 1 COMMENT '是否记录操作日志 1=是 2=否',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`api_id`) USING BTREE,
    UNIQUE KEY `idx_api_path_action` (`path`,`action`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统API';

-- ----------------------------
-- Table: sys_role_menu 角色菜单关联
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_role_menu` (
    `id`        bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `role_id`   bigint(20) NOT NULL COMMENT '角色id',
    `menu_id`   bigint(20) NOT NULL COMMENT '菜单id',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_role_menu` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';

-- ============================================
-- 模拟数据
-- ============================================

-- 1. admin: 超管
INSERT IGNORE INTO `sys_admin` (`id`,`username`,`nickname`,`password`,`role_id`,`dept_id`,`status`,`online`,`created_at`,`updated_at`)
VALUES
(1,'admin','超级管理员','e4b09aa4e68ef5f53d1c2eb3f1927a03$f8a089c4e0b3286ce665d05f54404091dc3b9455a0ff0533fc70c3f6edf91d35',1,1,1,1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(2,'zhangsan','张三','e4b09aa4e68ef5f53d1c2eb3f1927a03$f8a089c4e0b3286ce665d05f54404091dc3b9455a0ff0533fc70c3f6edf91d35',2,2,1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(3,'lisi','李四','e4b09aa4e68ef5f53d1c2eb3f1927a03$f8a089c4e0b3286ce665d05f54404091dc3b9455a0ff0533fc70c3f6edf91d35',2,2,1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(4,'wangwu','王五','e4b09aa4e68ef5f53d1c2eb3f1927a03$f8a089c4e0b3286ce665d05f54404091dc3b9455a0ff0533fc70c3f6edf91d35',3,3,1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP());

-- 2. role: 角色
INSERT IGNORE INTO `sys_role` (`id`,`parent_id`,`name`,`code`,`remark`,`status`,`is_del`,`level`,`created_at`)
VALUES
(1,0,'超级管理员','super_admin','系统最高权限角色',1,2,1,UNIX_TIMESTAMP()),
(2,1,'运营管理员','ops_admin','运营管理角色',1,2,2,UNIX_TIMESTAMP()),
(3,1,'数据查看员','data_viewer','仅查看数据角色',1,2,2,UNIX_TIMESTAMP());

-- 3. menu: 基础菜单
INSERT IGNORE INTO `sys_menu` (`menu_id`,`menu_name`,`title`,`parent_id`,`path`,`icon`,`component`,`order_num`,`menu_type`,`visible`,`status`,`created_at`)
VALUES
(1,'系统管理','系统管理',0,'/system','system','Layout',1,'C',1,1,UNIX_TIMESTAMP()),
(2,'账号管理','账号管理',1,'/system/admin','user','system/admin/index',1,'M',1,1,UNIX_TIMESTAMP()),
(3,'角色管理','角色管理',1,'/system/role','peoples','system/role/index',2,'M',1,1,UNIX_TIMESTAMP()),
(4,'菜单管理','菜单管理',1,'/system/menu','tree-table','system/menu/index',3,'M',1,1,UNIX_TIMESTAMP()),
(5,'接口管理','接口管理',1,'/system/api','api','system/api/index',4,'M',1,1,UNIX_TIMESTAMP()),
(6,'产品溯源','产品溯源',0,'/trace','connection','Layout',2,'C',1,1,UNIX_TIMESTAMP()),
(7,'产品管理','产品管理',6,'/trace/product','list','trace/product/index',1,'M',1,1,UNIX_TIMESTAMP()),
(8,'批次管理','批次管理',6,'/trace/batch','list','trace/batch/index',2,'M',1,1,UNIX_TIMESTAMP()),
(9,'溯源记录','溯源记录',6,'/trace/record','list','trace/record/index',3,'M',1,1,UNIX_TIMESTAMP()),
(10,'供应链管理','供应链管理',0,'/supply','connection','Layout',3,'C',1,1,UNIX_TIMESTAMP()),
(11,'供应商管理','供应商管理',10,'/supply/supplier','list','supply/supplier/index',1,'M',1,1,UNIX_TIMESTAMP()),
(12,'订单管理','订单管理',10,'/supply/order','list','supply/order/index',2,'M',1,1,UNIX_TIMESTAMP()),
(13,'仓库管理','仓库管理',0,'/warehouse','list','Layout',4,'C',1,1,UNIX_TIMESTAMP()),
(14,'仓库信息','仓库信息',13,'/warehouse/list','list','warehouse/index',1,'M',1,1,UNIX_TIMESTAMP()),
(15,'库存管理','库存管理',13,'/inventory/list','list','warehouse/inventory',2,'M',1,1,UNIX_TIMESTAMP()),
(16,'区块链管理','区块链管理',0,'/blockchain','link','Layout',5,'C',1,1,UNIX_TIMESTAMP()),
(17,'区块查询','区块查询',16,'/blockchain/blocks','list','blockchain/blocks',1,'M',1,1,UNIX_TIMESTAMP()),
(18,'合约管理','合约管理',16,'/blockchain/contract','list','blockchain/contract',2,'M',1,1,UNIX_TIMESTAMP());

-- 4. sys_api: 系统API
INSERT IGNORE INTO `sys_api` (`api_id`,`title`,`path`,`type`,`action`,`logger`,`created_at`)
VALUES
(1,'管理员列表','/api/system/admin/list','BUS','GET',1,UNIX_TIMESTAMP()),
(2,'管理员详情','/api/system/admin/info','BUS','GET',2,UNIX_TIMESTAMP()),
(3,'添加管理员','/api/system/admin/add','BUS','POST',1,UNIX_TIMESTAMP()),
(4,'编辑管理员','/api/system/admin/edit','BUS','PUT',1,UNIX_TIMESTAMP()),
(5,'删除管理员','/api/system/admin/del','BUS','DELETE',1,UNIX_TIMESTAMP()),
(6,'修改密码','/api/system/admin/password','BUS','PUT',1,UNIX_TIMESTAMP()),
(7,'修改状态','/api/system/admin/status','BUS','PUT',1,UNIX_TIMESTAMP()),
(8,'踢下线','/api/system/admin/clickout','BUS','PUT',1,UNIX_TIMESTAMP()),
(9,'接口列表','/api/system/api/list','SYS','GET',2,UNIX_TIMESTAMP()),
(10,'添加接口','/api/system/api/add','SYS','POST',1,UNIX_TIMESTAMP()),
(11,'编辑接口','/api/system/api/edit','SYS','PUT',1,UNIX_TIMESTAMP()),
(12,'删除接口','/api/system/api/del','SYS','DELETE',1,UNIX_TIMESTAMP()),
(13,'接口下拉','/api/system/api/dropdown','DEF','GET',2,UNIX_TIMESTAMP()),
(14,'接口路径下拉','/api/system/api/path','DEF','GET',2,UNIX_TIMESTAMP()),
(15,'角色列表','/api/system/role/list','BUS','GET',2,UNIX_TIMESTAMP()),
(16,'角色详情','/api/system/role/info','BUS','GET',2,UNIX_TIMESTAMP()),
(17,'添加角色','/api/system/role/add','BUS','POST',1,UNIX_TIMESTAMP()),
(18,'编辑角色','/api/system/role/edit','BUS','PUT',1,UNIX_TIMESTAMP()),
(19,'删除角色','/api/system/role/del','BUS','DELETE',1,UNIX_TIMESTAMP()),
(20,'修改角色状态','/api/system/role/status','BUS','PUT',1,UNIX_TIMESTAMP()),
(21,'角色下拉','/api/system/role/dropdown','DEF','GET',2,UNIX_TIMESTAMP()),
(22,'角色手机号','/api/system/role/is_mobile','BUS','PUT',1,UNIX_TIMESTAMP()),
(23,'菜单列表','/api/system/menu/list','SYS','GET',2,UNIX_TIMESTAMP()),
(24,'菜单详情','/api/system/menu/info','SYS','GET',2,UNIX_TIMESTAMP()),
(25,'添加菜单','/api/system/menu/add','SYS','POST',1,UNIX_TIMESTAMP()),
(26,'编辑菜单','/api/system/menu/edit','SYS','PUT',1,UNIX_TIMESTAMP()),
(27,'删除菜单','/api/system/menu/del','SYS','DELETE',1,UNIX_TIMESTAMP()),
(28,'角色菜单(权限)','/api/system/menu/role','DEF','GET',2,UNIX_TIMESTAMP()),
(29,'菜单下拉','/api/system/menu/dropdown','DEF','GET',2,UNIX_TIMESTAMP());

-- 5. sys_role_menu: 角色菜单关联
INSERT IGNORE INTO `sys_role_menu` (`role_id`,`menu_id`,`created_at`) VALUES
(1,1,UNIX_TIMESTAMP()),(1,2,UNIX_TIMESTAMP()),(1,3,UNIX_TIMESTAMP()),(1,4,UNIX_TIMESTAMP()),(1,5,UNIX_TIMESTAMP()),
(1,6,UNIX_TIMESTAMP()),(1,7,UNIX_TIMESTAMP()),(1,8,UNIX_TIMESTAMP()),(1,9,UNIX_TIMESTAMP()),
(1,10,UNIX_TIMESTAMP()),(1,11,UNIX_TIMESTAMP()),(1,12,UNIX_TIMESTAMP()),
(1,13,UNIX_TIMESTAMP()),(1,14,UNIX_TIMESTAMP()),(1,15,UNIX_TIMESTAMP()),
(1,16,UNIX_TIMESTAMP()),(1,17,UNIX_TIMESTAMP()),(1,18,UNIX_TIMESTAMP()),
(2,1,UNIX_TIMESTAMP()),(2,2,UNIX_TIMESTAMP()),(2,3,UNIX_TIMESTAMP()),
(2,6,UNIX_TIMESTAMP()),(2,7,UNIX_TIMESTAMP()),(2,8,UNIX_TIMESTAMP()),(2,9,UNIX_TIMESTAMP()),
(2,10,UNIX_TIMESTAMP()),(2,11,UNIX_TIMESTAMP()),(2,12,UNIX_TIMESTAMP()),
(3,6,UNIX_TIMESTAMP()),(3,7,UNIX_TIMESTAMP()),(3,9,UNIX_TIMESTAMP());

-- ============================================
-- Table: sys_dept 部门表
-- ============================================
CREATE TABLE IF NOT EXISTS `sys_dept` (
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `parent_id`  bigint(20) DEFAULT 0 COMMENT '上级部门id',
    `name`       varchar(64) DEFAULT NULL COMMENT '部门名称',
    `order_num`  int(11) DEFAULT 0 COMMENT '排序',
    `leader`     varchar(32) DEFAULT NULL COMMENT '负责人',
    `phone`      varchar(32) DEFAULT NULL COMMENT '联系电话',
    `email`      varchar(64) DEFAULT NULL COMMENT '邮箱',
    `status`     tinyint(4) DEFAULT 1 COMMENT '状态 1=正常 2=停用',
    `is_del`     tinyint(4) DEFAULT 2 COMMENT '删除 1=是 2=否',
    `created_at` bigint(13) DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';

-- 部门模拟数据
INSERT IGNORE INTO `sys_dept` (`id`,`parent_id`,`name`,`order_num`,`leader`,`phone`,`email`,`status`,`is_del`,`created_at`,`updated_at`)
VALUES
(1,0,'总公司',1,'张总','13800000000','ceo@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(2,1,'技术部',1,'李总','13800000001','cto@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(3,1,'市场部',2,'王总','13800000002','cmo@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(4,2,'前端组',1,'赵工','13800000010','fe@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(5,2,'后端组',2,'钱工','13800000011','be@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(6,3,'推广组',1,'孙经理','13800000020','market@company.com',1,2,UNIX_TIMESTAMP(),UNIX_TIMESTAMP());
