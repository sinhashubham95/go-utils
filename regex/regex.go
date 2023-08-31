package regex

import "regexp"

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
func FindIndex(d []byte, r *regexp.Regexp, n int) [][]int {
	return r.FindAllIndex(d, n)
}

// FindIndexWithPattern is like FindIndex but instead of a regex, it takes a pattern to match the content.
func FindIndexWithPattern(d []byte, pattern string, n int) [][]int {
	return regexp.MustCompile(pattern).FindAllIndex(d, n)
}

// FindIndexForString is like FindIndex but the source d is a string.
func FindIndexForString(d string, r *regexp.Regexp, n int) [][]int {
	return r.FindAllStringIndex(d, n)
}

// FindIndexForStringWithPattern is like FindIndexForString but instead of a regex, it takes a pattern to match the content.
func FindIndexForStringWithPattern(d string, pattern string, n int) [][]int {
	return regexp.MustCompile(pattern).FindAllStringIndex(d, n)
}

// FindAllIndex returns a slice of all successive matching indices in d of the regular expression.
func FindAllIndex(d []byte, r *regexp.Regexp) [][]int {
	return r.FindAllIndex(d, -1)
}

// FindAllIndexWithPattern is like FindAllIndex but instead of a regex, it takes a pattern to match the content.
func FindAllIndexWithPattern(d []byte, pattern string) [][]int {
	return regexp.MustCompile(pattern).FindAllIndex(d, -1)
}

// FindAllIndexForString is like FindAllIndex but the source d is a string.
func FindAllIndexForString(d string, r *regexp.Regexp) [][]int {
	return r.FindAllStringIndex(d, -1)
}

// FindAllIndexForStringWithPattern is like FindAllIndexForString but instead of a regex, it takes a pattern to match the content.
func FindAllIndexForStringWithPattern(d string, pattern string) [][]int {
	return regexp.MustCompile(pattern).FindAllStringIndex(d, -1)
}

// FindFirstIndex returns the leftmost matching index in d of the regular expression.
// A return value of nil indicates no match.
func FindFirstIndex(d []byte, r *regexp.Regexp) []int {
	return r.FindIndex(d)
}

// FindFirstIndexWithPattern is like FindFirstIndex but instead of a regex, it takes a pattern to match the content.
func FindFirstIndexWithPattern(d []byte, pattern string) []int {
	return regexp.MustCompile(pattern).FindIndex(d)
}

// FindFirstIndexForString is like FindFirstIndex but the source d is a string.
func FindFirstIndexForString(d string, r *regexp.Regexp) []int {
	return r.FindStringIndex(d)
}

// FindFirstIndexForStringWithPattern is like FindFirstIndexForString but instead of a regex, it takes a pattern to match the content.
func FindFirstIndexForStringWithPattern(d string, pattern string) []int {
	return regexp.MustCompile(pattern).FindStringIndex(d)
}

// FindLastIndex returns the rightmost matching index in d of the regular expression.
// A return value of nil indicates no match.
func FindLastIndex(d []byte, r *regexp.Regexp) []int {
	matches := r.FindAllIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastIndexWithPattern is like FindLastIndex but instead of a regex, it takes a pattern to match the content.
func FindLastIndexWithPattern(d []byte, pattern string) []int {
	matches := regexp.MustCompile(pattern).FindAllIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastIndexForString is like FindLastIndex but the source d is a string.
func FindLastIndexForString(d string, r *regexp.Regexp) []int {
	matches := r.FindAllStringIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// FindLastIndexForStringWithPattern is like FindLastIndexForString but instead of a regex, it takes a pattern to match the content.
func FindLastIndexForStringWithPattern(d string, pattern string) []int {
	matches := regexp.MustCompile(pattern).FindAllStringIndex(d, -1)
	l := len(matches)
	if l == 0 {
		return nil
	}
	return matches[l-1]
}

// RemoveAll is used to remove all the matches of the regex in the given data.
func RemoveAll(d []byte, r *regexp.Regexp) []byte {
	return r.ReplaceAll(d, nil)
}

// RemoveAllWithPattern is used to remove all the matches of the pattern in the given data.
func RemoveAllWithPattern(d []byte, pattern string) []byte {
	return regexp.MustCompile(pattern).ReplaceAll(d, nil)
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

// RemoveAllString is used to remove all the matches of the regex in the given string.
func RemoveAllString(s string, r *regexp.Regexp) string {
	return r.ReplaceAllString(s, "")
}

// RemoveAllStringWithPattern is used to remove all the matches of the pattern in the given string.
func RemoveAllStringWithPattern(s, pattern string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(s, "")
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
