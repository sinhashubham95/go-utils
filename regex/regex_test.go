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
