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

func TestIndexOfAny(t *testing.T) {

}

func TestIndexOfAnyIgnoreCase(t *testing.T) {

}
