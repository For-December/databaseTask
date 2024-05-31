package define

// Task1 （1）查询员工人数超过500的所有上市公司的公司代码、公司名称、注册地址；
const Task1 = `
SELECT org_code, org_name, registered_address
FROM companies
WHERE employee_count > 500;
`

// Task2 （2）查询“星光农机”公司的所有高管的基本信息
const Task2 = `
SELECT e.*
FROM executives e
JOIN companies c ON e.org_code = c.org_code
WHERE c.org_name LIKE '%星光农机%';
`

// Task3 （3）查询注册地址在湖北和湖南的所有上市公司
const Task3 = `
SELECT org_code, org_name, registered_address
FROM companies
WHERE registered_address LIKE '%湖北%' OR registered_address LIKE '%湖南%';
`

// Task4 （4）查询至少在2家以上上市公司任职“独立董事”的高管姓名
const Task4 = `
SELECT name
FROM executives
WHERE position = '独立董事'
GROUP BY name
HAVING COUNT(org_code) >= 2;
`

// Task5 （5）查询至少和“何德军”所在所有公司的相同的所有高管姓名
const Task5 = `
SELECT DISTINCT e2.name
FROM executives e1
JOIN executives e2 ON e1.org_code = e2.org_code
WHERE e1.name = '何德军' AND e2.name <> '何德军';

`

// Task6 （6）列出所有上市公司，统计每个上市公司的高管人数，并按照人数排序倒序排序
const Task6 = `
SELECT companies.org_code, companies.org_name, COUNT(*) AS executive_count
FROM executives
JOIN companies ON executives.org_code = companies.org_code
GROUP BY companies.org_code,companies.org_code
ORDER BY executive_count DESC;
`

// Task7 （7）列出所有上市公司，统计每个上市公司不同类型职务的高管人数，并按照注册资金倒序排序
const Task7 = `
SELECT c.org_code, c.org_name, c.registered_capital, e.position, COUNT(*) AS position_count
FROM executives e
JOIN companies c ON e.org_code = c.org_code
GROUP BY c.org_code, c.org_name, c.registered_capital, e.position
ORDER BY c.registered_capital DESC, position_count DESC;
`

// Task8 （8）查询所有“武汉大学”校友的高管列表以及其所属公司代码、公司名称
const Task8 = `
SELECT e.*, c.org_code, c.org_name
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
JOIN companies c ON e.org_code = c.org_code
WHERE u.university_name = '武汉大学';
`

// Task9 （9）查询所有高管来自于“武汉大学”校友的公司列表
const Task9 = `
SELECT DISTINCT c.org_code, c.org_name
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
JOIN companies c ON e.org_code = c.org_code
WHERE u.university_name = '武汉大学';
`

// Task10 （10）查询每一所学校的校友的高管列表以及其所属公司代码、公司名称
const Task10 = `
SELECT u.university_name, e.*, c.org_code, c.org_name
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
JOIN companies c ON e.org_code = c.org_code
ORDER BY u.university_name;
`

// Task11 （11）批量插入查询结果到“高管-学校的校友关联”的表中
const Task11 = `
INSERT INTO alumni_associations (executive_id, university_id)
SELECT e.executive_id, u.university_id
FROM executives e
JOIN universities u ON e.resume LIKE CONCAT('%', u.university_name, '%');
`

// Task12 （12）查询武汉大学的所有高管校友列表
const Task12 = `
SELECT e.*
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
WHERE u.university_name = '武汉大学';
`

// Task13 （13）去重查询武汉大学的所有高管校友列表，并输出该校友任职公司数量
const Task13 = `
SELECT e.name, e.sex, e.age, COUNT(DISTINCT e.org_code) AS company_count
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
WHERE u.university_name = '武汉大学'
GROUP BY e.name, e.sex, e.age;

`

// Task14 （14）列出所有高校，统计每个高校的高管人数，并根据高管人数进行倒序排序
const Task14 = `
SELECT u.university_name, COUNT(*) AS executive_count
FROM universities u
JOIN alumni_associations a ON u.university_id = a.university_id
JOIN executives e ON a.executive_id = e.executive_id
GROUP BY u.university_name
ORDER BY executive_count DESC;
`

// Task15 （15）列出所有高校，统计每个高校的每一类高管人数，并根据总的高管人数进行倒序排序
const Task15 = `
SELECT u.university_name, e.position, COUNT(*) AS position_count
FROM universities u
JOIN alumni_associations a ON u.university_id = a.university_id
JOIN executives e ON a.executive_id = e.executive_id
GROUP BY u.university_name, e.position
ORDER BY COUNT(*) DESC;
`

// Task16 （16）创建武汉大学高管校友视图
const Task16 = `
CREATE VIEW WuHanUniversityExecutives AS
SELECT e.*
FROM executives e
JOIN alumni_associations a ON e.executive_id = a.executive_id
JOIN universities u ON a.university_id = u.university_id
WHERE u.university_name = '武汉大学';
`

// Task17 （17）创建用户并授权查询权限
const Task17 = `
CREATE USER 'readonly_user'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT ON *.* TO 'readonly_user'@'localhost';
`
