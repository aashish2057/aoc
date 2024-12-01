package util

import (
	"fmt"
	"os"
)

func ReadFile(filename string) []byte {
    path := fmt.Sprintf("./inputs/%s", filename)

    file, error := os.ReadFile(path)
    CheckError(error)

    return file
}
