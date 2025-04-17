package repo

import (
	"sync"
	"time"

	"github.com/sarc-tech/backend/internal/models"
)

func (r *RepoHandler) GetSession(token string) (session models.Session, err error) {
	ok, userId := r.sessionTTL.Get(token)
	if ok {
		session.ID = token
		session.UserID = userId
		return
	}
	err = r.Db.Get(&session, "SELECT * FROM core_session WHERE id = $1", token)
	r.sessionTTL.Set(session.ID, session.UserID)
	return 
}

type sessionUser struct {
	userId int
	lastAccess int64
}

type sessionTTLMap struct {
	m   map[string]*sessionUser
	l	sync.Mutex
}

func newSessionTTLMap(maxTTL int) (s *sessionTTLMap) {
	s = &sessionTTLMap{m: make(map[string]*sessionUser)}
	go func() {
		for now := range time.Tick(time.Second) {
			s.l.Lock()
			for k, v := range s.m {
				if now.Unix() - v.lastAccess > int64(maxTTL) {
					delete(s.m, k)
				}
			}
			s.l.Unlock()
		}
	}()
	return
}

func (s *sessionTTLMap) Set(token string, userId int) {
	s.l.Lock()
	defer s.l.Unlock()
	su, ok := s.m[token]
	if !ok {
		su = &sessionUser{userId: userId}
		s.m[token] = su
	}
	su.lastAccess = time.Now().Unix()
}

func (s *sessionTTLMap) Get(token string) (bool, int) {
	s.l.Lock()
	defer s.l.Unlock()
	if su, ok := s.m[token]; ok {
		su.lastAccess = time.Now().Unix()
		return true, su.userId
	}
	return false, 0
}