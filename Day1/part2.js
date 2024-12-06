const fs = require("fs");
const filepath = "input.txt";
const data = fs.readFileSync(filepath, "utf8");

const lines = data.split("\n");

const left = [];
const right = [];

lines.forEach((line) => {
  const [leftNum, rightNum] = line
    .trim()
    .split(/\s+/)
    .map((num) => parseInt(num, 10));
  if (!isNaN(leftNum)) {
    left.push(leftNum);
  }
  if (!isNaN(rightNum)) {
    right.push(rightNum);
  }
});

// create a dictionary
// iterate through right, if right in dict, dict[right] += 1
// iterate through left, if left in dict, sum += left * dict[left]

const dict = {};
right.forEach((num) => {
  if (!isNaN(num)) {
    dict[num] = (dict[num] || 0) + 1;
  }
});

let sum = 0;
left.forEach((num) => {
  if (dict[num]) {
    sum += num * dict[num];
  }
});
console.log(sum);
