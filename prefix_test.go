package prefix

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	begin    string
	end      string
	expected []string
}{
	{
		begin: "4", end: "7",
		expected: []string{"4", "5", "6", "7"},
	},
	{
		begin: "12345", end: "12345",
		expected: []string{"12345"},
	},

	{
		begin: "78000", end: "78999",
		expected: []string{"78"},
	},

	{
		begin: "2010", end: "2040",
		expected: []string{"201", "202", "203", "2040"},
	},

	{
		begin: "20000", end: "20345",
		expected: []string{"200", "201", "202", "2030", "2031", "2032", "2033", "20340", "20341", "20342", "20343", "20344", "20345"},
	},

	{
		begin: "20010", end: "20345",
		expected: []string{"2001", "2002", "2003", "2004", "2005", "2006", "2007", "2008", "2009",
			"201", "202", "2030", "2031", "2032", "2033", "20340", "20341", "20342", "20343", "20344", "20345"},
	},

	{
		begin: "010", end: "345",
		expected: []string{"01", "02", "03", "04", "05", "06", "07", "08", "09",
			"1", "2", "30", "31", "32", "33", "340", "341", "342", "343", "344", "345"},
	},

	{
		begin: "344440", end: "899976",
		expected: []string{
			"34444", "34445", "34446", "34447", "34448", "34449",
			"3445", "3446", "3447", "3448", "3449",
			"345", "346", "347", "348", "349",
			"35", "36", "37", "38", "39",
			"4", "5", "6", "7",
			"80", "81", "82", "83", "84", "85", "86", "87", "88",
			"890", "891", "892", "893", "894", "895", "896", "897", "898",
			"8990", "8991", "8992", "8993", "8994", "8995", "8996", "8997", "8998",
			"89990", "89991", "89992", "89993", "89994", "89995", "89996",
			"899970", "899971", "899972", "899973", "899974", "899975", "899976"},
	},

	{
		begin: "30892190000000", end: "30892239999999",
		expected: []string{"3089219", "3089220", "3089221", "3089222", "3089223"},
	},

	{
		begin: "3203", end: "4099",
		expected: []string{"3203", "3204", "3205", "3206", "3207", "3208", "3209",
			"321", "322", "323", "324", "325", "326", "327", "328", "329",
			"33", "34", "35", "36", "37", "38", "39",
			"40"},
	},
}

func TestPrefix(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.begin, tc.end), func(t *testing.T) {
			require.True(t, simpeChek(tc.begin, tc.end, tc.expected), "Need check test conditions")

			actual, err := FromRange(tc.begin, tc.end)

			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkPrefix(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for i := range testCases {
			FromRange(testCases[i].begin, testCases[i].end)
		}
	}
}

// Simple checking by calculate difference
func simpeChek(begin string, end string, prefixes []string) bool {
	beginInt, _ := strconv.ParseUint(begin, 10, 64)
	endInt, _ := strconv.ParseUint(end, 10, 64)
	expected := endInt - beginInt + 1

	l := len(begin)

	var actual uint64
	for _, p := range prefixes {
		actual += uint64(math.Pow10(l - len(p)))
	}

	return expected == actual
}
