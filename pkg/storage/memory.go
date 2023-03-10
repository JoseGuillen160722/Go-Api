package storage

// "database/sql"
// "fmt"
// "log"
// "sync"
// "time"

// _ "github.com/denisenkom/go-mssqldb"
// _ "github.com/lib/pq"

// var (
// 	db   *sql.DB
// 	once sync.Once
// )

// func NewPostgresDB() {
// 	once.Do(func() {
// 		var err error
// 		db, err = sql.Open("postgres", "postgres://ND_dev:pg_dev_nd@192.168.0.14:5400/ND_dev?sslmode=disable")
// 		//db, err = sql.Open("sqlserver", "sqlserver://sian:S14n2017@SRVNDDEV/SRVNDBDDEV?database=Nucleo_RH&encrypt=disable&connection+timeout=30")
// 		if err != nil {
// 			log.Fatalf("Connection failed db: %v", err)
// 		}
// 		// defer db.Close()
// 		if err = db.Ping(); err != nil {
// 			log.Fatalf("Connection failed, no Ping db: %v", err)
// 		}
// 	})

// 	fmt.Println("Conectado a Postgres")
// }

// // Pool retorna una única instancia de DB
// func PostgresPool() *sql.DB {
// 	return db
// }
import (
	_ "github.com/JoseGuillen160722/Go-Api"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type Memory struct {
	CurrentId int
	Persons   map[int]model.Person
}
