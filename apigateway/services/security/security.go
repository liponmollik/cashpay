package security

import (
	"testing"
)

func TestAuthenticationManager_AuthenticateUser(t *testing.T) {
	// Initialize AuthenticationManager
	authManager := NewAuthenticationManager()

	// Define test credentials
	username := "test_user"
	password := "test_password"

	// Attempt authentication with valid credentials
	validUser := authManager.AuthenticateUser(username, password)

	// Expect authentication success
	if !validUser {
		t.Errorf("Authentication failed with valid credentials.")
	}

	// Attempt authentication with invalid credentials
	invalidUser := authManager.AuthenticateUser(username, "wrong_password")

	// Expect authentication failure
	if invalidUser {
		t.Errorf("Authentication succeeded with invalid credentials.")
	}
}

func TestAuthorizationManager_CheckPermission(t *testing.T) {
	// Initialize AuthorizationManager
	authzManager := NewAuthorizationManager()

	// Define test user and permissions
	user := "test_user"
	allowedResource := "/api/resource1"
	disallowedResource := "/api/resource2"

	// Check permission for allowed resource
	allowed := authzManager.CheckPermission(user, allowedResource)

	// Expect permission granted
	if !allowed {
		t.Errorf("Permission denied for allowed resource.")
	}

	// Check permission for disallowed resource
	disallowed := authzManager.CheckPermission(user, disallowedResource)

	// Expect permission denied
	if disallowed {
		t.Errorf("Permission granted for disallowed resource.")
	}
}

func TestEncryptionManager_EncryptDecrypt(t *testing.T) {
	// Initialize EncryptionManager
	encryptionManager := NewEncryptionManager()

	// Define test data
	originalData := "Sensitive data to encrypt"

	// Encrypt the data
	encryptedData := encryptionManager.Encrypt(originalData)

	// Decrypt the data
	decryptedData := encryptionManager.Decrypt(encryptedData)

	// Expect decrypted data to match original data
	if decryptedData != originalData {
		t.Errorf("Decrypted data does not match original data.")
	}
}

func TestSecurityAlerts_DetectSuspiciousActivity(t *testing.T) {
	// Initialize SecurityAlerts
	securityAlerts := NewSecurityAlerts()

	// Simulate suspicious activity
	suspiciousActivityDetected := securityAlerts.DetectSuspiciousActivity()

	// Expect suspicious activity detection
	if !suspiciousActivityDetected {
		t.Errorf("Suspicious activity was not detected.")
	}
}

func TestSecurityLogging_LogIncident(t *testing.T) {
	// Initialize SecurityLogging
	securityLogger := NewSecurityLogging()

	// Simulate security incident
	incident := "Unauthorized access attempt detected."

	// Log security incident
	securityLogger.LogIncident(incident)

	// Check if the incident is logged
	loggedIncident := securityLogger.GetLoggedIncidents()

	// Expect the incident to be logged
	if len(loggedIncident) == 0 || loggedIncident[0] != incident {
		t.Errorf("Security incident was not logged.")
	}
}

func TestSecurityConfiguration_ConfigureSecurity(t *testing.T) {
	// Initialize SecurityConfiguration
	securityConfig := NewSecurityConfiguration()

	// Simulate security configuration
	securityConfig.ConfigureSecurity()

	// Check if security configuration is applied
	configApplied := securityConfig.IsSecurityConfigured()

	// Expect security configuration to be applied
	if !configApplied {
		t.Errorf("Security configuration was not applied.")
	}
}
