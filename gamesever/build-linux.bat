SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
 
 
go build -o ./run/usercenter-api ./app/usercenter/cmd/api/usercenter.go
  
go build -o ./run/basic-api  ./app/basic/cmd/api/basic.go

go build -o ./run/game-rpc  ./app/game/cmd/rpc/game.go

go build -o ./run/game-api  ./app/game/cmd/api/game.go
 
