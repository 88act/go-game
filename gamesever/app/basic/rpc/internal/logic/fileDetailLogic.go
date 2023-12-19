package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDetailLogic {
	return &FileDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FileDetail //文件详情
func (l *FileDetailLogic) FileDetail(in *pb.FileDetailReq) (*pb.FileDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileDetailResp{}, nil
}
