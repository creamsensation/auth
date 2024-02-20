package auth

import (
	"fmt"
	
	"github.com/creamsensation/quirk"
)

var (
	pgUserFields = []quirk.Field{
		{Name: quirk.Id, Props: "serial primary key"},
		{Name: UserActive, Props: "bool not null default false"},
		{Name: UserRoles, Props: "varchar[]"},
		{Name: UserEmail, Props: "varchar(255) not null"},
		{Name: UserPassword, Props: "varchar(128) not null"},
		{Name: UserTfa, Props: "bool not null default false"},
		{Name: UserTfaSecret, Props: "varchar(255)"},
		{Name: UserTfaCodes, Props: "varchar(255)"},
		{Name: UserTfaUrl, Props: "varchar(255)"},
		{Name: quirk.Vectors, Props: "tsvector not null default ''"},
		{Name: UserLastActivity, Props: "timestamp not null default current_timestamp"},
		{Name: quirk.CreatedAt, Props: "timestamp not null default current_timestamp"},
		{Name: quirk.UpdatedAt, Props: "timestamp not null default current_timestamp"},
	}
)

func CreateTable(q *quirk.Quirk) error {
	fields := make([]quirk.Field, 0)
	switch q.DB.DriverName() {
	case quirk.Postgres:
		for _, f := range pgUserFields {
			fields = append(fields, f)
		}
	}
	q.Q(
		fmt.Sprintf(
			`CREATE TABLE IF NOT EXISTS %s (%s)`,
			usersTable,
			quirk.CreateTableStructure(fields),
		),
	)
	return q.Exec()
}

func MustCreateTable(q *quirk.Quirk) {
	err := CreateTable(q)
	if err != nil {
		panic(err)
	}
}

func DropTable(q *quirk.Quirk) error {
	return q.Q(fmt.Sprintf(`DROP TABLE IF EXISTS %s CASCADE`, usersTable)).Exec()
}

func MustDropTable(q *quirk.Quirk) {
	err := DropTable(q)
	if err != nil {
		panic(err)
	}
}
