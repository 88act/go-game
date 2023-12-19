package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) BuyGoodsLogic {
	return BuyGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyGoodsLogic) BuyGoods(req *types.IdReq) (resp *types.OkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
