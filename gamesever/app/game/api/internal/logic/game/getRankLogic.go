package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRankLogic {
	return GetRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankLogic) GetRank() (resp *types.RankResp, err error) {
	// todo: add your logic here and delete this line

	return
}
