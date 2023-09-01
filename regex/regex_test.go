package regex_test

import (
	"github.com/sinhashubham95/go-utils/structures/pair"
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

func TestFind(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("food")}, regex.Find([]byte(`seafood fool`), regexp.MustCompile(`foo.?`), 1))
	assert.Nil(t, regex.Find([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`), 1))
}

func TestFindWithPattern(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("food")}, regex.FindWithPattern([]byte(`seafood fool`), `foo.?`, 1))
	assert.Nil(t, regex.FindWithPattern([]byte(`seafood fool`), `xyz.?`, 1))
}

func TestFindString(t *testing.T) {
	assert.Equal(t, []string{"food"}, regex.FindString(`seafood fool`, regexp.MustCompile(`foo.?`), 1))
	assert.Nil(t, regex.FindString(`seafood fool`, regexp.MustCompile(`xyz.?`), 1))
}

func TestFindStringWithPattern(t *testing.T) {
	assert.Equal(t, []string{"food"}, regex.FindStringWithPattern(`seafood fool`, `foo.?`, 1))
	assert.Nil(t, regex.FindStringWithPattern(`seafood fool`, `xyz.?`, 1))
}

func TestFindAll(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("food"), []byte("fool")}, regex.FindAll([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindAll([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindAllWithPattern(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("food"), []byte("fool")}, regex.FindAllWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindAllWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindAllString(t *testing.T) {
	assert.Equal(t, []string{"food", "fool"}, regex.FindAllString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindAllString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindAllStringWithPattern(t *testing.T) {
	assert.Equal(t, []string{"food", "fool"}, regex.FindAllStringWithPattern(`seafood fool`, `foo.?`))
	assert.Nil(t, regex.FindAllStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindFirst(t *testing.T) {
	assert.Equal(t, []byte("food"), regex.FindFirst([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindFirst([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindFirstWithPattern(t *testing.T) {
	assert.Equal(t, []byte("food"), regex.FindFirstWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindFirstWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindFirstString(t *testing.T) {
	assert.Equal(t, "food", regex.FindFirstString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Empty(t, regex.FindFirstString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindFirstStringWithPattern(t *testing.T) {
	assert.Equal(t, "food", regex.FindFirstStringWithPattern(`seafood fool`, `foo.?`))
	assert.Empty(t, regex.FindFirstStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindLast(t *testing.T) {
	assert.Equal(t, []byte("fool"), regex.FindLast([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindLast([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindLastWithPattern(t *testing.T) {
	assert.Equal(t, []byte("fool"), regex.FindLastWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindLastWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindLastString(t *testing.T) {
	assert.Equal(t, "fool", regex.FindLastString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Empty(t, regex.FindLastString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindLastStringWithPattern(t *testing.T) {
	assert.Equal(t, "fool", regex.FindLastStringWithPattern(`seafood fool`, `foo.?`))
	assert.Empty(t, regex.FindLastStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindIndex(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7)},
		regex.FindIndex([]byte(`seafood fool`), regexp.MustCompile(`foo.?`), 1))
	assert.Nil(t, regex.FindIndex([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`), 1))
}

func TestFindIndexWithPattern(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7)}, regex.FindIndexWithPattern([]byte(`seafood fool`), `foo.?`, 1))
	assert.Nil(t, regex.FindIndexWithPattern([]byte(`seafood fool`), `xyz.?`, 1))
}

func TestFindIndexForString(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7)},
		regex.FindIndexForString(`seafood fool`, regexp.MustCompile(`foo.?`), 1))
	assert.Nil(t, regex.FindIndexForString(`seafood fool`, regexp.MustCompile(`xyz.?`), 1))
}

func TestFindIndexForStringWithPattern(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7)},
		regex.FindIndexForStringWithPattern(`seafood fool`, `foo.?`, 1))
	assert.Nil(t, regex.FindIndexForStringWithPattern(`seafood fool`, `xyz.?`, 1))
}

func TestFindAllIndex(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7), pair.New(8, 12)},
		regex.FindAllIndex([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindAllIndex([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindAllIndexWithPattern(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7), pair.New(8, 12)},
		regex.FindAllIndexWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindAllIndexWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindAllIndexForString(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7), pair.New(8, 12)},
		regex.FindAllIndexForString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindAllIndexForString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindAllIndexForStringWithPattern(t *testing.T) {
	assert.Equal(t, []*pair.Pair[int, int]{pair.New(3, 7), pair.New(8, 12)},
		regex.FindAllIndexForStringWithPattern(`seafood fool`, `foo.?`))
	assert.Nil(t, regex.FindAllIndexForStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindFirstIndex(t *testing.T) {
	assert.Equal(t, pair.New(3, 7),
		regex.FindFirstIndex([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindFirstIndex([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindFirstIndexWithPattern(t *testing.T) {
	assert.Equal(t, pair.New(3, 7), regex.FindFirstIndexWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindFirstIndexWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindFirstIndexForString(t *testing.T) {
	assert.Equal(t, pair.New(3, 7),
		regex.FindFirstIndexForString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindFirstIndexForString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindFirstIndexForStringWithPattern(t *testing.T) {
	assert.Equal(t, pair.New(3, 7),
		regex.FindFirstIndexForStringWithPattern(`seafood fool`, `foo.?`))
	assert.Nil(t, regex.FindFirstIndexForStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindLastIndex(t *testing.T) {
	assert.Equal(t, pair.New(8, 12),
		regex.FindLastIndex([]byte(`seafood fool`), regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindLastIndex([]byte(`seafood fool`), regexp.MustCompile(`xyz.?`)))
}

func TestFindLastIndexWithPattern(t *testing.T) {
	assert.Equal(t, pair.New(8, 12), regex.FindLastIndexWithPattern([]byte(`seafood fool`), `foo.?`))
	assert.Nil(t, regex.FindLastIndexWithPattern([]byte(`seafood fool`), `xyz.?`))
}

func TestFindLastIndexForString(t *testing.T) {
	assert.Equal(t, pair.New(8, 12),
		regex.FindLastIndexForString(`seafood fool`, regexp.MustCompile(`foo.?`)))
	assert.Nil(t, regex.FindLastIndexForString(`seafood fool`, regexp.MustCompile(`xyz.?`)))
}

func TestFindLastIndexForStringWithPattern(t *testing.T) {
	assert.Equal(t, pair.New(8, 12),
		regex.FindLastIndexForStringWithPattern(`seafood fool`, `foo.?`))
	assert.Nil(t, regex.FindLastIndexForStringWithPattern(`seafood fool`, `xyz.?`))
}

func TestFindSubMatches(t *testing.T) {
	assert.Equal(t, [][][]byte{{[]byte("option1: value1"), []byte("option1"), []byte("value1")}},
		regex.FindSubMatches([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`), 1))
	assert.Nil(t, regex.FindSubMatches([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile("xyz*"), 1))
}

func TestFindSubMatchesWithPattern(t *testing.T) {
	assert.Equal(t, [][][]byte{{[]byte("option1: value1"), []byte("option1"), []byte("value1")}},
		regex.FindSubMatchesWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`, 1))
	assert.Nil(t, regex.FindSubMatchesWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), "xyz*", 1))
}

func TestFindSubMatchesForString(t *testing.T) {
	assert.Equal(t, [][]string{{"option1: value1", "option1", "value1"}},
		regex.FindSubMatchesForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`), 1))
	assert.Nil(t, regex.FindSubMatchesForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile("xyz*"), 1))
}

func TestFindSubMatchesForStringWithPattern(t *testing.T) {
	assert.Equal(t, [][]string{{"option1: value1", "option1", "value1"}},
		regex.FindSubMatchesForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`, 1))
	assert.Nil(t, regex.FindSubMatchesForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, "xyz*", 1))
}

func TestFindAllSubMatches(t *testing.T) {
	assert.Equal(t, [][][]byte{{[]byte("option1: value1"), []byte("option1"), []byte("value1")},
		{[]byte("option2: value2"), []byte("option2"), []byte("value2")},
		{[]byte("option3: value3"), []byte("option3"), []byte("value3")}},
		regex.FindAllSubMatches([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindAllSubMatches([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile("xyz*")))
}

func TestFindAllSubMatchesWithPattern(t *testing.T) {
	assert.Equal(t, [][][]byte{{[]byte("option1: value1"), []byte("option1"), []byte("value1")},
		{[]byte("option2: value2"), []byte("option2"), []byte("value2")},
		{[]byte("option3: value3"), []byte("option3"), []byte("value3")}},
		regex.FindAllSubMatchesWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindAllSubMatchesWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), "xyz*"))
}

func TestFindAllSubMatchesForString(t *testing.T) {
	assert.Equal(t, [][]string{{"option1: value1", "option1", "value1"},
		{"option2: value2", "option2", "value2"},
		{"option3: value3", "option3", "value3"}},
		regex.FindAllSubMatchesForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindAllSubMatchesForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile("xyz*")))
}

func TestFindAllSubMatchesForStringWithPattern(t *testing.T) {
	assert.Equal(t, [][]string{{"option1: value1", "option1", "value1"},
		{"option2: value2", "option2", "value2"},
		{"option3: value3", "option3", "value3"}},
		regex.FindAllSubMatchesForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindAllSubMatchesForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, "xyz*"))
}

func TestFindFirstSubMatch(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("option1: value1"), []byte("option1"), []byte("value1")},
		regex.FindFirstSubMatch([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindFirstSubMatch([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile("xyz*")))
}

func TestFindFirstSubMatchWithPattern(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("option1: value1"), []byte("option1"), []byte("value1")},
		regex.FindFirstSubMatchWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindFirstSubMatchWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), "xyz*"))
}

func TestFindFirstSubMatchForString(t *testing.T) {
	assert.Equal(t, []string{"option1: value1", "option1", "value1"},
		regex.FindFirstSubMatchForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindFirstSubMatchForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile("xyz*")))
}

func TestFindFirstSubMatchForStringWithPattern(t *testing.T) {
	assert.Equal(t, []string{"option1: value1", "option1", "value1"},
		regex.FindFirstSubMatchForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindFirstSubMatchForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, "xyz*"))
}

func TestFindLastSubMatch(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("option3: value3"), []byte("option3"), []byte("value3")},
		regex.FindLastSubMatch([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindLastSubMatch([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), regexp.MustCompile("xyz*")))
}

func TestFindLastSubMatchWithPattern(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("option3: value3"), []byte("option3"), []byte("value3")},
		regex.FindLastSubMatchWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindLastSubMatchWithPattern([]byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`), "xyz*"))
}

func TestFindLastSubMatchForString(t *testing.T) {
	assert.Equal(t, []string{"option3: value3", "option3", "value3"},
		regex.FindLastSubMatchForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)))
	assert.Nil(t, regex.FindLastSubMatchForString(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, regexp.MustCompile("xyz*")))
}

func TestFindLastSubMatchForStringWithPattern(t *testing.T) {
	assert.Equal(t, []string{"option3: value3", "option3", "value3"},
		regex.FindLastSubMatchForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`))
	assert.Nil(t, regex.FindLastSubMatchForStringWithPattern(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`, "xyz*"))
}
