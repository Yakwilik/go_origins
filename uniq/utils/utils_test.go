package utils

import (
	"reflect"
	"testing"
)

func TestSkipChars(t *testing.T) {
	words := []string{"1", "22", "333", "4444", "55555", "123456", "1234567"}
	expectedWords := []string{"", "2", "33", "444", "5555", "23456", "234567"}

	for i := range words {
		if SkipChars(words[i], 1) != expectedWords[i] {
			t.Errorf("%s != %s – expected result", SkipChars(words[i], 1), expectedWords[i])
		}
	}
	if SkipChars("123456789", 5) != SkipChars(SkipChars("123456789", 2), 3) {
		t.Errorf("SkipChars(\"123456789\") != SkipChars(SkipChars(\"123456789\", 2), 3)")
	}

	if SkipChars("line", 999) != "" {
		t.Errorf("skipping more chars than string contain must return empty string")
	}
}

func TestSkipWords(t *testing.T) {
	line := "one two three"
	expectedLine := "two three"
	if SkipWords(line, 1) != expectedLine {
		t.Errorf("%s != expected line %s", SkipWords(line, 1), expectedLine)
	}

	line = " one two three"
	expectedLine = "two three"
	if SkipWords(line, 1) != expectedLine {
		t.Errorf("%s != expected line %s", SkipWords(line, 1), expectedLine)
	}

	line = " one two three "
	expectedLine = "two three "
	if SkipWords(line, 1) != expectedLine {
		t.Errorf("%s != expected line %s", SkipWords(line, 1), expectedLine)
	}

	line = " one two three"
	expectedLine = "three"
	if SkipWords(line, 2) != expectedLine {
		t.Errorf("%s != expected line %s", SkipWords(line, 2), expectedLine)
	}

	line = " one two three"
	expectedLine = ""
	if SkipWords(line, 10) != expectedLine {
		t.Errorf("%s != expected line %s", SkipWords(line, 10), expectedLine)
	}
}

func TestGetStringsWithMetCount(t *testing.T) {
	strs := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}
	rslt := []StringWithMetCount{
		{Str: "I love music.", MetCount: 3},
		{Str: "", MetCount: 1},
		{Str: "I love music of Kartik.", MetCount: 2},
		{Str: "Thanks.", MetCount: 1},
		{Str: "I love music of Kartik.", MetCount: 2},
	}

	if !reflect.DeepEqual(GetStringsWithMetCount(strs, false, 0, 0), rslt) {
		t.Fatalf("Struct must be equal")
	}

	strs = []string{
		"I love music.",
		"B love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"Q love music of Kartik.",
		"Thanks.",
		"B love music of Kartik.",
		"C love music of Kartik.",
	}
	rslt = []StringWithMetCount{
		{Str: "I love music.", MetCount: 3},
		{Str: "", MetCount: 1},
		{Str: "I love music of Kartik.", MetCount: 2},
		{Str: "Thanks.", MetCount: 1},
		{Str: "B love music of Kartik.", MetCount: 2},
	}
	if !reflect.DeepEqual(GetStringsWithMetCount(strs, false, 0, 1), rslt) {
		t.Fatalf("Struct must be equal")
	}

	strs = []string{
		"I love music.",
		"B Love music.",
		"C loVe mUsic.",
		"",
		"I love MUSIC of Kartik.",
		"Q love music of Kartik.",
		"Thanks.",
		"B love music of Kartik.",
		"C lOve music of Kartik.",
	}
	rslt = []StringWithMetCount{
		{Str: "I love music.", MetCount: 3},
		{Str: "", MetCount: 1},
		{Str: "I love MUSIC of Kartik.", MetCount: 2},
		{Str: "Thanks.", MetCount: 1},
		{Str: "B love music of Kartik.", MetCount: 2},
	}
	if !reflect.DeepEqual(GetStringsWithMetCount(strs, true, 0, 1), rslt) {
		t.Fatalf("Struct must be equal")
	}

	strs = []string{
		"I love music.",
		"babaB yove music.",
		"lalaC soVe mUsic.",
		"",
		"I love MUSIC of Kartik.",
		"GoodByeQ ;ove music of Kartik.",
		"Thanks.",
		"B love music of Kartik.",
		"QEQC lOve music of Kartik.",
	}
	rslt = []StringWithMetCount{
		{Str: "I love music.", MetCount: 3},
		{Str: "", MetCount: 1},
		{Str: "I love MUSIC of Kartik.", MetCount: 2},
		{Str: "Thanks.", MetCount: 1},
		{Str: "B love music of Kartik.", MetCount: 2},
	}
	if !reflect.DeepEqual(GetStringsWithMetCount(strs, true, 1, 1), rslt) {
		t.Fatalf("Struct must be equal")
	}
}

func TestGetStringCompareFunc(t *testing.T) {
	compareFunc := GetStringCompareFunc(true)

	res := compareFunc("Hello", "Hello")
	if !res {
		t.Fatalf("Compare function works wrong")
	}

	res = compareFunc("123", "111")
	if res {
		t.Fatalf("Compare function works wrong")
	}

	res = compareFunc("Hello", "heLLo")
	if !res {
		t.Errorf("Compare function must ignore registr")
	}

	compareFunc = GetStringCompareFunc(false)

	res = compareFunc("Hello", "Hello")
	if !res {
		t.Fatalf("Compare function works wrong")
	}

	res = compareFunc("123", "111")
	if res {
		t.Fatalf("Compare function works wrong")
	}

	res = compareFunc("Hello", "heLLo")
	if res {
		t.Errorf("Compare function must not ignore registr")
	}

}
