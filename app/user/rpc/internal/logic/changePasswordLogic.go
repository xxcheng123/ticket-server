package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"ticket-server/common/errs"
	"time"

	"ticket-server/app/user/rpc/internal/svc"
	"ticket-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *pb.ChangePasswordReq) (*pb.ChangePasswordResp, error) {
	uid, _ := strconv.ParseInt(in.Id, 10, 0)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		logx.Errorf("%+v", err)
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errs.UserNotExist
		}
		return nil, errs.MySQLInternalError
	}
	if user.LastChangeTime.UnixMilli() > in.LoginTime {
		logx.Errorf("user[%s] expired", user.Username)
		return nil, errs.UserExpired
	}
	user.Password = in.Password
	user.LastChangeTime = time.Now()
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		logx.Errorf("%+v", err)
		return nil, errs.MySQLInternalError
	}

	return &pb.ChangePasswordResp{
		Ok: true,
	}, nil
}
