#include <iostream>
#include <string>

using namespace std;

int main() {
  string s1;
  cin >> s1;
  string s2;
  cin >> s2;
  int c1, c2;
  int res = 0;
  for (int i = 0; i < s1.length(); i++) {
    c1 = s1[i];
    c2 = s2[i];
    if (c1 <= 90) {
      c1 += 32;
    }
    if (c2 <= 90) {
      c2 += 32;
    }
    if (c1 > c2) {
      res = 1;
      break;
    } else if (c2 > c1) {
      res = -1;
      break;
    }
  }
  cout << res;
  return 0;
}
