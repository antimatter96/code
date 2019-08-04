#include <iostream>
#include <vector>

using namespace std;

int main() {
  int n;
  cin >> n;
  int k;
  cin >> k;
  vector<int> arr(n);
  for (int i = 0; i < n; i++) {
    cin >> arr[i];
  }
  int res = 0;
  int target = arr[k - 1];
  for (int i = 0; i < n; i++) {
    if (arr[i] <= 0) {
      break;
    }
    if (arr[i] >= target) {
      res++;
    }
  }
  cout << res;
  return 0;
}
