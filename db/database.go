package db

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"os"
)

var g_db *sql.DB

func InitDb() error {
	var connectionString string
	if os.Getenv("GIN_MODE") == "release" {
		const (
			// Initialize connection constants.
			HOST     = "localhost"
			PORT     = 5432
			DATABASE = "postgres"
			USER     = "postgres"
			PASSWORD = "postgres1234"
		)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=required",
			HOST, PORT, USER, PASSWORD, DATABASE)
	} else {
		const (
			// Initialize connection constants.
			HOST     = "localhost"
			PORT     = 5432
			DATABASE = "postgres"
			USER     = "postgres"
			PASSWORD = "postgres1234"
		)

		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			HOST, PORT, USER, PASSWORD, DATABASE)
	}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	g_db = db
	fmt.Println("Connect to database success!")

	err = createTable()
	if err != nil {
		return err
	}
	fmt.Println("Create Table success!")
	return nil
}

func createType() error {
	ret, err := g_db.Exec("CREATE TYPE  (id serial PRIMARY KEY, datas double precision[][]);")

	return nil
}

func createTable() error {
	// Drop previous table of same name if one exists.
	ret, err := g_db.Exec("DROP TABLE IF EXISTS inventory;")
	if err != nil {
		return err
	}
	fmt.Println(ret)

	// Create table.
	ret, err = g_db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, datas double precision[][]);")
	if err != nil {
		return err
	}
	fmt.Println(ret)

	datas := [][]float64{
		{0.1, 0.2, 0.3, 0.4, 0.5, 0.6},
		{1.1, 1.2, 0.3, 0.4, 0.5, 0.6},
		{2.1, 2.2, 0.3, 0.4, 0.5, 0.6},
		{3.1, 3.2, 0.3, 0.4, 0.5, 0.6},
		{4.1, 4.2, 0.3, 0.4, 0.5, 0.6},
		{5.1, 5.2, 0.3, 0.4, 0.5, 0.6},
		{6.1, 6.2, 0.3, 0.4, 0.5, 0.6},
		{6.1, 6.2, 0.3, 0.4, 0.5, 0.6},
		{6.1, 6.2, 0.3, 0.4, 0.5, 0.6},
	}

	// Insert some data into table.
	var id int
	sql_statement := "INSERT INTO inventory (datas) VALUES ($1) RETURNING id"
	err = g_db.QueryRow(sql_statement, pq.Array(datas)).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println(id)

	for i, array := range datas {
		for j, v := range array {
			datas[i][j] = v + float64(j+i)
		}
	}

	err = g_db.QueryRow(sql_statement, pq.Array(datas)).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println(id)

	// Read rows from table.
	sql_statement = "SELECT * from inventory;"
	rows, err := g_db.Query(sql_statement)
	if err != nil {
		return err
	}

	for rows.Next() {
		var buffer [9][6]float64
		switch err := rows.Scan(&id, pq.Array(&buffer)); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Println(id, len(buffer))
		default:
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func arrayToSqlValue(datas [][]float64) string {
	value := "{"
	for i, array := range datas {
		value += "{"
		for j, v := range array {
			if j == len(array)-1 {
				value += fmt.Sprintf("%f", v)
			} else {
				value += fmt.Sprintf("%f,", v)
			}
		}

		if i == len(datas)-1 {
			value += "}"
		} else {
			value += "},"
		}
	}
	value += "}"

	return value
}
