package ini

import (
	. "../jasmine"
	"fmt"
	"testing"
)

func Test_(t *testing.T) {
	Describe("ini", func() {
		It("should read section name and value", func() {
			env, _ := Read("inifile/env.ini")
			develop := env["develop"]
			Expect(develop["database"]).ToBe("db1")
			Expect(develop["logLevel"]).ToBe("TRACE")
		})
		It("should read section name and value", func() {
			config, _ := Read("inifile/database.ini")
			Env := "develop"
			Expect(config[Env]["static"]).ToBe("/webdesign")
		})
		It("should not read non exist file", func() {
			_, err := Read("inifile/nonexist.ini")
			Expect(err != nil).ToBeTruthy()
		})
	})
}
