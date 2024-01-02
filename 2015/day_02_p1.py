import re

with open("2015/day_02_rd.txt", "r") as file:
    text = file.read()

numbers = re.findall("[0-9]+", text)

answer = []

for i in range(0, len(numbers), 3):
    if i + 2 < len(numbers):
        l = int(numbers[i])
        w = int(numbers[i + 1])
        h = int(numbers[i + 2])

#       totalLength
        tL = 2 * (l * w)
        tW = 2 * (w * h)
        tH = 2 * (h * l)

        slack = [tL/2, tW/2, tH/2]
        slack.sort()
        slackT = slack[0]

        squareFeet = tL + tW + tH + slackT
        
        answer.append(squareFeet)

print (sum(answer))