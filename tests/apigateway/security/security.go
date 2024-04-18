package security_test

import (
	"testing"

	"github.com/your_organization/security" // Import the security package
)

func TestAuthenticationManager_AuthenticateUser_ValidCredentials(t *testing.T) {
	authManager := security.NewAuthenticationManager()

	username := "test_user"
	password := "test_password"

	validUser := authManager.AuthenticateUser(username, password)

	if !validUser {
		t.Errorf("Authentication failed with valid credentials.")
	}
}

func TestAuthenticationManager_AuthenticateUser_InvalidCredentials(t *testing.T) {
	authManager := security.NewAuthenticationManager()

	username := "test_user"
	invalidPassword := "wrong_password"

	validUser := authManager.AuthenticateUser(username, invalidPassword)

	if validUser {
		t.Errorf("Authentication succeeded with invalid credentials.")
	}
}

func TestAuthorizationManager_CheckPermission_AllowedResource(t *testing.T) {
	authzManager := security.NewAuthorizationManager()

	user := "test_user"
	allowedResource := "/api/resource1"

	allowed := authzManager.CheckPermission(user, allowedResource)

	if !allowed {
		t.Errorf("Permission denied for allowed resource.")
	}
}

func TestAuthorizationManager_CheckPermission_DisallowedResource(t *testing.T) {
	authzManager := security.NewAuthorizationManager()

	user := "test_user"
	disallowedResource := "/api/resource2"

	disallowed := authzManager.CheckPermission(user, disallowedResource)

	if disallowed {
		t.Errorf("Permission granted for disallowed resource.")
	}
}

func TestEncryptionManager_EncryptDecrypt(t *testing.T) {
	encryptionManager := security.NewEncryptionManager()

	originalData := "Sensitive data to encrypt"

	encryptedData := encryptionManager.Encrypt(originalData)

	decryptedData := encryptionManager.Decrypt(encryptedData)

	if decryptedData != originalData {
		t.Errorf("Decrypted data does not match original data.")
	}
}

func TestSecurityAlerts_DetectSuspiciousActivity(t *testing.T) {
	securityAlerts := security.NewSecurityAlerts()

	suspiciousActivityDetected := securityAlerts.DetectSuspiciousActivity()

	if !suspiciousActivityDetected {
		t.Errorf("Suspicious activity was not detected.")
	}
}

func TestSecurityLogging_LogIncident(t *testing.T) {
	securityLogger := security.NewSecurityLogging()

	incident := "Unauthorized access attempt detected."

	securityLogger.LogIncident(incident)

	loggedIncident := securityLogger.GetLoggedIncidents()

	if len(loggedIncident) == 0 || loggedIncident[0] != incident {
		t.Errorf("Security incident was not logged.")
	}
}

func TestSecurityConfiguration_ConfigureSecurity(t *testing.T) {
	securityConfig := security.NewSecurityConfiguration()

	securityConfig.ConfigureSecurity()

	configApplied := securityConfig.IsSecurityConfigured()

	if !configApplied {
		t.Errorf("Security configuration was not applied.")
	}
}
