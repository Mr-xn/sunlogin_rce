# 前言
没啥好说的,直接放工具和有漏洞的版本吧,大家本地玩玩很不错，切勿用于非法测试。

测试向日葵客户端漏洞版本:11.0.0.33162

漏洞客户端在源码中提供.

端口扫描用的是:https://github.com/jboursiquot/portscan

# 漏洞版本(根据反馈获得)
11.1.1

10.3.0.27372

11.0.0.33162
# 命令说明
-h 指定目标

-t 选择扫描或者命令执行  默认scan

-p 设置扫描端口范围 默认4w到65535

-c 需要执行的命令

# 相关截图
扫描
![scan](https://github.com/TRYblog/sunlogin_rce_/blob/main/scan.png "sacn")
rce
![rce](https://github.com/TRYblog/sunlogin_rce_/blob/main/exp.png "rce")

# 时间
20220216
