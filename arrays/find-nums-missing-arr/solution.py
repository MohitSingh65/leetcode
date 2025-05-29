# Problem: Find all numbrs disappered in an array

from typing import List


def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
    numbers = set(nums)

    missingNumbers = []

    for i in range(1, len(nums) + 1):
        if i not in numbers:
            missingNumbers.append(i)

    return missingNumbers
