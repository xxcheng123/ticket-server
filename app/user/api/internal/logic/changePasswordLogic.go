package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"regexp"
	"ticket-server/app/user/rpc/pb"

	"ticket-server/app/user/api/internal/svc"
	"ticket-server/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.ChangePasswordResp, err error) {
	password := req.Password
	pExp := regexp.MustCompile("^[a-zA-Z0-9_-]{6,16}$")
	if !pExp.MatchString(password) {
		return nil, errors.New("密码格式错误")
	}
	id := l.ctx.Value("uid").(string)
	loginTime, _ := l.ctx.Value("loginTime").(json.Number).Int64()
	result, err := l.svcCtx.UserRpc.ChangePassword(l.ctx, &pb.ChangePasswordReq{
		Id:        id,
		Password:  password,
		LoginTime: loginTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.ChangePasswordResp{
		Ok: result.Ok,
		Id: id,
	}, nil
}
