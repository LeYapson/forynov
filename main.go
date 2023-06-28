package main

import (
    "log"
    "main/database"
    "main/routes"

    "github.com/gofiber/fiber/v2"
)

func main() {
    database.ConnectDb()
    app:= fiber.New()

    setupRoutes(app)
    app.Static("/", "./view")

    log.Fatal(app.Listen(":8080"))
}

func welcome(c * fiber.Ctx) error {
    return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app * fiber.App) {
    //welcome endpoint
    app.Get("/api", welcome)
        //User endpoints
    app.Post("/api/users", routes.CreateUser)
    app.Get("/api/users", routes.GetUsers)
    app.Get("/api/users/:id", routes.GetUser)
    app.Put("/api/users/:id", routes.UpdateUser)
    app.Delete("/api/users/:id", routes.DeleteUser)
        //moderator endpoints
    app.Post("/api/moderator", routes.CreateModerator)
    app.Get("/api/moderator", routes.GetModerator)
    app.Get("/api/moderator/:id", routes.GetModerator)
    app.Put("/api/moderator/:id", routes.UpdateModerator)
    app.Delete("/api/moderator/:id", routes.DeleteModerator)
        //UserProfile endpoints
    app.Post("/api/userProfiles", routes.CreateUserProfile)
    app.Get("/api/userProfiles", routes.GetUserProfiles)
    app.Get("/api/userProfiles/:id", routes.GetUserProfile)
    app.Put("/api/userProfiles/:id", routes.UpdateUserProfile)
    app.Delete("/api/userProfiles/:id", routes.DeleteUserProfile)
        //Subject endpoints
    app.Post("/api/subject", routes.CreateSubject)
    app.Get("/api/subject", routes.GetSubjects)
    app.Get("/api/subject/:id", routes.GetSubject)
    app.Delete("/api/subject/:id", routes.DeleteSubject)
        //message endpoints
    app.Post("/api/message", routes.CreateMessage)
    app.Get("/api/message", routes.GetMessages)
    app.Get("/api/message/:id", routes.GetMessage)
    app.Delete("/api/message/:id", routes.DeleteMessage)
}