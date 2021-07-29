package account

import (
	"database/sql"
	"time"
)

type UserDB struct {
	UserID     sql.NullString `db:"user_id"`
	UserName   sql.NullString `db:"username"`
	Password   sql.NullString `db:"password"`
	Salt       sql.NullString `db:"salt"`
	ProfilePic sql.NullString `db:"profile_pic"`
	CreatedAt  time.Time      `db:"created_at"`
}
