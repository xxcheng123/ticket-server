package logic

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/status"
	"regexp"
	"ticket-server/app/user/api/internal/svc"
	"ticket-server/app/user/api/internal/types"
	"ticket-server/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisteraResp, err error) {
	username := req.Username
	password := req.Password
	uExp := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]{5,11}$")
	if !uExp.MatchString(username) {
		return nil, errors.New("用户名格式错误")
	}
	pExp := regexp.MustCompile("^[a-zA-Z0-9_-]{6,16}$")
	if !pExp.MatchString(password) {
		return nil, errors.New("密码格式错误")
	}
	result, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: username,
		Password: password,
	})
	if err != nil {

		if rpcErr, ok := status.FromError(err); ok {
			fmt.Println(rpcErr.Message(), rpcErr.Code())
		}
		return nil, err
	}
	return &types.RegisteraResp{
		Id: result.Id,
		Ok: true,
	}, nil
}
