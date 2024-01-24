package logic

import (
	"context"
	"encoding/json"
	"ticket-server/app/user/rpc/pb"

	"ticket-server/app/user/api/internal/svc"
	"ticket-server/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.InfoResp, err error) {
	id := l.ctx.Value("uid").(string)
	loginTime, _ := l.ctx.Value("loginTime").(json.Number).Int64()

	result, err := l.svcCtx.UserRpc.Info(l.ctx, &pb.InfoReq{
		Id:        id,
		LoginTime: loginTime,
	})
	if err != nil {
		return nil, err
	}

	return &types.InfoResp{
		Id:       id,
		Username: result.Username,
		Balance:  result.Balance,
	}, nil
}
