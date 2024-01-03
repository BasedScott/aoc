from collections import Counter

x, y = 0,0
path = []

def cMapping(text):
    mapping = {
        '^': (0,+1),
        'v': (0,-1),
        '<': (-1,0),
        '>': (+1,0),
    }
    return mapping.get(text)

with open("2015/data/day03.txt", "r") as file:
    text = file.read()
    for c in text:
        coord = cMapping(c)
        if coord:
            x += coord[0]
            y += coord[1]
            path.append((x,y))

print (len(list(set(path)))+1)