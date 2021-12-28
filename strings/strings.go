package strings

import (
	"fmt"
	"strings"
)

func Abbreviate(s, abbreviateMarker string, offset, maxWidth int) string {
	if IsNotEmpty(s) && IsEmpty(abbreviateMarker) && maxWidth > 0 {
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

func AppendIfMissing(s, suffix string, suffixes ...string) string {
	return appendIfMissing(s, suffix, false, suffixes...)
}

func AppendIfMissingIgnoreCase(s, suffix string, suffixes ...string) string {
	return appendIfMissing(s, suffix, true, suffixes...)
}

func EndsWith(s, suffix string) bool {
	return endsWith(s, suffix, false)
}

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

func Center(s string, size int, padCharacter uint8) string {
	if size > 0 {
		l := len(s)
		pads := size - l
		if pads > 0 {
			return RightPad(LeftPad(s, l+pads/2, padCharacter), size, padCharacter)
		}
	}
	return s
}

func CenterString(s string, size int, padString string) string {
	if IsEmpty(padString) {
		padString = defaultPadString
	}
	l := len(s)
	pads := size - l
	if pads <= 0 {
		return s
	}
	return RightPadString(LeftPadString(s, l+pads/2, padString), size, padString)
}

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

func EndsWithIgnoreCase(s, suffix string) bool {
	return endsWith(s, suffix, true)
}

func IsAnyEmpty(ss ...string) bool {
	for _, s := range ss {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
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
	return s[offset:] == suffix
}
