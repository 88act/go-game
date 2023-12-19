package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBattleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBattleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBattleLogic {
	return GetBattleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBattleLogic) GetBattle() (resp *types.BattleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
