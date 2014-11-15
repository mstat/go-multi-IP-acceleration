go-multi-IP-acceleration
========================
go-multi-IP-acceleration是用golang实现的一个socks5代理，它可以把流量随机从多个本地IP发出，在某些网络环境下可以达到加速网络的效果

使用说明：
下载golang开发环境，然后编译对应平台的二进文件
用nmap等工具扫描出一个空余IP列表
然后在本地机器加上多个IP,(windows在设置IP的对话框选择高级，然后加上多个IP)，linux自己搜索相关资料。
然后把goproxy.ini放到和二进制文件一个目录，并修改配置文件
配置文件：
  port 为http代理功能保留，暂空
  port_socks5 代表监听端口
  local ip 本机IP地址列表
  allowed ip 允许使用代理的IP列表，暂未实现

然后运行， 打开迅雷，进入设置->高级设置->代理设置，然后添加代理，服务器 127.0.0.1 端口，和port_socks5一致，类型socks5.点测试，如果成功的话就ok了。建议把任务默认属性->原始地线程数调到10


TODO：
  身份验证
  http代理
  自动翻墙
  自动检测IP冲突
