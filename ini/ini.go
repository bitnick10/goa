package ini

import (
	"io/ioutil"
	"strings"
)

type Ini map[string]Section
type Section map[string]string

func Read(filename string) (ini Ini, err error) {
	ini = make(Ini)

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ini, err
	}
	content := string(bytes)

	lines := strings.Split(content, "\n")
	var section string
	var line line_t
	for _, v := range lines {
		line = line_t(v)
		if line.isSectionLine() {
			section = line.getSectionName()
			ini[section] = make(Section)
		}
		if line.isValueLine() {
			name, value := line.getNameAndValue()
			ini[section][name] = value
		}
	}
	return ini, nil
}

type line_t string

func (line line_t) isSectionLine() bool {
	return strings.Contains(string(line), "[")
}

func (line line_t) getSectionName() string {
	str1 := strings.Split(string(line), "]")[0]
	ret := strings.Split(str1, "[")[1]
	return ret
}

func (line line_t) isValueLine() bool {
	return strings.Contains(string(line), "=")
}

func (line line_t) getNameAndValue() (name, value string) {
	str := strings.Split(string(line), "=")
	name = strings.TrimSpace(str[0])
	value = strings.TrimSpace(str[1])
	return name, value
}
