package handler

import (
	"net/http"

	"go-zero-tpl/internal/logic"
	"go-zero-tpl/internal/svc"
	"go-zero-tpl/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ZeroHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZeroLogic(r.Context(), ctx)
		resp, err := l.Zero(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
