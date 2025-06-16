package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/natevaub/focus-companion/backend/db/api/handlers"
	"github.com/natevaub/focus-companion/backend/db/api/routes"
	"github.com/natevaub/focus-companion/backend/db/api/service"
	db "github.com/natevaub/focus-companion/backend/db/generated"
)

// func printAuthors(authors []generated.Author) {
//   for _, author := range authors {
//     bio := "No bio"
//     if author.Bio.Valid {
//       bio = author.Bio.String
//     }
//     fmt.Printf("Author: %s, Bio: %s\n", author.Name, bio)
//   }
// }

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Focus Companion API",
	})

	// Connect to database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	// Initialize services and handlers
	queries := db.New(conn)
	userService := service.NewUserService(queries)
	userHandler := handlers.NewUserHandler(userService)

	// Register routes
	routes.RegisterRoutes(app, userHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// Declare Dummy Data for 3 authors and create authors

	// authorsNames := []string{"John Doe", "Nathan Vaubien", "Jimzer"}
	// authorsBios := []string{"a software engineer", "a software engineer", "is a software engineer"}
	// for i := 0; i < len(authorsNames); i++ {
	//   // Create bio as pgtype.Text
	//   bio := pgtype.Text{
	//     String: authorsBios[i],
	//     Valid:  true,
	//   }

	//   params := generated.CreateAuthorParams{
	//     Name: authorsNames[i],
	//     Bio:  bio,
	//   }

	//   author, err := queries.CreateAuthor(ctx, params)
	//   if err != nil {
	//     fmt.Printf("Error creating author: %v\n", err)
	//     os.Exit(1)
	//   }
	//   fmt.Printf("Created author: %s\n", author.Name)
	// }

	// // List all authors
	// authors, err := queries.ListAuthors(ctx)
	// if err != nil {
	//   fmt.Printf("Error listing authors: %v\n", err)
	//   os.Exit(1)
	// }

	// fmt.Println("\nAll authors in database:")
	// printAuthors(authors)
}
