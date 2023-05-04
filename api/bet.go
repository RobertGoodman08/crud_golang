package api

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "venv/database"
    "venv/models"
)



func CreateBet(db *sql.DB) gin.HandlerFunc { // Создание ставки на спорт 
	return func(c *gin.Context) {
		var bet models.Bet
		err := c.BindJSON(&bet)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db, _ := database.Connect()
	
		defer db.Close()

		result, err := db.Exec("INSERT INTO bets (user_id, sport_id, amount) VALUES ($1, $2, $3)", bet.UserID, bet.SportID, bet.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bet creation failed"})
			return
		}

		c.Status(http.StatusOK)
	}
}


func UpdateBet(db *sql.DB) gin.HandlerFunc { // Редактирование ставки
	return func(c *gin.Context) {
		var bet models.Bet
		err := c.BindJSON(&bet)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}


		db, _ := database.Connect()

		defer db.Close()


		_, err = db.Exec("UPDATE bets SET user_id=$1, sport_id=$2, amount=$3 WHERE id=$4",
							bet.UserID, bet.SportID, bet.Amount, bet.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, bet)

	}
}



func DeleteBet(db *sql.DB) gin.HandlerFunc { // Удаление ставки
	return func(c *gin.Context) {
		betID := c.Param("id")

		db, _ := database.Connect()

		defer db.Close()

		_, err := db.Exec("DELETE FROM bets WHERE id=$1", betID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bet delete successfully"})
	}
}