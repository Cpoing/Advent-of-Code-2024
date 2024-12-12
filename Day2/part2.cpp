#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <string>
#include <cmath>
using namespace std;

bool isValidDifference(const vector<int>& nums);
bool isAscending(const vector<int>& nums);
bool isDescending(const vector<int>& nums);
bool isSafeWithDampener(const vector<int>& originalNums);

bool isValidDifference(const vector<int>& nums) {
  for (size_t i = 1; i < nums.size(); i++) {
    int diff = abs(nums[i] - nums[i - 1]);
    if (diff < 1 || diff > 3) {
      return false;
    }
  }
  return true;
}

bool isAscending(const vector<int>& nums) {
    for (size_t i = 1; i < nums.size(); ++i) {
        if (nums[i] < nums[i - 1]) {
            return false;
        }
    }
    return true;
}

bool isDescending(const vector<int>& nums) {
    for (size_t i = 1; i < nums.size(); ++i) {
        if (nums[i] > nums[i - 1]) {
            return false;
        }
    }
    return true;
}

bool isSafeWithDampener(const vector<int>& originalNums) {
    if ((isAscending(originalNums) || isDescending(originalNums)) && isValidDifference(originalNums)) {
        return true;
    }

    for (size_t remove = 0; remove < originalNums.size(); ++remove) {
        vector<int> nums = originalNums;
        nums.erase(nums.begin() + remove);

        if ((isAscending(nums) || isDescending(nums)) && isValidDifference(nums)) {
            return true;
        }
    }

    return false;
}

int main() {
    ifstream file("input.txt");
    string line;
    int safe = 0; 

    while (getline(file, line)) {
        istringstream iss(line);
        vector<int> nums;
        int num;
        while (iss >> num) {
            nums.push_back(num);
        }

        if (isSafeWithDampener(nums)) {
            safe++;
        } else {
            cout << "Unsafe report." << endl;
        }
    }
    file.close();
    
    cout << safe << endl;
    return 0;
}
