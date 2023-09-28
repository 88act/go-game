#  go-webrtc-im
 
go-webrtc-im  一个基于go webrtc 的视频会议于聊天IM项目 \
包含完整的服务器端 go webrtc/rtmp  推流服务器， go websocket im服务器  \
包含完整的vue3 typescript客户端 \
消息支持protocol buffer 和json 格式  \
支持消息路由  \
支持消息协议编号 \
简单明了，方便扩展  \


demo 实例 是 一个客户端 js html 发送 json 格式， 服务器端用protobuf 协议的混合demo 。 \
方便网页简单的发json数据 \
html使用json发送消息， 而服务器端可以方便的使用 protobuf 生成的类 \

```shell
go get github.com/88act/go-webrtc-im
```

## demo

简单在线聊天chat ， 维护在线用户状态，上线 ，下线通知，发送消息等。 可以方便的扩展为游戏服务器框架，或者即时聊天IM 服务器框架 

 

### client
客户端 
chat.html

```javascript
 <script>
    var ws = new WebSocket("ws://localhost:8000");  
    //连接打开时触发 
    ws.onopen = function(evt) {   
        console.log("连接成功。。。。")
        ws.send('101001|{"name":"web1111","group":"webBrowser","appKey":"eqwe23123123","ip":"202.202.202.202","ipLocal":"10.0.0.1","port":8080}')
		 
    };  
    //接收到消息时触发  
    ws.onmessage = function(evt) {  
        console.log("接收 e: " + evt.data);  
    };  
    //连接关闭时触发  
    ws.onclose = function(evt) {  
        console.log("Connection closed.");  
    }; 
    
    var intervalId, timeoutId;
    intervalId = 1;
    timeoutId = setInterval(function () {
        intervalId++;
        console.log(intervalId);
        console.log("发送。。。",intervalId); 
         ws.send('100001|{"status":1,"msg":"' + intervalId + '"}')  
    }, 6000);
    
    
  </script>
```

nginx 反向代理设置

```
upstream proxy_server {
        server 127.0.0.1:8000;
}

server {
        listen 443 ssl;
        server_name xxx.xxx.com;
        location /wss/ {
                proxy_pass http://proxy_server;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header Connection "upgrade";
                proxy_connect_timeout 30d;
                proxy_send_timeout 30d;
                proxy_read_timeout 30d;
        }
        keepalive_timeout 999999999s;
        ssl_certificate cert/xxx.pem;  #需要将cert-file-name.pem替换成已上传的证书文件的名称。
        ssl_certificate_key cert/xxx.key; #需要将cert-file-name.key替换成已上传的证书密钥文件的名
    
}

```
