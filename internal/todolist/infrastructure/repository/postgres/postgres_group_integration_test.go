package postgres

import "testing"

func TestCheckConnection(t *testing.T) {
	user := "postgres"
	password := "postgres"
	dbname := "stepanlahov"

	_ = NewPostgresGroup(user, password, dbname)
}
