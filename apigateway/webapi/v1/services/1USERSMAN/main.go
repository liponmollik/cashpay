package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin/v2"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    // Add other fields as needed
}

func main() {
    // Initialize the database connection
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize the Gin router
    router := gin.Default()

    // Grouping mobile endpoints under the specified mobile API gateway and version
    mobAPI := router.Group("/apigateway/mobapi/v1")
    {
        // Example endpoint to get all users for mobile
        mobAPI.GET("/services/users", func(c *gin.Context) {
            var users []User
            rows, err := db.Query("SELECT id, name FROM users")
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            defer rows.Close()
            for rows.Next() {
                var user User
                if err := rows.Scan(&user.ID, &user.Name); err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                    return
                }
                users = append(users, user)
            }
            c.JSON(http.StatusOK, users)
        })

        // Add more mobile endpoints as needed
    }

    // Grouping web endpoints under the specified web API gateway and version
    webAPI := router.Group("/apigateway/webapi/v1")
    {
        // Example endpoint to get all users for web
        webAPI.GET("/services/users", func(c *gin.Context) {
            var users []User
            rows, err := db.Query("SELECT id, name FROM users")
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            defer rows.Close()
            for rows.Next() {
                var user User
                if err := rows.Scan(&user.ID, &user.Name); err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                    return
                }
                users = append(users, user)
            }
            c.JSON(http.StatusOK, users)
        })

        // Add more web endpoints as needed
    }

    // Start the server
    port := ":8080"
    fmt.Printf("Server running on port %s\n", port)
    if err := router.Run(port); err != nil {
        log.Fatal(err)
    }
}
