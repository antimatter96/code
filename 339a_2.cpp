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

  int startAt = 0;
  for (int num = 0; num < 3; num++) {
    for (int i = 0, j = startAt; i < arr[num]; i++, j += 2) {
      s[j] = 48 + 1 + num;
    }
    startAt += (2 * arr[num]);
  }

  cout << s;
  return 0;
}
