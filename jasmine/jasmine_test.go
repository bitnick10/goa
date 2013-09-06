package jasmine

import (
	"testing"
)

func Test_jasmine(t *testing.T) {
	// PrintSpec()
	Describe("When first call Describe", func() {
		It("level 0 children's number should be 1", func() {
			Expect(len(g_rootSpec.children)).ToBe(1)
		})
		Describe("When inner call Describe ", func() {
			It("level 1 children's number should be 2", func() {
				Expect(len(g_rootSpec.children[0].children)).ToBe(2)
				Expect("for test error").ToBe("always error")
			})
		})
		Describe("When again inner call Describe ", func() {
			It("level 1 children's number should be 3", func() {
				Expect(len(g_rootSpec.children[0].children)).ToBe(3)
			})
		})
	})
}
