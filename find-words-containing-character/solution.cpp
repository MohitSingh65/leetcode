#include <vector>
#include <string>
using namespace std;

class Solution {
public:
  vector<int> findWordsContaining(vector<string> &words, char x) {
    vector<int> empty = {};
    if (sizeof(words) == 0) {
      return empty;
    }

    vector<int> result;
    for (int i = 0; i < words.size(); ++i) {
      if (words[i].find(x) != std::string::npos) {
        result.push_back(i);
      }
    }

    return result;
  }
};
