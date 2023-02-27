# car-rental-management-system

该项目是gin项目的模版，可以以该模版为基础进行开发

使用的库：
```
github.com/dgrijalva/jwt-go 
github.com/fsnotify/fsnotify 
github.com/gin-gonic/gin 
github.com/go-playground/validator/v10 
github.com/go-redis/redis/v8 
github.com/spf13/viper 
go.uber.org/zap 
golang.org/x/crypto 
gopkg.in/natefinch/lumberjack.v2 
gorm.io/driver/mysql 
gorm.io/gorm 
```


## go后台程序运行方式


1.使用config.yaml.example为模版配置config.yaml文件

2.在car-rental-management-system目录下运行`go run main.go`


```
文件/目录名称              说明
app/common               公共模块（请求、响应结构体等）
app/controllers          业务调度器
app/middleware           中间件
app/models               数据库结构体
app/services             业务层
bootstrap                项目启动初始化
config                   配置结构体
global                   全局变量
routes                   路由定义
static                   静态资源（允许外部访问）storage                  系统日志、文件等静态资源）utils                    工具函数
config.yaml              配置文件
main.go                  项目启动文件

```           