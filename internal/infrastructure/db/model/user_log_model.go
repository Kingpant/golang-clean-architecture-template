package model

import (
	"time"

	"github.com/uptrace/bun"
)

type UserLogActionType string

const (
	UserLogActionTypeCreate UserLogActionType = "CREATE"
	UserLogActionTypeUpdate UserLogActionType = "UPDATE"
)

type UserLog struct {
	bun.BaseModel `bun:"table:user_logs,alias:ul"`
	ID            string            `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserID        string            `bun:"user_id,notnull"`
	Action        UserLogActionType `bun:"action,type:user_log_action_type,notnull"`
	CreatedAt     time.Time         `bun:"created_at,notnull,default:current_timestamp,update_time"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}
