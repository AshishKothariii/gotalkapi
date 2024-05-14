# Testing Go Talk Api

## 1. Clone this GitHub repository to your local machine using the following command:

- git clone https://github.com/AshishKothariii/gotalkapi.git

## 2. Opening the Project in Visual studio code

### 2.1 Launch Visual studio.

- Choose "Open an existing folder" from the welcome screen.

- Navigate to the directory where you cloned the repository and select the folder.

- Click "OK" to open the folder.

## 3. Setup env variables:

- Navigate to gotalkapi/.env file .
- Add proper environment variables
- In terminal Do go mod init Your module name.
- go get "github.com/joho/godotenv".
- go get "github.com/gin-gonic/gin".
- go get "go.mongodb.org/mongo-driver/mongo".
- go get "github.com/golang-jwt/jwt/v4".
- go get "github.com/gorilla/websocket".
- go get "golang.org/x/crypto/bcrypt"

- Save the file after making the changes.

## 4. Building and Deploying the App:

- In terminal just type go run .

## 5. Testing the App:

- Go to postman and test the api.
- You can also use a frontend and connect to make your own app on this backend.
