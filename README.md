1. 临时搭建的目录，具体启动操作参看 Makefile 文件
2. 操作说明：
   - go mod tidy
   - 启动服务端：make server
   - 启动客户端：make client

3目录介绍：
   - cmd 存储的客户端和服务端的启动代码，也就是main.main
   - config 配置文件相关
   - docs 文档相关，存放着proto
   - init 启动相关配置
   - internal 不经常变动的代码
   - pkg 业务代码、拦截器等
   - tools 工具相关，这里放着证书密钥
   - website 网站相关