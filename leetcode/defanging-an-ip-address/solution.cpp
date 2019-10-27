class Solution {
public:
    string defangIPaddr(string address) {
      string ans;
      int len = address.length();
      ans.reserve(len + 15);
      
      for(int i = 0; i < len; i++) {
        if(address[i] == '.') {
          //ans.push_back('[');
          ans.append("[.]");
          //ans.push_back(']');
        } else {
          ans.push_back(address[i]);
        }
      }
      return ans;
    }
};

