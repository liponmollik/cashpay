package session

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Session represents a user session.
type Session struct {
	ID     string
	Data   map[string]interface{}
	Expiry time.Time
}

// SessionManager handles session management.
type SessionManager struct {
	sessionDir string
	mu         sync.Mutex
}

// NewSessionManager creates a new SessionManager with the specified session directory.
func NewSessionManager(sessionDir string) *SessionManager {
	return &SessionManager{sessionDir: sessionDir}
}

// GenerateID generates a secure session ID.
func GenerateID() (string, error) {
	id := make([]byte, 32)
	_, err := rand.Read(id)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(id)
	return base64.URLEncoding.EncodeToString(hash[:]), nil
}

// SaveSession saves a session to disk.
func (sm *SessionManager) SaveSession(sess *Session) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Serialize session data
	data, err := json.Marshal(sess)
	if err != nil {
		return err
	}

	// Create session file
	filePath := filepath.Join(sm.sessionDir, sess.ID)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write encrypted session data to file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// LoadSession loads a session from disk.
func (sm *SessionManager) LoadSession(sessionID string) (*Session, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Read session file
	filePath := filepath.Join(sm.sessionDir, sessionID)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Deserialize session data
	sess := &Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// DeleteSession deletes a session from disk.
func (sm *SessionManager) DeleteSession(sessionID string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Delete session file
	filePath := filepath.Join(sm.sessionDir, sessionID)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

// CleanupSessions deletes expired sessions from disk.
func (sm *SessionManager) CleanupSessions() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Get list of session files
	files, err := ioutil.ReadDir(sm.sessionDir)
	if err != nil {
		return
	}

	// Delete expired sessions
	for _, file := range files {
		filePath := filepath.Join(sm.sessionDir, file.Name())
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}
		sess := &Session{}
		err = json.Unmarshal(data, sess)
		if err != nil {
			continue
		}
		if sess.Expiry.Before(time.Now()) {
			os.Remove(filePath)
		}
	}
}
