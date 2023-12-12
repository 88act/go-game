package game

import (
	"context"

	"go-cms/app/game/cmd/api/internal/svc"
	"go-cms/app/game/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeScoreLogic {
	return ChangeScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeScoreLogic) ChangeScore(req *types.ChangeScoreReq) (resp *types.ChangeScoreResp, err error) {
	// todo: add your logic here and delete this line

	return
}
