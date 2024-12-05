--- Day 2: Red-Nosed Reports ---

Part 1
---

You are given a file, where each line is a report, each report is a list of levels represented as numbers seperated by spaces.
Determine which reports are safe. A report is safe is all levels are increasing or all decreasing and any 2 adjacent levels
differ by atleast 1 and at most 3.

1. Read file into list of list of ints.
2. Create variable to track totalSafeReports
3. Loop through list of list of ints.
4. Check if report is safe
    4a. Check if list is increasing or decreasing
    4b. Check if all levels differ by atleast 1 or at most 3
5. If report is safe add one to totalSafeReports
6. Return totalSafeReports

Part 2
---

1. If report is unsafe, keep removing index and re checking if safe. If you find report is safe by removing one level break
