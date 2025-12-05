-- 使用 floatisland 数据库
USE floatisland;

-- 支出分类
INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '餐饮美食', 'fa-utensils', 'orange', 1, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='餐饮美食' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '购物消费', 'fa-bag-shopping', 'purple', 2, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='购物消费' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '交通出行', 'fa-bus', 'blue', 3, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='交通出行' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '住房物业', 'fa-house', 'green', 4, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='住房物业' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '医疗健康', 'fa-heartbeat', 'red', 5, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='医疗健康' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'expense', '文化娱乐', 'fa-gamepad', 'pink', 6, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='expense' AND name='文化娱乐' AND is_system=TRUE
);


-- 收入分类
INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'income', '工资薪水', 'fa-sack-dollar', 'indigo', 1, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='income' AND name='工资薪水' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'income', '理财收益', 'fa-arrow-trend-up', 'red', 2, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='income' AND name='理财收益' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'income', '兼职外快', 'fa-briefcase', 'green', 3, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='income' AND name='兼职外快' AND is_system=TRUE
);

INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active, created_at, updated_at)
SELECT 0, 'income', '礼金红包', 'fa-gift', 'pink', 4, TRUE, TRUE, NOW(), NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE type='income' AND name='礼金红包' AND is_system=TRUE
);
