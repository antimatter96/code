#include <iostream>
#include <string>

using namespace std;

int main() {
  string target = "hello";
  int foundTill = -1;
  int breakCondition = target.length() - 1;
  string s;
  cin >> s;

  for (int i = 0; i < s.length(); i++) {
    if(s[i] == target[foundTill+1]) {
      foundTill++;
    } else {
      continue;
    }
    if(foundTill == breakCondition) {
      break;
    };
  }

  if(foundTill == breakCondition) {
    cout << "YES";
  } else {
    cout << "NO";
  }

  return 0;
}
