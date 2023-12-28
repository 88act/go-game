package user

import (
	"context"

	"go-game/app/usercenter/api/internal/svc"
	"go-game/app/usercenter/api/internal/types"
	"go-game/app/usercenter/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(userId int64) (resp *types.UserInfoResp, err error) {
	rpcReq := usercenter.GetUserInfoReq{Id: userId}
	if rpcResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &rpcReq); err == nil {
		user := types.MemUser{}
		_ = copier.Copy(&user, rpcResp.User)
		return &types.UserInfoResp{UserInfo: user}, nil
	} else {
		return nil, err
	}
}
