package strings_test

import (
	"context"
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

func TestChop(t *testing.T) {
	assert.Equal(t, "naruto", strings.Chop("naruto!"))
	assert.Equal(t, "naruto", strings.Chop("naruto\r\n"))
	assert.Equal(t, "naruto\n", strings.Chop("naruto\n\r"))
}

func TestCompare(t *testing.T) {
	assert.Equal(t, -1, strings.Compare("abc", "xyz"))
	assert.Equal(t, 1, strings.Compare("xyz", "abc"))
	assert.Equal(t, 0, strings.Compare("abc", "abc"))
	assert.Equal(t, 1, strings.Compare("naruto", "Naruto"))
}

func TestCompareIgnoreCase(t *testing.T) {
	assert.Equal(t, -1, strings.CompareIgnoreCase("abc", "xyz"))
	assert.Equal(t, 1, strings.CompareIgnoreCase("xyz", "abc"))
	assert.Equal(t, 0, strings.CompareIgnoreCase("abc", "abc"))
	assert.Equal(t, 0, strings.CompareIgnoreCase("naruto", "Naruto"))
}

func TestContains(t *testing.T) {
	assert.True(t, strings.Contains("naruto", 'n'))
	assert.False(t, strings.Contains("naruto", 'N'))
	assert.False(t, strings.Contains("naruto", 'x'))
}

func TestContainsIgnoreCase(t *testing.T) {
	assert.True(t, strings.ContainsIgnoreCase("naruto", 'n'))
	assert.True(t, strings.ContainsIgnoreCase("naruto", 'N'))
	assert.False(t, strings.ContainsIgnoreCase("naruto", 'x'))
}

func TestContainsString(t *testing.T) {
	assert.True(t, strings.ContainsString("naruto", "n"))
	assert.False(t, strings.ContainsString("naruto", "N"))
	assert.False(t, strings.ContainsString("naruto", "x"))
}

func TestContainsStringIgnoreCase(t *testing.T) {
	assert.True(t, strings.ContainsStringIgnoreCase("naruto", "n"))
	assert.True(t, strings.ContainsStringIgnoreCase("naruto", "N"))
	assert.False(t, strings.ContainsStringIgnoreCase("naruto", "x"))
}

func TestContainsAny(t *testing.T) {
	assert.True(t, strings.ContainsAny("naruto", 'x', 'y', 'n'))
	assert.False(t, strings.ContainsAny("naruto", 'N', 'A', 'R'))
	assert.False(t, strings.ContainsAny("naruto", 'x', 'y', 'z'))
}

func TestContainsAnyIgnoreCase(t *testing.T) {
	assert.True(t, strings.ContainsAnyIgnoreCase("naruto", 'x', 'y', 'n'))
	assert.True(t, strings.ContainsAnyIgnoreCase("naruto", 'N', 'A', 'R'))
	assert.False(t, strings.ContainsAnyIgnoreCase("naruto", 'x', 'y', 'z'))
}

func TestContainsNone(t *testing.T) {
	assert.False(t, strings.ContainsNone("naruto", 'x', 'y', 'n'))
	assert.True(t, strings.ContainsNone("naruto", 'N', 'A', 'R'))
	assert.True(t, strings.ContainsNone("naruto", 'x', 'y', 'z'))
}

func TestContainsNoneIgnoreCase(t *testing.T) {
	assert.False(t, strings.ContainsNoneIgnoreCase("naruto", 'x', 'y', 'n'))
	assert.False(t, strings.ContainsNoneIgnoreCase("naruto", 'N', 'A', 'R'))
	assert.True(t, strings.ContainsNoneIgnoreCase("naruto", 'x', 'y', 'z'))
}

func TestContainsOnly(t *testing.T) {
	assert.True(t, strings.ContainsOnly("aaabbbccc", 'a', 'b', 'c'))
	assert.False(t, strings.ContainsOnly("aaabbbccc", 'A', 'B', 'C'))
	assert.False(t, strings.ContainsOnly("aaabbbcccd", 'a', 'b', 'c'))
}

func TestContainsOnlyIgnoreCase(t *testing.T) {
	assert.True(t, strings.ContainsOnlyIgnoreCase("aaabbbccc", 'a', 'b', 'c'))
	assert.True(t, strings.ContainsOnlyIgnoreCase("aaabbbccc", 'A', 'B', 'C'))
	assert.False(t, strings.ContainsOnlyIgnoreCase("aaabbbcccd", 'a', 'b', 'c'))
}

func TestContainsWhitespace(t *testing.T) {
	assert.False(t, strings.ContainsWhitespace("naruto"))
	assert.False(t, strings.ContainsWhitespace(""))
	assert.True(t, strings.ContainsWhitespace("naruto rocks"))
}

func TestCount(t *testing.T) {
	assert.Zero(t, strings.Count("naruto", 'x'))
	assert.Zero(t, strings.Count("naruto", 'N'))
	assert.Equal(t, 1, strings.Count("naruto", 'n'))
}

func TestCountIgnoreCase(t *testing.T) {
	assert.Zero(t, strings.CountIgnoreCase("naruto", 'x'))
	assert.Equal(t, 1, strings.CountIgnoreCase("naruto", 'N'))
	assert.Equal(t, 1, strings.CountIgnoreCase("naruto", 'n'))
}

func TestDefaultIfBlank(t *testing.T) {
	assert.Equal(t, "naruto", strings.DefaultIfBlank("naruto", "rocks"))
	assert.Equal(t, "naruto", strings.DefaultIfBlank("", "naruto"))
	assert.Equal(t, "naruto", strings.DefaultIfBlank(" ", "naruto"))
	assert.Equal(t, "naruto", strings.DefaultIfBlank("     ", "naruto"))
	assert.Equal(t, "  a  ", strings.DefaultIfBlank("  a  ", "naruto"))
}

func TestDefaultIfEmpty(t *testing.T) {
	assert.Equal(t, "naruto", strings.DefaultIfEmpty("naruto", "rocks"))
	assert.Equal(t, "naruto", strings.DefaultIfEmpty("", "naruto"))
	assert.Equal(t, " ", strings.DefaultIfEmpty(" ", "naruto"))
	assert.Equal(t, "     ", strings.DefaultIfEmpty("     ", "naruto"))
	assert.Equal(t, "  a  ", strings.DefaultIfEmpty("  a  ", "naruto"))
}

func TestDeleteWhitespaces(t *testing.T) {
	assert.Equal(t, "naruto", strings.DeleteWhitespaces("naruto"))
	assert.Equal(t, "naruto", strings.DeleteWhitespaces("  na r ut   o    "))
	assert.Equal(t, "", strings.DeleteWhitespaces("      "))
	assert.Equal(t, "a", strings.DeleteWhitespaces("    a  "))
}

func TestDifference(t *testing.T) {
	assert.Equal(t, "boruto", strings.Difference("naruto", "boruto"))
	assert.Equal(t, "", strings.Difference("naruto", "naruto"))
	assert.Equal(t, "", strings.Difference("naruto rocks", "naruto"))
	assert.Equal(t, "dd", strings.Difference("aaabbbccc", "aaabbbcccdd"))
	assert.Equal(t, "dcccdd", strings.Difference("aaabbbccc", "aaabbbdcccdd"))
}

func TestEndsWith(t *testing.T) {
	assert.True(t, strings.EndsWith("naruto", "naruto"))
	assert.False(t, strings.EndsWith("naruto", "NARUTO"))
	assert.True(t, strings.EndsWith("", ""))
	assert.False(t, strings.EndsWith("", "naruto"))
	assert.False(t, strings.EndsWith("", "xyz"))
}

func TestEndsWithIgnoreCase(t *testing.T) {
	assert.True(t, strings.EndsWithIgnoreCase("naruto", "naruto"))
	assert.True(t, strings.EndsWithIgnoreCase("naruto", "NARUTO"))
	assert.True(t, strings.EndsWithIgnoreCase("", ""))
	assert.False(t, strings.EndsWithIgnoreCase("", "naruto"))
	assert.False(t, strings.EndsWithIgnoreCase("", "xyz"))
}

func TestEndsWithAny(t *testing.T) {
	assert.True(t, strings.EndsWithAny("naruto", "xyz", "naruto"))
	assert.False(t, strings.EndsWithAny("naruto", "xyz", "NARUTO"))
	assert.True(t, strings.EndsWithAny("", "xyz", ""))
	assert.False(t, strings.EndsWithAny("", "xyz", "naruto"))
	assert.False(t, strings.EndsWithAny("", "xyz", "xyz"))
}

func TestEndsWithAnyIgnoreCase(t *testing.T) {
	assert.True(t, strings.EndsWithAnyIgnoreCase("naruto", "xyz", "naruto"))
	assert.True(t, strings.EndsWithAnyIgnoreCase("naruto", "xyz", "NARUTO"))
	assert.True(t, strings.EndsWithAnyIgnoreCase("", "xyz", ""))
	assert.False(t, strings.EndsWithAnyIgnoreCase("", "xyz", "naruto"))
	assert.False(t, strings.EndsWithAnyIgnoreCase("", "xyz", "xyz"))
}

func TestEquals(t *testing.T) {
	assert.True(t, strings.Equals("naruto", "naruto"))
	assert.False(t, strings.Equals("naruto", "NARUTO"))
	assert.True(t, strings.Equals("", ""))
	assert.False(t, strings.Equals("naruto", ""))
	assert.False(t, strings.Equals("", "naruto"))
}

func TestEqualsIgnoreCase(t *testing.T) {
	assert.True(t, strings.EqualsIgnoreCase("naruto", "naruto"))
	assert.True(t, strings.EqualsIgnoreCase("naruto", "NARUTO"))
	assert.True(t, strings.EqualsIgnoreCase("", ""))
	assert.False(t, strings.EqualsIgnoreCase("naruto", ""))
	assert.False(t, strings.EqualsIgnoreCase("", "naruto"))
}

func TestEqualsAny(t *testing.T) {
	assert.True(t, strings.EqualsAny("naruto", "xyz", "naruto"))
	assert.False(t, strings.EqualsAny("naruto", "xyz", "NARUTO"))
	assert.True(t, strings.EqualsAny("", "xyz", ""))
	assert.False(t, strings.EqualsAny("naruto", "xyz", ""))
	assert.False(t, strings.EqualsAny("", "xyz", "naruto"))
}

func TestEqualsAnyIgnoreCase(t *testing.T) {
	assert.True(t, strings.EqualsAnyIgnoreCase("naruto", "xyz", "naruto"))
	assert.True(t, strings.EqualsAnyIgnoreCase("naruto", "xyz", "NARUTO"))
	assert.True(t, strings.EqualsAnyIgnoreCase("", "xyz", ""))
	assert.False(t, strings.EqualsAnyIgnoreCase("naruto", "xyz", ""))
	assert.False(t, strings.EqualsAnyIgnoreCase("", "xyz", "naruto"))
}

func TestFirstNonBlank(t *testing.T) {
	assert.Equal(t, "naruto", strings.FirstNonBlank("", "  ", " ", "    ", "naruto", " ", "", ""))
	assert.Equal(t, "", strings.FirstNonBlank("", "  ", " ", "    ", " ", "", ""))
}

func TestFirstNonEmpty(t *testing.T) {
	assert.Equal(t, "  ", strings.FirstNonEmpty("", "  ", " ", "    ", "naruto", " ", "", ""))
	assert.Equal(t, "  ", strings.FirstNonEmpty("", "  ", " ", "    ", " ", "", ""))
	assert.Equal(t, "", strings.FirstNonEmpty("", "", "", ""))
}

func TestGetCommonPrefix(t *testing.T) {
	assert.Equal(t, "naruto", strings.GetCommonPrefix("naruto", "naruto", "naruto", "naruto rocks"))
	assert.Equal(t, "", strings.GetCommonPrefix("naruto", "naruto", "naruto", "naruto rocks", "boruto"))
}

func TestGetDigits(t *testing.T) {
	assert.Equal(t, "12311111", strings.GetDigits("abc$1231$ab111bc1b"))
	assert.Equal(t, "", strings.GetDigits("naruto"))
}

func TestGetIfBlank(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "naruto", strings.GetIfBlank(ctx, "naruto", nil))
	assert.Equal(t, "", strings.GetIfBlank(ctx, "   ", nil))
	assert.Equal(t, "naruto", strings.GetIfBlank(ctx, "   ", func(_ context.Context) string {
		return "naruto"
	}))
}

func TestGetIfEmpty(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "naruto", strings.GetIfEmpty(ctx, "naruto", nil))
	assert.Equal(t, "", strings.GetIfEmpty(ctx, "", nil))
	assert.Equal(t, "   ", strings.GetIfEmpty(ctx, "   ", func(_ context.Context) string {
		return "naruto"
	}))
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOf("naruto", 'n'))
	assert.Equal(t, -1, strings.IndexOf("naruto", 'N'))
	assert.Equal(t, -1, strings.IndexOf("naruto", 'x'))
}

func TestIndexOfIgnoreCase(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfIgnoreCase("naruto", 'n'))
	assert.Equal(t, 0, strings.IndexOfIgnoreCase("naruto", 'N'))
	assert.Equal(t, -1, strings.IndexOfIgnoreCase("naruto", 'x'))
}

func TestIndexOfString(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfString("naruto rocks", "naruto"))
	assert.Equal(t, -1, strings.IndexOfString("naruto rocks", "boruto"))
	assert.Equal(t, 0, strings.IndexOfString("naruto rocks", ""))
	assert.Equal(t, -1, strings.IndexOfString("naruto rocks", "NARUTO"))
}

func TestIndexOfStringIgnoreCase(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfStringIgnoreCase("naruto rocks", "naruto"))
	assert.Equal(t, -1, strings.IndexOfStringIgnoreCase("naruto rocks", "boruto"))
	assert.Equal(t, 0, strings.IndexOfStringIgnoreCase("naruto rocks", ""))
	assert.Equal(t, 0, strings.IndexOfStringIgnoreCase("naruto rocks", "NARUTO"))
}

func TestIndexOfAny(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfAny("naruto", 'x', 'y', 'a', 'n'))
	assert.Equal(t, -1, strings.IndexOfAny("", 'x', 'y', 'a', 'n'))
	assert.Equal(t, 1, strings.IndexOfAny("naruto", 'x', 'y', 'a', 'N'))
}

func TestIndexOfAnyIgnoreCase(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfAnyIgnoreCase("naruto", 'x', 'y', 'a', 'n'))
	assert.Equal(t, -1, strings.IndexOfAnyIgnoreCase("", 'x', 'y', 'a', 'n'))
	assert.Equal(t, 0, strings.IndexOfAnyIgnoreCase("naruto", 'x', 'y', 'a', 'N'))
}

func TestIndexOfAnyBut(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfAnyBut("naruto", 's', 'h'))
	assert.Equal(t, -1, strings.IndexOfAnyBut("naruto", 'n', 'a', 'r', 'u', 't', 'o'))
	assert.Equal(t, 5, strings.IndexOfAnyBut("naruto", 'n', 'a', 'r', 'u', 't', 'O'))
}

func TestIndexOfAnyButIgnoreCase(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfAnyButIgnoreCase("naruto", 's', 'h'))
	assert.Equal(t, -1, strings.IndexOfAnyButIgnoreCase("naruto", 'n', 'a', 'r', 'u', 't', 'O'))
	assert.Equal(t, 5, strings.IndexOfAnyButIgnoreCase("naruto", 'n', 'a', 'r', 'u', 'T'))
}

func TestIndexOfDifference(t *testing.T) {
	assert.Equal(t, 0, strings.IndexOfDifference("naruto", "boruto", "sakura", "hinata"))
	assert.Equal(t, -1, strings.IndexOfDifference("nnn", "nnnn", "nnnnn", "nnnnnn"))
	assert.Equal(t, 2, strings.IndexOfDifference("nnn", "nnnn", "nnnnn", "nnnnnn", "nna"))
	assert.Equal(t, -1, strings.IndexOfDifference("nnn", "nnnn", "nnnnn", "nnnnnn", "nnnna"))
}

func TestIndexOfStringStartingAt(t *testing.T) {
	assert.Equal(t, 2, strings.IndexOfStringStartingAt("naruto", "r", 2))
	assert.Equal(t, -1, strings.IndexOfStringStartingAt("naruto", "r", 3))
}

func TestLastIndexOf(t *testing.T) {
	assert.Equal(t, 4, strings.LastIndexOf("aabaa", 'a'))
	assert.Equal(t, 3, strings.LastIndexOf("aabbaa", 'b'))
}

func TestLastIndexOfIgnoreCase(t *testing.T) {
	assert.Equal(t, 4, strings.LastIndexOfIgnoreCase("aabaa", 'A'))
	assert.Equal(t, 3, strings.LastIndexOfIgnoreCase("aabbaa", 'B'))
}

func TestLastIndexOfString(t *testing.T) {
	assert.Equal(t, 4, strings.LastIndexOfString("aabaa", "a"))
	assert.Equal(t, 3, strings.LastIndexOfString("aabbaa", "b"))
}

func TestLastIndexOfStringIgnoreCase(t *testing.T) {
	assert.Equal(t, 4, strings.LastIndexOfStringIgnoreCase("aabaa", "A"))
	assert.Equal(t, 3, strings.LastIndexOfStringIgnoreCase("aabbaa", "B"))
}

func TestLastIndexOfAny(t *testing.T) {
	assert.Equal(t, 5, strings.LastIndexOfAny("naruto", 'n', 'a', 'r', 'u', 't', 'o'))
	assert.Equal(t, 4, strings.LastIndexOfAny("naruto", 'n', 'a', 'r', 'u', 't'))
	assert.Equal(t, 0, strings.LastIndexOfAny("naruto", 'n'))
}

func TestLastIndexOfAnyIgnoreCase(t *testing.T) {
	assert.Equal(t, 5, strings.LastIndexOfAnyIgnoreCase("naruto", 'n', 'a', 'r', 'u', 't', 'O'))
	assert.Equal(t, 4, strings.LastIndexOfAnyIgnoreCase("naruto", 'n', 'a', 'r', 'u', 'T'))
	assert.Equal(t, 0, strings.LastIndexOfAnyIgnoreCase("naruto", 'N'))
}

func TestLastIndexOfAnyBut(t *testing.T) {
	assert.Equal(t, 5, strings.LastIndexOfAnyBut("naruto", 's', 'h'))
	assert.Equal(t, -1, strings.LastIndexOfAnyBut("naruto", 'n', 'a', 'r', 'u', 't', 'o'))
	assert.Equal(t, 0, strings.LastIndexOfAnyBut("naruto", 'N', 'a', 'r', 'u', 't', 'o'))
}

func TestLastIndexOfAnyButIgnoreCase(t *testing.T) {
	assert.Equal(t, 5, strings.LastIndexOfAnyButIgnoreCase("naruto", 's', 'h'))
	assert.Equal(t, -1, strings.LastIndexOfAnyButIgnoreCase("naruto", 'n', 'a', 'r', 'u', 't', 'O'))
	assert.Equal(t, 0, strings.LastIndexOfAnyButIgnoreCase("naruto", 'a', 'r', 'u', 'T', 'o'))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, strings.IsBlank("     "))
	assert.True(t, strings.IsBlank(""))
	assert.False(t, strings.IsBlank("   n   "))
}

func TestIsAllBlank(t *testing.T) {
	assert.True(t, strings.IsAllBlank("     ", "   ", ""))
	assert.True(t, strings.IsAllBlank(""))
	assert.False(t, strings.IsAllBlank("", "", "    ", "   n   "))
}

func TestIsNoneBlank(t *testing.T) {
	assert.False(t, strings.IsNoneBlank("     ", "   ", ""))
	assert.False(t, strings.IsNoneBlank(""))
	assert.False(t, strings.IsNoneBlank("", "", "    ", "   n   "))
	assert.True(t, strings.IsNoneBlank("naruto", "boruto", "hinata", "   n   "))
}

func TestIsAnyBlank(t *testing.T) {
	assert.True(t, strings.IsAnyBlank("     ", "   ", ""))
	assert.True(t, strings.IsAnyBlank(""))
	assert.True(t, strings.IsAnyBlank("", "", "    ", "   n   "))
	assert.False(t, strings.IsAnyBlank("naruto", "boruto", "hinata", "   n   "))
}

func TestIsEmpty(t *testing.T) {
	assert.False(t, strings.IsEmpty("     "))
	assert.True(t, strings.IsEmpty(""))
	assert.False(t, strings.IsEmpty("   n   "))
}

func TestIsAllEmpty(t *testing.T) {
	assert.False(t, strings.IsAllEmpty("     ", "   ", ""))
	assert.True(t, strings.IsAllEmpty(""))
	assert.False(t, strings.IsAllEmpty("", "", "    ", "   n   "))
}

func TestIsNoneEmpty(t *testing.T) {
	assert.False(t, strings.IsNoneEmpty("     ", "   ", ""))
	assert.False(t, strings.IsNoneEmpty(""))
	assert.False(t, strings.IsNoneEmpty("", "", "    ", "   n   "))
	assert.True(t, strings.IsNoneEmpty("naruto", "boruto", "hinata", "   n   "))
}

func TestIsAnyEmpty(t *testing.T) {
	assert.True(t, strings.IsAnyEmpty("     ", "   ", ""))
	assert.True(t, strings.IsAnyEmpty(""))
	assert.True(t, strings.IsAnyEmpty("", "", "    ", "   n   "))
	assert.False(t, strings.IsAnyEmpty("naruto", "boruto", "hinata", "   n   "))
}

func TestIsLowerCase(t *testing.T) {
	assert.True(t, strings.IsLowerCase("naruto"))
	assert.True(t, strings.IsLowerCase(""))
	assert.False(t, strings.IsLowerCase("Naruto"))
}

func TestIsAllLowerCase(t *testing.T) {
	assert.True(t, strings.IsAllLowerCase("naruto", "boruto"))
	assert.True(t, strings.IsAllLowerCase("", "n"))
	assert.False(t, strings.IsAllLowerCase("Naruto"))
}

func TestIsAnyLowerCase(t *testing.T) {
	assert.True(t, strings.IsAnyLowerCase("naruto"))
	assert.True(t, strings.IsAnyLowerCase(""))
	assert.False(t, strings.IsAnyLowerCase("Naruto"))
}

func TestIsUpperCase(t *testing.T) {
	assert.True(t, strings.IsUpperCase("NARUTO"))
	assert.True(t, strings.IsUpperCase(""))
	assert.False(t, strings.IsUpperCase("Naruto"))
}

func TestIsAllUpperCase(t *testing.T) {
	assert.True(t, strings.IsAllUpperCase("NARUTO", "BORUTO"))
	assert.True(t, strings.IsAllUpperCase("", "N"))
	assert.False(t, strings.IsAllUpperCase("Naruto"))
}

func TestIsAnyUpperCase(t *testing.T) {
	assert.True(t, strings.IsAnyUpperCase("NARUTO"))
	assert.True(t, strings.IsAnyUpperCase(""))
	assert.False(t, strings.IsAnyUpperCase("Naruto"))
}

func TestIsMixedCase(t *testing.T) {
	assert.False(t, strings.IsMixedCase("NARUTO"))
	assert.False(t, strings.IsMixedCase(""))
	assert.True(t, strings.IsMixedCase("Naruto"))
}

func TestIsAllMixedCase(t *testing.T) {
	assert.False(t, strings.IsAllMixedCase("NARUTO", "BORUTO"))
	assert.False(t, strings.IsAllMixedCase("", "N"))
	assert.True(t, strings.IsAllMixedCase("Naruto"))
}

func TestIsAnyMixedCase(t *testing.T) {
	assert.False(t, strings.IsAnyMixedCase("NARUTO"))
	assert.False(t, strings.IsAnyMixedCase(""))
	assert.True(t, strings.IsAnyMixedCase("Naruto"))
}

func TestIsAlpha(t *testing.T) {
	assert.True(t, strings.IsAlpha("naruto"))
	assert.False(t, strings.IsAlpha("naruto   "))
	assert.False(t, strings.IsAlpha("naruto123"))
	assert.False(t, strings.IsAlpha(""))
}

func TestIsNumeric(t *testing.T) {
	assert.False(t, strings.IsNumeric("naruto123"))
	assert.True(t, strings.IsNumeric("123"))
	assert.False(t, strings.IsNumeric("123    "))
	assert.False(t, strings.IsNumeric(""))
}

func IsAlphaNumeric(t *testing.T) {
	assert.True(t, strings.IsAlphaNumeric("naruto123"))
	assert.False(t, strings.IsAlphaNumeric("naruto123    "))
	assert.True(t, strings.IsAlphaNumeric("123"))
	assert.True(t, strings.IsAlphaNumeric("naruto"))
	assert.False(t, strings.IsAlphaNumeric(""))
}

func TestIsAlphaSpace(t *testing.T) {
	assert.True(t, strings.IsAlphaSpace("naruto"))
	assert.True(t, strings.IsAlphaSpace("naruto   "))
	assert.False(t, strings.IsAlphaSpace("naruto123"))
	assert.False(t, strings.IsAlphaSpace(""))
}

func TestIsNumericSpace(t *testing.T) {
	assert.False(t, strings.IsNumericSpace("naruto123"))
	assert.True(t, strings.IsNumericSpace("123"))
	assert.True(t, strings.IsNumericSpace("123    "))
	assert.False(t, strings.IsNumericSpace(""))
}

func IsAlphaNumericSpace(t *testing.T) {
	assert.True(t, strings.IsAlphaNumericSpace("naruto123"))
	assert.True(t, strings.IsAlphaNumeric("naruto123    "))
	assert.True(t, strings.IsAlphaNumericSpace("123"))
	assert.True(t, strings.IsAlphaNumericSpace("naruto"))
	assert.False(t, strings.IsAlphaNumericSpace(""))
}

func TestIsAsciiPrintable(t *testing.T) {
	assert.True(t, strings.IsAsciiPrintable("naruto"))
	assert.True(t, strings.IsAsciiPrintable("NARUTO"))
	assert.True(t, strings.IsAsciiPrintable("NARUTO1223456576"))
	assert.True(t, strings.IsAsciiPrintable("NARUTO    "))
	assert.True(t, strings.IsAsciiPrintable("@NARUTO@"))
	assert.True(t, strings.IsAsciiPrintable("`@NARUTO@`"))
	assert.True(t, strings.IsAsciiPrintable("`@NARUTO@`"))
	assert.False(t, strings.IsAsciiPrintable(string(rune(0))))
}

func TestJoinByChar(t *testing.T) {
	assert.Equal(t, "n a r u t o", strings.JoinByChar([]uint8{'n', 'a', 'r', 'u', 't', 'o'}, ' '))
}

func TestJoinStringsByChar(t *testing.T) {
	assert.Equal(t, "naruto rocks", strings.JoinStringsByChar([]string{"naruto", "rocks"}, ' '))
}

func TestJoinStringsByString(t *testing.T) {
	assert.Equal(t, "naruto rocks", strings.JoinStringsByString([]string{"naruto", "rocks"}, " "))
}

func TestLeft(t *testing.T) {
	assert.Equal(t, "", strings.Left("naruto", -1))
	assert.Equal(t, "naruto", strings.Left("naruto", 10))
	assert.Equal(t, "nar", strings.Left("naruto", 3))
}

func TestLeftPad(t *testing.T) {
	assert.Equal(t, "naruto", strings.LeftPad("naruto", 1, ' '))
	assert.Equal(t, "naruto", strings.LeftPad("naruto", 1, ' '))
	assert.Equal(t, "XXXXnaruto", strings.LeftPad("naruto", 10, 'X'))
}

func TestLeftPadString(t *testing.T) {
	assert.Equal(t, "naruto", strings.LeftPadString("naruto", 1, " "))
	assert.Equal(t, "naruto", strings.LeftPadString("naruto", 1, " "))
	assert.Equal(t, "XXXXnaruto", strings.LeftPadString("naruto", 10, "X"))
}

func TestLowerCase(t *testing.T) {
	assert.Equal(t, "naruto", strings.LowerCase("Naruto"))
	assert.Equal(t, "naruto", strings.LowerCase("NARUTO"))
	assert.Equal(t, "naruto", strings.LowerCase("naruto"))
	assert.Equal(t, "", strings.LowerCase(""))
}

func TestMid(t *testing.T) {
	assert.Equal(t, "ru", strings.Mid("naruto", 2, 2))
	assert.Equal(t, "nar", strings.Mid("naruto", 0, 3))
	assert.Equal(t, "naruto", strings.Mid("naruto", 0, 10))
	assert.Equal(t, "", strings.Mid("naruto", 10, 2))
	assert.Equal(t, "", strings.Mid("naruto", 2, -1))
}

func TestNormalizeSpace(t *testing.T) {
	assert.Equal(t, "naruto rocks", strings.NormalizeSpace("    naruto    rocks   "))
	assert.Equal(t, "naruto rocks", strings.NormalizeSpace("    naruto rocks "))
	assert.Equal(t, "naruto rocks", strings.NormalizeSpace("    naruto rocks    "))
	assert.Equal(t, "naruto rocks", strings.NormalizeSpace("naruto rocks    "))
	assert.Equal(t, "naruto rocks", strings.NormalizeSpace("naruto rocks"))
}

func TestOrdinalIndexOf(t *testing.T) {
	assert.Equal(t, 2, strings.OrdinalIndexOf("nnnnnnnn", "n", 3))
	assert.Equal(t, -1, strings.OrdinalIndexOf("nnnnnnnn", "a", 3))
}

func TestOverlay(t *testing.T) {
	assert.Equal(t, "naruto", strings.Overlay("naruto", "aru", 1, 4))
}

func TestPrependIfMissing(t *testing.T) {
	assert.Equal(t, "naruto rocks", strings.PrependIfMissing(" rocks", "naruto", "xyz"))
	assert.Equal(t, "naruto", strings.PrependIfMissing("naruto", "", "xyz"))
	assert.Equal(t, "naruto", strings.PrependIfMissing("naruto", "rocks", "nar"))
}

func TestPrependIfMissingIgnoreCase(t *testing.T) {
	assert.Equal(t, "naruto rocks", strings.PrependIfMissingIgnoreCase(" rocks", "naruto", "XYZ"))
	assert.Equal(t, "naruto", strings.PrependIfMissingIgnoreCase("naruto", "", "xyz"))
	assert.Equal(t, "naruto", strings.PrependIfMissingIgnoreCase("naruto", "rocks", "NAR", "xyz"))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, "naruto", strings.Remove("ncacrcuctco", 'c'))
	assert.Equal(t, "naruto", strings.Remove("naruto", 'x'))
}

func TestRemoveIgnoreCase(t *testing.T) {
	assert.Equal(t, "naruto", strings.RemoveIgnoreCase("ncacrcuctco", 'C'))
	assert.Equal(t, "naruto", strings.RemoveIgnoreCase("naruto", 'X'))
}

func TestRemoveString(t *testing.T) {
	assert.Equal(t, "naruto", strings.RemoveString("ncacrcuctco", "c"))
	assert.Equal(t, "naruto", strings.RemoveString("naruto", "x"))
}

func TestRemoveStringIgnoreCase(t *testing.T) {
	assert.Equal(t, "naruto", strings.RemoveStringIgnoreCase("ncacrcuctco", "C"))
	assert.Equal(t, "naruto", strings.RemoveStringIgnoreCase("naruto", "X"))
}

func TestRemoveStart(t *testing.T) {
	assert.Equal(t, "uto", strings.RemoveStart("naruto", "nar"))
	assert.Equal(t, "naruto", strings.RemoveStart("naruto", "xyz"))
}

func TestRemoveStartIgnoreCase(t *testing.T) {
	assert.Equal(t, "uto", strings.RemoveStartIgnoreCase("naruto", "NAR"))
	assert.Equal(t, "naruto", strings.RemoveStartIgnoreCase("naruto", "XYZ"))
}

func TestRemoveEnd(t *testing.T) {
	assert.Equal(t, "nar", strings.RemoveEnd("naruto", "uto"))
	assert.Equal(t, "naruto", strings.RemoveEnd("naruto", "xyz"))
}

func TestRemoveEndIgnoreCase(t *testing.T) {
	assert.Equal(t, "nar", strings.RemoveEndIgnoreCase("naruto", "UTO"))
	assert.Equal(t, "naruto", strings.RemoveEndIgnoreCase("naruto", "XYZ"))
}

func TestRight(t *testing.T) {
	assert.Equal(t, "", strings.Right("naruto", -1))
	assert.Equal(t, "naruto", strings.Right("naruto", 10))
	assert.Equal(t, "uto", strings.Right("naruto", 3))
}

func TestRightPad(t *testing.T) {
	assert.Equal(t, "naruto", strings.RightPad("naruto", 1, ' '))
	assert.Equal(t, "naruto", strings.RightPad("naruto", 1, ' '))
	assert.Equal(t, "narutoXXXX", strings.RightPad("naruto", 10, 'X'))
}

func TestRotate(t *testing.T) {
	assert.Equal(t, "utonar", strings.Rotate("naruto", 3))
	assert.Equal(t, "utonar", strings.Rotate("naruto", 9))
	assert.Equal(t, "naruto", strings.Rotate("naruto", 0))
	assert.Equal(t, "naruto", strings.Rotate("naruto", 6))
}

func TestSplit(t *testing.T) {
	assert.Len(t, strings.Split("naruto rocks", ' '), 2)
	assert.Len(t, strings.Split("", ' '), 0)
	assert.Nil(t, strings.Split("", ' '))
	assert.Len(t, strings.Split("naruto rocks", 'a'), 2)
}

func TestSplitByString(t *testing.T) {
	assert.Len(t, strings.SplitByString("naruto rocks", " "), 2)
	assert.Len(t, strings.SplitByString("", " "), 0)
	assert.Nil(t, strings.SplitByString("", " "))
	assert.Len(t, strings.SplitByString("naruto rocks", "a"), 2)
}

func TestSplitN(t *testing.T) {
	assert.Len(t, strings.SplitN("naruto rocks", ' ', -1), 2)
	assert.Len(t, strings.SplitN("", ' ', -1), 0)
	assert.Nil(t, strings.SplitN("", ' ', -1))
	assert.Len(t, strings.SplitN("naruto rocks", 'a', -1), 2)
	assert.Len(t, strings.SplitN("naruto rocks", ' ', 1), 1)
	assert.Len(t, strings.SplitN("", ' ', 1), 0)
	assert.Nil(t, strings.SplitN("", ' ', 1))
	assert.Len(t, strings.SplitN("naruto rocks", 'a', 1), 1)
	assert.Len(t, strings.SplitN("naruto rocks", ' ', 0), 0)
	assert.Len(t, strings.SplitN("", ' ', 0), 0)
	assert.Nil(t, strings.SplitN("", ' ', 0))
	assert.Len(t, strings.SplitN("naruto rocks", 'a', 0), 0)
}

func TestSplitNByString(t *testing.T) {
	assert.Len(t, strings.SplitNByString("naruto rocks", " ", -1), 2)
	assert.Len(t, strings.SplitNByString("", " ", -1), 0)
	assert.Nil(t, strings.SplitNByString("", " ", -1))
	assert.Len(t, strings.SplitNByString("naruto rocks", "a", -1), 2)
	assert.Len(t, strings.SplitNByString("naruto rocks", " ", 1), 1)
	assert.Len(t, strings.SplitNByString("", " ", 1), 0)
	assert.Nil(t, strings.SplitNByString("", " ", 1))
	assert.Len(t, strings.SplitNByString("naruto rocks", "a", 1), 1)
	assert.Len(t, strings.SplitNByString("naruto rocks", " ", 0), 0)
	assert.Len(t, strings.SplitNByString("", " ", 0), 0)
	assert.Nil(t, strings.SplitNByString("", " ", 0))
	assert.Len(t, strings.SplitNByString("naruto rocks", "a", 0), 0)
}

func TestSplitWithTrim(t *testing.T) {
	assert.Len(t, strings.SplitWithTrim("naruto rocks", ' '), 2)
	assert.Len(t, strings.SplitWithTrim("naruto      rocks", ' '), 2)
	assert.Len(t, strings.SplitWithTrim("", ' '), 0)
	assert.Nil(t, strings.SplitWithTrim("", ' '))
	assert.Len(t, strings.SplitWithTrim("naruto rocks", 'a'), 2)
}

func TestSplitByStringWithTrim(t *testing.T) {
	assert.Len(t, strings.SplitByStringWithTrim("naruto rocks", " "), 2)
	assert.Len(t, strings.SplitByStringWithTrim("", " "), 0)
	assert.Nil(t, strings.SplitByStringWithTrim("", " "))
	assert.Len(t, strings.SplitByStringWithTrim("naruto rocks", "a"), 2)
}

func TestSplitNWithTrim(t *testing.T) {
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", ' ', -1), 2)
	assert.Len(t, strings.SplitNWithTrim("", ' ', -1), 0)
	assert.Nil(t, strings.SplitNWithTrim("", ' ', -1))
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", 'a', -1), 2)
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", ' ', 1), 1)
	assert.Len(t, strings.SplitNWithTrim("", ' ', 1), 0)
	assert.Nil(t, strings.SplitNWithTrim("", ' ', 1))
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", 'a', 1), 1)
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", ' ', 0), 0)
	assert.Len(t, strings.SplitNWithTrim("", ' ', 0), 0)
	assert.Nil(t, strings.SplitNWithTrim("", ' ', 0))
	assert.Len(t, strings.SplitNWithTrim("naruto rocks", 'a', 0), 0)
}

func TestSplitNByStringWithTrim(t *testing.T) {
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", " ", -1), 2)
	assert.Len(t, strings.SplitNByStringWithTrim("", " ", -1), 0)
	assert.Nil(t, strings.SplitNByStringWithTrim("", " ", -1))
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", "a", -1), 2)
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", " ", 1), 1)
	assert.Len(t, strings.SplitNByStringWithTrim("", " ", 1), 0)
	assert.Nil(t, strings.SplitNByStringWithTrim("", " ", 1))
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", "a", 1), 1)
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", " ", 0), 0)
	assert.Len(t, strings.SplitNByStringWithTrim("", " ", 0), 0)
	assert.Nil(t, strings.SplitNByStringWithTrim("", " ", 0))
	assert.Len(t, strings.SplitNByStringWithTrim("naruto rocks", "a", 0), 0)
}

func TestSplitWithTrimCutSet(t *testing.T) {
	assert.Len(t, strings.SplitWithTrimCutSet("naruto rocks", ' ', " "), 2)
	assert.Len(t, strings.SplitWithTrimCutSet("naruto      rocks", ' ', " "), 2)
	assert.Len(t, strings.SplitWithTrimCutSet("", ' ', " "), 0)
	assert.Nil(t, strings.SplitWithTrimCutSet("", ' ', " "))
	assert.Len(t, strings.SplitWithTrimCutSet("naruto rocks", 'a', " "), 2)
}

func TestSplitByStringWithTrimCutSet(t *testing.T) {
	assert.Len(t, strings.SplitByStringWithTrimCutSet("naruto rocks", " ", " "), 2)
	assert.Len(t, strings.SplitByStringWithTrimCutSet("", " ", " "), 0)
	assert.Nil(t, strings.SplitByStringWithTrimCutSet("", " ", " "))
	assert.Len(t, strings.SplitByStringWithTrimCutSet("naruto rocks", "a", "x"), 2)
	assert.Len(t, strings.SplitByStringWithTrimCutSet("naruto rocks", "a", "x"), 2)
	assert.Len(t, strings.SplitByStringWithTrimCutSet("naruto rocks", "a", "n"), 1)
}

func TestSplitNWithTrimCutSet(t *testing.T) {
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", ' ', -1, " "), 2)
	assert.Len(t, strings.SplitNWithTrimCutSet("", ' ', -1, " "), 0)
	assert.Nil(t, strings.SplitNWithTrimCutSet("", ' ', -1, " "))
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", 'a', -1, " "), 2)
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", ' ', 1, " "), 1)
	assert.Len(t, strings.SplitNWithTrimCutSet("", ' ', 1, " "), 0)
	assert.Nil(t, strings.SplitNWithTrimCutSet("", ' ', 1, " "))
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", 'a', 1, " "), 1)
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", ' ', 0, " "), 0)
	assert.Len(t, strings.SplitNWithTrimCutSet("", ' ', 0, " "), 0)
	assert.Nil(t, strings.SplitNWithTrimCutSet("", ' ', 0, " "))
	assert.Len(t, strings.SplitNWithTrimCutSet("naruto rocks", 'a', 0, " "), 0)
}

func TestSplitNByStringWithTrimCutSet(t *testing.T) {
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", " ", -1, " "), 2)
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("", " ", -1, " "), 0)
	assert.Nil(t, strings.SplitNByStringWithTrimCutSet("", " ", -1, " "))
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", "a", -1, " "), 2)
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", " ", 1, " "), 1)
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("", " ", 1, " "), 0)
	assert.Nil(t, strings.SplitNByStringWithTrimCutSet("", " ", 1, " "))
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", "a", 1, " "), 1)
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", " ", 0, " "), 0)
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("", " ", 0, " "), 0)
	assert.Nil(t, strings.SplitNByStringWithTrimCutSet("", " ", 0, " "))
	assert.Len(t, strings.SplitNByStringWithTrimCutSet("naruto rocks", "a", 0, " "), 0)
}

func TestStartsWith(t *testing.T) {
	assert.True(t, strings.StartsWith("naruto", "n"))
	assert.False(t, strings.StartsWith("naruto", "nap"))
	assert.True(t, strings.StartsWith("naruto", ""))
}

func TestStartsWithAny(t *testing.T) {
	assert.True(t, strings.StartsWithAny("naruto", "n"))
	assert.False(t, strings.StartsWithAny("naruto", "nap"))
	assert.True(t, strings.StartsWithAny("naruto", ""))
}

func TestStartsWithIgnoreCase(t *testing.T) {
	assert.True(t, strings.StartsWithIgnoreCase("naruto", "N"))
	assert.False(t, strings.StartsWithIgnoreCase("naruto", "nap"))
	assert.True(t, strings.StartsWithIgnoreCase("naruto", ""))
}

func TestStartsWithAnyIgnoreCase(t *testing.T) {
	assert.True(t, strings.StartsWithAnyIgnoreCase("naruto", "N"))
	assert.False(t, strings.StartsWithAnyIgnoreCase("naruto", "nap"))
	assert.True(t, strings.StartsWithAnyIgnoreCase("naruto", ""))
}
