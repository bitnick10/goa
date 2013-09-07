package session

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type S struct {
	sessionmap map[string]*Content
}

type Content struct {
	ObjectId       string
	LastAccessTime time.Time
}

func New(duration time.Duration) (session *S) {
	session = &S{make(map[string]*Content)}
	ticker := time.NewTicker(duration)
	go func() {
		for {
			<-ticker.C
			session.Sweep(duration)
		}
	}()
	return session
}
func (session *S) Get(key string) *Content {
	return session.sessionmap[key]
}
func (session *S) Len() int {
	return len(session.sessionmap)
}

func (session *S) Sweep(duration time.Duration) {
	for key, value := range session.sessionmap {
		if value.LastAccessTime.Add(duration).Before(time.Now()) {
			delete(session.sessionmap, key)
		}
	}
}
func (session *S) Update(sessionId string) {
	if session.sessionmap[sessionId] != nil {
		session.sessionmap[sessionId].LastAccessTime = time.Now()
	}

}
func (session *S) Add(objectId string) (id string) {
	h := md5.New()
	io.WriteString(h, objectId)
	io.WriteString(h, randString())
	id = hex.EncodeToString(h.Sum(nil))

	for session.sessionmap[id] != nil {
		io.WriteString(h, strconv.Itoa(rd()))
		id = hex.EncodeToString(h.Sum(nil))
	}
	session.sessionmap[id] = &Content{objectId, time.Now()}
	return id
}
func rd() int {
	rand.Seed(time.Now().UTC().UnixNano())
	// returns a non-negative pseudo-random int
	return rand.Int()
}
func randString() string {
	return strconv.Itoa(rd())
}
