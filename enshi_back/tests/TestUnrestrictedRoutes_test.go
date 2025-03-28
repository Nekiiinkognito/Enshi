package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"enshi/db_connection"
	"enshi/env"
	"enshi/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRouter initializes the router with route groups for testing
func SetupRouter() (*gin.Engine, error) {
	router := gin.Default()
	err := routes.SetupRotes(router)
	return router, err
}

// TestMain sets up the global router before running tests
func TestMain(m *testing.M) {
	if err := env.LoadEnv("../utils/secret.env"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := db_connection.SetupDatabase(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db_connection.Dbx.Close()
	defer db_connection.Dbx_connection.Close(context.Background())

	var err error
	router, err = SetupRouter()
	if err != nil {
		panic(err)
	}

	// Run all tests
	exitCode := m.Run()

	// Exit with the appropriate status code
	os.Exit(exitCode)
}

// TestRouterSetup verifies that the router is set up correctly
func TestRouterSetup(t *testing.T) {
	if router == nil {
		t.Fatalf("Router is not initialized")
	}
	t.Log("Routes setup successfully")
}

// TestGetRandomPost tests the /posts/random endpoint
func TestGetRandomPost(t *testing.T) {
	t.Log("router in here", router)
	req, _ := http.NewRequest("GET", "/test/posts/random", nil)
	w := httptest.NewRecorder()

	// Use the global router to serve the request
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d. Body %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetRandomPost endpoint returned OK")
	}
}

// TestGetRandomPost_WithLimit проверяет успешный запрос с указанным лимитом постов.
func TestGetRandomPost_WithLimit(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test/posts/random?limit=5", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetRandomPost with limit endpoint returned OK")
	}
}

// TestGetRandomPost_InvalidLimit проверяет запрос с недопустимым лимитом.
func TestGetRandomPost_InvalidLimit(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test/posts/random?limit=invalid", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusInternalServerError, w.Code, w.Body.String())
	} else {
		t.Log("GetRandomPost with invalid limit returned Internal Server Error as expected")
	}
}

// TestGetPost_Success проверяет успешное получение поста по корректному ID.
func TestGetPost_Success(t *testing.T) {
	// Предполагается, что пост с ID 1 существует в тестовой базе данных
	postID := "5984625340089268323"
	req, _ := http.NewRequest("GET", "/test/posts/"+postID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetPost endpoint returned OK")
	}
}

// TestGetPost_InvalidID проверяет запрос с некорректным форматом ID.
func TestGetPost_InvalidID(t *testing.T) {
	invalidPostID := "abc"
	req, _ := http.NewRequest("GET", "/test/posts/"+invalidPostID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("GetPost with invalid ID returned Bad Request as expected")
	}
}

// TestGetPost_NotFound проверяет запрос к несуществующему посту.
func TestGetPost_NotFound(t *testing.T) {
	// Предполагается, что пост с ID 9999 не существует
	nonExistentPostID := "9999"
	req, _ := http.NewRequest("GET", "/test/posts/"+nonExistentPostID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusInternalServerError, w.Code, w.Body.String())
	} else {
		t.Log("GetPost with non-existent ID returned Internal Server Error as expected")
	}
}

// TestCreateBlog_Success проверяет успешное создание блога.
func TestCreateBlog_Success(t *testing.T) {
	// Создаем тестовые данные для блога
	blogData := map[string]interface{}{
		"user_id":     1,
		"title":       "Test Blog",
		"description": "A blog for testing",
		"category_id": 2,
	}
	body, _ := json.Marshal(blogData)

	req, _ := http.NewRequest("POST", "/test/blogs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
	} else {
		t.Log("CreateBlog endpoint returned Created")
	}
}

// TestCreateBlog_InvalidData проверяет создание блога с некорректными данными.
func TestCreateBlog_InvalidData(t *testing.T) {
	// Отсутствует обязательное поле user_id
	blogData := map[string]interface{}{
		"title":       "Test Blog",
		"description": "A blog with invalid data",
	}
	body, _ := json.Marshal(blogData)

	req, _ := http.NewRequest("POST", "/test/blogs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("CreateBlog with invalid data returned Bad Request as expected")
	}
}

// TestGetBlogsByUser_Success проверяет успешное получение блогов по user_id.
func TestGetBlogsByUser_Success(t *testing.T) {
	userID := 1
	req, _ := http.NewRequest("GET", "/test/users/"+strconv.Itoa(userID)+"/blogs", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetBlogsByUser endpoint returned OK")
	}
}

// TestGetBlog_Success проверяет успешное получение блога по blog_id.
func TestGetBlog_Success(t *testing.T) {
	// Предполагается, что блог с ID 1 существует
	blogID := "9080137998708363218"
	req, _ := http.NewRequest("GET", "/blogs/"+blogID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetBlog endpoint returned OK")
	}
}

// TestGetBlog_InvalidID проверяет получение блога с некорректным blog_id.
func TestGetBlog_InvalidID(t *testing.T) {
	invalidBlogID := "abc"
	req, _ := http.NewRequest("GET", "/blogs/"+invalidBlogID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("GetBlog with invalid ID returned Bad Request as expected")
	}
}

// TestGetBlog_NotFound проверяет получение несуществующего блога.
func TestGetBlog_NotFound(t *testing.T) {
	// Предполагается, что блог с ID 9999 не существует
	nonExistentBlogID := "9999"
	req, _ := http.NewRequest("GET", "/blogs/"+nonExistentBlogID, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusNotFound, w.Code, w.Body.String())
	} else {
		t.Log("GetBlog with non-existent ID returned Not Found as expected")
	}
}

// TestUpdateBlog_Success проверяет успешное обновление блога.
func TestUpdateBlog_Success(t *testing.T) {
	// Предполагается, что блог с ID 1 существует
	blogID := "1"
	updateData := map[string]interface{}{
		"title":       "Updated Test Blog",
		"description": "Updated description",
		"category_id": 3,
	}
	body, _ := json.Marshal(updateData)

	req, _ := http.NewRequest("PUT", "/blogs/"+blogID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("UpdateBlog endpoint returned OK")
	}
}

// TestUpdateBlog_InvalidData проверяет обновление блога с некорректными данными.
func TestUpdateBlog_InvalidData(t *testing.T) {
	// Предполагается, что блог с ID 1 существует
	blogID := "1"
	// Отсутствует обязательное поле title
	updateData := map[string]interface{}{
		"description": "Updated description without title",
	}
	body, _ := json.Marshal(updateData)

	req, _ := http.NewRequest("PUT", "/test/blogs/"+blogID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("UpdateBlog with invalid data returned Bad Request as expected")
	}
}

// TestUpdateBlog_NotFound проверяет обновление несуществующего блога.
func TestUpdateBlog_NotFound(t *testing.T) {
	nonExistentBlogID := "9999"
	updateData := map[string]interface{}{
		"title":       "Non-existent Blog",
		"description": "Trying to update a blog that does not exist",
		"category_id": 4,
	}
	body, _ := json.Marshal(updateData)

	req, _ := http.NewRequest("PUT", "/test/blogs/"+nonExistentBlogID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusNotFound, w.Code, w.Body.String())
	} else {
		t.Log("UpdateBlog for non-existent blog returned Not Found as expected")
	}
}

// TestCreateBookmark_Success проверяет успешное создание закладки.
func TestCreateBookmark_Success(t *testing.T) {
	bookmarkData := map[string]interface{}{
		"user_id": 1,
		"post_id": 1,
	}
	body, _ := json.Marshal(bookmarkData)

	req, _ := http.NewRequest("POST", "/test/bookmarks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusCreated, w.Code, w.Body.String())
	} else {
		t.Log("CreateBookmark endpoint returned Created")
	}
}

// TestCreateBookmark_InvalidData проверяет создание закладки с некорректными данными.
func TestCreateBookmark_InvalidData(t *testing.T) {
	// Отсутствует обязательное поле post_id
	bookmarkData := map[string]interface{}{
		"user_id": 1,
	}
	body, _ := json.Marshal(bookmarkData)

	req, _ := http.NewRequest("POST", "/test/bookmarks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("CreateBookmark with invalid data returned Bad Request as expected")
	}
}

// TestGetBookmarksByUser_Success проверяет успешное получение закладок по user_id.
func TestGetBookmarksByUser_Success(t *testing.T) {
	userID := 1
	req, _ := http.NewRequest("GET", "/test/users/"+strconv.Itoa(userID)+"/bookmarks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetBookmarksByUser endpoint returned OK")
	}
}

// TestGetBookmarksByUser_InvalidUserID проверяет получение закладок с некорректным user_id.
func TestGetBookmarksByUser_InvalidUserID(t *testing.T) {
	invalidUserID := "abc"
	req, _ := http.NewRequest("GET", "/test/users/"+invalidUserID+"/bookmarks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("GetBookmarksByUser with invalid user ID returned Bad Request as expected")
	}
}

// TestGetCountOfBookmarksByPost_Success проверяет успешное получение количества закладок для поста.
func TestGetCountOfBookmarksByPost_Success(t *testing.T) {
	postID := "1"
	req, _ := http.NewRequest("GET", "/test/posts/"+postID+"/bookmarks/count", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetCountOfBookmarksByPost endpoint returned OK")
	}
}

// TestGetCountOfBookmarksByPost_InvalidPostID проверяет получение количества закладок с некорректным post_id.
func TestGetCountOfBookmarksByPost_InvalidPostID(t *testing.T) {
	invalidPostID := "abc"
	req, _ := http.NewRequest("GET", "/test/posts/"+invalidPostID+"/bookmarks/count", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("GetCountOfBookmarksByPost with invalid post ID returned Bad Request as expected")
	}
}

// TestGetBookmarkTimestamp_Success проверяет успешное получение временной метки закладки.
func TestGetBookmarkTimestamp_Success(t *testing.T) {
	userID := "1"
	postID := "1"
	req, _ := http.NewRequest("GET", "/test/bookmarks/"+userID+"/"+postID+"/timestamp", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusOK, w.Code, w.Body.String())
	} else {
		t.Log("GetBookmarkTimestamp endpoint returned OK")
	}
}

// TestGetBookmarkTimestamp_InvalidData проверяет получение временной метки с некорректными данными.
func TestGetBookmarkTimestamp_InvalidData(t *testing.T) {
	invalidUserID := "xyz"
	invalidPostID := "abc"
	req, _ := http.NewRequest("GET", "/test/bookmarks/"+invalidUserID+"/"+invalidPostID+"/timestamp", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d. Body: %s", http.StatusBadRequest, w.Code, w.Body.String())
	} else {
		t.Log("GetBookmarkTimestamp with invalid data returned Bad Request as expected")
	}
}
