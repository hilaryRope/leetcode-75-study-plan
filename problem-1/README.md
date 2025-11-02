# LeetCode 75 - Problem 1: Merge Strings Alternately

## Problem Statement
You are given two strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

### Example 1:
```
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
```

### Example 2:
```
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s
```

### Example 3:
```
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d
```

## Solution Approach
### Intuition
We can merge the strings by iterating through both strings simultaneously and alternately picking characters from each string until we reach the end of the shorter string. Then, we append the remaining characters from the longer string to the result.

### Algorithm
1. Initialize an empty result string.
2. Determine the maximum length between the two strings.
3. Iterate from 0 to the maximum length - 1:
   - If the current index is within the length of `word1`, append the character to the result.
   - If the current index is within the length of `word2`, append the character to the result.
4. Return the merged string.

### Complexity Analysis
- **Time Complexity:** O(max(n, m)), where n and m are the lengths of `word1` and `word2` respectively.
- **Space Complexity:** O(n + m) for the result string.

## Solution Code
```python
def mergeAlternately(word1: str, word2: str) -> str:
    result = []
    max_len = max(len(word1), len(word2))
    
    for i in range(max_len):
        if i < len(word1):
            result.append(word1[i])
        if i < len(word2):
            result.append(word2[i])
            
    return ''.join(result)
```

## Test Cases
```python
def test_mergeAlternately():
    assert mergeAlternately("abc", "pqr") == "apbqcr"
    assert mergeAlternately("ab", "pqrs") == "apbqrs"
    assert mergeAlternately("abcd", "pq") == "apbqcd"
    assert mergeAlternately("", "xyz") == "xyz"
    assert mergeAlternately("123", "") == "123"
    print("All test cases passed!")

test_mergeAlternately()
```

## Constraints
- `1 <= word1.length, word2.length <= 100`
- `word1` and `word2` consist of lowercase English letters.