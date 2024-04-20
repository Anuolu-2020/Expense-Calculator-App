package middleware

import (
	"context"
	"log"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func InitSession() *scs.SessionManager {
	var sessionManager *scs.SessionManager

	env := pkg.Env{}

	// Establish connection to CockroachDB.
	db, err := pgxpool.New(context.Background(), env.GetSessionDBUrl())
	if err != nil {
		log.Println("Error connecting to session store")
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS sessions (
		token TEXT PRIMARY KEY,
		data BYTEA NOT NULL,
		expiry TIMESTAMPTZ NOT NULL
	);`)
	if err != nil {
		log.Fatal("Error inserting session:", err)
	}

	// Create an index on expiry column
	_, err = db.Exec(
		context.Background(),
		`CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);`,
	)
	if err != nil {
		log.Fatal("error creating index: ", err)
	}

	// Clear db
	/*
		log.Println("Clearing session db...")
		_, err = db.Exec(`DELETE FROM sessions`)
		if err != nil {
			log.Fatal("error clearing db: ", err)
		}
		log.Println("Successfully cleared session db")
	*/

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = false
	sessionManager.Store = pgxstore.New(db)
	log.Println("Connected to session store successfully")

	return sessionManager
}
