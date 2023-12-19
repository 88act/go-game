package basic

import (
	"net/http"

	"go-game/common/result"

	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewFileDetailLogic(r.Context(), svcCtx)
		resp, err := l.FileDetail(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
