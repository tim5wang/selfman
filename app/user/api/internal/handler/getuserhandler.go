package handler

import (
	"net/http"

	"github.com/tim5wang/selfman/app/user/api/internal/logic"
	"github.com/tim5wang/selfman/app/user/api/internal/svc"
	"github.com/tim5wang/selfman/app/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
