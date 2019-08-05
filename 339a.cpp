#include <iostream>
#include <string>

using namespace std;

int main() {
  string s;
  cin >> s;

  int arr[3] = {0, 0, 0};
  for (int i = 0; i < s.length(); i++) {
    if (s[i] > 48 && s[i] < 52) {
      arr[s[i] - 48 - 1]++;
    }
  }

  for (int i = 0, j = 0; i < arr[0]; i++, j += 2) {
    s[j] = 48 + 1;
  }
  for (int i = 0, j = (2 * arr[0]); i < arr[1]; i++, j += 2) {
    s[j] = 48 + 2;
  }
  for (int i = 0, j = (2 * (arr[0] + arr[1])); i < arr[2]; i++, j += 2) {
    s[j] = 48 + 3;
  }

  cout << s;
  return 0;
}
