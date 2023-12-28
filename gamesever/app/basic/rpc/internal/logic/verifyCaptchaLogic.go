package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 图形码
func (l *VerifyCaptchaLogic) VerifyCaptcha(in *pb.VerifyCodeReq) (*pb.VerifyCodeResp, error) {

	codeStr, _ := l.svcCtx.Redis.Get(in.Key)
	if codeStr == in.Code {
		return &pb.VerifyCodeResp{Status: 1}, nil
	} else {
		return &pb.VerifyCodeResp{Status: 0}, nil
	}
}
