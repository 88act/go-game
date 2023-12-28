package basic

import (
	"net/http"
	"strconv"

	"go-game/common/result"

	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//文件上传一般会采用 POST multipart/form-data 的形式，处理这类请求要调用 r.ParseMultipartForm，无论是显式调用，
		//还是在 r.FormFile 里面的隐式调用 那 32 Mb 是对文件上传大小的限制吗 ，上传的文件们按顺序存入内存中，累加大小不得超出 32Mb

		// basicFile := model.BasicFile{}
		// r.ParseMultipartForm(32 << 20) //32 Mb
		// r.ParseForm()
		// if r.MultipartForm != nil {
		// 	//userId, _ := strconv.Atoi(r.MultipartForm.Value["userId"][0])

		// 	basicFile.Md5 = r.MultipartForm.Value["md5"][0]
		// 	basicFile.Sha1 = r.MultipartForm.Value["sha1"][0]
		// 	basicFile.Size, _ = strconv.Atoi(r.MultipartForm.Value["size"][0])
		// 	basicFile.FileType = r.MultipartForm.Value["fileType"][0]
		// 	catid, _ := strconv.Atoi(r.MultipartForm.Value["catId"][0])
		// 	basicFile.CatId = int64(catid)
		// }
		// l := basic.NewUploadLogic(r.Context(), svcCtx, r)
		// resp, err := l.Upload(basicFile)
		// result.HttpResult(r, w, resp, err)

		// if err := httpx.Parse(r, &req); err != nil {
		// 	result.ParamErrorResult(r, w, err)
		// 	return
		// }
		//fmt.Println("MyUploadHandler开始文件上传....")
		var req types.UploadReq
		//basicFile := model.BasicFile{}
		r.ParseMultipartForm(32 << 20) //32 Mb
		r.ParseForm()
		if r.MultipartForm != nil {
			// userId, _ := strconv.Atoi(r.MultipartForm.Value["userId"][0])
			//basicFile.UserId = int64(userId)
			req.Md5 = r.MultipartForm.Value["md5"][0]
			req.Sha1 = r.MultipartForm.Value["sha1"][0]
			req.Size, _ = strconv.Atoi(r.MultipartForm.Value["size"][0])
			req.FileType = r.MultipartForm.Value["fileType"][0]
			catid, _ := strconv.Atoi(r.MultipartForm.Value["catId"][0])
			req.CatId = int64(catid)
			beCut, _ := strconv.Atoi(r.MultipartForm.Value["beCut"][0])
			req.BeCut = beCut
		}
		//logx.Errorf("beCut = ", req.BeCut)
		l := basic.NewUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.Upload(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
