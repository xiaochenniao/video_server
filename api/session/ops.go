package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/ntils"
)


var sessionMap *sync.Map

func nowInMilli() int64 {
	return 	time.Now().UnixNano()/1000000
}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB()  {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := ntils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30 * 60 * 1000 //过期时间为30分

	ss := &defs.SimpleSession{Username:un, TTL:ttl}

	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}