// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package role

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiCreateReq                  = sys.ApiCreateReq
	ApiData                       = sys.ApiData
	ApiDeleteReq                  = sys.ApiDeleteReq
	ApiIndexReq                   = sys.ApiIndexReq
	ApiIndexResp                  = sys.ApiIndexResp
	ApiUpdateReq                  = sys.ApiUpdateReq
	AreaInfo                      = sys.AreaInfo
	AreaInfoDeleteReq             = sys.AreaInfoDeleteReq
	AreaInfoIndexReq              = sys.AreaInfoIndexReq
	AreaInfoIndexResp             = sys.AreaInfoIndexResp
	AreaInfoReadReq               = sys.AreaInfoReadReq
	AreaInfoTreeReq               = sys.AreaInfoTreeReq
	AreaInfoTreeResp              = sys.AreaInfoTreeResp
	AuthApiInfo                   = sys.AuthApiInfo
	ConfigResp                    = sys.ConfigResp
	DateRange                     = sys.DateRange
	JwtToken                      = sys.JwtToken
	LoginLogCreateReq             = sys.LoginLogCreateReq
	LoginLogIndexData             = sys.LoginLogIndexData
	LoginLogIndexReq              = sys.LoginLogIndexReq
	LoginLogIndexResp             = sys.LoginLogIndexResp
	Map                           = sys.Map
	MenuCreateReq                 = sys.MenuCreateReq
	MenuData                      = sys.MenuData
	MenuDeleteReq                 = sys.MenuDeleteReq
	MenuIndexReq                  = sys.MenuIndexReq
	MenuIndexResp                 = sys.MenuIndexResp
	MenuUpdateReq                 = sys.MenuUpdateReq
	OperLogCreateReq              = sys.OperLogCreateReq
	OperLogIndexData              = sys.OperLogIndexData
	OperLogIndexReq               = sys.OperLogIndexReq
	OperLogIndexResp              = sys.OperLogIndexResp
	PageInfo                      = sys.PageInfo
	PageInfo_OrderBy              = sys.PageInfo_OrderBy
	Point                         = sys.Point
	ProjectInfo                   = sys.ProjectInfo
	ProjectInfoDeleteReq          = sys.ProjectInfoDeleteReq
	ProjectInfoIndexReq           = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp          = sys.ProjectInfoIndexResp
	ProjectInfoReadReq            = sys.ProjectInfoReadReq
	Response                      = sys.Response
	RoleApiAuthReq                = sys.RoleApiAuthReq
	RoleApiIndexReq               = sys.RoleApiIndexReq
	RoleApiIndexResp              = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq         = sys.RoleApiMultiUpdateReq
	RoleCreateReq                 = sys.RoleCreateReq
	RoleDeleteReq                 = sys.RoleDeleteReq
	RoleIndexData                 = sys.RoleIndexData
	RoleIndexReq                  = sys.RoleIndexReq
	RoleIndexResp                 = sys.RoleIndexResp
	RoleMenuUpdateReq             = sys.RoleMenuUpdateReq
	RoleUpdateReq                 = sys.RoleUpdateReq
	UserAuthArea                  = sys.UserAuthArea
	UserAuthAreaIndexReq          = sys.UserAuthAreaIndexReq
	UserAuthAreaIndexResp         = sys.UserAuthAreaIndexResp
	UserAuthAreaMultiUpdateReq    = sys.UserAuthAreaMultiUpdateReq
	UserAuthProject               = sys.UserAuthProject
	UserAuthProjectIndexReq       = sys.UserAuthProjectIndexReq
	UserAuthProjectIndexResp      = sys.UserAuthProjectIndexResp
	UserAuthProjectMultiUpdateReq = sys.UserAuthProjectMultiUpdateReq
	UserCheckTokenReq             = sys.UserCheckTokenReq
	UserCheckTokenResp            = sys.UserCheckTokenResp
	UserCreateResp                = sys.UserCreateResp
	UserDeleteReq                 = sys.UserDeleteReq
	UserIndexReq                  = sys.UserIndexReq
	UserIndexResp                 = sys.UserIndexResp
	UserInfo                      = sys.UserInfo
	UserLoginReq                  = sys.UserLoginReq
	UserLoginResp                 = sys.UserLoginResp
	UserReadReq                   = sys.UserReadReq
	UserRegister1Req              = sys.UserRegister1Req
	UserRegister1Resp             = sys.UserRegister1Resp
	UserRegister2Req              = sys.UserRegister2Req

	Role interface {
		RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error)
		RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error)
		RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error)
	}

	defaultRole struct {
		cli zrpc.Client
	}

	directRole struct {
		svcCtx *svc.ServiceContext
		svr    sys.RoleServer
	}
)

func NewRole(cli zrpc.Client) Role {
	return &defaultRole{
		cli: cli,
	}
}

func NewDirectRole(svcCtx *svc.ServiceContext, svr sys.RoleServer) Role {
	return &directRole{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRole) RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleCreate(ctx, in, opts...)
}

func (d *directRole) RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleCreate(ctx, in)
}

func (m *defaultRole) RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleIndex(ctx, in, opts...)
}

func (d *directRole) RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error) {
	return d.svr.RoleIndex(ctx, in)
}

func (m *defaultRole) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (d *directRole) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleUpdate(ctx, in)
}

func (m *defaultRole) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (d *directRole) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleDelete(ctx, in)
}

func (m *defaultRole) RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleMenuUpdate(ctx, in, opts...)
}

func (d *directRole) RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleMenuUpdate(ctx, in)
}

func (m *defaultRole) RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleApiAuth(ctx, in, opts...)
}

func (d *directRole) RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleApiAuth(ctx, in)
}

func (m *defaultRole) RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleApiMultiUpdate(ctx, in, opts...)
}

func (d *directRole) RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleApiMultiUpdate(ctx, in)
}

func (m *defaultRole) RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleApiIndex(ctx, in, opts...)
}

func (d *directRole) RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error) {
	return d.svr.RoleApiIndex(ctx, in)
}