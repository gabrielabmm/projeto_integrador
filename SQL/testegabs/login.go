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

// Estrutura para armazenar as informações de um usuário
type User struct {
	ID       int
	email    string
	Password string
}

// Função para conectar ao banco de dados PostgreSQL
func setupDatabase() {
	var err error
	connStr := "postgres://seu_usuario:sua_senha@localhost/seu_banco?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
}

// Função para verificar se o email existe no banco de dados
func checkEmail(email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Email não encontrado
		}
		return nil, err // Outro erro
	}
	return &user, nil
}

// Função para realizar login
func loginUser(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")

	// Verifica se o email existe no banco de dados
	user, err := checkEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar o email"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	// Verifica se a senha fornecida corresponde à senha armazenada
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	// Se o login for bem-sucedido, redireciona para o painel de controle ou outra página interna
	c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido", "user": user.Email})
}

func main() {
	// Configuração do banco de dados
	setupDatabase()
	defer db.Close()

	// Inicialização do servidor Gin
	router := gin.Default()

	// Rota para a tela de login
	router.POST("/login", loginUser)

	// Inicialização do servidor
	router.Run(":8080")
}
