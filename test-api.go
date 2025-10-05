package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	baseURL := "http://localhost:8080"
	
	fmt.Println("🚀 Testing Go Spring CRUD API")
	fmt.Println("===============================")
	
	// Wait a moment for server to be ready
	time.Sleep(2 * time.Second)
	
	// Test 1: GET all articles
	fmt.Println("\n1️⃣ Testing GET /articles (Get all articles)")
	resp, err := http.Get(baseURL + "/articles")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	// Test 2: GET single article
	fmt.Println("\n2️⃣ Testing GET /articles/1 (Get single article)")
	resp, err = http.Get(baseURL + "/articles/1")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	// Test 3: POST new article
	fmt.Println("\n3️⃣ Testing POST /articles (Create new article)")
	newArticle := Article{
		Title:   "Testing CRUD Operations",
		Desc:    "A guide to testing REST APIs",
		Content: "This article demonstrates how to test CRUD operations in a REST API built with Go.",
	}
	
	jsonData, err := json.Marshal(newArticle)
	if err != nil {
		fmt.Printf("❌ Error marshaling JSON: %v\n", err)
		return
	}
	
	resp, err = http.Post(baseURL+"/articles", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	// Parse the response to get the created article ID
	var createResponse Response
	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		fmt.Printf("❌ Error parsing response: %v\n", err)
		return
	}
	
	// Get the created article ID
	articleData, ok := createResponse.Data.(map[string]interface{})
	if !ok {
		fmt.Println("❌ Error getting article data")
		return
	}
	
	articleID := int(articleData["id"].(float64))
	
	// Test 4: PUT update article
	fmt.Printf("\n4️⃣ Testing PUT /articles/%d (Update article)\n", articleID)
	updateArticle := Article{
		Title:   "Updated: Testing CRUD Operations",
		Desc:    "An updated guide to testing REST APIs",
		Content: "This article has been updated to demonstrate how to test CRUD operations in a REST API built with Go.",
	}
	
	jsonData, err = json.Marshal(updateArticle)
	if err != nil {
		fmt.Printf("❌ Error marshaling JSON: %v\n", err)
		return
	}
	
	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/articles/%d", baseURL, articleID), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	// Test 5: DELETE article
	fmt.Printf("\n5️⃣ Testing DELETE /articles/%d (Delete article)\n", articleID)
	req, err = http.NewRequest("DELETE", fmt.Sprintf("%s/articles/%d", baseURL, articleID), nil)
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}
	
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	// Test 6: GET all articles again to see final state
	fmt.Println("\n6️⃣ Final state - GET /articles (Get all articles)")
	resp, err = http.Get(baseURL + "/articles")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Response: %s\n", string(body))
	
	fmt.Println("\n🎉 CRUD API testing completed!")
	fmt.Println("📁 Data persisted to: articles.gob")
}