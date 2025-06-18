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
	ID    int
	Email string
}

func conectar() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ProjetoIntegrador sslmode=disable"
	return sql.Open("postgres", connStr)
}

// Função para verificar se o email existe no banco de dados
func checkEmail(email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Email não encontrado
		}
		return nil, err // Outro erro
	}
	return &user, nil
}

// Função para atualizar a senha do usuário
func updatePassword(email, password string) error {
	_, err := db.Exec("UPDATE users SET password = $1 WHERE email = $2", password, email)
	return err
}

func main() {
	// Configuração do banco de dados
	setupDatabase()
	defer db.Close()

	// Inicialização do servidor Gin
	router := gin.Default()

	// Rota para verificar email e redirecionar para a tela de definir senha
	router.POST("/verificar_email", func(c *gin.Context) {
		email := c.DefaultPostForm("email", "")

		user, err := checkEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar o email"})
			return
		}

		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email não encontrado"})
			return
		}

		// Redireciona para a tela de definir senha
		c.HTML(http.StatusOK, "primeiroacesso.html", gin.H{
			"email": email,
		})
	})

	// Rota para definir a nova senha
	router.POST("/definir_senha", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Atualiza a senha do usuário
		err := updatePassword(email, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a senha"})
			return
		}

		// Redireciona para a página de login ou outra página de sucesso
		c.JSON(http.StatusOK, gin.H{"message": "Senha atualizada com sucesso"})
	})

	// Configuração da template (HTML)
	router.LoadHTMLFiles("primeiroacesso.html")

	// Inicialização do servidor
	router.Run(":8080")
}
