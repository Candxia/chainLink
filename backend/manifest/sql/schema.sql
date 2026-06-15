-- 链环系统数据库初始化脚本
CREATE DATABASE IF NOT EXISTS cl_system DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE cl_system;

-- ==================== 用户与权限 ====================
CREATE TABLE IF NOT EXISTS sys_user (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE COMMENT '用户名',
    password VARCHAR(128) NOT NULL COMMENT '密码',
    email VARCHAR(128) DEFAULT '' COMMENT '邮箱',
    phone VARCHAR(32) DEFAULT '' COMMENT '手机号',
    role VARCHAR(32) DEFAULT 'user' COMMENT '角色',
    avatar VARCHAR(256) DEFAULT '' COMMENT '头像',
    status TINYINT DEFAULT 1 COMMENT '状态: 1启用 0禁用',
    enterprise_id BIGINT UNSIGNED DEFAULT 0 COMMENT '所属企业ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_role (role),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 默认管理员账号: admin / admin123
INSERT INTO sys_user (username, password, email, role, status) VALUES
('admin', MD5(CONCAT('admin123', 'cl_system_salt')), 'admin@clsystem.com', 'super_admin', 1),
('demo', MD5(CONCAT('demo123', 'cl_system_salt')), 'demo@clsystem.com', 'user', 1);

CREATE TABLE IF NOT EXISTS sys_role (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL COMMENT '角色名称',
    code VARCHAR(64) NOT NULL UNIQUE COMMENT '角色编码',
    permissions TEXT COMMENT '权限JSON',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

INSERT INTO sys_role (name, code, permissions) VALUES
('超级管理员', 'super_admin', '["*"]'),
('企业管理员', 'admin', '["trace:create","trace:read","supply:manage","user:read"]'),
('质检员', 'inspector', '["trace:create","trace:read"]'),
('普通用户', 'user', '["trace:read"]');

CREATE TABLE IF NOT EXISTS enterprise (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL COMMENT '企业名称',
    code VARCHAR(64) UNIQUE COMMENT '企业编码',
    contact VARCHAR(64) DEFAULT '' COMMENT '联系人',
    phone VARCHAR(32) DEFAULT '' COMMENT '联系电话',
    address VARCHAR(512) DEFAULT '' COMMENT '地址',
    logo VARCHAR(256) DEFAULT '' COMMENT '企业Logo',
    status TINYINT DEFAULT 1 COMMENT '状态',
    user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '管理员用户ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='企业表';

-- ==================== 产品溯源 ====================
CREATE TABLE IF NOT EXISTS product (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL COMMENT '产品名称',
    category VARCHAR(128) DEFAULT '' COMMENT '产品分类',
    spec VARCHAR(256) DEFAULT '' COMMENT '规格型号',
    unit VARCHAR(32) DEFAULT '个' COMMENT '单位',
    description TEXT COMMENT '产品描述',
    image VARCHAR(512) DEFAULT '' COMMENT '产品图片',
    enterprise_id BIGINT UNSIGNED DEFAULT 0 COMMENT '所属企业',
    status VARCHAR(32) DEFAULT 'active' COMMENT '状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_name (name),
    INDEX idx_category (category),
    INDEX idx_enterprise (enterprise_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品表';

CREATE TABLE IF NOT EXISTS batch (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    batch_no VARCHAR(128) NOT NULL COMMENT '批次号',
    product_id BIGINT UNSIGNED NOT NULL COMMENT '产品ID',
    quantity INT DEFAULT 0 COMMENT '数量',
    produce_date VARCHAR(32) DEFAULT '' COMMENT '生产日期',
    expire_date VARCHAR(32) DEFAULT '' COMMENT '到期日期',
    status VARCHAR(32) DEFAULT 'active' COMMENT '状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_batch_no (batch_no),
    INDEX idx_product (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='批次表';

CREATE TABLE IF NOT EXISTS trace_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED NOT NULL COMMENT '产品ID',
    batch_id BIGINT UNSIGNED DEFAULT 0 COMMENT '批次ID',
    record_type VARCHAR(64) NOT NULL COMMENT '记录类型(生产/加工/质检/物流)',
    content TEXT COMMENT '记录内容',
    operator VARCHAR(128) DEFAULT '' COMMENT '操作人',
    location VARCHAR(256) DEFAULT '' COMMENT '操作地点',
    tx_hash VARCHAR(256) DEFAULT '' COMMENT '交易哈希',
    block_number BIGINT UNSIGNED DEFAULT 0 COMMENT '区块高度',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_product (product_id),
    INDEX idx_batch (batch_id),
    INDEX idx_tx_hash (tx_hash),
    INDEX idx_record_type (record_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='溯源记录表';

-- ==================== 供应链管理 ====================
CREATE TABLE IF NOT EXISTS supplier (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL COMMENT '供应商名称',
    code VARCHAR(64) UNIQUE COMMENT '供应商编码',
    contact VARCHAR(64) DEFAULT '' COMMENT '联系人',
    phone VARCHAR(32) DEFAULT '' COMMENT '联系电话',
    email VARCHAR(128) DEFAULT '' COMMENT '邮箱',
    address VARCHAR(512) DEFAULT '' COMMENT '地址',
    category VARCHAR(128) DEFAULT '' COMMENT '供应类别',
    qual_level VARCHAR(32) DEFAULT 'B' COMMENT '资质等级(A/B/C)',
    status VARCHAR(32) DEFAULT 'pending' COMMENT '状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_name (name),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='供应商表';

CREATE TABLE IF NOT EXISTS order_info (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_no VARCHAR(128) NOT NULL COMMENT '订单编号',
    product_id BIGINT UNSIGNED DEFAULT 0 COMMENT '产品ID',
    batch_id BIGINT UNSIGNED DEFAULT 0 COMMENT '批次ID',
    supplier_id BIGINT UNSIGNED DEFAULT 0 COMMENT '供应商ID',
    quantity INT DEFAULT 0 COMMENT '数量',
    total_amount DECIMAL(15,2) DEFAULT 0 COMMENT '总金额',
    status VARCHAR(32) DEFAULT 'pending' COMMENT '状态',
    remark TEXT COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_order_no (order_no),
    INDEX idx_supplier (supplier_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE IF NOT EXISTS warehouse_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT UNSIGNED DEFAULT 0 COMMENT '关联订单',
    product_id BIGINT UNSIGNED DEFAULT 0 COMMENT '产品ID',
    batch_id BIGINT UNSIGNED DEFAULT 0 COMMENT '批次ID',
    op_type VARCHAR(16) NOT NULL COMMENT '操作类型(in/out)',
    quantity INT DEFAULT 0 COMMENT '数量',
    operator VARCHAR(128) DEFAULT '' COMMENT '操作人',
    remark VARCHAR(512) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_order (order_id),
    INDEX idx_product (product_id),
    INDEX idx_op_type (op_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='仓储记录表';

CREATE TABLE IF NOT EXISTS logistics_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT UNSIGNED NOT NULL COMMENT '订单ID',
    node_name VARCHAR(128) NOT NULL COMMENT '节点名称',
    location VARCHAR(256) DEFAULT '' COMMENT '位置',
    status VARCHAR(32) DEFAULT 'created' COMMENT '物流状态',
    operator VARCHAR(128) DEFAULT '' COMMENT '操作人',
    remark VARCHAR(512) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物流记录表';

-- ==================== 区块链管理 ====================
CREATE TABLE IF NOT EXISTS contract (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(256) NOT NULL COMMENT '合约地址',
    name VARCHAR(128) DEFAULT '' COMMENT '合约名称',
    abi LONGTEXT COMMENT 'ABI定义',
    deploy_tx VARCHAR(256) DEFAULT '' COMMENT '部署交易哈希',
    deploy_time VARCHAR(64) DEFAULT '' COMMENT '部署时间',
    owner VARCHAR(256) DEFAULT '' COMMENT '合约所有者',
    status VARCHAR(32) DEFAULT 'active' COMMENT '状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE INDEX idx_address (address)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='智能合约表';

-- ==================== 系统管理 ====================
CREATE TABLE IF NOT EXISTS sys_config (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `key` VARCHAR(128) NOT NULL UNIQUE COMMENT '配置键',
    `value` TEXT COMMENT '配置值',
    `desc` VARCHAR(512) DEFAULT '' COMMENT '说明',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统配置表';

INSERT INTO sys_config (key, value, `desc`) VALUES
('system_name', '链环系统', '系统名称'),
('system_logo', '', '系统Logo'),
('blockchain_enabled', 'true', '是否启用区块链'),
('trace_required', 'true', '溯源是否必填'),
('notification_interval', '3600', '通知检查间隔(秒)');

CREATE TABLE IF NOT EXISTS sys_log (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '用户ID',
    username VARCHAR(64) DEFAULT '' COMMENT '用户名',
    type VARCHAR(64) DEFAULT '' COMMENT '日志类型',
    action VARCHAR(256) DEFAULT '' COMMENT '操作内容',
    ip VARCHAR(64) DEFAULT '' COMMENT 'IP地址',
    user_agent VARCHAR(512) DEFAULT '' COMMENT 'UserAgent',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_type (type),
    INDEX idx_user (user_id),
    INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';

CREATE TABLE IF NOT EXISTS notification (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '用户ID',
    title VARCHAR(256) NOT NULL COMMENT '标题',
    content TEXT COMMENT '内容',
    type VARCHAR(32) DEFAULT 'system' COMMENT '类型',
    is_read TINYINT DEFAULT 0 COMMENT '是否已读',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user (user_id),
    INDEX idx_is_read (is_read)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知表';
