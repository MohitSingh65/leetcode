from typing import List


class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
        if not words:
            return []

        result = []
        for i in range(len(words)):
            if x in words[i]:
                result.append(i)
        return result
