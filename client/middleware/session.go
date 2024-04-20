package middleware

import (
	"database/sql"
	"log"
	"time"

	"github.com/alexedwards/scs/cockroachdbstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/lib/pq"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func InitSession() *scs.SessionManager {
	var sessionManager *scs.SessionManager

	env := pkg.Env{}

	// Establish connection to CockroachDB.
	db, err := sql.Open("postgres", env.GetSessionDBUrl())
	if err != nil {
		log.Println("Error connecting to session store")
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
		token TEXT PRIMARY KEY,
		data BYTEA NOT NULL,
		expiry TIMESTAMPTZ NOT NULL
	);`)
	if err != nil {
		log.Fatal("Error inserting session:", err)
	}

	// Create an index on expiry column
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);`)
	if err != nil {
		log.Fatal("error creating index: ", err)
	}

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = false
	sessionManager.Store = cockroachdbstore.New(db)
	log.Println("Connected to session store successfully")

	return sessionManager
}
