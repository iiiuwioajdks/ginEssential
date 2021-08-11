# ginEssential
## main函数
main函数不包含业务逻辑，仅仅调用初始化数据库，调用配置文件，启动框架的操作

## config包
config包下是appication.yml文件，主要存放各类工具的配置，如mysql，redis，server port等等

## common包
common包是基础包配置，如数据库初始化函数，jwt获取函数等等的存放

## controller包
cotroller 是各类handler函数的实现，在这里实现crud等操作

## model包
存放后端数据库实体类

## dto包
封装返回前端的数据，也就是与model对应的，返回前端的结构体（实体类）

## middleware包
存放中间件，也就是路由调用handle之前处理的事情
AuthMiddleWare 是token验证中间件

## response包
封装统一的返回函数，Response（）和两个常见的 Success 和 Fail

## util包
封装工具类

## routers包
路由包，存放路由组等