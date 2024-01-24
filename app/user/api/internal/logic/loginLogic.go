package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"ticket-server/app/user/rpc/pb"
	"time"

	"ticket-server/app/user/api/internal/svc"
	"ticket-server/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	result, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	n := time.Now()
	claims["exp"] = n.Add(time.Hour).Unix()
	claims["iat"] = n.Unix()
	claims["uid"] = result.Id
	claims["loginTime"] = n.UnixMilli()
	claims["username"] = req.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, _ := token.SignedString([]byte("okhttp123"))
	return &types.LoginResp{
		Token:      tk,
		Id:         result.Id,
		ExpireTime: int64(time.Hour / time.Second),
	}, nil
}
