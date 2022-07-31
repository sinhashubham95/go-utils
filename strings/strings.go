package strings

import (
	"fmt"
	"strings"
	"unicode"
)

// Abbreviate is used to abbreviate the given string with the given replacement character.
// An abbreviation is a shortened form of a word or phrase.
func Abbreviate(s, abbreviateMarker string, offset, maxWidth int) string {
	if !IsEmpty(s) && IsEmpty(abbreviateMarker) && maxWidth > 0 {
		return Substring(s, 0, maxWidth)
	}
	if IsAnyEmpty(s, abbreviateMarker) {
		return s
	}
	abbreviateMarkerLength := len(abbreviateMarker)
	minAbbreviateWidth := abbreviateMarkerLength + 1
	minAbbreviateWidthOffset := 2*abbreviateMarkerLength + 1
	if maxWidth < minAbbreviateWidth {
		panic(fmt.Sprintf("minimum abbreviation width is %d", minAbbreviateWidth))
	}
	l := len(s)
	if l <= maxWidth {
		return s
	}
	if offset > l {
		offset = l
	}
	if l-offset < maxWidth-abbreviateMarkerLength {
		offset = l - (maxWidth - abbreviateMarkerLength)
	}
	if offset <= abbreviateMarkerLength+1 {
		return s[0:(maxWidth-abbreviateMarkerLength)] + abbreviateMarker
	}
	if maxWidth < minAbbreviateWidthOffset {
		panic(fmt.Sprintf("minimum abbreviation width with offset is %d", minAbbreviateWidthOffset))
	}
	if offset+maxWidth-abbreviateMarkerLength < l {
		return abbreviateMarker + Abbreviate(s[offset:], abbreviateMarker, 0, maxWidth-abbreviateMarkerLength)
	}
	return abbreviateMarker + s[l-(maxWidth-abbreviateMarkerLength):]
}

// AbbreviateMiddle is used to abbreviate a string to the passed length by replacing the middle characters with the
// supplied replacement string.
//
// The method works only if the following conditions are met.
// The length of the supplied string should be greater than the length of the abbreviated string.
// The length of the abbreviated string should be greater than zero.
func AbbreviateMiddle(s, middle string, length int) string {
	l := len(s)
	lm := len(middle)
	if !IsAnyEmpty(s, middle) && length < l && length >= lm+2 {
		targetStringLength := length - lm
		startOffset := targetStringLength/2 + targetStringLength%2
		endOffset := l - targetStringLength/2
		return s[0:startOffset] + middle + s[endOffset:]
	}
	return s
}

// AppendIfMissing is used to append a given suffix at the end of a string if it does not already end with one
// on the given list.
func AppendIfMissing(s, suffix string, suffixes ...string) string {
	return appendIfMissing(s, suffix, false, suffixes...)
}

// AppendIfMissingIgnoreCase is used to append a given suffix at the end of a string if it does not already end with one
// on the given list ignoring case.
func AppendIfMissingIgnoreCase(s, suffix string, suffixes ...string) string {
	return appendIfMissing(s, suffix, true, suffixes...)
}

// Capitalize used to convert the first character of the given string to upper case.
// The remaining characters of the string are not changed.
func Capitalize(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	f := s[:1]
	u := strings.ToUpper(f)
	if f == u {
		return s
	}
	if l == 1 {
		return u
	}
	return u + s[1:]
}

// Center centers a string in a larger string of given size padding around the given string the given pad character.
func Center(s string, size int, padCharacter uint8) string {
	l := len(s)
	pads := size - l
	if pads > 0 {
		return RightPad(LeftPad(s, l+pads/2, padCharacter), size, padCharacter)
	}
	return s
}

// CenterString centers a string in a larger string of given size padding around the given string the given pad string.
func CenterString(s string, size int, padString string) string {
	if IsEmpty(padString) {
		padString = defaultPadString
	}
	l := len(s)
	pads := size - l
	if pads > 0 {
		return RightPadString(LeftPadString(s, l+pads/2, padString), size, padString)
	}
	return s
}

// Chomp is used to remove the last occurring newline character in a given string. If the string ends with multiple
// newline characters, then only the last occurring newline character is removed. The newline characters are \r, \n or \r\n.
func Chomp(s string) string {
	if IsEmpty(s) {
		return s
	}
	l := len(s)
	if l == 1 && s != carriageReturn && s != lineBreak {
		return s
	}
	if l == 1 {
		return empty
	}
	last := l - 1
	if s[last] == lineBreakCharacter {
		if s[last-1] == carriageReturnCharacter {
			last -= 1
		}
	} else if s[last] != carriageReturnCharacter {
		last += 1
	}
	return Substring(s, 0, last)
}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

// CompareIgnoreCase returns an integer comparing two strings ignoring case lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func CompareIgnoreCase(a, b string) int {
	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
}

// Contains is used to check if the character is present in the string or not
func Contains(s string, searchChar uint8) bool {
	return strings.IndexByte(s, searchChar) >= 0
}

// ContainsIgnoreCase is used to check if the character is present in the string or not ignoring case
func ContainsIgnoreCase(s string, searchChar uint8) bool {
	return strings.IndexByte(strings.ToLower(s), toLower(searchChar)) >= 0
}

// ContainsString is used to check if the string is present in the given string or not
func ContainsString(s, search string) bool {
	return strings.Contains(s, search)
}

// ContainsStringIgnoreCase is used to check if the string is present in the string or not ignoring case
func ContainsStringIgnoreCase(s, search string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(search))
}

// ContainsAny is used to check if any of the characters is present in the string or not
func ContainsAny(s string, searchChars ...uint8) bool {
	joined := join(searchChars...)
	return strings.ContainsAny(s, joined)
}

// ContainsAnyIgnoreCase is used to check if any of the characters is present in the string or not ignoring case
func ContainsAnyIgnoreCase(s string, searchChars ...uint8) bool {
	joined := join(searchChars...)
	return strings.ContainsAny(strings.ToLower(s), strings.ToLower(joined))
}

// ContainsNone is used to check if none of the characters is present in the string or not
func ContainsNone(s string, searchChars ...uint8) bool {
	return !ContainsAny(s, searchChars...)
}

// ContainsNoneIgnoreCase is used to check if none of the characters is present in the string or not ignoring case
func ContainsNoneIgnoreCase(s string, searchChars ...uint8) bool {
	return !ContainsAnyIgnoreCase(s, searchChars...)
}

// ContainsOnly is used to check if the strings contains only the given valid characters
func ContainsOnly(s string, validChars ...uint8) bool {
	joined := join(validChars...)
	for _, c := range s {
		if !Contains(joined, uint8(c)) {
			return false
		}
	}
	return true
}

// ContainsOnlyIgnoreCase is used to check if the strings contains only the given valid characters ignoring case
func ContainsOnlyIgnoreCase(s string, validChars ...uint8) bool {
	joined := join(validChars...)
	for _, c := range s {
		if !ContainsIgnoreCase(joined, uint8(c)) {
			return false
		}
	}
	return true
}

// CountMatches is used to count the number of occurrences of a character or a substring in a larger string/text.
// Here, character or substring matching is case-sensitive in nature.
func CountMatches(s string, ch uint8) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < l; i += 1 {
		if s[l] == ch {
			cnt++
		}
	}
	return cnt
}

// DefaultIfBlank takes 2 parameters a string and a default string. If the passed string is blank, then it returns the
// default string otherwise it returns the passed string.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func DefaultIfBlank(s, df string) string {
	if IsBlank(s) {
		return df
	}
	return s
}

// DefaultIfEmpty takes 2 parameters a string and a default string. If the passed string is empty, then it returns the
// default string otherwise it returns the passed string.
func DefaultIfEmpty(s, df string) string {
	if IsEmpty(s) {
		return df
	}
	return s
}

// DeleteWhitespaces removes all the whitespace characters in a given string.
func DeleteWhitespaces(s string) string {
	var b strings.Builder
	for _, c := range s {
		if !unicode.IsSpace(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// Difference is used to compare two strings and return the remaining characters of the second string that differ
// from the first string.
func Difference(a, b string) string {
	if a == b {
		return ""
	}
	al := len(a)
	bl := len(b)
	var i int
	for i = 0; i < al && i < bl && a[i] == b[i]; i += 1 {
	}
	if i >= bl && i >= al {
		return ""
	}
	return Substring(b, i, bl)
}

// EndsWith checks whether the given string ends with the given string/suffix. This method is case-sensitive when
// comparing the suffix with the end of the string.
func EndsWith(s, suffix string) bool {
	return endsWith(s, suffix, false)
}

// EndsWithIgnoreCase checks if the given string ends with the given string/suffix. This method is case-insensitive
// when comparing the suffix with the end of the string.
func EndsWithIgnoreCase(s, suffix string) bool {
	return endsWith(s, suffix, true)
}

// EndsWithAny checks whether the given string ends with the given any of the given strings/suffixes.
// This method is case-sensitive when comparing the suffix with the end of the string.
func EndsWithAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if EndsWith(s, suffix) {
			return true
		}
	}
	return false
}

// EndsWithAnyIgnoreCase checks whether the given string ends with the given any of the given strings/suffixes.
// This method is case-insensitive when comparing the suffix with the end of the string.
func EndsWithAnyIgnoreCase(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if EndsWithIgnoreCase(s, suffix) {
			return true
		}
	}
	return false
}

// Equals is used to check if the 2 strings are equal or not.
func Equals(a, b string) bool {
	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}
	for i := 0; i < al; i += 1 {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualsIgnoreCase is used to check if the 2 strings are equal or not ignoring case.
func EqualsIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}

// EqualsAny is used to check if a string is equal to any of the given strings or not
func EqualsAny(a string, bb ...string) bool {
	for _, b := range bb {
		if Equals(a, b) {
			return true
		}
	}
	return false
}

// EqualsAnyIgnoreCase is used to check if a string is equal to any of the given strings or not ignoring case.
func EqualsAnyIgnoreCase(a string, bb ...string) bool {
	for _, b := range bb {
		if EqualsIgnoreCase(a, b) {
			return true
		}
	}
	return false
}

// FirstNonBlank returns the first element in the given list of elements that is not empty, null, or whitespace only.
// If all the elements are blank or the list is null or empty, null is returned.
func FirstNonBlank(ss ...string) string {
	for _, s := range ss {
		if !IsBlank(s) {
			return s
		}
	}
	return ""
}

// FirstNonEmpty returns the first element in a given list of elements that is not empty. If all the elements are null
// or empty then null is returned.
func FirstNonEmpty(ss ...string) string {
	for _, s := range ss {
		if !IsEmpty(s) {
			return s
		}
	}
	return ""
}

// IsBlank is used to check if the given string is blank.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func IsBlank(s string) bool {
	return IsEmpty(strings.TrimSpace(s))
}

// IsAnyBlank is used to check if any of the given strings is blank or not.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func IsAnyBlank(ss ...string) bool {
	for _, s := range ss {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsEmpty is used to check if a given string is empty or not.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsAnyEmpty is used to check if any of the given strings is empty or not.
func IsAnyEmpty(ss ...string) bool {
	for _, s := range ss {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

func LeftPad(s string, size int, padCharacter uint8) string {
	l := len(s)
	pads := size - l
	if pads <= 0 {
		return s
	}
	if pads > maxPadsRepeatSize {
		return LeftPadString(s, size, string(padCharacter))
	}
	return Repeat(padCharacter, pads) + s
}

func LeftPadString(s string, size int, padString string) string {
	if IsEmpty(padString) {
		padString = defaultPadString
	}
	l := len(s)
	pl := len(padString)
	pads := size - l
	if pads <= 0 {
		return s
	}
	if pl == 1 && pads <= 8192 {
		return LeftPad(s, size, padString[0])
	}
	if pl == pads {
		return padString + s
	}
	if pads < pl {
		return padString[0:pads] + s
	}
	p := make([]uint8, pads)
	for i := 0; i < pads; i += 1 {
		p[i] = padString[i%pl]
	}
	return string(p) + s
}

func Repeat(character uint8, repeat int) string {
	if repeat <= 0 {
		return ""
	}
	s := make([]uint8, repeat)
	for i := 0; i < repeat; i += 1 {
		s[i] = character
	}
	return string(s)
}

func RightPad(s string, size int, padCharacter uint8) string {
	l := len(s)
	pads := size - l
	if pads <= 0 {
		return s
	}
	if pads > maxPadsRepeatSize {
		return RightPadString(s, size, string(padCharacter))
	}
	return s + Repeat(padCharacter, pads)
}

func RightPadString(s string, size int, padString string) string {
	if IsEmpty(padString) {
		padString = defaultPadString
	}
	l := len(s)
	pl := len(padString)
	pads := size - l
	if pads <= 0 {
		return s
	}
	if pl == 1 && pads <= 8192 {
		return RightPad(s, size, padString[0])
	}
	if pl == pads {
		return s + padString
	}
	if pads < pl {
		return s + padString[0:pads]
	}
	p := make([]uint8, pads)
	for i := 0; i < pads; i += 1 {
		p[i] = padString[i%pl]
	}
	return s + string(p)
}

func SubstringTillEnd(s string, start int) string {
	l := len(s)
	if start < 0 {
		start += l
	}
	if start < 0 {
		start = 0
	}
	if start > l {
		return ""
	}
	return s[start:]
}

func Substring(s string, start, end int) string {
	l := len(s)
	if end < 0 {
		end += l
	}
	if start < 0 {
		start += l
	}
	if end > l {
		end = l
	}
	if start > end {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	return s[start:end]
}

func appendIfMissing(s, suffix string, ignoreCase bool, suffixes ...string) string {
	if !IsEmpty(suffix) && !endsWith(s, suffix, ignoreCase) {
		for _, su := range suffixes {
			if endsWith(s, su, ignoreCase) {
				return s
			}
		}
		return s + suffix
	}
	return s
}

func endsWith(s, suffix string, ignoreCase bool) bool {
	l := len(s)
	sl := len(suffix)
	if sl > l {
		return false
	}
	offset := l - sl
	if ignoreCase {
		return strings.EqualFold(s[offset:], suffix)
	}
	return s[offset:] == suffix
}

func join(chars ...uint8) string {
	var b strings.Builder
	for _, c := range chars {
		b.WriteByte(c)
	}
	return b.String()
}

func toLower(ch uint8) uint8 {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 'a' - 'A'
	}
	return ch
}
