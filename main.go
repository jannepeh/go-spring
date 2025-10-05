package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// In-memory storage with file persistence
var articles []Article
var nextID int = 1
var articlesMutex sync.RWMutex
const dataFile = "articles.gob"

// Initialize database (load from file or create sample data)
func initDatabase() {
	// Try to load existing data
	if err := loadArticles(); err != nil {
		fmt.Println("No existing data found, creating sample articles...")
		createSampleData()
		saveArticles()
	}
	
	fmt.Printf("Database initialized with %d articles!\n", len(articles))
}

// Load articles from file
func loadArticles() error {
	file, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	
	articlesMutex.Lock()
	defer articlesMutex.Unlock()
	
	var data struct {
		Articles []Article
		NextID   int
	}
	
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	
	articles = data.Articles
	nextID = data.NextID
	
	fmt.Println("Articles loaded from file!")
	return nil
}

// Save articles to file
func saveArticles() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	
	articlesMutex.RLock()
	defer articlesMutex.RUnlock()
	
	data := struct {
		Articles []Article
		NextID   int
	}{
		Articles: articles,
		NextID:   nextID,
	}
	
	return encoder.Encode(data)
}

// Create sample data
func createSampleData() {
	articlesMutex.Lock()
	defer articlesMutex.Unlock()
	
	sampleArticles := []Article{
		{
			ID:      1,
			Title:   "Introduction to Go",
			Desc:    "Learn the basics of Go programming language",
			Content: "Go is a statically typed, compiled programming language designed at Google. It's syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.",
			Created: time.Now().Add(-24 * time.Hour),
			Updated: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:      2,
			Title:   "Building REST APIs with Go",
			Desc:    "A comprehensive guide to creating REST APIs in Go",
			Content: "REST APIs are a fundamental part of modern web development. Go provides excellent support for building fast and efficient web services with its built-in net/http package and third-party routers like Gorilla Mux.",
			Created: time.Now().Add(-12 * time.Hour),
			Updated: time.Now().Add(-12 * time.Hour),
		},
		{
			ID:      3,
			Title:   "Database Integration in Go",
			Desc:    "How to connect Go applications with databases",
			Content: "Go supports various databases including SQLite, PostgreSQL, MySQL, MariaDB, and more. This article covers best practices for database integration in Go applications.",
			Created: time.Now().Add(-6 * time.Hour),
			Updated: time.Now().Add(-6 * time.Hour),
		},
	}
	
	articles = sampleArticles
	nextID = 4
}

// GET /articles - Get all articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articlesMutex.RLock()
	defer articlesMutex.RUnlock()

	response := Response{
		Message: "Articles retrieved successfully",
		Data:    articles,
	}

	json.NewEncoder(w).Encode(response)
}

// GET /articles/{id} - Get single article
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	articlesMutex.RLock()
	defer articlesMutex.RUnlock()

	for _, article := range articles {
		if article.ID == id {
			response := Response{
				Message: "Article retrieved successfully",
				Data:    article,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Article not found", http.StatusNotFound)
}

// POST /articles - Create new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if article.Title == "" || article.Desc == "" || article.Content == "" {
		http.Error(w, "Title, description, and content are required", http.StatusBadRequest)
		return
	}

	articlesMutex.Lock()
	defer articlesMutex.Unlock()

	// Set ID and timestamps
	article.ID = nextID
	nextID++
	article.Created = time.Now()
	article.Updated = time.Now()

	// Add to articles slice
	articles = append(articles, article)

	// Save to file
	go func() {
		if err := saveArticles(); err != nil {
			log.Printf("Warning: Failed to save articles: %v", err)
		}
	}()

	response := Response{
		Message: "Article created successfully",
		Data:    article,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// PUT /articles/{id} - Update article
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var updateData Article
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	articlesMutex.Lock()
	defer articlesMutex.Unlock()

	// Find and update the article
	for i, article := range articles {
		if article.ID == id {
			// Update fields if provided
			if updateData.Title != "" {
				articles[i].Title = updateData.Title
			}
			if updateData.Desc != "" {
				articles[i].Desc = updateData.Desc
			}
			if updateData.Content != "" {
				articles[i].Content = updateData.Content
			}
			articles[i].Updated = time.Now()

			// Save to file
			go func() {
				if err := saveArticles(); err != nil {
					log.Printf("Warning: Failed to save articles: %v", err)
				}
			}()

			response := Response{
				Message: "Article updated successfully",
				Data:    articles[i],
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Article not found", http.StatusNotFound)
}

// DELETE /articles/{id} - Delete article
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	articlesMutex.Lock()
	defer articlesMutex.Unlock()

	// Find and delete the article
	for i, article := range articles {
		if article.ID == id {
			// Remove article from slice
			articles = append(articles[:i], articles[i+1:]...)

			// Save to file
			go func() {
				if err := saveArticles(); err != nil {
					log.Printf("Warning: Failed to save articles: %v", err)
				}
			}()

			response := Response{
				Message: "Article deleted successfully",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Article not found", http.StatusNotFound)
}

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to the Go Spring API with persistent file storage! Use /articles for CRUD operations.",
	}
	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", getAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	fmt.Println("Server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("GET    /articles     - Get all articles")
	fmt.Println("GET    /articles/{id} - Get single article")
	fmt.Println("POST   /articles     - Create new article")
	fmt.Println("PUT    /articles/{id} - Update article")
	fmt.Println("DELETE /articles/{id} - Delete article")
	fmt.Println()
	fmt.Printf("Data is persisted to file: %s\n", dataFile)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	// Initialize database
	initDatabase()

	// Start the server
	handleRequests()
}