package regex

// SubMatches are matches of parenthesized subexpressions (also known as capturing groups) within the regular expression,
// numbered from left to right in order of opening parenthesis. SubMatch 0 is the match of the entire expression,
// subMatch 1 is the match of the first parenthesized subexpression, and so on.

import (
	"regexp"

	"github.com/sinhashubham95/go-utils/structures/pair"
)

// Expand appends template to the result and during append it replaces variables in the template with n number of
// matches drawn from the source data d.
//
// In the template, a variable is denoted by a substring of the form $name or ${name}, where name is a non-empty
// sequence of letters, digits, and underscores. A purely numeric name like $1 refers to the sub-match with the
// corresponding index; other names refer to capturing parentheses named with the (?P<name>...) syntax.
// A reference to an out of range or unmatched index or a name that is not present in the regular expression is
// replaced with an empty slice.
//
// In the $name form, name is taken to be as long as possible: $1x is equivalent to ${1x}, not ${1}x, and,
// $10 is equivalent to ${10}, not ${1}0.
//
// To insert a literal $ in the output, use $$ in the template.
func Expand(d, template []byte, r *regexp.Regexp, n int) []byte {
	subMatches := r.FindAllSubmatchIndex(d, n)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.Expand(re, template, d, m)
	}
	return re
}

// ExpandWithPattern is like Expand but instead of a regex, it takes a pattern to match the content.
func ExpandWithPattern(d, template []byte, pattern string, n int) []byte {
	r := regexp.MustCompile(pattern)
	subMatches := r.FindAllSubmatchIndex(d, n)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.Expand(re, template, d, m)
	}
	return re
}

// ExpandString is like Expand but the template and source are strings.
func ExpandString(d, template string, r *regexp.Regexp, n int) string {
	subMatches := r.FindAllStringSubmatchIndex(d, n)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.ExpandString(re, template, d, m)
	}
	return string(re)
}

// ExpandStringWithPattern is like ExpandString but instead of a regex, it takes a pattern to match the content.
func ExpandStringWithPattern(d, template string, pattern string, n int) string {
	r := regexp.MustCompile(pattern)
	subMatches := r.FindAllStringSubmatchIndex(d, n)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.ExpandString(re, template, d, m)
	}
	return string(re)
}

// ExpandAll appends template to the result and during append it replaces variables in the template with all the
// matches drawn from the source data d.
//
// In the template, a variable is denoted by a substring of the form $name or ${name}, where name is a non-empty
// sequence of letters, digits, and underscores. A purely numeric name like $1 refers to the sub-match with the
// corresponding index; other names refer to capturing parentheses named with the (?P<name>...) syntax.
// A reference to an out of range or unmatched index or a name that is not present in the regular expression is
// replaced with an empty slice.
//
// In the $name form, name is taken to be as long as possible: $1x is equivalent to ${1x}, not ${1}x, and,
// $10 is equivalent to ${10}, not ${1}0.
//
// To insert a literal $ in the output, use $$ in the template.
func ExpandAll(d, template []byte, r *regexp.Regexp) []byte {
	subMatches := r.FindAllSubmatchIndex(d, -1)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.Expand(re, template, d, m)
	}
	return re
}

// ExpandAllWithPattern is like ExpandAll but instead of a regex, it takes a pattern to match the content.
func ExpandAllWithPattern(d, template []byte, pattern string) []byte {
	r := regexp.MustCompile(pattern)
	subMatches := r.FindAllSubmatchIndex(d, -1)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.Expand(re, template, d, m)
	}
	return re
}

// ExpandAllString is like ExpandAll but the template and source are strings.
func ExpandAllString(d, template string, r *regexp.Regexp) string {
	subMatches := r.FindAllStringSubmatchIndex(d, -1)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.ExpandString(re, template, d, m)
	}
	return string(re)
}

// ExpandAllStringWithPattern is like ExpandAllString but instead of a regex, it takes a pattern to match the content.
func ExpandAllStringWithPattern(d, template string, pattern string) string {
	r := regexp.MustCompile(pattern)
	subMatches := r.FindAllStringSubmatchIndex(d, -1)
	re := make([]byte, 0)
	for _, m := range subMatches {
		re = r.ExpandString(re, template, d, m)
	}
	return string(re)
}

// Find returns a slice of max n successive matches in d of the regular expression.
// A return value of nil indicates no match.
func Find(d []byte, r *regexp.Regexp, n int) [][]byte {
	return r.FindAll(d, n)
}

// FindWithPattern is like Find but instead of a regex, it takes a pattern to match the content.
func FindWithPattern(d []byte, pattern string, n int) [][]byte {
	return regexp.MustCompile(pattern).FindAll(d, n)
}

// FindString is like Find but the source d is a string.
func FindString(d string, r *regexp.Regexp, n int) []string {
	return r.FindAllString(d, n)
}

// FindStringWithPattern is like FindString but instead of a regex, it takes a pattern to match the content.
func FindStringWithPattern(d string, pattern string, n int) []string {
	return regexp.MustCompile(pattern).FindAllString(d, n)
}

// FindAll returns a slice of all successive matches in d of the regular expression.
func FindAll(d []byte, r *regexp.Regexp) [][]byte {
	return r.FindAll(d, -1)
}

// FindAllWithPattern is like FindAll but instead of a regex, it takes a pattern to match the content.
func FindAllWithPattern(d []byte, pattern string) [][]byte {
	return regexp.MustCompile(pattern).FindAll(d, -1)
}

// FindAllString is like FindAll but the source d is a string.
func FindAllString(d string, r *regexp.Regexp) []string {
	return r.FindAllString(d, -1)
}

// FindAllStringWithPattern is like FindAllString but instead of a regex, it takes a pattern to match the content.
func FindAllStringWithPattern(d string, pattern string) []string {
	return regexp.MustCompile(pattern).FindAllString(d, -1)
}

// FindFirst returns the leftmost match in d of the regular expression.
// A return value of nil indicates no match.
func FindFirst(d []byte, r *regexp.Regexp) []byte {
	return r.Find(d)
}

// FindFirstWithPattern is like FindFirst but instead of a regex, it takes a pattern to match the content.
func FindFirstWithPattern(d []byte, pattern string) []byte {
	return regexp.MustCompile(pattern).Find(d)
}

// FindFirstString is like FindFirst but the source d is a string.
func FindFirstString(d string, r *regexp.Regexp) string {
	return r.FindString(d)
}

// FindFirstStringWithPattern is like FindFirstString but instead of a regex, it takes a pattern to match the content.
func FindFirstStringWithPattern(d string, pattern string) string {
	return regexp.MustCompile(pattern).FindString(d)
}

// FindLast returns the rightmost match in d of the regular expression.
// A return value of nil indicates no match.
func FindLast(d []byte, r *regexp.Regexp) []byte {
	matches := r.FindAll(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastWithPattern is like FindLast but instead of a regex, it takes a pattern to match the content.
func FindLastWithPattern(d []byte, pattern string) []byte {
	matches := regexp.MustCompile(pattern).FindAll(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastString is like FindLast but the source d is a string.
func FindLastString(d string, r *regexp.Regexp) string {
	matches := r.FindAllString(d, -1)
	l := len(matches)
	if l == 0 {
		return ""
	}
	return matches[l-1]
}

// FindLastStringWithPattern is like FindLastString but instead of a regex, it takes a pattern to match the content.
func FindLastStringWithPattern(d string, pattern string) string {
	matches := regexp.MustCompile(pattern).FindAllString(d, -1)
	l := len(matches)
	if l == 0 {
		return ""
	}
	return matches[l-1]
}

// FindIndex returns a slice of max n successive matching indices in d of the regular expression.
// A return value of nil indicates no match.
func FindIndex(d []byte, r *regexp.Regexp, n int) []*pair.Pair[int, int] {
	indices := r.FindAllIndex(d, n)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindIndexWithPattern is like FindIndex but instead of a regex, it takes a pattern to match the content.
func FindIndexWithPattern(d []byte, pattern string, n int) []*pair.Pair[int, int] {
	indices := regexp.MustCompile(pattern).FindAllIndex(d, n)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindIndexForString is like FindIndex but the source d is a string.
func FindIndexForString(d string, r *regexp.Regexp, n int) []*pair.Pair[int, int] {
	indices := r.FindAllStringIndex(d, n)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindIndexForStringWithPattern is like FindIndexForString but instead of a regex, it takes a pattern to match the content.
func FindIndexForStringWithPattern(d string, pattern string, n int) []*pair.Pair[int, int] {
	indices := regexp.MustCompile(pattern).FindAllStringIndex(d, n)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindAllIndex returns a slice of all successive matching indices in d of the regular expression.
func FindAllIndex(d []byte, r *regexp.Regexp) []*pair.Pair[int, int] {
	indices := r.FindAllIndex(d, -1)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindAllIndexWithPattern is like FindAllIndex but instead of a regex, it takes a pattern to match the content.
func FindAllIndexWithPattern(d []byte, pattern string) []*pair.Pair[int, int] {
	indices := regexp.MustCompile(pattern).FindAllIndex(d, -1)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindAllIndexForString is like FindAllIndex but the source d is a string.
func FindAllIndexForString(d string, r *regexp.Regexp) []*pair.Pair[int, int] {
	indices := r.FindAllStringIndex(d, -1)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindAllIndexForStringWithPattern is like FindAllIndexForString but instead of a regex, it takes a pattern to match the content.
func FindAllIndexForStringWithPattern(d string, pattern string) []*pair.Pair[int, int] {
	indices := regexp.MustCompile(pattern).FindAllStringIndex(d, -1)
	l := len(indices)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l)
	for i, idx := range indices {
		result[i] = pair.NewFromCollection(idx)
	}
	return result
}

// FindFirstIndex returns the leftmost matching index in d of the regular expression.
// A return value of nil indicates no match.
func FindFirstIndex(d []byte, r *regexp.Regexp) *pair.Pair[int, int] {
	return pair.NewFromCollection(r.FindIndex(d))
}

// FindFirstIndexWithPattern is like FindFirstIndex but instead of a regex, it takes a pattern to match the content.
func FindFirstIndexWithPattern(d []byte, pattern string) *pair.Pair[int, int] {
	return pair.NewFromCollection(regexp.MustCompile(pattern).FindIndex(d))
}

// FindFirstIndexForString is like FindFirstIndex but the source d is a string.
func FindFirstIndexForString(d string, r *regexp.Regexp) *pair.Pair[int, int] {
	return pair.NewFromCollection(r.FindStringIndex(d))
}

// FindFirstIndexForStringWithPattern is like FindFirstIndexForString but instead of a regex, it takes a pattern to match the content.
func FindFirstIndexForStringWithPattern(d string, pattern string) *pair.Pair[int, int] {
	return pair.NewFromCollection(regexp.MustCompile(pattern).FindStringIndex(d))
}

// FindLastIndex returns the rightmost matching index in d of the regular expression.
// A return value of nil indicates no match.
func FindLastIndex(d []byte, r *regexp.Regexp) *pair.Pair[int, int] {
	matches := r.FindAllIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return pair.NewFromCollection(matches[l-1])
}

// FindLastIndexWithPattern is like FindLastIndex but instead of a regex, it takes a pattern to match the content.
func FindLastIndexWithPattern(d []byte, pattern string) *pair.Pair[int, int] {
	matches := regexp.MustCompile(pattern).FindAllIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return pair.NewFromCollection(matches[l-1])
}

// FindLastIndexForString is like FindLastIndex but the source d is a string.
func FindLastIndexForString(d string, r *regexp.Regexp) *pair.Pair[int, int] {
	matches := r.FindAllStringIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return pair.NewFromCollection(matches[l-1])
}

// FindLastIndexForStringWithPattern is like FindLastIndexForString but instead of a regex, it takes a pattern to match the content.
func FindLastIndexForStringWithPattern(d string, pattern string) *pair.Pair[int, int] {
	matches := regexp.MustCompile(pattern).FindAllStringIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return pair.NewFromCollection(matches[l-1])
}

// FindSubMatches returns a slice of slices holding the text of the n number of
// sub-matches of the regular expression in d and the matches, if any, of its
// subexpressions, as defined by the 'SubMatch' descriptions in the package
// comment.
// A return value of nil indicates no match.
func FindSubMatches(d []byte, r *regexp.Regexp, n int) [][][]byte {
	return r.FindAllSubmatch(d, n)
}

// FindSubMatchesWithPattern is like FindSubMatches but instead of a regex, it takes a pattern to match the content.
func FindSubMatchesWithPattern(d []byte, pattern string, n int) [][][]byte {
	return regexp.MustCompile(pattern).FindAllSubmatch(d, n)
}

// FindSubMatchesForString is like FindSubMatches but the source d is a string.
func FindSubMatchesForString(d string, r *regexp.Regexp, n int) [][]string {
	return r.FindAllStringSubmatch(d, n)
}

// FindSubMatchesForStringWithPattern is like FindSubMatchesForString but instead of a regex, it takes a pattern to match the content.
func FindSubMatchesForStringWithPattern(d string, pattern string, n int) [][]string {
	return regexp.MustCompile(pattern).FindAllStringSubmatch(d, n)
}

// FindAllSubMatches returns a slice of slices holding the text of all
// sub-matches of the regular expression in d and the matches, if any, of its
// subexpressions, as defined by the 'SubMatch' descriptions in the package
// comment.
// A return value of nil indicates no match.
func FindAllSubMatches(d []byte, r *regexp.Regexp) [][][]byte {
	return r.FindAllSubmatch(d, -1)
}

// FindAllSubMatchesWithPattern is like FindAllSubMatches but instead of a regex, it takes a pattern to match the content.
func FindAllSubMatchesWithPattern(d []byte, pattern string) [][][]byte {
	return regexp.MustCompile(pattern).FindAllSubmatch(d, -1)
}

// FindAllSubMatchesForString is like FindAllSubMatches but the source d is a string.
func FindAllSubMatchesForString(d string, r *regexp.Regexp) [][]string {
	return r.FindAllStringSubmatch(d, -1)
}

// FindAllSubMatchesForStringWithPattern is like FindAllSubMatchesForString but instead of a regex,
// it takes a pattern to match the content.
func FindAllSubMatchesForStringWithPattern(d string, pattern string) [][]string {
	return regexp.MustCompile(pattern).FindAllStringSubmatch(d, -1)
}

// FindFirstSubMatch returns a slice of slices holding the text of the leftmost sub-match of the regular
// expression in b and the matches, if any, of its subexpressions, as defined by the 'SubMatch' descriptions
// in the package comment.
// A return value of nil indicates no match.
func FindFirstSubMatch(d []byte, r *regexp.Regexp) [][]byte {
	return r.FindSubmatch(d)
}

// FindFirstSubMatchWithPattern is like FindFirstSubMatch but instead of a regex, it takes a pattern to match the content.
func FindFirstSubMatchWithPattern(d []byte, pattern string) [][]byte {
	return regexp.MustCompile(pattern).FindSubmatch(d)
}

// FindFirstSubMatchForString is like FindFirstSubMatch but the source d is a string.
func FindFirstSubMatchForString(d string, r *regexp.Regexp) []string {
	return r.FindStringSubmatch(d)
}

// FindFirstSubMatchForStringWithPattern is like FindFirstSubMatchForString but instead of a regex,
// it takes a pattern to match the content.
func FindFirstSubMatchForStringWithPattern(d string, pattern string) []string {
	return regexp.MustCompile(pattern).FindStringSubmatch(d)
}

// FindLastSubMatch returns a slice of slices holding the text of the rightmost sub-match of the regular
// expression in b and the matches, if any, of its subexpressions, as defined by the 'SubMatch' descriptions
// in the package comment.
// A return value of nil indicates no match.
func FindLastSubMatch(d []byte, r *regexp.Regexp) [][]byte {
	matches := r.FindAllSubmatch(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastSubMatchWithPattern is like FindLastSubMatch but instead of a regex, it takes a pattern to match the content.
func FindLastSubMatchWithPattern(d []byte, pattern string) [][]byte {
	matches := regexp.MustCompile(pattern).FindAllSubmatch(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastSubMatchForString is like FindLastSubMatch but the source d is a string.
func FindLastSubMatchForString(d string, r *regexp.Regexp) []string {
	matches := r.FindAllStringSubmatch(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastSubMatchForStringWithPattern is like FindLastSubMatchForString but instead of a regex, it takes a pattern to match the content.
func FindLastSubMatchForStringWithPattern(d string, pattern string) []string {
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindSubMatchingIndices returns a slice of slices holding the index pairs identifying the
// n number of sub-matches of the regular expression in d and the matches, if any, of
// its subexpressions, as defined by the 'SubMatch' descriptions in the package comment.
// A return value of nil indicates no match.
func FindSubMatchingIndices(d []byte, r *regexp.Regexp, n int) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(r.FindAllSubmatchIndex(d, n))
}

// FindSubMatchingIndicesWithPattern is like FindSubMatchingIndices but instead of a regex, it takes a pattern to match the content.
func FindSubMatchingIndicesWithPattern(d []byte, pattern string, n int) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(regexp.MustCompile(pattern).FindAllSubmatchIndex(d, n))
}

// FindSubMatchingIndicesForString is like FindSubMatchingIndices but the source d is a string.
func FindSubMatchingIndicesForString(d string, r *regexp.Regexp, n int) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(r.FindAllStringSubmatchIndex(d, n))
}

// FindSubMatchingIndicesForStringWithPattern is like FindSubMatchingIndicesForString but instead of a regex,
// it takes a pattern to match the content.
func FindSubMatchingIndicesForStringWithPattern(d string, pattern string, n int) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(regexp.MustCompile(pattern).FindAllStringSubmatchIndex(d, n))
}

// FindAllSubMatchingIndices returns a slice of slices holding the index pairs identifying the
// all the sub-matches of the regular expression in d and the matches, if any, of
// its subexpressions, as defined by the 'SubMatch' descriptions in the package comment.
// A return value of nil indicates no match.
func FindAllSubMatchingIndices(d []byte, r *regexp.Regexp) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(r.FindAllSubmatchIndex(d, -1))
}

// FindAllSubMatchingIndicesWithPattern is like FindAllSubMatchingIndices but instead of a regex,
// it takes a pattern to match the content.
func FindAllSubMatchingIndicesWithPattern(d []byte, pattern string) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(regexp.MustCompile(pattern).FindAllSubmatchIndex(d, -1))
}

// FindAllSubMatchingIndicesForString is like FindAllSubMatchingIndices but the source d is a string.
func FindAllSubMatchingIndicesForString(d string, r *regexp.Regexp) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(r.FindAllStringSubmatchIndex(d, -1))
}

// FindAllSubMatchingIndicesForStringWithPattern is like FindAllSubMatchingIndicesForString but instead of a regex,
// it takes a pattern to match the content.
func FindAllSubMatchingIndicesForStringWithPattern(d string, pattern string) [][]*pair.Pair[int, int] {
	return getPairedIndicesFromCollections(regexp.MustCompile(pattern).FindAllStringSubmatchIndex(d, -1))
}

// FindFirstSubMatchingIndex returns a slice holding the index pairs identifying the
// leftmost sub-match of the regular expression in b and the matches, if any, of
// its subexpressions, as defined by the 'SubMatch'.
// A return value of nil indicates no match.
func FindFirstSubMatchingIndex(d []byte, r *regexp.Regexp) []*pair.Pair[int, int] {
	return getPairedIndicesFromCollection(r.FindSubmatchIndex(d))
}

// FindFirstSubMatchingIndexWithPattern is like FindFirstSubMatchingIndex but instead of a regex,
// it takes a pattern to match the content.
func FindFirstSubMatchingIndexWithPattern(d []byte, pattern string) []*pair.Pair[int, int] {
	return getPairedIndicesFromCollection(regexp.MustCompile(pattern).FindSubmatchIndex(d))
}

// FindFirstSubMatchingIndexForString is like FindFirstSubMatchingIndex but the source d is a string.
func FindFirstSubMatchingIndexForString(d string, r *regexp.Regexp) []*pair.Pair[int, int] {
	return getPairedIndicesFromCollection(r.FindStringSubmatchIndex(d))
}

// FindFirstSubMatchingIndexForStringWithPattern is like FindFirstSubMatchingIndexForString but instead of a regex,
// it takes a pattern to match the content.
func FindFirstSubMatchingIndexForStringWithPattern(d string, pattern string) []*pair.Pair[int, int] {
	return getPairedIndicesFromCollection(regexp.MustCompile(pattern).FindStringSubmatchIndex(d))
}

// FindLastSubMatchingIndex returns a slice holding the index pairs identifying the
// rightmost sub-match of the regular expression in b and the matches, if any, of
// its subexpressions, as defined by the 'SubMatch'.
// A return value of nil indicates no match.
func FindLastSubMatchingIndex(d []byte, r *regexp.Regexp) []*pair.Pair[int, int] {
	matches := r.FindAllSubmatchIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return getPairedIndicesFromCollection(matches[l-1])
}

// FindLastSubMatchingIndexWithPattern is like FindLastSubMatchingIndex but instead of a regex,
// it takes a pattern to match the content.
func FindLastSubMatchingIndexWithPattern(d []byte, pattern string) []*pair.Pair[int, int] {
	matches := regexp.MustCompile(pattern).FindAllSubmatchIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return getPairedIndicesFromCollection(matches[l-1])
}

// FindLastSubMatchingIndexForString is like FindLastSubMatchingIndex but the source d is a string.
func FindLastSubMatchingIndexForString(d string, r *regexp.Regexp) []*pair.Pair[int, int] {
	matches := r.FindAllStringSubmatchIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return getPairedIndicesFromCollection(matches[l-1])
}

// FindLastSubMatchingIndexForStringWithPattern is like FindLastSubMatchingIndexForString but instead of a regex,
// it takes a pattern to match the content.
func FindLastSubMatchingIndexForStringWithPattern(d string, pattern string) []*pair.Pair[int, int] {
	matches := regexp.MustCompile(pattern).FindAllStringSubmatchIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return getPairedIndicesFromCollection(matches[l-1])
}

// Match reports whether the byte slice d contains any match of the regular expression r.
func Match(d []byte, r *regexp.Regexp) bool {
	return r.Match(d)
}

// MatchWithPattern is like Match but instead of a regex, it takes a pattern to match the content.
func MatchWithPattern(d []byte, pattern string) bool {
	return regexp.MustCompile(pattern).Match(d)
}

// MatchString is like Match but the source d is a string.
func MatchString(d string, r *regexp.Regexp) bool {
	return r.MatchString(d)
}

// MatchStringWithPattern is like MatchString but instead of a regex, it takes a pattern to match the content.
func MatchStringWithPattern(d string, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(d)
}

// RemoveAll is used to remove all the matches of the regex in the given data.
func RemoveAll(d []byte, r *regexp.Regexp) []byte {
	return r.ReplaceAll(d, nil)
}

// RemoveAllWithPattern is used to remove all the matches of the pattern in the given data.
func RemoveAllWithPattern(d []byte, pattern string) []byte {
	return regexp.MustCompile(pattern).ReplaceAll(d, nil)
}

// RemoveAllString is used to remove all the matches of the regex in the given string.
func RemoveAllString(s string, r *regexp.Regexp) string {
	return r.ReplaceAllString(s, "")
}

// RemoveAllStringWithPattern is used to remove all the matches of the pattern in the given string.
func RemoveAllStringWithPattern(s, pattern string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(s, "")
}

// ReplaceAll is used to replace all the matches of the regex in the given data with the replacement.
// Inside replacement, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first sub-match.
func ReplaceAll(d []byte, r *regexp.Regexp, replacement []byte) []byte {
	return r.ReplaceAll(d, replacement)
}

// ReplaceAllWithPattern is used to replace all the matches of the pattern in the given data with the replacement.
// Inside replacement, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first sub-match.
func ReplaceAllWithPattern(d []byte, pattern string, replacement []byte) []byte {
	return regexp.MustCompile(pattern).ReplaceAll(d, replacement)
}

// ReplaceAllString is used to replace all the matches of the regex in the given string with the replacement.
// Inside replacement, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first sub-match.
func ReplaceAllString(s string, r *regexp.Regexp, replacement string) string {
	return r.ReplaceAllString(s, replacement)
}

// ReplaceAllStringWithPattern is used to replace all the matches of the pattern in the given string with the replacement.
// Inside replacement, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first sub-match.
func ReplaceAllStringWithPattern(s, pattern, replacement string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(s, replacement)
}

func getPairedIndicesFromCollections(matches [][]int) [][]*pair.Pair[int, int] {
	if len(matches) == 0 {
		return nil
	}
	result := make([][]*pair.Pair[int, int], len(matches))
	for i, m := range matches {
		result[i] = getPairedIndicesFromCollection(m)
	}
	return result
}

func getPairedIndicesFromCollection(m []int) []*pair.Pair[int, int] {
	l := len(m)
	if l == 0 {
		return nil
	}
	result := make([]*pair.Pair[int, int], l/2)
	j := 0
	k := 1
	x := 0
	for ; j < l && k < l; x += 1 {
		result[x] = pair.New(m[j], m[k])
		j += 2
		k += 2
	}
	return result
}
