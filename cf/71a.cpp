#include <iostream>
#include <string>

using namespace std;

int main() {
  int n;
  cin >> n;
  string s;
  int len;
  while (n > 0) {
    cin >> s;
    if (s.length() > 10) {
      len = s.length();
      cout << s[0] << len - 2 << s[len - 1];
    } else {
      cout << s;
    }
    cout << '\n';
    n--;
  }
  return 0;
}
