package game

import (
	"context"

	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPackLogic {
	return GetPackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackLogic) GetPack() (resp *types.PackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
