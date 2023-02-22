package bootstrap

import "database/sql"

type Application struct {
	Env    *Env
	DBConn *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DBConn = NewDBConnection(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.DBConn)
}
