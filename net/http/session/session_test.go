package session

import (
	. "github.com/Bitnick2002/goa/jasmine"
	"testing"
	"time"
)

func Test_(t *testing.T) {
	var id string
	duration := time.Millisecond * 10
	session := New(duration)
	Describe("Session", func() {
		It("should add session", func() {
			session.Add("5216ea39d8b6d1a51d071d32")
			session.Add("5216ea39d8b6d1a51d071d32")
			session.Add("5216ea39d8b6d1a51d071d33")
			session.Add("5216ea39d8b6d1a51d071d34")
			id = session.Add("5216ea39d8b6d1a51d071d35")
			Expect(len(session)).ToBe(5)
		})
		It("should update session", func() {
			time.Sleep(duration / 2)
			session.Update(id)
			time.Sleep(duration - duration/10)
			Expect(len(session)).ToBe(1)
			Expect(session[id].ObjectId).ToBe("5216ea39d8b6d1a51d071d35")
		})
		Describe("when session has been swept", func() {
			It("should have no seesion", func() {
				time.Sleep(duration)
				Expect(len(session)).ToBe(0)
			})
		})
	})
}
