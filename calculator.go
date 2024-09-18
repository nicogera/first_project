package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите выражение: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    parts := strings.Split(input, " ")
    if len(parts) != 3 {
        fmt.Println("Ошибка: неверный формат ввода")
        return
    }

    a := parts[0]
    operator := parts[1]
    b := parts[2]

    if isRoman(a) && isRoman(b) {
        result, err := calculateRoman(a, b, operator)
        if err != nil {
            fmt.Println("Ошибка:", err)
        } else {
            fmt.Println("Результат:", result)
        }
    } else if isArabic(a) && isArabic(b) {
        result, err := calculateArabic(a, b, operator)
        if err != nil {
            fmt.Println("Ошибка:", err)
        } else {
            fmt.Println("Результат:", result)
        }
    } else {
        fmt.Println("Ошибка: используйте либо арабские, либо римские цифры")
    }
}

// Проверяет, является ли строка римской цифрой
func isRoman(s string) bool {
    romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
    for _, numeral := range romanNumerals {
        if s == numeral {
            return true
        }
    }
    return false
}

// Проверяет, является ли строка арабской цифрой и находится ли в диапазоне от 1 до 10
func isArabic(s string) bool {
    num, err := strconv.Atoi(s)
    if err != nil {
        return false
    }
    return num >= 1 && num <= 10
}

// Выполняет арифметическую операцию с римскими цифрами
func calculateRoman(a, b, operator string) (string, error) {
    romanToArabic := map[string]int{
        "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
        "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
    }
    arabicToRoman := []string{
        "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
    }

    aVal := romanToArabic[a]
    bVal := romanToArabic[b]
    var result int

    switch operator {
    case "+":
        result = aVal + bVal
    case "-":
        result = aVal - bVal
    case "*":
        result = aVal * bVal
    case "/":
        result = aVal / bVal
    default:
        return "", fmt.Errorf("неизвестная операция")
    }

    if result <= 0 {
        return "", fmt.Errorf("результат меньше единицы")
    }

    return arabicToRoman[result], nil
}

// Выполняет арифметическую операцию с арабскими цифрами
func calculateArabic(a, b, operator string) (string, error) {
    aVal, _ := strconv.Atoi(a)
    bVal, _ := strconv.Atoi(b)
    var result int

    switch operator {
    case "+":
        result = aVal + bVal
    case "-":
        result = aVal - bVal
    case "*":
        result = aVal * bVal
    case "/":
        result = aVal / bVal
    default:
        return "", fmt.Errorf("неизвестная операция")
    }

    return strconv.Itoa(result), nil
}
