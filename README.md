# Calculator Service

This is a simple web service that evaluates arithmetic expressions.

## Endpoints

The service exposes one endpoint:

*   **`POST /api/v1/calculate`**

    *   **Request Body (JSON):**
        ```json
        {
          "expression": "arithmetic expression"
        }
        ```
    *   **Response (JSON):**
        *   **Success (HTTP 200):**
            ```json
            {
              "result": "evaluation result"
            }
            ```
        *   **Invalid Input (HTTP 422):**
            ```json
            {
                "error": "Expression is not valid"
            }
            ```
        *   **Internal Server Error (HTTP 500):**
            ```json
            {
              "error": "Internal server error"
            }
            ```

## Examples

### Successful Calculation

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
### To use the code.

1.  Create the directory structure as shown above.
2.  Copy the code into the appropriate files.
3. Run `go mod init calculator-service` in root
4.  Run `go mod tidy`
5.  Run `go run ./cmd/calc_service/main.go` to start the server
6. Use curl or a browser to test the service.
