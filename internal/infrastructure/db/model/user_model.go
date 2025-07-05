package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name      string    `bun:"name,notnull"`
	Email     string    `bun:"email,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp,update_time"`
}
