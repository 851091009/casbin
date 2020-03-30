# 说明
casbin 在生产环境，root 无法通过权限验证。生产环境是 centos8 , 错误产生是，把编译后放到 centos 中出现这个 BUG

## 1. 先试用 postgres 创建一个名字为 go 的数据库，导入 large_admin_casbin.sql 文件

## 2.  cabin 验证目录：middleware/casbin/
```cgo
// 验证权限
result, err := expansionCasbin.Enforcer.Enforce("1", "/v1/admin/index/index", "GET")
````

## 注意，请 go build main.go 用postman 访问：

127.0.0.1:8095/v1/admin/index/index