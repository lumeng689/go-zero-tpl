package logic

import (
	"context"

	"go-zero-tpl/internal/svc"
	"go-zero-tpl/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ZeroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZeroLogic(ctx context.Context, svcCtx *svc.ServiceContext) ZeroLogic {
	return ZeroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZeroLogic) Zero(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
