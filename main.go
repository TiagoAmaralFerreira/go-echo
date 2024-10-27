package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Conexão com o banco de dados MySQL
	connStr := "root:root@tcp(localhost:3306)/finance" // ajuste a senha e o nome do banco de dados
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		e.Logger.Fatal("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Testa a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		e.Logger.Fatal("Error pinging the database:", err)
		return
	}
	fmt.Println("Successfully connected to the database!")

	// Endpoint para retornar a mensagem de sucesso
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Successfully connected to the database!")
	})

	// Endpoint para listar usuários
	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error executing query")
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name, email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				return c.String(http.StatusInternalServerError, "Error scanning row")
			}
			users = append(users, map[string]interface{}{
				"id":    id,
				"name":  name,
				"email": email,
			})
		}

		if err := rows.Err(); err != nil {
			return c.String(http.StatusInternalServerError, "Error iterating over rows")
		}

		return c.JSON(http.StatusOK, users)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }
