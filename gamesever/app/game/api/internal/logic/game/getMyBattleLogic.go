package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyBattleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyBattleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMyBattleLogic {
	return GetMyBattleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyBattleLogic) GetMyBattle() (resp *types.BattleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
