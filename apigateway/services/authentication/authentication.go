package authentication

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Authenticator contains functions to listen for incoming requests and authenticate users.
type Authenticator struct {
	// Add any necessary fields here
	userPhone int
	userPIN   int
}

// CredentialExtractor functions to extract authentication credentials from request headers.
type CredentialExtractor struct {
	// Request represents the HTTP request.
	Request *http.Request

	DB *gorm.DB // GORM instance for database operations

	// Logger represents the logger utility.
	Logger *log.Logger

	// DeviceType represents the type of the device.
	DeviceType string

	// DeviceLocation represents the location of the device.
	DeviceLocation string

	// DeviceIP represents the IP address of the device.
	DeviceIP string
}

// CredentialValidator functions to validate authentication credentials against predefined standards.
type CredentialValidator struct {
	// Add any necessary fields here
	userAccessToken string
	userType        string
	userID          int
	userOrgId       int
	userAPIkey      string
}

// Authorizer functions to authorize access to API resources based on the authenticated identity and requested resource.
type Authorizer struct {
	// Add any necessary fields here
	userType  string
	userID    int
	userOrgId int
}

// ErrorHandler represents a utility for handling errors in the application.
type ErrorHandler struct {
	Logger *log.Logger // Logger instance for logging errors

}

/*verifies the JWT token. It checks if the Authorization
header is present, parses the token using the jwt.Parse
function, and validates the token using a secret key*/

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}

// Redirect all routing
func createRouter() *mux.Router {
	// create a new router
	router := mux.NewRouter()

	// register the microservices with the router
	router.HandleFunc("/mobapi", mobHandler)
	router.HandleFunc("/webapi", ordersHandler)
	router.HandleFunc("/comapi", ordersHandler)
	router.HandleFunc("/", ordersHandler)

	return router
}

// MultiFactorAuthenticator functions to implement multi-factor authentication (MFA) for enhanced security.
type MultiFactorAuthenticator struct {
	// Add any necessary fields here
}

// IdentityProviderIntegration functions to integrate with external identity providers like LDAP, SAML, or OpenID Connect.
type IdentityProviderIntegration struct {
	// Add any necessary fields here
}

// NewAuthenticator creates a new instance of Authenticator.
func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

// Listen starts listening for incoming API requests on the designated endpoint.
func (a *Authenticator) Listen(endpoint string) error {
	// Implementation goes here
	return nil
}

// ExtractCredentials extracts authentication credentials from request headers.
func (c *CredentialExtractor) ExtractCredentials(req *http.Request) (username, password string, err error) {
	// Implementation goes here
	return "", "", nil
}

// ValidateCredentials validates authentication credentials against predefined standards.
func (c *CredentialValidator) ValidateCredentials(pin int8, password string) error {
	// Implementation
}

func AuthenticateUser(userPIN, userAccount, userAccessToken, userAPIKey string) (bool, error) {
	// Encrypt userPIN with MD5 and substring to 10 numeric characters
	hashedPIN := hashAndSubstring(userPIN)

	// Setup database connection
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/cashpay_db")
	if err != nil {
		return false, fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Query database to check user credentials
	var storedPINHash, storedAPIKey, storedAccessToken string
	err = db.QueryRow("SELECT PINHash, APIKey, AccessToken FROM app_users WHERE userAccount = ?", userAccount).Scan(&storedPINHash, &storedAPIKey, &storedAccessToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User not found
		}
		return false, fmt.Errorf("error querying database: %v", err)
	}

	// Check hashed and subscripted userPIN matches with the one in the table
	if hashedPIN != storedPINHash {
		return false, nil // UserPIN does not match
	}

	// Check userAPIKey
	if userAPIKey != storedAPIKey {
		return false, nil // UserAPIKey does not match
	}

	// Check AccessToken is not blank
	if userAccessToken == "" {
		return false, nil // AccessToken is blank
	}

	// Save access token to the database
	if err := saveAccessToken(db, storedPINHash, userAccessToken); err != nil {
		return false, fmt.Errorf("error saving access token to database: %v", err)
	}

	return true, nil // Authentication successful
}

func hashAndSubstring(pin string) string {
	hasher := md5.New()
	hasher.Write([]byte(pin))
	hashedPIN := fmt.Sprintf("%x", hasher.Sum(nil))
	return hashedPIN[:10] // Substring to first 10 characters
}

func saveAccessToken(db *sql.DB, userID, accessToken string) error {
	_, err := db.Exec("INSERT INTO app_users_token (user_id, access_token, login_date) VALUES (?, ?, ?)",
		userID, accessToken, time.Now())
	return err
}
