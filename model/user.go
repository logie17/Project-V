package model

type User struct {
	Id           int64
	Username     sql.NullString `sql:not null`
	Name         string  `sql:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
