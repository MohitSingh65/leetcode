class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        words = s.strip().split()
        lastWord = words[-1]
        return len(lastWord)
