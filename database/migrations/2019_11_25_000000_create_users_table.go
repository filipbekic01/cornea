package migrations

import (
	"github.com/filipbekic01/cornea/database"
)

type CreateUsersTable struct {
	*database.Kernel
}

func (m *CreateUsersTable) Up() {
	m.CreateTable("users", true)
	m.Table("users").ChangeColumn("name", "first_name")
}
