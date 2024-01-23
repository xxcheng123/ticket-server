package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"ticket-server/common/errs"

	"ticket-server/app/user/rpc/internal/svc"
	"ticket-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *pb.InfoReq) (*pb.InfoResp, error) {
	uid, _ := strconv.ParseInt(in.Id, 10, 0)
	loginTime := in.LoginTime
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		logx.Errorf("%+v", errors.Wrap(err, "查找用户"))
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errs.UserNotExist
		}
		return nil, errs.MySQLInternalError
	}
	fmt.Println("111", user.LastChangeTime.UnixMilli())
	if user.LastChangeTime.UnixMilli() > loginTime {
		logx.Errorf("user[%s] expired", user.Username)
		return nil, errs.UserExpired
	}
	return &pb.InfoResp{
		Ok:       true,
		Id:       strconv.FormatInt(user.Id, 10),
		Username: user.Username,
		Balance:  user.Balance,
	}, nil
}
