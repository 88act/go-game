package game

import (
	"context"

	"go-cms/app/game/cmd/api/internal/svc"
	"go-cms/app/game/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetNotifyLogic {
	return GetNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNotifyLogic) GetNotify() (resp *types.NotifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
