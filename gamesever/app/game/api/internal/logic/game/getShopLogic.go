package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetShopLogic {
	return GetShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopLogic) GetShop() (resp *types.ShopResp, err error) {
	// todo: add your logic here and delete this line

	return
}
