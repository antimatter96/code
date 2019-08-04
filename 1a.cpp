#include <iostream>

using namespace std;

int main() {
  long long int n, m;
  long long int a;
  cin >> n;
  cin >> m;
  cin >> a;

  int temp;
  temp = n % a;
  if (temp != 0) {
    n += a - temp;
  }
  temp = m % a;
  if (temp != 0) {
    m += a - temp;
  }
  long long int res =  (m * n) / (a * a);
  cout << res;
  return 0;
}
