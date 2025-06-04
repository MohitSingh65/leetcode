from typing import List


class Solution:
    def divideString(self, s: str, k: int, fill: str) -> List[str]:
        result = []
        for i in range(0, len(s), k):
            group = s[i : i + k]
            if len(group) < k:
                group = group.ljust(k, fill)
            result.append(group)
        return result
