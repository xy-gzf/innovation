# innovation
创新创业信息管理系统后端



## 版本

- **go：** v1.14.0+
- **开发工具：** goland



## 使用

1. 通过开发工具启动

   通过goland打开项目文件，在goland中设置go代理  GOPROXY=https://goproxy.io/

   goland自动下载项目依赖包后，通过运行main.go文件即可启动项目

   <img src="https://img-blog.csdnimg.cn/20210621134958370.png" alt="image-20210621134819284" style="zoom:67%;" />

2. **通过命令行启动**

   ```shell
   # 1.设置代理
   go env -w GOPROXY=https://goproxy.cn,direct
   # 2.下载依赖包
   go mod tidy
   # 3.打包项目 为可执行文件
   go build
   # 4.运行可执行文件
   
   ## 不打包，直接运行
   go run main.go
   ```

   

## 项目结构

- Innovation
  - api                          传统意义上的service层，通过路由调用，然后调用持久层
  - config                    各中间件的配置信息，为项目配置信息做依赖
  - core                        项目的启动配置信息               
  - globle                    配置数据库、redis、日志等全局参数
  - initialize                 项目启动初始化工作（数据库表基础信息录入、路由启动等）
  - log                          日志文件夹，存放日志文件
  - middleware          中间件，自定义jwt，http请求头部等
  - model                    模型层，对数据进行绑定，restful风格
  - router                    路由层，监听http请求
  - service                   实际执行任务的持久层，对数据库进行读写
  - source                   数据库要初始化的数据源（数据不需要通过脚本导入，运行项目自动导入）
  - uploads                 上传文件所在路径
  - utils                        工具类
  - config.yaml           项目配置信息，端口号、数据库等配置
  - go.mod                  项目依赖
  - main.go                 程序入口
  - README.md         项目必读

