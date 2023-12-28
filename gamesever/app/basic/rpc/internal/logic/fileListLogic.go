package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FileList 文件列表
func (l *FileListLogic) FileList(in *pb.FileListReq) (*pb.FileListResp, error) {

	return nil, nil
}
