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

left.sort();
right.sort();
let sum = 0;
for (let i = 0; i < left.length; i++) {
  sum += Math.abs(left[i] - right[i]);
}

console.log(sum);
