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
        panic("Ошибка: неверный формат ввода")
    }

    a := parts[0]
    operator := parts[1]
    b := parts[2]

    if isRoman(a) && isRoman(b) {
        result, err := calculateRoman(a, b, operator)
        if err != nil {
            panic(err)
        } else {
            fmt.Println("Результат:", result)
        }
    } else if isArabic(a) && isArabic(b) {
        result, err := calculateArabic(a, b, operator)
        if err != nil {
            panic(err)
        } else {
            fmt.Println("Результат:", result)
        }
    } else {
        panic("Ошибка: используйте либо арабские, либо римские цифры")
    }
}

func isRoman(s string) bool {
    romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
    for _, numeral := range romanNumerals {
        if s == numeral {
            return true
        }
    }
    return false
}

func isArabic(s string) bool {
    num, err := strconv.Atoi(s)
    if err != nil {
        return false
    }
    return num >= 1 && num <= 10
}

func calculateRoman(a, b, operator string) (string, error) {
    romanToArabic := map[string]int{
        "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
        "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
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

    return arabicToRoman(result), nil
}

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

func arabicToRoman(num int) string {
    val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    roman := ""
    for i := 0; i < len(val); i++ {
        for num >= val[i] {
            num -= val[i]
            roman += syb[i]
        }
    }
    return roman
}
