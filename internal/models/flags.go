package models

type Flags struct {
	Help           bool
	Env            string
	GenMigration   string
	ListMigrations bool
	Create         bool
	Drop           bool
	Migrate        bool
	MigrateDry     bool
	FullReset      bool
	Summary        bool
}
