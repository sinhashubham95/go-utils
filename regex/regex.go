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
