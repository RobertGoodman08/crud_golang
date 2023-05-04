package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "venv/database"
    "venv/api"
)

func main() {
    // Загрузка переменных окружения из файла .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err.Error())
    }

    // Подключение к базе данных
    _, err = database.Connect()
    if err != nil {
        log.Fatalf("Error connecting to database: %s", err.Error())
    }

    
    router := gin.Default()

    db := database.DB

    router.GET("/", api.IndexSport(db))  // Получение-полный список матчей спортивных

    router.POST("/register", api.Register(db)) // Регистрация
    router.POST("/login", api.Login(db)) // Авторизация 

    router.POST("/createsport", api.CreateSport(db)) // Создание спортивного матча
    router.GET("/sport/search", api.SearchSportByName(db)) // Поиск матчей
    router.PUT("/updatesports/:id", api.UpdateSport(db)) // Редактирование матча
    router.DELETE("/deletesport/:id", api.DeleteSport(db)) // Удаление матча


    // Ставки
    router.POST("/createbet", api.CreateBet(db)) // Создание ставки на спорт 
    router.PUT("/updatebet/:id", api.UpdateBet(db)) // Редактирование ставки
    router.DELETE("/deletebet/:id", api.DeleteBet(db)) // Удаление ставки




    // Запуск сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server started on port %s", port)
    err = http.ListenAndServe(":"+port, router)
    if err != nil {
        log.Fatalf("Error starting server: %s", err.Error())
    }
}
