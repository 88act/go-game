package basic

import (
	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"
	"go-game/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MyFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewMyFileListLogic(r.Context(), svcCtx)
		resp, err := l.MyFileList(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
