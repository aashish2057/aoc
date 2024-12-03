package util

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func Abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
