package basic

import (
	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"
	"go-game/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFileByKeyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewGetFileByKeyLogic(r.Context(), svcCtx)
		resp, err := l.GetFileByKey(&req)
		result.HttpResult(r, w, resp, err, req)

	}
}
