#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <string>
#include <cmath>
using namespace std;

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

int main() {
    ifstream file("input.txt");
    string line;
    int safe;

    while (getline(file, line)) {
        istringstream iss(line);
        vector<int> nums;
        int num;

        while (iss >> num) {
            nums.push_back(num);
        }

        if (isAscending(nums) and isValidDifference(nums)) {
            safe++;
        } else if (isDescending(nums) and isValidDifference(nums)) {
            safe++;
        } else {
            cout << "Unsafe Report." << endl;
        }

    }

    file.close();
    cout << safe << endl;
    return 0;

}
