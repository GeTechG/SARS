package sessions

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var sessionManager *scs.SessionManager
var db *sql.DB

func InitSessions() error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("root:root@tcp(localhost:%s)/sars_sessions?parseTime=true", os.Getenv("DB_PORT")))
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
