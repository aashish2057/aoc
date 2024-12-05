--- Day 4: Ceres Search ---

Part 1
---
Given a file with letters, find all occurances of the word XMAS. The word can be horizontal, vertical, diagona, backwards, and
overlapping

1. Process file into 2d matrix
2. Create pointer to search
3. Create variable to hold possible ways the pointer can move (0,1), (1,0), (0,-1), (-1,0), (1, -1), (-1, 1), (1, 1), (-1,-1)
4. If any adjacent positions to current pointer have a valid next character, keep moving in that direction until all characters,
    are found or an invalid character is found
5. Move pointer one position

Part 2
---

1. Look for A, since its the middle.
2. Check all the corners for all possabilities
3. If it matches return 1 else return 0
