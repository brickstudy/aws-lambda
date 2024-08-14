package domain

import "database/sql"

type Media struct {
	Idx      uint64
	Source_  string
	Category string
	Url      string
	ClientID sql.NullString
	ClientPW sql.NullString
	Token    sql.NullString
}
