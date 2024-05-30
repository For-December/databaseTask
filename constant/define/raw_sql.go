package define

// Task1 （1）查询员工人数超过500的所有上市公司的公司代码、公司名称、注册地址；
const Task1 = `
SELECT CompanyCode, CompanyName, RegisteredAddress
FROM Company
WHERE EmployeeCount > 500;

`
