package sessions

import (
	"database/sql"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

var sessionManager *scs.SessionManager
var db *sql.DB

func InitSessions() error {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DB_SESSION_URL"))
	if err != nil {
		return err
	}

	sessionManager = scs.New()
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteNoneMode
	sessionManager.Store = mysqlstore.New(db)
	return nil
}

func Shutdown() {
	_ = db.Close()
}

func GetSessions() *scs.SessionManager {
	return sessionManager
}
