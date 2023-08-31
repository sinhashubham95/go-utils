package regex

import "regexp"

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
