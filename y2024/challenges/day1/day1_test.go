package day1

import (
    "testing"
)

func TestGetTotalDistance(t *testing.T) {
    historian1 := []int{1, 2, 3, 3, 3, 4}
    historian2 := []int{3, 3, 3, 4, 5, 9}
    expectedTotalDistance := 11

    actualTotalDistance := getTotalDistance(historian1, historian2);

    if expectedTotalDistance != actualTotalDistance {
        t.Errorf("Expected %v, got %v", expectedTotalDistance, actualTotalDistance)
    }
}

func TestGetSimilarityScore(t *testing.T) {
    historian1 := []int{1, 2, 3, 3, 3, 4}
    numberOfOccurances := map[int]int{3:3, 4:1, 5:1, 9:1}
    expectedSimilarityScore := 31

    actualSimilarityScore := getSimilarityScore(historian1, numberOfOccurances);

    if expectedSimilarityScore != actualSimilarityScore {
        t.Errorf("Expected %v, got %v", expectedSimilarityScore, actualSimilarityScore)
    }

}
