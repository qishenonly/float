-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    type ENUM('expense','income','transfer') NOT NULL COMMENT '交易类型: expense-支出, income-收入, transfer-转账',
    category_id BIGINT COMMENT '分类ID',
    account_id BIGINT COMMENT '账户ID',
    to_account_id BIGINT COMMENT '转入账户ID (仅转账时使用)',
    
    amount DECIMAL(15, 2) NOT NULL COMMENT '金额',
    currency VARCHAR(10) DEFAULT 'CNY' COMMENT '货币单位',
    
    title VARCHAR(200) COMMENT '标题/商家名称',
    description TEXT COMMENT '备注说明',
    location VARCHAR(200) COMMENT '地点',
    
    transaction_date DATE NOT NULL COMMENT '交易日期',
    transaction_time TIME COMMENT '交易时间',
    
    -- 关联信息
    bill_id BIGINT COMMENT '关联账单ID',
    wishlist_id BIGINT COMMENT '关联心愿单ID',
    
    -- 附加信息
    tags JSON COMMENT '标签数组',
    images JSON COMMENT '图片URL数组',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE SET NULL,
    FOREIGN KEY (to_account_id) REFERENCES accounts(id) ON DELETE SET NULL,
    
    INDEX idx_user_date (user_id, transaction_date),
    INDEX idx_user_type (user_id, type),
    INDEX idx_category (category_id),
    INDEX idx_account (account_id),
    INDEX idx_date (transaction_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易记录表';
