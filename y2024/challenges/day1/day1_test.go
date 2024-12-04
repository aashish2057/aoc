package day1

import (
    "testing"
    "reflect"
)

func TestProcessFile(t *testing.T) {
    filePath := "./test.txt"

    expectedLocations1 := []int{1, 2, 3, 3, 3, 4}
    expectedLocations2 := []int{3, 3, 3, 4, 5, 9}

    actualLocations1, actualLocations2 := processFile(filePath)

    if reflect.DeepEqual(actualLocations1, expectedLocations1) {
        t.Errorf("Expected: %v, Actual: %v, when calling processFile with input: %v",
            expectedLocations1, actualLocations1, filePath)
    }
    if reflect.DeepEqual(actualLocations2, expectedLocations2) {
        t.Errorf("Expected: %v, Actual: %v, when calling processFile with input: %v",
            expectedLocations2, actualLocations2, filePath)
    }
}

func TestCalculateTotalDistance(t *testing.T) {
    locations1 := []int{1, 2, 3, 3, 3, 4}
    locations2 := []int{3, 3, 3, 4, 5, 9}
    expectedTotalDistance := 11

    actualTotalDistance := calculateTotalDistance(locations1, locations2)

    if actualTotalDistance != expectedTotalDistance {
        t.Errorf("Expected: %d, Actual: %d, when calling calculateTotalDistance with: %v and %v",
            expectedTotalDistance, actualTotalDistance, locations1, locations2)
    }
}

func TestCalculateSimilarityScore(t *testing.T) {
    locations1 := []int{1, 2, 3, 3, 3, 4}
    locations2 := []int{3, 3, 3, 4, 5, 9}
    mapLocations2 := getMapOccuranceOfLocationIds(locations2)
    expectedSimilarityScore := 31
    actualSimilarityScore := calculateSimilarityScore(locations1, mapLocations2)

    if actualSimilarityScore != expectedSimilarityScore {
        t.Errorf("Expected: %d, Actual: %d, when calling calculateSimilarityScore with: %v and %v",
            expectedSimilarityScore, actualSimilarityScore, locations1, mapLocations2)
    }

}


