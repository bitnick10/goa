package session

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type S map[string]*Content

type Content struct {
	ObjectId       string
	LastAccessTime time.Time
}

func New(duration time.Duration) (session S) {
	session = make(S)
	ticker := time.NewTicker(duration)
	go func() {
		for {
			<-ticker.C
			session.Sweep(duration)
		}
	}()
	return session
}

func (session S) Sweep(duration time.Duration) {
	for key, value := range session {
		if value.LastAccessTime.Add(duration).Before(time.Now()) {
			delete(session, key)
		}
	}
}
func (session S) Update(sessionId string) {
	if session[sessionId] != nil {
		session[sessionId].LastAccessTime = time.Now()
	}

}
func (session S) Add(objectId string) (id string) {
	h := md5.New()
	io.WriteString(h, objectId)
	io.WriteString(h, randString())
	id = hex.EncodeToString(h.Sum(nil))

	for session[id] != nil {
		io.WriteString(h, strconv.Itoa(rd()))
		id = hex.EncodeToString(h.Sum(nil))
	}
	session[id] = &Content{objectId, time.Now()}
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
