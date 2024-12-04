--- Day 1: Historian Hysteria ---


Part 1
---
You are given a file where each line contains 2 numbers seperated by 3 spaces. Determine the total distance between the two columns
of data.

Pull each column of locationId into a list of numbers and sort both of them in ascending order. At each index find the difference,
between both lists and add up the differences at all the indexes to get the total distance between the two lists.

1. Store each column of locationIds in a list
2. Sort both lists
3. Create a variable to keep track of total distance
4. Loop through the lists and add the absolute value of the differenc to the total distance
5. Return the total distance

Part 2
---
Calculate the similarity score by getting the product of a locationId in the first list and the number of occurances of that
locationId in the 2nd list. Sum all the products for all locationIds in the first list to get simliarity score.

1. Create a map of 2nd list where key is the locationId and value is the number of occurances of that locationId
2. Create variable to keep track of similarity score
3. Loop through first list and multiply the locationId by the value when access the map with the locationId as key
4. Add to similairty
5. Return similarity score


