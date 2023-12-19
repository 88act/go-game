package logic

import (
	"context"

	"go-game/app/basic/rpc/internal/svc"
	"go-game/app/basic/rpc/pb"
	"go-game/common/utils"

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

	resp := pb.FileListResp{}
	for _, v := range in.Ids {
		fileInfo := pb.FileInfo{Guid: v}
		if !utils.IsEmpty(v) {
			fileInfo.Path, _ = l.svcCtx.BasicFileSev.GetPathByGuid(l.ctx, v)
		}
		resp.List = append(resp.List, &fileInfo)
	}
	return &resp, nil
}
