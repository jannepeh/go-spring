@echo off
echo Testing Go Spring CRUD API
echo.

echo 1. Testing Homepage...
curl -s http://localhost:8080/
echo.
echo.

echo 2. Creating a test article...
curl -s -X POST http://localhost:8080/articles ^
  -H "Content-Type: application/json" ^
  -d "{\"title\":\"Test Article\",\"desc\":\"This is a test article\",\"content\":\"This is the content of the test article.\"}"
echo.
echo.

echo 3. Getting all articles...
curl -s http://localhost:8080/articles
echo.
echo.

echo API tests completed!
echo.
echo To test UPDATE and DELETE operations:
echo - Copy an article ID from the response above
echo - Use these commands (replace {id} with actual ID):
echo.
echo UPDATE: curl -X PUT http://localhost:8080/articles/{id} -H "Content-Type: application/json" -d "{\"title\":\"Updated Title\"}"
echo DELETE: curl -X DELETE http://localhost:8080/articles/{id}
echo.
pause