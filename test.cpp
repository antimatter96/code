#include <iostream>
#include <vector>
#include <string>

using namespace std;

int main()
{

    vector<string> msg {"Hello"};
    //msg.push_back("Hello");
    //[ "Hello", "C++", "World", "from", "VS Code!"];

    for (int i = 0; i < msg.size(); i++)
    {
        cout << msg[i] << " " << i << " ";
    }
    cout << endl;
}