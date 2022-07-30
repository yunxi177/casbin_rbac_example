本示例是用 go 语言写的，使用 gorm 适配器，将 policy 存入数据库中的。

## 使用方式
### clone 项目
```shell
git clone git@github.com:yunxi177/casbin_rbac_example.git
```
### 安装依赖
```shell
go mod tidy
```
### 修改配置
打开 main.go 文件，将下面数据库连接修改为你的数据库连接
```go
gormadapter.NewAdapter("mysql", "username:password@tcp(host:3306)/", "your policy table name")
```

### 运行项目
```shell
go run main.go
```
### 常见问题
#### Error 1071: Specified key was too long; max key length is 1000 bytes
此时数据库表是已经建好的，但是建索引的时候报错了，可以进入数据库把表的引擎改为 InnoDB，然后重新运行。