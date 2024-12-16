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

**To Use the Code:**

1.  Create the directory structure as shown above.
2.  Copy the code into the appropriate files.
3. Run `go mod init calculator-service` in root
4.  Run `go mod tidy`
5.  Run `go run ./cmd/calc_service/main.go` to start the server
6. Use curl or a browser to test the service.

**Important Notes:**

*   **Expression Evaluation:** The provided `calculator.Calculate` function is a very basic placeholder. You must implement your own logic for evaluating more complex arithmetic expressions correctly, using the parser and evaluator logic you developed in your previous module.
*   **Error Handling:**  The error handling in this code is basic. In a real application, you might want to add more sophisticated logging, error reporting, and handling of various edge cases.
*   **Input Validation:**  The code does basic JSON validation and also the character set of the expression.  Add more robust validation logic.
*   **Scalability:**  For a production service, you'd need to think about scalability, security, and deployment.
*   **Testing:** It is recommended to create unit tests to test each part of the application.

This setup should satisfy the requirements of your task.  Remember to replace the placeholder evaluation logic with your actual arithmetic evaluator.
