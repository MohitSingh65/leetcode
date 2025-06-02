#include <bits/stdc++.h>
#include <map>
#include <string>

class Solution {
    public:
    std::vector<std::string> letterCombinations(std::string digits) {
        if(digits.empty()){
            return {};
        }

        std::map<char, std::string> 
        digit_to_char {{'2', "abc"}, {'3', "def"}, {'4', "ghi"}, {'5', "jkl"}, {'6', "mno"}, {'7', "pqrs"}, {'8', "tuv"}, {'9', "wxyz"},
        };

        std::vector<std::string> result;
        std::string current_combination;


        std::function<void(int)> backtrack = [&](int index) {
            if (index == digits.size()) {
                result.push_back(current_combination);
                return;
            }

            for (char c: digit_to_char[digits[index]]) {
                current_combination.push_back(c);
                backtrack(index + 1);
                current_combination.pop_back();
            }
        };

        backtrack(0);
        return result;
    };
};
