# irisv12-test

与lib-ui对应的服务端代码，使用iris+gorm+mysql搭建的一个restful项目模板

#### 更新

1. 优化代码结构
2. json传参
3. md5处理

```
conf  配置文件
controllers  控制器 入参处理 api的入口
datasource 数据库配置 
models  结构体
db  sql数据文件 postman接口文件
repo 数据库的操作
route  注册路由
service 业务逻辑代码
utils  工具类
config.json 配置文件的映射
main.go 主程序入口
```

### 启动项目

```
1.安装依赖 go get
2.go run main.go
```