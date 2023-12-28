package user

import (
	"context"

	"go-game/app/usercenter/api/internal/svc"
	"go-game/app/usercenter/api/internal/types"
	"go-game/app/usercenter/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	rpcReq := pb.LoginReq{Username: req.Username, Password: req.Password, LoginType: 1}
	if rpcResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &rpcReq); err == nil {
		var myResp types.LoginResp
		_ = copier.Copy(&myResp, rpcResp)
		return &myResp, nil
	} else {
		return nil, err
	}
}
