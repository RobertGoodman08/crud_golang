package api

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "venv/database"
    "venv/models"
)


func Register(db *sql.DB) gin.HandlerFunc { // Регистрация пользователя
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db, _ := database.Connect()
	
		defer db.Close()

		var exists bool
		err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}

		_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}

func Login(db *sql.DB) gin.HandlerFunc { // Авторизация пользователя
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		db, _ := database.Connect()
		
		defer db.Close()
		
		var storedPassword string
		err = db.QueryRow("SELECT password FROM users WHERE email = $1", user.Email).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user.Password != storedPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}

		c.Status(http.StatusOK)
	}
}


