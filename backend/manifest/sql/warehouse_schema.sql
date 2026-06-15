-- 车辆回收系统 - 仓库管理模块
-- 表结构 + 模拟数据

USE cl_system;

-- ==================== 仓库信息 ====================
CREATE TABLE IF NOT EXISTS warehouse (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(64) NOT NULL UNIQUE COMMENT '仓库编码',
    name VARCHAR(128) NOT NULL COMMENT '仓库名称',
    address VARCHAR(256) DEFAULT '' COMMENT '仓库地址',
    company VARCHAR(128) DEFAULT '' COMMENT '所属公司',
    status TINYINT DEFAULT 1 COMMENT '状态: 1启用 0停用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='仓库表';

-- 仓库区域
CREATE TABLE IF NOT EXISTS warehouse_area (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '所属仓库ID',
    name VARCHAR(64) NOT NULL COMMENT '区域名称',
    description VARCHAR(256) DEFAULT '' COMMENT '区域说明',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='仓库区域表';

-- 仓库货架
CREATE TABLE IF NOT EXISTS warehouse_shelf (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    area_id BIGINT UNSIGNED NOT NULL COMMENT '所属区域ID',
    name VARCHAR(64) NOT NULL COMMENT '货架名称',
    description VARCHAR(256) DEFAULT '' COMMENT '货架说明',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_area (area_id),
    FOREIGN KEY (area_id) REFERENCES warehouse_area(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='货架表';

-- 仓库盘存
CREATE TABLE IF NOT EXISTS stocktake (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    stocktake_no VARCHAR(64) NOT NULL UNIQUE COMMENT '盘存编号',
    type TINYINT DEFAULT 0 COMMENT '类型: 0定期盘点 1不定期盘点',
    status TINYINT DEFAULT 0 COMMENT '状态: 0待盘点 1无差异 2有差异',
    operator VARCHAR(64) DEFAULT '' COMMENT '盘点人员',
    stocktake_time DATETIME DEFAULT NULL COMMENT '盘点时间',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_status (status),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='盘存表';

-- ==================== 库存管理 ====================

-- 整车库存
CREATE TABLE IF NOT EXISTS inventory_car (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    vin VARCHAR(64) NOT NULL COMMENT 'VIN码',
    brand VARCHAR(64) DEFAULT '' COMMENT '品牌',
    model VARCHAR(64) DEFAULT '' COMMENT '车型',
    plate_no VARCHAR(32) DEFAULT '' COMMENT '车牌号',
    color VARCHAR(32) DEFAULT '' COMMENT '颜色',
    year INT DEFAULT 0 COMMENT '出厂年份',
    status TINYINT DEFAULT 1 COMMENT '状态: 1在库 2已拆解 3已出库',
    shelf_id BIGINT UNSIGNED DEFAULT NULL COMMENT '货架位置',
    entry_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '入库时间',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_vin (vin),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='整车库存表';

-- 原材料库存
CREATE TABLE IF NOT EXISTS inventory_raw_material (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    material_no VARCHAR(64) NOT NULL UNIQUE COMMENT '原材料编号',
    material_type ENUM('steel','copper','aluminum','plastic','rubber','glass','other') NOT NULL COMMENT '原材料类型',
    quantity DECIMAL(12,2) DEFAULT 0 COMMENT '库存数量(kg)',
    shelf_id BIGINT UNSIGNED DEFAULT NULL COMMENT '货架位置',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (material_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='原材料库存表';

-- 危固废管理
CREATE TABLE IF NOT EXISTS inventory_waste (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    waste_no VARCHAR(64) NOT NULL UNIQUE COMMENT '废料编号',
    waste_type ENUM('hazardous','solid') NOT NULL COMMENT '废料类型',
    quantity DECIMAL(12,2) DEFAULT 0 COMMENT '数量(kg)',
    shelf_id BIGINT UNSIGNED DEFAULT NULL COMMENT '货架位置',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (waste_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='危固废库存表';

-- 配件-溯源件
CREATE TABLE IF NOT EXISTS inventory_part_traceable (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    part_name VARCHAR(128) NOT NULL COMMENT '配件名称',
    part_type VARCHAR(64) DEFAULT '' COMMENT '配件类型',
    car_model VARCHAR(64) DEFAULT '' COMMENT '车型',
    car_series VARCHAR(64) DEFAULT '' COMMENT '车系',
    description TEXT COMMENT '配件描述',
    vin VARCHAR(64) DEFAULT '' COMMENT '来源车辆VIN',
    shelf_id BIGINT UNSIGNED DEFAULT NULL COMMENT '货架位置',
    quantity INT DEFAULT 0 COMMENT '库存数量',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (part_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='配件溯源件表';

-- 配件-非溯源件
CREATE TABLE IF NOT EXISTS inventory_part_untraceable (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    part_name VARCHAR(128) NOT NULL COMMENT '配件名称',
    part_type VARCHAR(64) DEFAULT '' COMMENT '配件类型',
    quantity INT DEFAULT 0 COMMENT '库存数量',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (part_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='配件非溯源件表';

-- ==================== 操作记录 ====================

-- 入库记录
CREATE TABLE IF NOT EXISTS inbound_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    record_no VARCHAR(64) NOT NULL UNIQUE COMMENT '入库单号',
    inbound_type ENUM('car','raw_material','waste','part_traceable','part_untraceable') NOT NULL COMMENT '入库类型',
    inbound_method ENUM('dismantle','transfer','purchase','other') NOT NULL COMMENT '入库方式',
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '目标仓库',
    operator VARCHAR(64) DEFAULT '' COMMENT '操作人',
    total_quantity INT DEFAULT 0 COMMENT '入库总数目',
    status TINYINT DEFAULT 0 COMMENT '审核状态: 0待审核 1已通过 2未通过',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    inbound_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '入库时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (inbound_type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='入库记录表';

-- 出库记录
CREATE TABLE IF NOT EXISTS outbound_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    record_no VARCHAR(64) NOT NULL UNIQUE COMMENT '出库单号',
    outbound_type ENUM('car','raw_material','waste','part_traceable','part_untraceable') NOT NULL COMMENT '出库类型',
    outbound_method ENUM('sale','transfer','scrap','other') NOT NULL COMMENT '出库方式',
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '出库仓库',
    operator VARCHAR(64) DEFAULT '' COMMENT '操作人',
    total_quantity INT DEFAULT 0 COMMENT '出库总数量',
    status TINYINT DEFAULT 0 COMMENT '审核状态: 0待审核 1已通过 2未通过',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    outbound_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '出库时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_type (outbound_type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='出库记录表';

-- 调拨记录
CREATE TABLE IF NOT EXISTS transfer_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    record_no VARCHAR(64) NOT NULL UNIQUE COMMENT '调拨单号',
    transfer_type ENUM('car','raw_material','waste','part_traceable','part_untraceable') NOT NULL COMMENT '调拨类型',
    from_warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '调出仓库',
    to_warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '调入仓库',
    item_type_desc VARCHAR(128) DEFAULT '' COMMENT '商品类型描述',
    total_categories INT DEFAULT 0 COMMENT '调拨总品类',
    total_quantity INT DEFAULT 0 COMMENT '调拨总数量',
    status TINYINT DEFAULT 0 COMMENT '审核状态: 0待审核 1已通过 2未通过',
    operator VARCHAR(64) DEFAULT '' COMMENT '操作人',
    remark VARCHAR(256) DEFAULT '' COMMENT '备注',
    transfer_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '调拨时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_from_warehouse (from_warehouse_id),
    INDEX idx_to_warehouse (to_warehouse_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='调拨记录表';

-- ==================== 模拟数据 ====================

-- 仓库
INSERT INTO warehouse (id, code, name, address, company, status) VALUES
(1, 'WH-2026001', '双流拆解厂一号仓库', '成都市双流区九江街道拆解产业园A区', '成都双流车辆拆解有限公司', 1),
(2, 'WH-2026002', '龙泉驿拆解厂仓库', '成都市龙泉驿区汽车拆解工业园B区', '成都龙泉车辆回收有限公司', 1),
(3, 'WH-2026003', '高新西区配件仓库', '成都市高新区西区配件仓储中心3栋', '成都高新供应链管理有限公司', 1);

-- 区域
INSERT INTO warehouse_area (id, warehouse_id, name, description) VALUES
(1, 1, 'A区', '整车停放区'),
(2, 1, 'B区', '拆解作业区'),
(3, 1, 'C区', '原材料存放区'),
(4, 1, 'D区', '危固废暂存区'),
(5, 1, 'AK区', '配件存放区'),
(6, 2, 'A区', '整车停放区'),
(7, 2, 'B区', '原材料存放区');

-- 货架
INSERT INTO warehouse_shelf (id, area_id, name, description) VALUES
(1, 1, 'A-1', '整车存放位1-20'),
(2, 1, 'A-2', '整车存放位21-40'),
(3, 1, 'A-3', '整车存放位41-60'),
(4, 3, 'C-1', '钢铁料区'),
(5, 3, 'C-2', '铜铝料区'),
(6, 3, 'C-3', '塑料橡胶区'),
(7, 4, 'D-1', '废油液存放区'),
(8, 4, 'D-2', '废电池存放区'),
(9, 5, 'AK-1', '发动机存放架'),
(10, 5, 'AK-2', '变速箱存放架'),
(11, 5, 'AK-3', '轮胎轮毂架');

-- 盘存记录
INSERT INTO stocktake (id, warehouse_id, stocktake_no, type, status, operator, stocktake_time, remark) VALUES
(1, 1, 'ST-20260601', 0, 1, '张三', '2026-06-01 10:00:00', '月度定期盘点，无差异'),
(2, 1, 'ST-20260610', 1, 2, '李四', '2026-06-10 14:30:00', '随机抽检，发现2件差异'),
(3, 2, 'ST-20260605', 0, 1, '王五', '2026-06-05 09:00:00', '月度盘点，无差异'),
(4, 3, 'ST-20260608', 0, 0, '', NULL, '待盘点');

-- 整车库存 (10台报废车辆)
INSERT INTO inventory_car (id, warehouse_id, vin, brand, model, plate_no, color, year, status, shelf_id, entry_date) VALUES
(1, 1, 'LSVBCN4637C000001', '大众', '帕萨特', '川A·00001', '黑色', 2015, 1, 1, '2026-05-10'),
(2, 1, 'LSVJF4657N000002', '大众', '朗逸', '川A·00002', '白色', 2018, 1, 1, '2026-05-12'),
(3, 1, 'LVGBE40K8AG000003', '丰田', '凯美瑞', '川A·00003', '银色', 2016, 1, 2, '2026-05-15'),
(4, 1, 'LHGCP2615A0000004', '本田', '雅阁', '川A·00004', '黑色', 2017, 1, 2, '2026-05-18'),
(5, 1, 'LVSFDFME0RF000005', '福特', '福克斯', '川A·00005', '红色', 2019, 2, 3, '2026-05-20'),
(6, 2, 'LSGGF53F5BH000006', '别克', '君威', '川A·00006', '蓝色', 2014, 1, 6, '2026-05-22'),
(7, 2, 'LFV2A11K163000007', '奥迪', 'A6L', '川A·00007', '黑色', 2020, 1, 6, '2026-05-25'),
(8, 2, 'WBAFB11060L000008', '宝马', '5系', '川A·00008', '银色', 2019, 1, 7, '2026-05-28'),
(9, 1, 'LE4FG8BB6FL000009', '奔驰', 'C级', '川A·00009', '白色', 2021, 1, 3, '2026-06-01'),
(10, 1, 'LJ8F2C5D7GB000010', '比亚迪', '汉', '川A·00010', '灰色', 2022, 1, 3, '2026-06-05');

-- 原材料库存
INSERT INTO inventory_raw_material (id, warehouse_id, material_no, material_type, quantity, shelf_id) VALUES
(1, 1, 'RM-20260001', 'steel', 8560.50, 4),
(2, 1, 'RM-20260002', 'steel', 4230.00, 4),
(3, 1, 'RM-20260003', 'copper', 185.30, 5),
(4, 1, 'RM-20260004', 'aluminum', 1230.00, 5),
(5, 1, 'RM-20260005', 'plastic', 670.80, 6),
(6, 1, 'RM-20260006', 'rubber', 890.20, 6),
(7, 2, 'RM-20260007', 'steel', 3200.00, 6),
(8, 2, 'RM-20260008', 'copper', 95.60, 6),
(9, 2, 'RM-20260009', 'aluminum', 560.00, 6),
(10, 1, 'RM-20260010', 'glass', 340.00, 6);

-- 危固废
INSERT INTO inventory_waste (id, warehouse_id, waste_no, waste_type, quantity, shelf_id) VALUES
(1, 1, 'HW-20260001', 'hazardous', 320.00, 7),
(2, 1, 'HW-20260002', 'hazardous', 180.50, 7),
(3, 1, 'SW-20260001', 'solid', 1250.00, 8),
(4, 1, 'SW-20260002', 'solid', 890.00, 8),
(5, 2, 'HW-20260003', 'hazardous', 150.00, 7);

-- 配件-溯源件
INSERT INTO inventory_part_traceable (id, warehouse_id, part_name, part_type, car_model, car_series, description, vin, shelf_id, quantity) VALUES
(1, 1, 'EA888发动机总成', '发动机', '帕萨特', 'Passat B8', '2.0T EA888三代发动机', 'LSVBCN4637C000001', 9, 1),
(2, 1, 'DQ380变速箱', '变速箱', '帕萨特', 'Passat B8', '7速湿式双离合变速箱总成', 'LSVBCN4637C000001', 10, 1),
(3, 1, '铁轮毂', '轮毂', '朗逸', 'Lavida', '15寸铁质轮毂（含轮胎）', 'LSVJF4657N000002', 11, 4),
(4, 1, '启动机总成', '电器', '凯美瑞', 'Camry', '12V启动电机', 'LVGBE40K8AG000003', 9, 1),
(5, 1, '发电机总成', '电器', '雅阁', 'Accord', '14V 120A发电机', 'LHGCP2615A0000004', 9, 1),
(6, 1, '前排座椅总成', '内饰', '福克斯', 'Focus', '黑色真皮电动座椅', 'LVSFDFME0RF000005', 10, 2);

-- 配件-非溯源件
INSERT INTO inventory_part_untraceable (id, warehouse_id, part_name, part_type, quantity) VALUES
(1, 1, '火花塞', '电器', 120),
(2, 1, '机油滤清器', '滤清器', 85),
(3, 1, '空气滤芯', '滤清器', 60),
(4, 1, '刹车片（前）', '制动', 200),
(5, 1, '刹车片（后）', '制动', 180),
(6, 1, '雨刮片', '车身', 150),
(7, 1, '灯泡（H7）', '电器', 300),
(8, 1, '保险丝套装', '电器', 500),
(9, 1, '车门把手', '车身', 45),
(10, 3, '轮胎螺丝', '底盘', 1000),
(11, 3, '防冻液', '养护', 200),
(12, 3, '玻璃水', '养护', 500);

-- 入库记录
INSERT INTO inbound_record (id, record_no, inbound_type, inbound_method, warehouse_id, operator, total_quantity, status, inbound_time) VALUES
(1, 'IN-20260510001', 'car', 'purchase', 1, '张三', 5, 1, '2026-05-10'),
(2, 'IN-20260520001', 'raw_material', 'dismantle', 1, '李四', 6, 1, '2026-05-20'),
(3, 'IN-20260601001', 'car', 'purchase', 1, '张三', 3, 1, '2026-06-01'),
(4, 'IN-20260602001', 'part_traceable', 'dismantle', 1, '李四', 6, 1, '2026-06-02'),
(5, 'IN-20260603001', 'raw_material', 'dismantle', 1, '王五', 4, 1, '2026-06-03'),
(6, 'IN-20260605001', 'car', 'purchase', 2, '赵六', 3, 1, '2026-06-05'),
(7, 'IN-20260607001', 'part_untraceable', 'purchase', 3, '钱七', 12, 0, '2026-06-07');

-- 出库记录
INSERT INTO outbound_record (id, record_no, outbound_type, outbound_method, warehouse_id, operator, total_quantity, status, outbound_time) VALUES
(1, 'OUT-20260525001', 'raw_material', 'sale', 1, '李四', 2000, 1, '2026-05-25'),
(2, 'OUT-20260601001', 'part_untraceable', 'sale', 1, '张三', 50, 1, '2026-06-01'),
(3, 'OUT-20260608001', 'raw_material', 'sale', 2, '赵六', 1000, 1, '2026-06-08'),
(4, 'OUT-20260610001', 'car', 'scrap', 1, '王五', 1, 2, '2026-06-10');

-- 调拨记录
INSERT INTO transfer_record (id, record_no, transfer_type, from_warehouse_id, to_warehouse_id, item_type_desc, total_categories, total_quantity, status, operator, transfer_time) VALUES
(1, 'TF-20260601001', 'raw_material', 1, 2, '钢铁、铝材', 2, 1500, 1, '李四', '2026-06-01'),
(2, 'TF-20260605001', 'part_untraceable', 1, 3, '轮胎螺丝、火花塞', 2, 300, 0, '张三', '2026-06-05'),
(3, 'TF-20260608001', 'waste', 1, 2, '危废', 1, 100, 1, '王五', '2026-06-08');
