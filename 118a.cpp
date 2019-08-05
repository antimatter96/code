#include <iostream>
#include <string>
#include <vector>

using namespace std;

int main() {
  string s;
  cin >> s;
  vector<int> arr;
  for (int i = 0; i < s.length(); i++) {
    if(s[i] <= 90) {
      s[i] += 32;
    }
    if(s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'y') {

    } else {
      arr.push_back(s[i] - 0);
    }
  }
  if (arr.size() == 0 ) {
    cout << '.';
  } else {
    for(int i = 0; i < arr.size(); i++) {
      cout << '.' << char(arr[i]);
    }
  }
  return 0;
}
