package sessions

import (
	"database/sql"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
)

var sessionManager *scs.SessionManager
var db *sql.DB

func InitSessions() error {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/sars_sessions?parseTime=true")
	if err != nil {
		return err
	}

	sessionManager = scs.New()
	sessionManager.Store = mysqlstore.New(db)
	return nil
}

func Shutdown() {
	_ = db.Close()
}

func GetSessions() *scs.SessionManager {
	return sessionManager
}
