package postgres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	dsql "database/sql"

	sql "github.com/alokic/gopkg/sql/v2"
)

var (
	// DriverName is type of sql driver used to access db.
	DriverName = "postgres"
)

// Logger is interface for logger in db.
type Logger interface {
	Log(...interface{}) error
}

// DB is sql.DB interface.
type DB interface {
	Exec(string, ...interface{}) (dsql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Transaction(func(*sql.Tx) (interface{}, error)) error
}

// BatchInsert inserts in batch.
func BatchInsert(db DB, tableName string, records []interface{}, fi *sql.FieldInfo) error {
	stmt, params := sql.BatchInsertStatement(tableName, records, fi)
	_, err := db.Exec(stmt, params...)
	return err
}

// Clear db.
func Clear(d DB, logger Logger) {
	for _, v := range []string{"access_policies_permissions", "role_access_policies", "user_roles", "users", "roles", "permissions", "access_policies", "services"} {
		logger.Log("Clear", v, "error", clearTable(d, v))
	}
}

// Seed seeds user db.
func Seed(d DB, folder string, logger Logger) {
	for _, v := range []string{"users", "services", "roles", "permissions", "access_policies", "user_roles", "access_policies_permissions", "role_access_policies"} {
		logger.Log("Seed", v, "error", seed(d, v, fmt.Sprintf("%s/%s.json", folder, v)))
	}
}

func seed(d DB, name, file string) error {
	var ms []map[string]interface{}

	err := readSeedFile(file, &ms)
	if err != nil {
		return err
	}

	for _, m := range ms {
		cols, params := createQuery(m)
		_, err = d.Exec(fmt.Sprintf("INSERT into %s (%s) VALUES (%s)", name, cols, params))
		if err != nil {
			return err
		}
	}
	return nil
}

// clearTable removes data from table.
func clearTable(d DB, name string) error {
	_, err := d.Exec(fmt.Sprintf("DELETE from %s", name))
	return err
}

func createQuery(m map[string]interface{}) (string, string) {
	var cols []string
	var params []string

	// if _, ok := m["created_at"]; !ok {
	// 	m["created_at"] = time.Now().UTC()
	// }

	// if _, ok := m["updated_at"]; !ok {
	// 	m["updated_at"] = time.Now().UTC()
	// }

	for k, v := range m {
		cols = append(cols, k)
		params = append(params, fmt.Sprintf("'%v'", v))
	}
	return strings.Join(cols, ","), strings.Join(params, ",")
}

func readSeedFile(file string, v interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
