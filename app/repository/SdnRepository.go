package repository

import (
	"database/sql"
	"fmt"
	"golangexample/component"
	"golangexample/entity"
	"strings"
)

func Init() {
	conn := component.DbConnection()
	tableStatement := `CREATE TABLE IF NOT EXISTS sdn (
		id SERIAL PRIMARY KEY,
		uid INT,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		sdn_type VARCHAR(255)
	)`
	indexStatement := `
	DO
	$$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_sdn_uid') THEN
				CREATE UNIQUE INDEX idx_sdn_uid ON sdn(uid);
			END IF;
		END
	$$
`

	result, err := conn.Exec(tableStatement)

	if err != nil {
		panic(err)
	}

	result, err = conn.Exec(indexStatement)

	if err != nil {
		panic(err)
	}

	println(result)
}

func Insert(conn *sql.DB, sdnEntity entity.SdnEntity) sql.Result {
	statement := `INSERT INTO sdn (uid, first_name, last_name, sdn_type) VALUES ($1, $2, $3, $4) ON CONFLICT (uid) DO NOTHING`
	result, err := conn.Exec(statement, sdnEntity.Uid, sdnEntity.FirstName, sdnEntity.LastName, sdnEntity.SdnType)

	if err != nil {
		panic(err.Error())
	}

	return result
}

func SearchStrong(conn *sql.DB, name string) []entity.SdnEntity {
	var rows *sql.Rows
	var err error

	nameParts := strings.Split(name, " ")

	if len(nameParts) > 1 {
		statement := `SELECT uid, first_name, last_name, sdn_type FROM sdn WHERE LOWER(first_name) = LOWER($1) AND LOWER(last_name) = LOWER($2)`
		rows, err = conn.Query(statement, nameParts[0], nameParts[1])
	} else if len(nameParts) == 1 {
		statement := `SELECT uid, first_name, last_name, sdn_type FROM sdn WHERE LOWER(first_name) = LOWER($1) OR LOWER(last_name) = LOWER($1)`
		rows, err = conn.Query(statement, nameParts[0])
	}
	collection := []entity.SdnEntity{}

	if err != nil {
		fmt.Println(err.Error())
	}

	if rows != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var uid int
		var first_name string
		var last_name string
		var sdn_type string

		err = rows.Scan(&uid, &first_name, &last_name, &sdn_type)
		if err != nil {
			println(fmt.Errorf("Fetch result failed: %s", err))
		}

		collection = append(collection, entity.SdnEntity{
			Uid:       uid,
			FirstName: first_name,
			LastName:  last_name,
			SdnType:   sdn_type,
		})
	}

	return collection
}

func SearchWeak(conn *sql.DB, name string) []entity.SdnEntity {
	var rows *sql.Rows
	var err error

	nameParts := strings.Split(name, " ")

	if len(nameParts) > 1 {
		statement := `SELECT uid, first_name, last_name, sdn_type FROM sdn WHERE first_name ILIKE $1 OR last_name ILIKE $2`
		rows, err = conn.Query(statement, "%"+nameParts[0]+"%", "%"+nameParts[1]+"%")
	} else if len(nameParts) == 1 {
		statement := `SELECT uid, first_name, last_name, sdn_type FROM sdn WHERE first_name ILIKE $1 OR last_name ILIKE $1`
		rows, err = conn.Query(statement, "%"+nameParts[0]+"%")
	}
	collection := []entity.SdnEntity{}

	if err != nil {
		fmt.Println(err.Error())
	}

	if rows != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var uid int
		var first_name string
		var last_name string
		var sdn_type string

		err = rows.Scan(&uid, &first_name, &last_name, &sdn_type)
		if err != nil {
			println(fmt.Errorf("Fetch result failed: %s", err))
		}

		collection = append(collection, entity.SdnEntity{
			Uid:       uid,
			FirstName: first_name,
			LastName:  last_name,
			SdnType:   sdn_type,
		})
	}

	return collection
}
