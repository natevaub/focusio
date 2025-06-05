package main

import (
  "context"
  "fmt"
  "os"
  "github.com/jackc/pgx/v5"
  "github.com/jackc/pgx/v5/pgtype"
  "github.com/natevaub/focus-companion/backend/sqlc/generated"
)

func printAuthors(authors []db.Author) {
  for _, author := range authors {
    bio := "No bio"
    if author.Bio.Valid {
      bio = author.Bio.String
    }
    fmt.Printf("Author: %s, Bio: %s\n", author.Name, bio)
  }
}

func main() {
      // app := fiber.New()

      //   app.Get("/ping", func(c *fiber.Ctx) error {
      //           return c.SendString("pong")
      //             })

      //     app.Listen(":8080")

      ctx := context.Background()

      conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/db?sslmode=disable")
      if err != nil {
        fmt.Printf("Unable to connect to database: %v\n", err)
        os.Exit(1)
      }
      defer conn.Close(ctx)

      queries := generated.New(conn)

      // Declare Dummy Data for 3 authors and create authors

      authorsNames := []string{"John Doe", "Nathan Vaubien", "Jimzer"}
      authorsBios := []string{"a software engineer", "a software engineer", "is a software engineer"}
      for i := 0; i < len(authorsNames); i++ {
        // Create bio as pgtype.Text
        bio := pgtype.Text{
          String: authorsBios[i],
          Valid:  true,
        }
        
        params := generated.CreateAuthorParams{
          Name: authorsNames[i],
          Bio:  bio,
        }
        
        author, err := queries.CreateAuthor(ctx, params)
        if err != nil {
          fmt.Printf("Error creating author: %v\n", err)
          os.Exit(1)
        }
        fmt.Printf("Created author: %s\n", author.Name)
      }

      // List all authors
      authors, err := queries.ListAuthors(ctx)
      if err != nil {
        fmt.Printf("Error listing authors: %v\n", err)
        os.Exit(1)
      }

      fmt.Println("\nAll authors in database:")
      printAuthors(authors)
}
