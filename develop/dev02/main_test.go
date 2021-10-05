package dev02_test

import (
	"testing"
	"wb_lvl2/develop/dev02"
)

func TestUnpack(t *testing.T) {
	// Тестовые данные
	testString := "a4bc2d5e"
	expectedString := "aaaabccddddde"

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestUnpack1(t *testing.T) {
	// Тестовые данные
	testString := "abcd"
	expectedString := "abcd"

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestUnpack2(t *testing.T) {
	// Тестовые данные
	testString := "45"
	expectedString := ""

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestUnpack3(t *testing.T) {
	// Тестовые данные
	testString := ""
	expectedString := ""

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestUnpack4(t *testing.T) {
	// Тестовые данные
	testString := `qwe\4\5`
	expectedString := "qwe45"

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestUnpack5(t *testing.T) {
	// Тестовые данные
	testString := `qwe\45`
	expectedString := "qwe44444"

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestUnpack6(t *testing.T) {
	// Тестовые данные
	testString := `qwe\\5`
	expectedString := `qwe\\\\\`

	// Тестирование
	test, err := dev02.Unpack(testString)

	// Проверка результатов
	if test != expectedString {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}
