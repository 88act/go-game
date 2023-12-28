cd /app/gogame
chmod +x ./usercenter-rpc
chmod +x ./usercenter-api
chmod +x ./basic-rpc
chmod +x ./basic-api 
chmod +x ./game-api
nohup ./usercenter-rpc -f usercenter-rpc.yaml &
nohup ./usercenter-api -f usercenter-api.yaml &
nohup ./basic-rpc -f basic-rpc.yaml &
nohup ./basic-api -f basic-api.yaml &
./game-api -f game-api.yaml 