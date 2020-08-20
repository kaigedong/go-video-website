package session

import (
	"sync"
	"time"

	"github.com/go-video-website/api/dbops"
	"github.com/go-video-website/api/defs"
	"github.com/go-video-website/api/utils"
)

// 1.服务器起来的时候，需要从db中获取session，存到缓存
// 2.用户登录/注册时，给用户产生一个session
// 3.需要判断session是否过期

// 1.不用redis减少复杂度
// 2.用户数据量少，in memery够了
// 3.sync.Map在1.9之后加入，能够并发读写，线程安全，读取性能非常好，写的时候，由于全局锁的原因有一定性能限制
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func noInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// 从数据库读取session，然后写入到mem
func LoadSessionsFromDB() {
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
	id, _ := utils.NewUUID()
	ct := noInMilli()
	ttl := ct + 30*60*1000 // Server side session valid time: 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
