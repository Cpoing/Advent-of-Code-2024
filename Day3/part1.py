import re

with open("./input.txt", "r") as file:
    lines = file.readlines()

pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
total = 0
for element in lines:
    multiplications = re.findall(pattern, element)
    for x, y in multiplications:
        total += int(x) * int(y)
print(total)

