package logic

import (
	"context"

	"go-game/app/usercenter/rpc/internal/svc"
	"go-game/app/usercenter/rpc/pb"
	"go-game/app/usercenter/rpc/usercenter"
	"go-game/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	if user, err := l.svcCtx.MemUserSev.Get(l.ctx, in.Id, ""); err == nil {
		if user.Status != 1 {
			return nil, errors.New("用户状态无效")
		}
		var respUser usercenter.MemUser
		_ = copier.Copy(&respUser, user)
		return &usercenter.GetUserInfoResp{
			User: &respUser,
		}, nil
	} else {
		return nil, err
	}
}
