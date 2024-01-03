import re

with open("2015/data/day02.txt", "r") as file:
    text = file.read()

numbers = re.findall("[0-9]+", text)

answer = []

for i in range(0, len(numbers), 3):
    if i + 2 < len(numbers):
        l = int(numbers[i])
        w = int(numbers[i + 1])
        h = int(numbers[i + 2])

        ribbonWrapList = [l,w,h]
        ribbonWrapList.sort()
        ribbonWrapSort = ribbonWrapList[0-1]

        ribbonWrap = ribbonWrapList[0] + ribbonWrapList[0] + ribbonWrapList[1] + ribbonWrapList[1]
        
        ribbonBow = l * w * h
        
        answer.append(ribbonWrap)
        answer.append(ribbonBow)

print (sum(answer))