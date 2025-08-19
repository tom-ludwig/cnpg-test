package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// db, err := connectWithPassword("cluster-example-rw.default.svc.cluster.local", "5432", "app", "app", "eo5jRy7HpTtZ5axMfSdTnKAe3XXLdR7h6BB3aK1JRVbuwEQlhQ0HUQgU9nPNXyQW", "/certs/ca.crt")
	// if err != nil {
	// 	log.Fatalf("Password auth failed: %v", err)
	// }

	db, err := connectWithMTLS(
		getEnv("PG_HOST", "cluster-example-rw.default.svc"),
		getEnv("PG_PORT", "5432"),
		getEnv("PG_DB", "app"),
		getEnv("PG_USER", "app"),
		getEnv("PG_SSLMODE", "verify-full"),
		getEnv("PG_CLIENT_CERT", "/certs/tls.crt"),
		getEnv("PG_CLIENT_KEY", "/certs/tls.key"),
		getEnv("PG_SSLROOTCERT", "/certs/ca.crt"),
	)
	if err != nil {
		log.Fatalf("mTLS auth failed: %v", err)
	}

	defer db.Close()

	log.Println("Connected to Postgres over TLS")

	setupHTTP(db)
}

func connectWithPassword(host, port, dbname, user, password, rootCert string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=verify-full sslrootcert=%s",
		host, port, dbname, user, password, rootCert,
	)
	return sqlx.Connect("postgres", dsn)
}

func connectWithMTLS(host, port, dbname, user, sslmode, clientCert, clientKey, rootCert string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s",
		host, port, dbname, user, sslmode, clientCert, clientKey, rootCert,
	)
	return sqlx.Connect("postgres", dsn)
}

func setupHTTP(db *sqlx.DB) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "DB not reachable", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Ok"))
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		var now time.Time
		if err := db.Get(&now, "SELECT NOW()"); err != nil {
			http.Error(w, "DB query failed", http.StatusInternalServerError)
			return
		}

		bytes, err := fmt.Fprintf(w, "DB time: %s", now.String())
		if err != nil {
			log.Printf("Error while returning date (%d): %s", bytes, err)
		}
	})

	port := getEnv("PORT", "8080")
	log.Printf("Server listening on :%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
