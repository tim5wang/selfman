package logic

import (
	"context"
	"errors"

	"github.com/tim5wang/selfman/app/user/api/internal/svc"
	"github.com/tim5wang/selfman/app/user/api/internal/types"
	"github.com/tim5wang/selfman/app/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.UserReq) (*types.UserRsp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}
	if user.Name != "test" {
		return nil, errors.New("用户不存在")
	}
	return &types.UserRsp{
		Id:   req.Id,
		Name: "test user",
	}, nil
}
