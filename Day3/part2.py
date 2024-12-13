import re

with open("./input.txt", "r") as file:
    lines = file.read()

# filter the input text,
# split the text by the substring "do()"
# for each substring in the split, split again by "dont()" and keep the first substring aka the substring before the dont()
# dont()mul(7,7)asdfgdo()mul(7,8)dasjhfiesdont()mul(6,6)
# split by "do()" [dont()mul(7,7)asdfgdo, mul(7,8)dasjhfiesdont()mul(6,6)]
# split by "dont()" and keep the 0th substring [mul(7,8)dasjhfies]
# result is a string of only the characters we care about

filtered = list()
for element in lines.split("do()"):
    filtered.append(element.split("don't()")[0])
print(filtered)

pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
total = 0
for element in filtered:
    multiplications = re.findall(pattern, element)
    for x, y in multiplications:
        total += int(x) * int(y)
print(total)


