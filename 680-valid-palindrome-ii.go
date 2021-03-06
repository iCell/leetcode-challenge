func validPalindrome(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        if s[left] != s[right] {
            break
        }
        left++
        right--
    }

    return isPalindrome(s, left, right-1) || isPalindrome(s, left+1, right)
}

func isPalindrome(s string, left, right int) bool {
    if left >= right {
        return true
    }
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func validPalindrome(s string) bool {

    var indexes []int

    i, j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            indexes = append(indexes, i)
            indexes = append(indexes, j)
            break
        }
        i++
        j--
    }

    if len(indexes) == 0 {
        return true
    }

    results := make([]bool, 0, 2)
    for _, index := range indexes {
        str := s
        if index == len(str)-1 {
            str = str[:index]
        } else {
            str = str[:index] + str[index+1:]
        }

        results = append(results, check(str))
    }

    for _, r := range results {
        if r == true {
            return true
        }
    }

    return false
}

func check(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

