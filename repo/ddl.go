package repo

import (
	"goconsole/config"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //postgres drivers for initialization
)

//Ddl ...
type Ddl struct {
}

//GetAllNamespace ...
func (r *Ddl) GetAllNamespace() ([]string, error) {

	db, err := sql.Open("postgres", config.Data.ConnectionString)
	if err != nil {
		log.Println("Error opening sql connection", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT nspname FROM pg_catalog.pg_namespace ORDER BY nspname")

	if err != nil {
		log.Fatalf("Error fetching namespace list: %s", err)
		return nil, err
	}

	var result []string

	for rows.Next() {
		var nspname string
		if err := rows.Scan(&nspname); err != nil {
			log.Fatal(err)
		}
		result = append(result, nspname)
	}

	return result, nil
}

//GetAllTablesIn ...
func (r *Ddl) GetAllTablesIn(namespace string, limit int, offset int) ([]string, error) {

	db, err := sql.Open("postgres", config.Data.ConnectionString)
	if err != nil {
		log.Println("Error opening sql connection", err)
		return nil, err
	}
	defer db.Close()

	log.Println("namespace==",namespace)

	log.Println(fmt.Sprintf(`SELECT table_name
	FROM information_schema.tables
	WHERE table_schema='%s'
	ORDER BY table_name`, namespace))

	rows, err := db.Query(fmt.Sprintf(`SELECT table_name
							FROM information_schema.tables
							WHERE table_schema='%s'
							ORDER BY table_name
							LIMIT %d
							OFFSET %d`, namespace, limit, offset))

	if err != nil {
		log.Fatalln("Error fetching tables list from namespace", namespace, "\nError: ", err)
	}

	var result []string

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		result = append(result, tableName)
	}

	return result, nil
}

//Ping ...  Ping the DB to verify if the able to connect to the db
func Ping() {

	db, err := sql.Open("postgres", config.Data.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully able to connect!")
	}

}
