package strings

import (
	"context"
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
	if offset <= minAbbreviateWidth {
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
	} else if s[last] == carriageReturnCharacter {
		if s[last-1] == lineBreakCharacter {
			last -= 1
		}
	} else {
		last += 1
	}
	return Substring(s, 0, last)
}

// Chop is used to remove the last character and the second last character if it's newline.
func Chop(s string) string {
	l := len(s)
	if l <= 2 {
		return empty
	}
	lastIdx := l - 1
	ret := Substring(s, 0, lastIdx)
	last := s[lastIdx]
	if last == '\n' && ret[lastIdx-1] == '\r' {
		return Substring(ret, 0, lastIdx-1)
	}
	return ret
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

// ContainsWhitespace is used to check if the string contains whitespace
func ContainsWhitespace(s string) bool {
	return Contains(s, ' ')
}

// Count is used to count the number of occurrences of a character or a substring in a larger string/text.
// Here, character or substring matching is case-sensitive in nature.
func Count(s string, ch uint8) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < l; i += 1 {
		if s[i] == ch {
			cnt++
		}
	}
	return cnt
}

// CountIgnoreCase is used to count the number of occurrences of a character or a substring in a larger string/text.
// Here, character or substring matching is not case-sensitive in nature.
func CountIgnoreCase(s string, ch uint8) int {
	return Count(LowerCase(s), toLower(ch))
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
	i := IndexOfDifference(a, b)
	if i == -1 {
		la := len(a)
		lb := len(b)
		if la < lb {
			return SubstringTillEnd(b, la)
		}
		return ""
	}
	return Substring(b, i, len(b))
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

// GetCommonPrefix compares all strings in an array of strings and returns the common starting character sequence
// that all of them share.
func GetCommonPrefix(ss ...string) string {
	l := len(ss)
	minL := getMinimumLength(ss...)
	var b strings.Builder
	for i := 0; i < minL; i += 1 {
		c := ss[0][i]
		match := true
		for j := 1; j < l; j += 1 {
			if c != ss[j][i] {
				match = false
				break
			}
		}
		if !match {
			break
		}
		b.WriteByte(c)
	}
	return b.String()
}

// GetDigits checks for Unicode digits in a string and returns a new string that contains all the digits in the given string.
// If the given string does not contain any digits, an empty string will be returned.
func GetDigits(s string) string {
	var b strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// GetIfBlank is used to call the function and get the data if the provided string is blank.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func GetIfBlank(ctx context.Context, s string, supplier func(ctx context.Context) string) string {
	if IsBlank(s) {
		if supplier == nil {
			return empty
		}
		return supplier(ctx)
	}
	return s
}

// GetIfEmpty is used to call the function and get the data if the provided string is empty.
func GetIfEmpty(ctx context.Context, s string, supplier func(ctx context.Context) string) string {
	if IsEmpty(s) {
		if supplier == nil {
			return empty
		}
		return supplier(ctx)
	}
	return s
}

// IndexOf is used to get the index of the first occurrence of the given character. If the character is not found,
// then it will return -1.
func IndexOf(s string, search uint8) int {
	return strings.IndexByte(s, search)
}

// IndexOfIgnoreCase is used to get the index of the first occurrence of the given character ignoring case.
// If the character is not found, then it will return -1.
func IndexOfIgnoreCase(s string, search uint8) int {
	return strings.IndexFunc(s, func(r rune) bool {
		return unicode.ToLower(r) == unicode.ToLower(rune(search))
	})
}

// IndexOfString is used to get the index of the first occurrence of the given string. If the character is not found,
// then it will return -1.
func IndexOfString(s, search string) int {
	return strings.Index(s, search)
}

// IndexOfStringIgnoreCase is used to get the index of the first occurrence of the given string ignoring case.
// If the string is not found, then it will return -1.
func IndexOfStringIgnoreCase(s, search string) int {
	return strings.Index(strings.ToLower(s), strings.ToLower(search))
}

// IndexOfAny is used to get the index of the first occurrence of one of the given characters. If the character is not found,
// then it will return -1.
func IndexOfAny(s string, searchChars ...uint8) int {
	return strings.IndexFunc(s, func(r rune) bool {
		for _, c := range searchChars {
			if rune(c) == r {
				return true
			}
		}
		return false
	})
}

// IndexOfAnyIgnoreCase is used to get the index of the first occurrence of one of the given characters ignoring case.
// If the character is not found, then it will return -1.
func IndexOfAnyIgnoreCase(s string, searchChars ...uint8) int {
	return strings.IndexFunc(s, func(r rune) bool {
		for _, c := range searchChars {
			if unicode.ToLower(r) == unicode.ToLower(rune(c)) {
				return true
			}
		}
		return false
	})
}

// IndexOfAnyBut is used to get the index of the first occurrence of any character other than one of the given characters.
// If the character is not found, then it will return -1.
func IndexOfAnyBut(s string, searchChars ...uint8) int {
	return strings.IndexFunc(s, func(r rune) bool {
		found := false
		for _, c := range searchChars {
			if rune(c) == r {
				found = true
				break
			}
		}
		return !found
	})
}

// IndexOfAnyButIgnoreCase is used to get the index of the first occurrence of any character other than one of the
// given characters ignoring case. If the character is not found, then it will return -1.
func IndexOfAnyButIgnoreCase(s string, searchChars ...uint8) int {
	return strings.IndexFunc(s, func(r rune) bool {
		found := false
		for _, c := range searchChars {
			if unicode.ToLower(rune(c)) == unicode.ToLower(r) {
				found = true
				break
			}
		}
		return !found
	})
}

// IndexOfDifference figures out the index of the first character sequence that differs from the given sequence.
// The comparison of the character sequences in the given text is case-sensitive.
func IndexOfDifference(ss ...string) int {
	l := len(ss)
	if l == 0 {
		return -1
	}
	minL := getMinimumLength(ss...)
	for i := 0; i < minL; i += 1 {
		c := ss[0][i]
		match := true
		for j := 1; j < l; j += 1 {
			if c != ss[j][i] {
				match = false
				break
			}
		}
		if !match {
			return i
		}
	}
	return -1
}

// IndexOfStringStartingAt is used to get the index of the first occurrence of the given string starting with index provided.
// If the character is not found, then it will return -1.
func IndexOfStringStartingAt(s, search string, start int) int {
	l := len(s)
	if start >= l {
		return -1
	}
	idx := IndexOfString(SubstringTillEnd(s, start), search)
	if idx < 0 {
		return idx
	}
	return idx + start
}

// LastIndexOf is used to get the index of the first occurrence of the given character. If the character is not found,
// then it will return -1.
func LastIndexOf(s string, search uint8) int {
	return strings.LastIndexByte(s, search)
}

// LastIndexOfIgnoreCase is used to get the index of the first occurrence of the given character ignoring case.
// If the character is not found, then it will return -1.
func LastIndexOfIgnoreCase(s string, search uint8) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		return unicode.ToLower(r) == unicode.ToLower(rune(search))
	})
}

// LastIndexOfString is used to get the index of the first occurrence of the given string. If the character is not found,
// then it will return -1.
func LastIndexOfString(s, search string) int {
	return strings.LastIndex(s, search)
}

// LastIndexOfStringIgnoreCase is used to get the index of the first occurrence of the given string ignoring case.
// If the string is not found, then it will return -1.
func LastIndexOfStringIgnoreCase(s, search string) int {
	return strings.LastIndex(strings.ToLower(s), strings.ToLower(search))
}

// LastIndexOfAny is used to get the index of the first occurrence of one of the given characters. If the character is not found,
// then it will return -1.
func LastIndexOfAny(s string, searchChars ...uint8) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		for _, c := range searchChars {
			if rune(c) == r {
				return true
			}
		}
		return false
	})
}

// LastIndexOfAnyIgnoreCase is used to get the index of the first occurrence of one of the given characters ignoring case.
// If the character is not found, then it will return -1.
func LastIndexOfAnyIgnoreCase(s string, searchChars ...uint8) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		for _, c := range searchChars {
			if unicode.ToLower(r) == unicode.ToLower(rune(c)) {
				return true
			}
		}
		return false
	})
}

// LastIndexOfAnyBut is used to get the index of the first occurrence of any character other than one of the given characters.
// If the character is not found, then it will return -1.
func LastIndexOfAnyBut(s string, searchChars ...uint8) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		found := false
		for _, c := range searchChars {
			if rune(c) == r {
				found = true
				break
			}
		}
		return !found
	})
}

// LastIndexOfAnyButIgnoreCase is used to get the index of the first occurrence of any character other than one of the
// given characters ignoring case. If the character is not found, then it will return -1.
func LastIndexOfAnyButIgnoreCase(s string, searchChars ...uint8) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		found := false
		for _, c := range searchChars {
			if unicode.ToLower(rune(c)) == unicode.ToLower(r) {
				found = true
				break
			}
		}
		return !found
	})
}

// IsBlank is used to check if the given string is blank.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func IsBlank(s string) bool {
	return IsEmpty(strings.TrimSpace(s))
}

// IsAllBlank is used to check if the all the given strings are blank.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func IsAllBlank(ss ...string) bool {
	for _, s := range ss {
		if !IsBlank(s) {
			return false
		}
	}
	return true
}

// IsNoneBlank is used to check if none of the given strings are blank.
//
// A string is considered to be blank if it satisfied one of the criteria below.
// Length of string is 0.
// String contains only whitespace characters.
func IsNoneBlank(ss ...string) bool {
	for _, s := range ss {
		if IsBlank(s) {
			return false
		}
	}
	return true
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

// IsAllEmpty is used to check if the all the given strings are empty.
func IsAllEmpty(ss ...string) bool {
	for _, s := range ss {
		if !IsEmpty(s) {
			return false
		}
	}
	return true
}

// IsNoneEmpty is used to check if none of the given strings are empty.
func IsNoneEmpty(ss ...string) bool {
	for _, s := range ss {
		if IsEmpty(s) {
			return false
		}
	}
	return true
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

// IsLowerCase is used to check if all characters in the string is in lower case.
func IsLowerCase(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAllLowerCase is used to check if all characters in the all the given strings is in lower case.
func IsAllLowerCase(ss ...string) bool {
	for _, s := range ss {
		if !IsLowerCase(s) {
			return false
		}
	}
	return true
}

// IsAnyLowerCase is used to check if all characters in any of the given strings is in lower case.
func IsAnyLowerCase(ss ...string) bool {
	for _, s := range ss {
		if IsLowerCase(s) {
			return true
		}
	}
	return false
}

// IsUpperCase is used to check if all characters in the string is in upper case.
func IsUpperCase(s string) bool {
	for _, c := range s {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

// IsAllUpperCase is used to check if all characters in the all the given strings is in upper case.
func IsAllUpperCase(ss ...string) bool {
	for _, s := range ss {
		if !IsUpperCase(s) {
			return false
		}
	}
	return true
}

// IsAnyUpperCase is used to check if all characters in any of the given strings is in upper case.
func IsAnyUpperCase(ss ...string) bool {
	for _, s := range ss {
		if IsUpperCase(s) {
			return true
		}
	}
	return false
}

// IsMixedCase is used to check if it contains both upper and lower case characters.
func IsMixedCase(s string) bool {
	isUpper := false
	isLower := false
	for _, c := range s {
		if unicode.IsUpper(c) {
			isUpper = true
		}
		if unicode.IsLower(c) {
			isLower = true
		}
	}
	return isUpper && isLower
}

// IsAllMixedCase is used to check if all the given strings contain both upper and lower case characters.
func IsAllMixedCase(ss ...string) bool {
	for _, s := range ss {
		if !IsMixedCase(s) {
			return false
		}
	}
	return true
}

// IsAnyMixedCase is used to check if any of the given strings contain both upper and lower case characters.
func IsAnyMixedCase(ss ...string) bool {
	for _, s := range ss {
		if IsMixedCase(s) {
			return true
		}
	}
	return false
}

// IsAlpha checks whether a given string contains only letters. The function returns false if the
// input string is empty.
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsNumeric checks whether a given string contains only digits. The function returns false if the
// input string is empty.
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric checks whether a given string contains only letters or digits. The function returns false if the
// input string is empty.
func IsAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaSpace checks whether a given string contains only letters or space. The function returns false if the
// input string is empty.
func IsAlphaSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsNumericSpace checks whether a given string contains only digits or space. The function returns false if the
// input string is empty.
func IsNumericSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAlphaNumericSpace checks whether a given string contains only letters or digits or space. The function returns false if the
// input string is empty.
func IsAlphaNumericSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAsciiPrintable is used to check if the given string contains only ASCII characters that are printable.
func IsAsciiPrintable(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !(c >= ' ' && c < 127) {
			return false
		}
	}
	return true
}

// JoinByChar is used to join the characters by a character.
func JoinByChar(cc []uint8, d uint8) string {
	var b strings.Builder
	for i, c := range cc {
		if i > 0 {
			b.WriteByte(d)
		}
		b.WriteByte(c)
	}
	return b.String()
}

// JoinStringsByChar is used to join the strings by a character.
func JoinStringsByChar(ss []string, d uint8) string {
	var b strings.Builder
	for i, s := range ss {
		if i > 0 {
			b.WriteByte(d)
		}
		b.WriteString(s)
	}
	return b.String()
}

// JoinStringsByString is used to join the strings by a string.
func JoinStringsByString(ss []string, ds string) string {
	return strings.Join(ss, ds)
}

// Left is used to get the specified number of leftmost characters of a string.
func Left(s string, l int) string {
	if l < 0 {
		return ""
	}
	if len(s) < l {
		return s
	}
	return Substring(s, 0, l)
}

// LeftPad is used to left pad the given character to the given string.
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

// LeftPadString is used to left pad the given string to the given string.
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

// LowerCase is used to convert a string to lower case.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// Mid is used to return the string starting from the given position index upto the given length.
func Mid(s string, pos, length int) string {
	l := len(s)
	if length < 0 && pos > l {
		return ""
	}
	if pos < 0 {
		pos = 0
	}
	if l <= pos+length {
		return SubstringTillEnd(s, pos)
	}
	return Substring(s, pos, pos+length)
}

// NormalizeSpace is used to return the whitespace normalized string by removing the leading and trailing whitespace
// and then replacing sequences of whitespace characters with a single space.
func NormalizeSpace(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	b := strings.Builder{}
	whitespacesCount := 0
	startWhitespaces := true
	for i := 0; i < l; i += 1 {
		isWhitespace := unicode.IsSpace(rune(s[i]))
		if isWhitespace {
			if whitespacesCount == 0 && !startWhitespaces {
				b.WriteRune(' ')
				whitespacesCount++
			}
			whitespacesCount++
		} else {
			startWhitespaces = false
			b.WriteByte(s[i])
			whitespacesCount = 0
		}
	}
	if whitespacesCount > 0 {
		return b.String()[:b.Len()-1]
	}
	return b.String()
}

// OrdinalIndexOf is used to return the index of the nth occurrence of the search string in the given string
func OrdinalIndexOf(s, search string, ordinal int) int {
	found := 0
	idx := -1
	for {
		if found >= ordinal {
			break
		}
		idx = IndexOfStringStartingAt(s, search, idx+1)
		if idx < 0 {
			return idx
		}
		found += 1
	}
	return idx
}

// Overlay overlays part of a String with another String.
// A negative index is treated as zero. An index greater than the string length is treated as the string length.
// The start index is always the smaller of the two indices.
func Overlay(s, overlay string, start, end int) string {
	l := len(s)
	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	if start > end {
		temp := start
		start = end
		end = temp
	}
	return Substring(s, 0, start) + overlay + SubstringTillEnd(s, end)
}

// PrependIfMissing is used to prepend a given prefix at the start of a string if it does not already start with one
// on the given list.
func PrependIfMissing(s, prefix string, prefixes ...string) string {
	return prependIfMissing(s, prefix, false, prefixes...)
}

// PrependIfMissingIgnoreCase is used to prepend a given prefix at the start of a string if it does not already start with one
// on the given list ignoring case.
func PrependIfMissingIgnoreCase(s, prefix string, prefixes ...string) string {
	return prependIfMissing(s, prefix, true, prefixes...)
}

// Remove is used to remove all occurrences of a given character in the string
func Remove(s string, remove uint8) string {
	l := len(s)
	var b strings.Builder
	for i := 0; i < l; i += 1 {
		if s[i] != remove {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// RemoveIgnoreCase is used to remove all occurrences of a given character in the string ignoring case
func RemoveIgnoreCase(s string, remove uint8) string {
	ls := strings.ToLower(s)
	r := unicode.ToLower(rune(remove))
	var b strings.Builder
	for i, c := range ls {
		if c != r {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// RemoveString is used to remove all occurrences of a given string in the string
func RemoveString(s, remove string) string {
	if IsEmpty(s) || IsEmpty(remove) {
		return s
	}
	return replace(s, remove, "", -1, false)
}

// RemoveStringIgnoreCase is used to remove all occurrences of a given string in the string ignoring case
func RemoveStringIgnoreCase(s, remove string) string {
	if IsEmpty(s) || IsEmpty(remove) {
		return s
	}
	return replace(s, remove, "", -1, true)
}

// RemoveStart is used to remove the occurrence of the given string from the given string's start
func RemoveStart(s, remove string) string {
	if startsWith(s, remove, false) {
		return SubstringTillEnd(s, len(remove))
	}
	return s
}

// RemoveStartIgnoreCase is used to remove the occurrence of the given string's start from the given string ignoring case
func RemoveStartIgnoreCase(s, remove string) string {
	if startsWith(s, remove, true) {
		return SubstringTillEnd(s, len(remove))
	}
	return s
}

// RemoveEnd is used to remove the occurrence of the given string from the given string's end
func RemoveEnd(s, remove string) string {
	if endsWith(s, remove, false) {
		return Substring(s, 0, len(s)-len(remove))
	}
	return s
}

// RemoveEndIgnoreCase is used to remove the occurrence of the given string from the given string's end ignoring case
func RemoveEndIgnoreCase(s, remove string) string {
	if endsWith(s, remove, true) {
		return Substring(s, 0, len(s)-len(remove))
	}
	return s
}

// Repeat is used to repeat the given character the given number of times.
func Repeat(c uint8, repeat int) string {
	if repeat <= 0 {
		return ""
	}
	s := make([]uint8, repeat)
	for i := 0; i < repeat; i += 1 {
		s[i] = c
	}
	return string(s)
}

// RepeatString is used to repeat the given string the given number of times.
func RepeatString(s string, repeat int) string {
	if repeat <= 0 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < repeat; i += 1 {
		b.WriteString(s)
	}
	return b.String()
}

// RepeatStringWithSeparator is used to repeat the given string the given number of times each separated by the given separator.
func RepeatStringWithSeparator(s, separator string, repeat int) string {
	if repeat <= 0 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < repeat; i += 1 {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(s)
	}
	return b.String()
}

// Replace is used to replace the given number of occurrences of a string in the given string with another string.
func Replace(s, search, replacement string, max int) string {
	return replace(s, search, replacement, max, false)
}

// ReplaceAll is used to replace all occurrences of a string in the given string with another string.
func ReplaceAll(s, search, replacement string) string {
	return Replace(s, search, replacement, -1)
}

// ReplaceEach is used to replace each of the
func ReplaceEach(s string, searchList, replacementList []string) string {
	if s == empty || len(searchList) == 0 {
		return s
	}
	if len(searchList) != len(replacementList) {
		panic("search and replacement list lengths don't match")
	}
	l := len(searchList)
	for i := 0; i < l; i += 1 {
		s = Replace(s, searchList[i], replacementList[i], 1)
	}
	return s
}

// ReplaceIgnoreCase is used to replace the given number of occurrences of a string in the given string with
// another string ignoring case.
func ReplaceIgnoreCase(s, search, replacement string, max int) string {
	return replace(s, search, replacement, max, true)
}

// Reverse is used to reverse the given string
func Reverse(s string) string {
	if s == empty {
		return s
	}
	ret := empty
	l := len(s)
	for i := l - 1; i >= 0; i -= 1 {
		ret += string(s[i])
	}
	return ret
}

// ReverseDelimited is used to reverse the given string split by the delimiter
func ReverseDelimited(s string, d uint8) string {
	split := strings.Split(s, string(d))
	l := len(split)
	retSplit := make([]string, l)
	for i := 0; i < l; i += 1 {
		retSplit[i] = split[l-i-1]
	}
	return JoinStringsByChar(retSplit, d)
}

// Right is used to get the specified number of rightmost characters of a string.
func Right(s string, l int) string {
	if l < 0 {
		return ""
	}
	sl := len(s)
	if sl < l {
		return s
	}
	return Substring(s, sl-l, sl)
}

// RightPad is used to right pad the given character to the given string.
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

// RightPadString is used to right pad the given string to the given string.
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

// Rotate is used to rotate the string by given number of characters
func Rotate(s string, shift int) string {
	if s == empty {
		return s
	}
	l := len(s)
	if shift > 0 && shift%l > 0 {
		offset := -(shift % l)
		return SubstringTillEnd(s, offset) + Substring(s, 0, offset)
	}
	return s
}

// Split is used to split the string into a list of strings about the separator provided
func Split(s string, separator uint8) []string {
	if s == empty {
		return nil
	}
	return strings.Split(s, string(separator))
}

// SplitByString is used to split the string by the separator provided
func SplitByString(s, separator string) []string {
	if s == empty {
		return nil
	}
	return strings.Split(s, separator)
}

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitN(s string, separator uint8, n int) []string {
	if s == empty {
		return nil
	}
	return strings.SplitN(s, string(separator), n)
}

// SplitNByString slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitNByString(s, separator string, n int) []string {
	if s == empty {
		return nil
	}
	return strings.SplitN(s, separator, n)
}

// SplitWithTrim is used to split and trim the splits
func SplitWithTrim(s string, separator uint8) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.Split(s, string(separator)), space)
}

// SplitByStringWithTrim is used to split the string by the separator provided and trim the splits
func SplitByStringWithTrim(s, separator string) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.Split(s, separator), space)
}

// SplitNWithTrim slices s into substrings separated by sep and returns a slice of
// the substrings between those separators and along with that trim the splits.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitNWithTrim(s string, separator uint8, n int) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.SplitN(s, string(separator), n), space)
}

// SplitNByStringWithTrim slices s into substrings separated by sep and returns a slice of
// the substrings between those separators and along with that trim the splits.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitNByStringWithTrim(s, separator string, n int) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.SplitN(s, separator, n), space)
}

// SplitWithTrimCutSet is used to split and trailing Unicode code points contained in cut set removed.
func SplitWithTrimCutSet(s string, separator uint8, set string) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.Split(s, string(separator)), set)
}

// SplitByStringWithTrimCutSet is used to split the string by the separator provided and trailing Unicode code points
// contained in cut set removed.
func SplitByStringWithTrimCutSet(s, separator string, set string) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.Split(s, separator), set)
}

// SplitNWithTrimCutSet slices s into substrings separated by sep and returns a slice of
// the substrings between those separators and trailing Unicode code points contained in cut set removed.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitNWithTrimCutSet(s string, separator uint8, n int, set string) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.SplitN(s, string(separator), n), set)
}

// SplitNByStringWithTrimCutSet slices s into substrings separated by sep and returns a slice of
// the substrings between those separators and trailing Unicode code points contained in cut set removed.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the remainder which won't be split.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitNByStringWithTrimCutSet(s, separator string, n int, set string) []string {
	if s == empty {
		return nil
	}
	return trimSplits(strings.SplitN(s, separator, n), set)
}

// StartsWith is used to check if the provided prefix is present at the beginning of the string or not.
func StartsWith(s, prefix string) bool {
	return startsWith(s, prefix, false)
}

// StartsWithIgnoreCase is used to check if the provided prefix is present at the beginning of the
// string or not ignoring case.
func StartsWithIgnoreCase(s, prefix string) bool {
	return startsWith(s, prefix, true)
}

// StartsWithAny is used to check if any of the provided prefixes is present at the beginning of the string or not.
func StartsWithAny(s string, prefixes ...string) bool {
	if len(prefixes) == 0 {
		return false
	}
	for _, prefix := range prefixes {
		if startsWith(s, prefix, false) {
			return true
		}
	}
	return false
}

// StartsWithAnyIgnoreCase is used to check if any of the provided prefixes is present at the beginning of the
// string or not ignoring case.
func StartsWithAnyIgnoreCase(s string, prefixes ...string) bool {
	if len(prefixes) == 0 {
		return false
	}
	for _, prefix := range prefixes {
		if startsWith(s, prefix, true) {
			return true
		}
	}
	return false
}

// StripStart is used to remove the given set of characters from the starting of the string
func StripStart(s, cc string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	var i int
	for i = 0; i < l; i += 1 {
		if !Contains(cc, s[i]) {
			break
		}
	}
	if i == l {
		return empty
	}
	return SubstringTillEnd(s, i)
}

// StripEnd is used to remove the given set of characters from the ending of the string
func StripEnd(s, cc string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	var i int
	for i = l - 1; i >= 0; i -= 1 {
		if !Contains(cc, s[i]) {
			break
		}
	}
	if i == -1 {
		return empty
	}
	return Substring(s, 0, i+1)
}

// Strip is used to remove the given set of characters from the starting and ending of the string
func Strip(s, cc string) string {
	return StripEnd(StripStart(s, cc), cc)
}

// StripAllStart is used to remove the given set of characters from the starting of the given list of strings
func StripAllStart(ss []string, cc string) []string {
	l := len(ss)
	if l == 0 {
		return ss
	}
	res := make([]string, l)
	for i, s := range ss {
		res[i] = StripStart(s, cc)
	}
	return res
}

// StripAllEnd is used to remove the given set of characters from the ending of the given list of strings
func StripAllEnd(ss []string, cc string) []string {
	l := len(ss)
	if l == 0 {
		return ss
	}
	res := make([]string, l)
	for i, s := range ss {
		res[i] = StripEnd(s, cc)
	}
	return res
}

// StripAll is used to remove the given set of characters from the starting and ending of the given list of strings
func StripAll(ss []string, cc string) []string {
	l := len(ss)
	if l == 0 {
		return ss
	}
	res := make([]string, l)
	for i, s := range ss {
		res[i] = StripEnd(StripStart(s, cc), cc)
	}
	return res
}

// Substring returns the string between the given start(inclusive) and the end(exclusive) index.
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

// SubstringTillEnd returns the string from the given start index till the end of the string.
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

// SubstringAfter returns the string after the first occurrence of the separator in the string.
func SubstringAfter(s string, separator uint8) string {
	if s == empty {
		return s
	}
	idx := IndexOf(s, separator)
	if idx == -1 {
		return empty
	}
	return SubstringTillEnd(s, idx+1)
}

// SubstringAfterString returns the string after the first occurrence of the separator in the string.
func SubstringAfterString(s, separator string) string {
	if s == empty {
		return s
	}
	idx := IndexOfString(s, separator)
	if idx == -1 {
		return empty
	}
	return SubstringTillEnd(s, idx+len(separator))
}

// SubstringAfterLast returns the string after the last occurrence of the separator in the string.
func SubstringAfterLast(s string, separator uint8) string {
	if s == empty {
		return s
	}
	idx := LastIndexOf(s, separator)
	if idx == -1 {
		return empty
	}
	return SubstringTillEnd(s, idx+1)
}

// SubstringAfterLastString returns the string after the last occurrence of the separator in the string.
func SubstringAfterLastString(s, separator string) string {
	if s == empty {
		return s
	}
	idx := LastIndexOfString(s, separator)
	if idx == -1 {
		return empty
	}
	return SubstringTillEnd(s, idx+len(separator))
}

// SubstringBefore returns the string after the first occurrence of the separator in the string.
func SubstringBefore(s string, separator uint8) string {
	if s == empty {
		return s
	}
	idx := IndexOf(s, separator)
	if idx == -1 {
		return empty
	}
	return Substring(s, 0, idx)
}

// SubstringBeforeString returns the string after the first occurrence of the separator in the string.
func SubstringBeforeString(s, separator string) string {
	if s == empty {
		return s
	}
	idx := IndexOfString(s, separator)
	if idx == -1 {
		return empty
	}
	return Substring(s, 0, idx)
}

// SubstringBeforeLast returns the string after the last occurrence of the separator in the string.
func SubstringBeforeLast(s string, separator uint8) string {
	if s == empty {
		return s
	}
	idx := LastIndexOf(s, separator)
	if idx == -1 {
		return empty
	}
	return Substring(s, 0, idx)
}

// SubstringBeforeLastString returns the string after the last occurrence of the separator in the string.
func SubstringBeforeLastString(s, separator string) string {
	if s == empty {
		return s
	}
	idx := LastIndexOfString(s, separator)
	if idx == -1 {
		return empty
	}
	return Substring(s, 0, idx)
}

// SwapCase is used to change the case of lower case characters to upper case and vice-versa.
func SwapCase(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	b := strings.Builder{}
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			b.WriteRune(c + 32)
		} else if c >= 'a' && c <= 'z' {
			b.WriteRune(c - 32)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// ToCodePoints returns the unicode code point for the characters of the string.
func ToCodePoints(s string) []int32 {
	l := len(s)
	if l == 0 {
		return nil
	}
	res := make([]int32, l)
	for i, c := range s {
		res[i] = c
	}
	return res
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cut set removed.
func Trim(s, cc string) string {
	return strings.Trim(s, cc)
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cut set removed.
//
// To remove a prefix, use TrimPrefix instead.
func TrimLeft(s, cc string) string {
	return strings.TrimLeft(s, cc)
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cut-set removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s, cc string) string {
	return strings.TrimRight(s, cc)
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Truncate is used to return the substring of the string starting from the offset position and having a width <= max width.
func Truncate(s string, offset, maxWidth int) string {
	if offset < 0 {
		panic("offset cannot be negative")
	}
	if maxWidth < 0 {
		panic("max width cannot be negative")
	}
	l := len(s)
	if l == 0 {
		return s
	}
	if offset >= l {
		return empty
	}
	if l > maxWidth {
		end := offset + maxWidth
		if l < end {
			end = l
		}
		return Substring(s, offset, end)
	}
	return SubstringTillEnd(s, offset)
}

// TruncateFromStart is used to return the substring of the string starting from the beginning and having a width <= max width.
func TruncateFromStart(s string, maxWidth int) string {
	return Truncate(s, 0, maxWidth)
}

// Uncapitalize used to convert the first character of the given string to lower case.
// The remaining characters of the string are not changed.
func Uncapitalize(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	f := s[:1]
	u := strings.ToLower(f)
	if f == u {
		return s
	}
	if l == 1 {
		return u
	}
	return u + s[1:]
}

// UpperCase is used to convert a string to upper case.
func UpperCase(s string) string {
	return strings.ToUpper(s)
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

func getMinimumLength(ss ...string) int {
	l := len(ss)
	if len(ss) == 0 {
		return 0
	}
	minL := len(ss[0])
	for i := 1; i < l; i += 1 {
		cl := len(ss[i])
		if cl < minL {
			minL = cl
		}
	}
	return minL
}

func prependIfMissing(s, prefix string, ignoreCase bool, prefixes ...string) string {
	if !IsEmpty(prefix) && !startsWith(s, prefix, ignoreCase) {
		for _, pr := range prefixes {
			if startsWith(s, pr, ignoreCase) {
				return s
			}
		}
		return prefix + s
	}
	return s
}

func replace(s, search, replacement string, max int, ignoreCase bool) string {
	if IsEmpty(s) || IsEmpty(search) || max == 0 {
		return s
	}
	if max == -1 {
		max = len(s)
	}
	if ignoreCase {
		search = strings.ToLower(search)
	}
	searchLength := len(search)
	idx := 0
	if ignoreCase {
		idx = IndexOfStringIgnoreCase(s, search)
	} else {
		idx = IndexOfString(s, search)
	}
	if idx == -1 {
		return s
	}
	return replace(strings.Replace(s, s[idx:idx+searchLength], replacement, 1), search, replacement, max-1, ignoreCase)
}

func startsWith(s, prefix string, ignoreCase bool) bool {
	l := len(s)
	pl := len(prefix)
	if pl > l {
		return false
	}
	if ignoreCase {
		return strings.EqualFold(s[:pl], prefix)
	}
	return s[:pl] == prefix
}

func trimSplits(splits []string, set string) []string {
	result := make([]string, 0, len(splits))
	for _, split := range splits {
		var r string
		if set == space {
			r = strings.TrimSpace(split)
		} else {
			r = strings.Trim(split, set)
		}
		if r == empty {
			continue
		}
		result = append(result, r)
	}
	return result
}
