#include <iostream>
#include <vector>

using namespace std;

int main() {
  int n;
  cin >> n;

  vector<int> arr = {4,   7,   44,  47,  74,  77,  444,
                     447, 474, 477, 744, 747, 774, 777};
  bool found = false;
  for (int i = 0; i < arr.size() && n >= arr[i]; i++) {
    if ( n % arr[i] == 0) {
      found = true;
      break;
    }
  }
  if(found == true) {
    cout << "YES";
  } else {
    cout << "NO";
  }
  return 0;
}
