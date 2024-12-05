package day3

import (
    "reflect"
    "testing"
)

func TestFindAllValidSubStrings1(t *testing.T) {
    memory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    expectedSubStrings := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

    actualSubStrings := findAllValidSubStrings(memory, `mul\(\d{1,3},\d{1,3}\)`)

    if !reflect.DeepEqual(actualSubStrings, expectedSubStrings) {
        t.Errorf("Expected: %v, Actual: %v, when calling findAllValidSubStrings with: %v",
            expectedSubStrings, actualSubStrings, memory)
    }

}

func TestFindAllValidSubStrings2(t *testing.T) {
    memory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    expectedSubStrings := []string{"mul(2,4)", "don't()","mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"}

    actualSubStrings := findAllValidSubStrings(memory, `(mul\(\d{1,3},\d{1,3}\))|(don't\(\))|(do\(\))`)

    if !reflect.DeepEqual(actualSubStrings, expectedSubStrings) {
        t.Errorf("Expected: %v, Actual: %v, when calling findAllValidSubStrings with: %v",
            expectedSubStrings, actualSubStrings, memory)
    }

}

func TestRemoveDisabledInstructions(t *testing.T) {
    validSubStrings := []string{"mul(2,4)", "don't()","mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"}
    expectedSubStrings := []string{"mul(2,4)", "mul(8,5)"}

    actualSubStrings := removeDisabledInstructions(validSubStrings)

    if !reflect.DeepEqual(actualSubStrings, expectedSubStrings) {
        t.Errorf("Expected: %v, Actual: %v, when calling removeDisabledInstructions with: %v",
            expectedSubStrings, actualSubStrings, validSubStrings)
    }
}

func TestCalculateProduct(t *testing.T) {
    tests := []struct {
        substring string
        expectedProduct int
    }{
        {"mul(2,4)", 8},
        {"mul(5,5)", 25},
        {"mul(11,8)", 88},
        {"mul(8,5)", 40},
    }

    for _, test := range tests {
        actualProduct := calculateProduct(test.substring)

        if actualProduct != test.expectedProduct {
            t.Errorf("Expected: %d, Actual: %d, when calling calculateProduct with: %s", test.expectedProduct,
                actualProduct, test.substring)
        }
    }
}
