package regex_test

import (
	"regexp"
	"testing"

	"github.com/sinhashubham95/go-utils/regex"
	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	assert.Equal(t, []byte("option1=value1\n"), regex.Expand([]byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`), []byte("$key=$value\n"), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`), 1))
}

func TestExpandWithPattern(t *testing.T) {
	assert.Equal(t, []byte("option1=value1\n"), regex.ExpandWithPattern([]byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`), []byte("$key=$value\n"), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`, 1))
}

func TestExpandString(t *testing.T) {
	assert.Equal(t, "option1=value1\n", regex.ExpandString(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`, "$key=$value\n", regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`), 1))
}

func TestExpandStringWithPattern(t *testing.T) {
	assert.Equal(t, "option1=value1\n", regex.ExpandStringWithPattern(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`, "$key=$value\n", `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`, 1))
}

func TestExpandAll(t *testing.T) {
	assert.Equal(t, []byte("option1=value1\noption2=value2\noption3=value3\n"), regex.ExpandAll([]byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`), []byte("$key=$value\n"), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
}

func TestExpandAllWithPattern(t *testing.T) {
	assert.Equal(t, []byte("option1=value1\noption2=value2\noption3=value3\n"), regex.ExpandAllWithPattern([]byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`), []byte("$key=$value\n"), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
}

func TestExpandAllString(t *testing.T) {
	assert.Equal(t, "option1=value1\noption2=value2\noption3=value3\n", regex.ExpandAllString(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`, "$key=$value\n", regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
}

func TestExpandAllStringWithPattern(t *testing.T) {
	assert.Equal(t, "option1=value1\noption2=value2\noption3=value3\n", regex.ExpandAllStringWithPattern(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`, "$key=$value\n", `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
}
