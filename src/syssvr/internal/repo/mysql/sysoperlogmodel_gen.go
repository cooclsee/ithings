// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysOperLogFieldNames          = builder.RawFieldNames(&SysOperLog{})
	sysOperLogRows                = strings.Join(sysOperLogFieldNames, ",")
	sysOperLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysOperLogFieldNames, "`id`", "`createdTime`", "`updatedTime`", "`deletedTime`"), ",")
	sysOperLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysOperLogFieldNames, "`id`", "`createdTime`", "`updatedTime`", "`deletedTime`"), "=?,") + "=?"
)

type (
	sysOperLogModel interface {
		Insert(ctx context.Context, data *SysOperLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysOperLog, error)
		Update(ctx context.Context, data *SysOperLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysOperLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysOperLog struct {
		Id           int64          `db:"id"`           // 编号
		OperUid      int64          `db:"operUid"`      // 用户id
		OperUserName string         `db:"operUserName"` // 操作人员名称
		OperName     string         `db:"operName"`     // 操作名称
		BusinessType int64          `db:"businessType"` // 业务类型（1新增 2修改 3删除 4查询 5其它）
		Uri          string         `db:"uri"`          // 请求地址
		OperIpAddr   string         `db:"operIpAddr"`   // 主机地址
		OperLocation string         `db:"operLocation"` // 操作地点
		Req          sql.NullString `db:"req"`          // 请求参数
		Resp         sql.NullString `db:"resp"`         // 返回参数
		Code         int64          `db:"code"`         // 返回状态（200成功 其它失败）
		Msg          string         `db:"msg"`          // 提示消息
		CreatedTime  time.Time      `db:"createdTime"`  // 操作时间
	}
)

func newSysOperLogModel(conn sqlx.SqlConn) *defaultSysOperLogModel {
	return &defaultSysOperLogModel{
		conn:  conn,
		table: "`sys_oper_log`",
	}
}

func (m *defaultSysOperLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysOperLogModel) FindOne(ctx context.Context, id int64) (*SysOperLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysOperLogRows, m.table)
	var resp SysOperLog
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysOperLogModel) Insert(ctx context.Context, data *SysOperLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysOperLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OperUid, data.OperUserName, data.OperName, data.BusinessType, data.Uri, data.OperIpAddr, data.OperLocation, data.Req, data.Resp, data.Code, data.Msg)
	return ret, err
}

func (m *defaultSysOperLogModel) Update(ctx context.Context, data *SysOperLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysOperLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OperUid, data.OperUserName, data.OperName, data.BusinessType, data.Uri, data.OperIpAddr, data.OperLocation, data.Req, data.Resp, data.Code, data.Msg, data.Id)
	return err
}

func (m *defaultSysOperLogModel) tableName() string {
	return m.table
}
