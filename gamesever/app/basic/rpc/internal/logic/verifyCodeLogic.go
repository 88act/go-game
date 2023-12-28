package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 图形码
func (l *VerifyCodeLogic) VerifyCode(in *pb.VerifyCodeReq) (*pb.VerifyCodeResp, error) {
	codeStr, _ := l.svcCtx.Redis.Get(in.Key)
	if codeStr == in.Code {
		return &pb.VerifyCodeResp{Status: 1}, nil
	} else {
		return &pb.VerifyCodeResp{Status: 0}, nil
	}
}
