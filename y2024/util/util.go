package util

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func Abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
