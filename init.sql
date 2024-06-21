INSERT INTO Department (name, area, description, parent_department_id, parent_departments)
SELECT '山川科技', '北京', '北京山川科技股份有限公司根部门', 9223372036854775806, '0'
WHERE NOT EXISTS (
    SELECT 1 FROM Department
);

-- 初始化角色数据
INSERT INTO Role (name, weight)
SELECT '开发者管理员', 0
WHERE NOT EXISTS (
    SELECT 1 FROM Role
);

-- 初始化用户数据
INSERT INTO User (account, name, department_id, role_id, password, status)
SELECT 'administrator', 'admin', 1, 1, '0okm)OKM', TRUE
WHERE NOT EXISTS (
    SELECT 1 FROM User
);
