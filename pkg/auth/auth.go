package auth

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// DatabaseConfig represents the structure of the database configuration read from the YAML file
type DatabaseConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// authService represents the authentication service
type authService struct {
	db *sql.DB // Database connection
}

// NewAuthService creates a new instance of the authentication service
func NewAuthService(db *sql.DB) *authService {
	return &authService{
		db: db,
	}
}

// apiGateway represents the API gateway
type apiGateway struct {
	authService *authService // Authentication service
}

// NewAPIGateway creates a new instance of API gateway
func NewAPIGateway(authService *authService) *apiGateway {
	return &apiGateway{
		authService: authService,
	}
}

// AuthenticateUser authenticates the user with the provided phone number and password
func (as *authService) AuthenticateUser(requestType, userPin, userPhone, userPassword string) bool {
	// Sanitize the user's provided password
	sanitizedPassword := sanitizePassword(userPassword)

	// Encrypt the user's provided PIN
	encryptedPin, err := encryptPIN(userPin)
	if err != nil {
		fmt.Printf("Error encrypting PIN: %v", err)
		return false
	}

	// Prepare the SQL query
	query := "SELECT user_pin, user_password FROM app_users WHERE user_phone_no=?"
	row := as.db.QueryRow(query, userPhone)

	var storedPin, storedPassword string
	// Execute the query and retrieve the stored PIN and password
	err = row.Scan(&storedPin, &storedPassword)
	if err != nil {
		fmt.Printf("Error retrieving stored PIN and password: %v", err)
		return false
	}

	// Compare the encrypted PIN and sanitized password with the stored values
	if encryptedPin == storedPin && sanitizedPassword == storedPassword {
		return true
	} else {
		return false
	}
}

// sanitizePassword sanitizes the provided password as substr(md5(password), -8, 8)
func sanitizePassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Take the last 8 characters of the hashed password
	sanitizedPassword := hashedPassword[len(hashedPassword)-8:]
	return sanitizedPassword
}

// Encrypt user PIN
func encryptPIN(userPin int) (string, error) {
	// Convert PIN to string
	pinStr := strconv.Itoa(userPin)

	// Calculate MD5 hash of the PIN string
	hash := md5.Sum([]byte(pinStr))
	hashStr := hex.EncodeToString(hash[:])

	// Take the first 8 characters of the hash
	first8Chars := hashStr[:8]

	// Convert hexadecimal string to decimal
	decimal, err := strconv.ParseUint(first8Chars, 16, 64)
	if err != nil {
		return "", err
	}

	// Convert decimal to string
	decimalStr := strconv.FormatUint(decimal, 10)

	// Pad with leading zeros if necessary to make it 10 digits
	encryptedPIN := fmt.Sprintf("%010s", decimalStr)

	return encryptedPIN, nil
}

// HandleRequest handles incoming API requests
func (ag *apiGateway) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Extract parameters from the request body
	requestType := r.Form.Get("requestType")
	phone := r.Form.Get("phone")
	password := r.Form.Get("password")

	// Authenticate the user
	authenticated := ag.authService.AuthenticateUser(requestType, secretkey, apiauthokey, phone, password)

	// Check if authentication succeeded
	if authenticated {
		// Authentication succeeded, handle the request based on requestType
		if requestType == "1" {
			http.Redirect(w, r, "/mobapi/v1/routes/", http.StatusFound)
		} else if requestType == "2" {
			http.Redirect(w, r, "/webapi/v1/routes/", http.StatusFound)
		} else {
			// Invalid requestType, return an error response
			errorMessage := fmt.Sprintf("Invalid requestType: %s", requestType)
			http.Error(w, errorMessage, http.StatusBadRequest)
		}
	} else {
		// Authentication failed, return an error response
		sanitizedPassword := sanitizePassword(password)
		errorMessage := fmt.Sprintf("Authentication failed for phone number: %s and password: %s", phone, sanitizedPassword)
		http.Error(w, errorMessage, http.StatusUnauthorized)
	}
}

// readConfig reads the database configuration from the specified YAML file
func readConfig(filename string) (*DatabaseConfig, error) {
	// Read the YAML configuration file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the YAML configuration
	var config DatabaseConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
