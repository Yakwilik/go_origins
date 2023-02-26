package handlers

import (
	"GolangCourse/uniq/options"
	"reflect"
	"testing"
)

func TestWithoutOptions(t *testing.T) {
	input := []string{
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
	expResultWithoutOptions := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
	}
	opts := options.Options{}

	if !reflect.DeepEqual(HandleLines(input, opts), expResultWithoutOptions) {
		t.Fatalf("wrong result")
	}
}

func TestStrMeetCountFlag(t *testing.T) {
	input := []string{
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
	opts := options.Options{EShowStrMeetCount: true}

	expResultWithCFlag := []string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}

	if !reflect.DeepEqual(HandleLines(input, opts), expResultWithCFlag) {
		t.Fatalf("wrong result")
	}
}

func TestShowNotUniqueStr(t *testing.T) {
	input := []string{
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
	opts := options.Options{EShowNotUniqueStr: true}

	expResultWithDFlag := []string{
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	if !reflect.DeepEqual(HandleLines(input, opts), expResultWithDFlag) {
		t.Fatalf("wrong result")
	}
}

func TestShowUniqueStr(t *testing.T) {
	input := []string{
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
	opts := options.Options{EShowUniqueStr: true}

	expResultWithUFlag := []string{
		"",
		"Thanks.",
	}
	if !reflect.DeepEqual(HandleLines(input, opts), expResultWithUFlag) {
		t.Fatalf("wrong result")
	}
}

func TestIgnoreRegisterFlag(t *testing.T) {
	inputIFlag := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik.",
	}
	opts := options.Options{IgnoreRegister: true}

	expResultWithIFlag := []string{
		"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik.",
	}
	if !reflect.DeepEqual(HandleLines(inputIFlag, opts), expResultWithIFlag) {
		t.Fatalf("wrong result")
	}
}

func TestIgnoreNumFields(t *testing.T) {
	inputFFlag := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	expResultWithFFlag := []string{
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	opts := options.Options{SkippedStringsCount: 1}
	if !reflect.DeepEqual(HandleLines(inputFFlag, opts), expResultWithFFlag) {
		t.Fatalf("wrong result")
	}
}

func TestIgnoreNumChars(t *testing.T) {
	inputSFlag := []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	expResultWithSFlag := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	opts := options.Options{SkippedCharsCount: 1}
	if !reflect.DeepEqual(HandleLines(inputSFlag, opts), expResultWithSFlag) {
		t.Fatalf("wrong result")
	}

}
