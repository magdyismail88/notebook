package bootstrap

import "database/sql"

// Open database connection
func NewDBConnection(env *Env) *sql.DB {
	conn, err := sql.Open(env.DatabaseAdapter, env.DatabaseName)
	if err != nil {
		panic(err)
	}
	return conn
}

// Close database connectino
func CloseDBConnection(conn *sql.DB) {
	if conn == nil {
		return
	}
	conn.Close()
}
