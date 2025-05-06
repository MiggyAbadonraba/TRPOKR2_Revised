// Файл: operations_test.go
package main

import (
	"testing"
)

// Тесты для сложения и вычитания (Разработчик 1)
func TestAddSubtract(t *testing.T) {
	tests := []struct {
		a, b float64
		add  float64
		sub  float64
	}{
		{2, 3, 5, -1},
		{-5, 3, -2, -8},
		{0, 0, 0, 0},
	}

	for _, tt := range tests {
		if got := Add(tt.a, tt.b); got != tt.add {
			t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.add)
		}
		if got := Sub(tt.a, tt.b); got != tt.sub {
			t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.sub)
		}
	}
}

// Тесты для умножения и деления (Разработчик 2)
func TestMultiplyDivide(t *testing.T) {
	tests := []struct {
		a, b    float64
		mul     float64
		div     float64
		wantErr bool
	}{
		{4, 2, 8, 2, false},
		{-3, 3, -9, -1, false},
		{5, 0, 0, 0, true}, // Деление на ноль
	}

	for _, tt := range tests {
		// Проверка умножения
		if got := Multiply(tt.a, tt.b); got != tt.mul {
			t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.mul)
		}

		// Проверка деления
		got, err := Divide(tt.a, tt.b)
		if tt.wantErr {
			if err == nil {
				t.Errorf("Divide(%v, %v) ожидалась ошибка", tt.a, tt.b)
			}
		} else {
			if got != tt.div {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.div)
			}
		}
	}
}

// Тесты для возведения в степень (Разработчик 3)
func TestPower(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{2, 3, 8},
		{5, 0, 1},    // Любое число в 0 степени
		{2, -1, 0.5}, // Отрицательная степень
	}

	for _, tt := range tests {
		if got := Power(tt.a, tt.b); got != tt.want {
			t.Errorf("Power(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
