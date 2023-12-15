package utils

func SwitchCharacters(str string, index1, index2 int) string {
    charArray := []rune(str)
    char1 := charArray[index1]
    char2 := charArray[index2]
    charArray[index1] = char2
    charArray[index2] = char1
    return string(charArray)
}

func SwitchStringCharacters(str1, str2 string, index1, index2 int) (string, string) {
    charArray1 := []rune(str1)
    charArray2 := []rune(str2)
    char1 := charArray1[index1]
    char2 := charArray2[index2]
    charArray1[index1] = char2
    charArray2[index2] = char1
    return string(charArray1), string(charArray2)
}

func ExtractDigits(str string) string {
    digits := ""
    for _, ch := range str {
        if ch >= '0' && ch <= '9' {
            digits += string(ch)
        }
    }
    return digits
}
