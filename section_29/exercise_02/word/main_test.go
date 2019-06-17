package word_test

import (
	"github.com/amundsenjunior/how-to-go/section_29/exercise_02/word"
	"strings"
	"testing"
)

var str1 = "First rule. No conversation lasts longer than 100 total words. I have used 9. You have used 20."
var wordCount = map[string]int{
	"First": 1,
	"rule.": 1,
	"No": 1,
	"conversation": 1,
	"lasts": 1,
	"longer": 1,
	"than": 1,
	"100": 1,
	"total": 1,
	"words.": 1,
	"I": 1,
	"have": 2,
	"used": 2,
	"9.": 1,
	"You": 1,
	"20.": 1,
}

func TestCount(t *testing.T) {
	sLength := len(strings.Split(str1, " "))
	count := word.Count(str1)
	if count != sLength {
		t.Errorf("Count func doesn't return correct value for the following string:\n%s\nExpected %v, got %v.\n", str1, count, sLength)
	}
}

func TestUseCount(t *testing.T) {
	wrdCnt := word.UseCount(str1)
	for k, v := range wrdCnt {
		expect := wordCount[k]
		if v != expect {
			t.Errorf("Number of words does not match for %s.\nExpected %v, got %v.\n", k, expect, v)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		word.Count(str1)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		word.UseCount(str1)
	}
}
