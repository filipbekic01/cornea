package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Table
var tableName = ""

// Model
type Model struct {
}

func (m *Model) TableName() string {
	return tableName
}

// Model GORM
type ModelGORM struct {
	*gorm.Model
}

func (m *ModelGORM) TableName() string {
	return tableName
}

// Kernel
type Kernel struct {
	*gorm.DB
}

func (k *Kernel) Table(name string) *Kernel {
	k.DB = k.DB.Table(name)
	return k
}

func (k *Kernel) CreateTable(name string, fields bool) {
	tableName = name

	if fields {
		k.DB.CreateTable(&ModelGORM{})
	} else {
		k.DB.CreateTable(&Model{})
	}
}

func (k *Kernel) AddColumn(newName string, tag string) {
	scope := k.NewScope(k.Value)

	newName = scope.Quote(newName)
	tag = scope.Quote(tag)

	scope.Raw(fmt.Sprintf("ALTER TABLE %v ADD COLUMN %v %v;", scope.QuotedTableName(), newName, tag)).Exec()
}

func (k *Kernel) ChangeColumn(name string, newName string) {
	scope := k.DB.NewScope(k.DB.Value)

	name = scope.Quote(name)
	newName = scope.Quote(newName)

	cmd := fmt.Sprintf("ALTER TABLE %v RENAME %v TO %v;", scope.QuotedTableName(), name, newName)

	fmt.Println(cmd)

	scope.Raw(cmd).Exec()
}
