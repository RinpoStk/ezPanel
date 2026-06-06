package service

import "github.com/gorilla/websocket"

// PtyManager manages PTY/terminal sessions.
type PtyManager struct {
	sessions map[string]*PtySession
}

// PtySession represents an active terminal session.
type PtySession struct {
	ID      string
	Conn    *websocket.Conn
	Closed  bool
}

// NewPtyManager creates a new PtyManager.
func NewPtyManager() *PtyManager {
	return &PtyManager{
		sessions: make(map[string]*PtySession),
	}
}

// Create creates a new PTY session.
func (m *PtyManager) Create(id string, conn *websocket.Conn) *PtySession {
	// TODO: implement actual PTY spawning
	s := &PtySession{
		ID:     id,
		Conn:   conn,
		Closed: false,
	}
	m.sessions[id] = s
	return s
}

// Close closes a PTY session by ID.
func (m *PtyManager) Close(id string) {
	// TODO: implement actual PTY cleanup
	if s, ok := m.sessions[id]; ok {
		s.Closed = true
		delete(m.sessions, id)
	}
}

// Get retrieves a PTY session by ID.
func (m *PtyManager) Get(id string) (*PtySession, bool) {
	s, ok := m.sessions[id]
	return s, ok
}
