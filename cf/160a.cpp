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

  sort(arr.begin(), arr.end());

  int tillNow = 0;
  int i = n - 1;
  int remaining = total;
  for (; i > -1; i--) {
    tillNow += arr[i];
    remaining -= arr[i];
    if (tillNow > remaining) {
      break;
    }
  }
  cout << n - i;
  return 0;
}
