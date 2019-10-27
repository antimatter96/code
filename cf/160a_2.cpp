#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int main() {
  int n;
  cin >> n;

  vector<int> arr(n);

  int total = 0;
  for (int i = 0; i < n; i++) {
    cin >> arr[i];
    total += arr[i];
  }

  sort(arr.begin(), arr.end(), std::greater<int>());

  int tillNow = 0;
  int i = 0;
  int remaining = total;
  for (; i < n; i++) {
    tillNow += arr[i];
    remaining -= arr[i];
    if (tillNow > remaining) {
      break;
    }
  }
  cout << i + 1;
  return 0;
}
