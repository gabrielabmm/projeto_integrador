package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID       int
	email    string
	Password string
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}


func checkEmail(email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err 
	}
	return &user, nil
}


func loginUser(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")

	
	user, err := checkEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar o email"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido", "user": user.Email})
}

func main() {
	
	setupDatabase()
	defer db.Close()

	
	router := gin.Default()

	
	router.POST("/login", loginUser)

	
	router.Run(":8080")
}
