package main

import (
	"errors"
	"log"
	"sync"
	"time"
)
var sessionTime = time.Now().Unix()

// SessionManager keeps track of all sessions from creation, updating
// to destroying.
type SessionManager struct {
	sessions map[string]Session
	mu sync.Mutex
}

// Session stores the session's data
type Session struct {
	Data map[string]interface{}
	time time.Time
}

// NewSessionManager creates a new sessionManager
func NewSessionManager() *SessionManager {
	m := &SessionManager{
		sessions: make(map[string]Session),
	}

	go func() {
		t := time.Tick(time.Second * 1)

		for {
			<-t
			m.mu.Lock()
			for k, v := range m.sessions {
				if v.time.Unix() < time.Now().Unix()-sessionTime {
					delete(m.sessions, k)
				}
			}
			m.mu.Unlock()
		}
			
		
		
	}()

	return m
}

// CreateSession creates a new session and returns the sessionID
func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.mu.Lock()
	m.sessions[sessionID] = Session{
		Data:  make(map[string]interface{}),
	}
	m.mu.Unlock()

	log.Println("Session created: ", sessionID)

	

	return sessionID, nil
}

// ErrSessionNotFound returned when sessionID not listed in
// SessionManager
var ErrSessionNotFound = errors.New("SessionID does not exists")

// GetSessionData returns data related to session if sessionID is
// found, errors otherwise
func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
	session, ok := m.sessions[sessionID]

	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.Data, nil
}

// UpdateSessionData overwrites the old session data with the new one
func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	_, ok := m.sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}
	m.mu.Lock()
	m.sessions[sessionID] = Session{
		Data: data,
		time: time.Now(),
	}
	m.mu.Unlock()
	return nil
}

func main() {
	// Create new sessionManager and new session
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	// Update session data
	data := make(map[string]interface{})
	data["website"] = "longhoang.de"

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to longhoang.de")

	// Retrieve data from manager again
	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}