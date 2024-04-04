package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"nrdev.se/node-walker/app"
)

//go:embed schema.sql
var dd1 string

func main() {
	ctx := context.Background()

	// Delete old database
	os.Remove("network.db")

	// Create a local sqlite database network.db
	db, err := sql.Open("sqlite3", "network.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	if _, err = db.ExecContext(ctx, dd1); err != nil {
		log.Fatal(err)
	}

	queries := app.New(db)

	CreateDemoLines(ctx, queries)
	CreateMoreLines(ctx, queries)

	// List all lines
	lines, err := queries.GetLines(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
	}

	// Render the network

}

func CreateMoreLines(ctx context.Context, queries *app.Queries) {
	// Insert Pipe C
	line, err := queries.CreateLine(ctx, app.CreateLineParams{
		Name: sql.NullString{String: "Pipe C", Valid: true},
		Type: sql.NullString{String: "Normal Pipe", Valid: true},
		NID:  sql.NullInt64{Int64: 3, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created line: %v\n", line)

	// Insert Connection CA for Pipe C
	node, err := queries.CreateNode(ctx, app.CreateNodeParams{
		Name:   sql.NullString{String: "Connection CA", Valid: true},
		NodeID: sql.NullInt64{Int64: 2, Valid: true},
		NID:    sql.NullInt64{Int64: 3, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created node: %v\n", node)
	// Insert Connection AC for Pipe A
	node, err = queries.CreateNode(ctx, app.CreateNodeParams{
		Name:   sql.NullString{String: "Connection CA", Valid: true},
		NodeID: sql.NullInt64{Int64: 2, Valid: true},
		NID:    sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created node: %v\n", node)
}

func CreateDemoLines(ctx context.Context, queries *app.Queries) {
	// Insert Pipe A
	line, err := queries.CreateLine(ctx, app.CreateLineParams{
		Name: sql.NullString{String: "Pipe A", Valid: true},
		Type: sql.NullString{String: "Normal Pipe", Valid: true},
		NID:  sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created line: %v\n", line)

	// Insert Pipe B
	line, err = queries.CreateLine(ctx, app.CreateLineParams{
		Name: sql.NullString{String: "Pipe B", Valid: true},
		Type: sql.NullString{String: "Normal Pipe", Valid: true},
		NID:  sql.NullInt64{Int64: 2, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created line: %v\n", line)

	// Insert Connection AB for Pipe A
	node, err := queries.CreateNode(ctx, app.CreateNodeParams{
		Name:   sql.NullString{String: "Connection AB", Valid: true},
		NodeID: sql.NullInt64{Int64: 1, Valid: true},
		NID:    sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created node: %v\n", node)

	// Insert Connection BA for Pipe B
	node, err = queries.CreateNode(ctx, app.CreateNodeParams{
		Name:   sql.NullString{String: "Connection BA", Valid: true},
		NodeID: sql.NullInt64{Int64: 1, Valid: true},
		NID:    sql.NullInt64{Int64: 2, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created node: %v\n", node)
}
