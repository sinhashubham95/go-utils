package strings_test

import (
	"github.com/sinhashubham95/go-utils/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbbreviate(t *testing.T) {
	assert.Equal(t, "na", strings.Abbreviate("naruto", "", 0, 2))
	assert.Equal(t, "na", strings.Abbreviate("naruto", "", 3, 2))
	assert.Equal(t, "", strings.Abbreviate("", "", 0, 2))
	assert.Panics(t, func() {
		strings.Abbreviate("naruto", "...", 0, 1)
	})
	assert.Equal(t, "na", strings.Abbreviate("naruto", "", 0, 2))
	assert.Equal(t, "naruto", strings.Abbreviate("naruto", "", 0, 10))
	assert.Equal(t, "na...", strings.Abbreviate("naruto", "...", 5, 5))
	assert.Equal(t, "na...", strings.Abbreviate("naruto", "...", 5, 5))
	assert.Equal(t, "na...", strings.Abbreviate("naruto", "...", 5, 5))
	assert.Panics(t, func() {
		strings.Abbreviate("naruto", "...", 5, 3)
	})
	assert.Equal(t, "na...", strings.Abbreviate("naruto", "...", 5, 5))
}

func TestAbbreviateMiddle(t *testing.T) {
	assert.Equal(t, "", strings.AbbreviateMiddle("", "", 0))
	assert.Equal(t, "na...s", strings.AbbreviateMiddle("naruto rocks", "...", 6))
	assert.Equal(t, "na...ks", strings.AbbreviateMiddle("naruto rocks", "...", 7))
}

func TestAppendIfMissing(t *testing.T) {
	assert.Equal(t, "naruto", strings.AppendIfMissing("naruto", "uto", "to", "o"))
	assert.Equal(t, "naruto", strings.AppendIfMissing("naruto", " rocks", "to", "o"))
	assert.Equal(t, "naruto", strings.AppendIfMissing("naruto", " rocks", "boruto", "o"))
	assert.Equal(t, "naruto rocks", strings.AppendIfMissing("naruto", " rocks", "boruto", "hinata"))
}

func TestAppendIfMissingIgnoreCase(t *testing.T) {
	assert.Equal(t, "naruto", strings.AppendIfMissingIgnoreCase("naruto", "UTO", "TO", "O"))
	assert.Equal(t, "naruto", strings.AppendIfMissingIgnoreCase("naruto", " rocks", "TO", "O"))
	assert.Equal(t, "naruto", strings.AppendIfMissingIgnoreCase("naruto", " rocks", "boruto", "O"))
	assert.Equal(t, "naruto rocks", strings.AppendIfMissingIgnoreCase("naruto", " rocks", "boruto", "hinata"))
}

func TestCapitalize(t *testing.T) {
	assert.Equal(t, "Naruto", strings.Capitalize("naruto"))
	assert.Equal(t, "NARUTO", strings.Capitalize("nARUTO"))
	assert.Equal(t, "Naruto", strings.Capitalize("Naruto"))
	assert.Equal(t, "NarutO", strings.Capitalize("NarutO"))
}

func TestCenter(t *testing.T) {
	assert.Equal(t, "naruto", strings.Center("naruto", 3, 'A'))
	assert.Equal(t, "AnarutoA", strings.Center("naruto", 8, 'A'))
	assert.Equal(t, "narutoA", strings.Center("naruto", 7, 'A'))
}

func TestCenterString(t *testing.T) {
	assert.Equal(t, "naruto", strings.CenterString("naruto", 3, "AAA"))
	assert.Equal(t, "AnarutoAA", strings.CenterString("naruto", 9, "AAA"))
	assert.Equal(t, "AnarutoA", strings.CenterString("naruto", 8, "AAA"))
	assert.Equal(t, "narutoA", strings.CenterString("naruto", 7, "AAA"))
}

func TestChomp(t *testing.T) {
	assert.Equal(t, "naruto", strings.Chomp("naruto"))
	assert.Equal(t, "naruto", strings.Chomp("naruto\n"))
	assert.Equal(t, "naruto", strings.Chomp("naruto\n\r"))
	assert.Equal(t, "naruto", strings.Chomp("naruto\r"))
	assert.Equal(t, "naruto", strings.Chomp("naruto\r\n"))
}
