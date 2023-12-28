  

### 1.介绍
 
go-game 是一款基于Go语言开发的支持分布式与微服务架构的游戏服务器。 支持单体、分布式等多种架构方案。  
本游戏框架居于 gozero+gorm 微服务框架开发，可复用go-zero的grpc调用，服务治理，消息队列，日志，链路追踪等功能   
通讯使用zinx框架支持 websocket ，http ，kcp ，tcp 多种通讯方式， 支持protobuf，json 协议。     
消息队列支持kafka、redis、nats   
日志支持文件模式，elasticsearch模式，支持prometheus性能监控等。  

client 提供一个居于  layabox typescript protobuf 的客户端 即时战斗 mmo rpg 游戏 demo。  
 
![架构图](gogame.jpg) 
 
### 2.背景现状

老一代的游戏服务器框架中，游戏相关的逻辑都放在单一服务器，单一服务器有性能瓶颈，特别是复杂的战斗运算等，通常一个区只能达到数千人的规模。 单体服务的代码复杂，往往伴有不少bug，维护和测试都相当困难，服务器出bug宕机时将导致整区停服。 一些老游戏服务器也会把聊天服务器，拍卖行服务器，商城服务器，工会服务器等拆分成独立的进程或独立服务器。但还是基于各自独立编码的方式，没有统一的标准，导致项目变的更加庞杂难以维护。


### 3.本项目优势
 
  本项目居于成熟的微服务器框架和通讯框架构   
  充分利用http短连接 和websocket 长连接的各自优点处理游戏场景  
  适合大世界模式游戏， 回合制，战旗 等类型的游戏开发，虚拟社区 ，微信小游戏等  
  也适合即时战斗类，帧同步类的 MMORPG类型游戏  
  适合即时通讯类的应用开发  

  ### 4.目录结构

  ```golang

   gameserver  //服务器
      doc
         game.sql // 数据库
      app
         usercenter            
            api
            rpc
            model
         basic             
            api 
            rpc
            model
         game            
            api
               internal
                  gameserver   // 游戏服务器 wss/tcp /kcp
                     pb/msg.proto // protobuf 定义文件
            rpc
            model
      devops  // 本地部署docker文件  包含 prometheus jaeger  kafka 等
         
      client      //客户端
        layabox  // layabox 客户端demo  
        h5      //    vue/h5 客户端 demo          
     

```
 ###  4. 本地替换官方的 github.com/zeromicro/go-zero 库
 从官方克隆一份 github.com/zeromicro/go-zero 源码 ，放在本地目录go-zero  
  在 go.mod 添加   
    replace github.com/zeromicro/go-zero v1.6.1 =>  ../go-zero  
 然后 下载   https://github.com/88act/go-zero/blob/master/core/stores/redis/redis.go  
 替换官方的 core/stores/redis/redis.go 文件  

 修改后的 redis  
   支持缓存 struct interface{}  []byte list 等复杂对象  
  支持 设置过期时间  
  支持 批量删除等  
 
  ### 5. 启动服务

启动环境 
cd  /devpos/   
docker-compose -f docker-compose.yaml  up -d  

进入容器  
docker exec -it kafka /bin/sh  
cd /opt/kafka/bin/  
创建topic  
 ./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic gogame-log  
 
修改权限sudo  
chmod 777 data/elasticsearch/data  
chown root deploy/filebeat/conf/filebeat.yml  

重启一次   
docker-compose restart go-stash  
docker-compose restart filebeat   

拷贝编译后的go二进制文件 到  
/devpos/go/gogame/目录  

启动 gogame  
docker-compose -f docker-gogame.yaml  up -d  



测试： 

测试机添加host ，指向测试服务器  
10.0.0.100  goenv.local   


访问接口测试 api接口   

https://goenv.local/usercenter/v1/user/login  
https://goenv.local/usercenter/v1/user/register  


ELV日志查询  
http://goenv.local:8980/  

prometheus 监控   
http://goenv.local:9090/  


链路追踪  
http://goenv.local:16686/search  
 

测试 websocket  

客户端连接 https://goenv.local/wss   






        


 