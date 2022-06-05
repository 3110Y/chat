package usermigration

import "github.com/3110Y/migrator"

func Migration20220606() migrator.MigrationInterface {
	m := migrator.NewMigration("20220606", "Добавляет Пользователя")
	m.Up("")
	m.Down("")
	return m
}
