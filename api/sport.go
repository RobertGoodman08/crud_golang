package api

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "venv/database"
    "venv/models"
)



func CreateSport(db *sql.DB) gin.HandlerFunc { // Создание спортивного матча
	return func(c *gin.Context) {
		var sport models.Sport
		err := c.BindJSON(&sport)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db, _ := database.Connect()
	
		defer db.Close()

		var sportID int
		err = db.QueryRow("INSERT INTO sports (name) VALUES ($1) RETURNING id", sport.Name).Scan(&sportID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sport.ID = sportID
		c.JSON(http.StatusOK, sport)
	}
}


func UpdateSport(db *sql.DB) gin.HandlerFunc { // Обновление спортивного матча
	return func(c *gin.Context) {
		var sport models.Sport
		err := c.BindJSON(&sport)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db, _ := database.Connect()
	
		defer db.Close()

		_, err = db.Exec("UPDATE sports SET name=$1 WHERE id=$2", sport.Name, sport.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sport)
	}
}



func DeleteSport(db *sql.DB) gin.HandlerFunc { // Удаление спортивного матча
	return func(c *gin.Context) {
		sportID := c.Param("id")

		db, _ := database.Connect()
	
		defer db.Close()

		_, err := db.Exec("DELETE FROM sports WHERE id=$1", sportID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sport deleted successfully"})
	}
}





func IndexSport(db *sql.DB) gin.HandlerFunc { // Получение-полный список матчей спортивных
	return func(c *gin.Context){
		db, _ := database.Connect()
	
		defer db.Close()

		rows, err := db.Query("SELECT id, name FROM sports")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		sports := []models.Sport{}
		for rows.Next() {
			var sport models.Sport
			err := rows.Scan(&sport.ID, &sport.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			sports = append(sports, sport)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sports)
	}
}


func SearchSportByName(db *sql.DB) gin.HandlerFunc { // Поиск матчей
	return func(c *gin.Context) {
		name := c.Query("name")

		db, _ := database.Connect()
	
		defer db.Close()

		rows, err := db.Query("SELECT id, name FROM sports WHERE name LIKE $1", "%"+name+"%")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sports := []models.Sport{}
		for rows.Next() {
			var sport models.Sport
			err := rows.Scan(&sport.ID, &sport.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			sports = append(sports, sport)
		}

		c.JSON(http.StatusOK, sports)
	}
}