package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Article struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// In-memory storage for articles
var articles []Article
var nextID int = 1

// CRUD Operations

// CREATE - Add new article
func createArticle(title, desc, content string) Article {
	article := Article{
		ID:      nextID,
		Title:   title,
		Desc:    desc,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
	nextID++
	articles = append(articles, article)
	fmt.Printf("✅ Created article: %s (ID: %d)\n", title, article.ID)
	return article
}

// READ - Get all articles
func getAllArticles() []Article {
	fmt.Printf("📚 Retrieved %d articles\n", len(articles))
	return articles
}

// READ - Get single article by ID
func getArticleByID(id int) *Article {
	for _, article := range articles {
		if article.ID == id {
			fmt.Printf("📖 Found article: %s (ID: %d)\n", article.Title, id)
			return &article
		}
	}
	fmt.Printf("❌ Article with ID %d not found\n", id)
	return nil
}

// UPDATE - Update article by ID
func updateArticle(id int, title, desc, content string) *Article {
	for i, article := range articles {
		if article.ID == id {
			if title != "" {
				articles[i].Title = title
			}
			if desc != "" {
				articles[i].Desc = desc
			}
			if content != "" {
				articles[i].Content = content
			}
			articles[i].Updated = time.Now()
			fmt.Printf("✏️ Updated article: %s (ID: %d)\n", articles[i].Title, id)
			return &articles[i]
		}
	}
	fmt.Printf("❌ Article with ID %d not found for update\n", id)
	return nil
}

// DELETE - Delete article by ID
func deleteArticle(id int) bool {
	for i, article := range articles {
		if article.ID == id {
			fmt.Printf("🗑️ Deleted article: %s (ID: %d)\n", article.Title, id)
			articles = append(articles[:i], articles[i+1:]...)
			return true
		}
	}
	fmt.Printf("❌ Article with ID %d not found for deletion\n", id)
	return false
}

// Pretty print articles
func printAllArticles() {
	fmt.Println("\n📚 ALL ARTICLES:")
	fmt.Println("================")
	for _, article := range articles {
		fmt.Printf("ID: %d\n", article.ID)
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Description: %s\n", article.Desc)
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Printf("Created: %s\n", article.Created.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated: %s\n", article.Updated.Format("2006-01-02 15:04:05"))
		fmt.Println("---")
	}
	fmt.Println()
}

func main() {
	fmt.Println("🚀 Go Spring CRUD Demo - Article Management System")
	fmt.Println("==================================================")

	// Initialize with sample data
	fmt.Println("\n1️⃣ CREATING SAMPLE ARTICLES:")
	createArticle(
		"Introduction to Go",
		"Learn the basics of Go programming language",
		"Go is a statically typed, compiled programming language designed at Google.",
	)
	
	createArticle(
		"Building REST APIs",
		"A guide to creating REST APIs in Go",
		"REST APIs are fundamental in modern web development. Go provides excellent support for building web services.",
	)

	createArticle(
		"Database Integration",
		"How to connect Go applications with databases",
		"Go supports various databases including MongoDB, PostgreSQL, MySQL, and more.",
	)

	// READ - Get all articles
	fmt.Println("\n2️⃣ READING ALL ARTICLES:")
	getAllArticles()
	printAllArticles()

	// READ - Get single article
	fmt.Println("3️⃣ READING SINGLE ARTICLE:")
	article := getArticleByID(2)
	if article != nil {
		data, _ := json.MarshalIndent(article, "", "  ")
		fmt.Printf("Article JSON:\n%s\n\n", data)
	}

	// CREATE - Add a new article
	fmt.Println("4️⃣ CREATING NEW ARTICLE:")
	createArticle(
		"Advanced Go Concepts",
		"Exploring advanced features of Go",
		"This article covers goroutines, channels, interfaces, and more advanced Go concepts.",
	)

	// UPDATE - Update an existing article
	fmt.Println("\n5️⃣ UPDATING ARTICLE:")
	updateArticle(1, "Introduction to Go Programming", "", "Go (also known as Golang) is a modern programming language that makes it easy to build simple, reliable, and efficient software.")

	// DELETE - Delete an article
	fmt.Println("\n6️⃣ DELETING ARTICLE:")
	deleteSuccess := deleteArticle(3)
	fmt.Printf("Deletion successful: %v\n", deleteSuccess)

	// Final state
	fmt.Println("\n7️⃣ FINAL STATE - ALL REMAINING ARTICLES:")
	getAllArticles()
	printAllArticles()

	fmt.Println("✨ Demo completed! This shows how CRUD operations work.")
	fmt.Println("🔗 The same logic is used in the REST API with HTTP endpoints.")
	
	fmt.Println("\n🌐 To test the REST API:")
	fmt.Println("1. Make sure MongoDB is running (for the full version)")
	fmt.Println("2. Run: go run main-simple.go (for in-memory version)")
	fmt.Println("3. Use the API endpoints with tools like curl or Postman")
}