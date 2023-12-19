package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) UseGoodsLogic {
	return UseGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseGoodsLogic) UseGoods(req *types.IdReq) (resp *types.OkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
