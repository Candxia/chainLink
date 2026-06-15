-- 访问日志表 - 记录每个API请求的完整信息
USE cl_system;

CREATE TABLE IF NOT EXISTS access_log (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    trace_id VARCHAR(64) DEFAULT '' COMMENT '链路追踪ID',
    operator VARCHAR(64) DEFAULT '' COMMENT '操作人(用户名)',
    user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '操作人ID',
    action_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    request_method VARCHAR(16) DEFAULT '' COMMENT '请求方法(GET/POST/PUT/DELETE)',
    request_path VARCHAR(256) DEFAULT '' COMMENT '请求路径',
    request_params TEXT COMMENT '请求参数(query+body)',
    api_name VARCHAR(256) DEFAULT '' COMMENT '接口名称(路由描述)',
    duration_ms INT DEFAULT 0 COMMENT '耗时(毫秒)',
    ip_address VARCHAR(64) DEFAULT '' COMMENT '客户端IP地址',
    user_agent VARCHAR(512) DEFAULT '' COMMENT '客户端User-Agent',
    auth_token VARCHAR(256) DEFAULT '' COMMENT '客户端Token(JWT)',
    response_code INT DEFAULT 0 COMMENT 'HTTP响应状态码',
    response_body TEXT COMMENT '返回详情(响应内容)',
    status TINYINT DEFAULT 1 COMMENT '状态: 1成功 0失败',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_operator (operator),
    INDEX idx_action_time (action_time),
    INDEX idx_request_path (request_path(64)),
    INDEX idx_duration (duration_ms),
    INDEX idx_ip (ip_address(32))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='访问日志表';
