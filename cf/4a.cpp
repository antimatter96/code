#include <iostream>

using namespace std;

int main() {
  int weight;
  cin >> weight;
  if (weight == 2) {
    cout << "No";
  } else if (weight % 2 == 0) {
    cout << "Yes";
  } else {
    cout << "No";
  }
  return 0;
}
