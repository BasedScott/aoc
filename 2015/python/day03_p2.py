
def mapping(text):
    return {
        '^': (0,+1),
        'v': (0,-1),
        '<': (-1,0),
        '>': (+1,0),
    }.get(text)

santaPath= [(0,0)]
roboPath = [(0,0)]

with open("2015/data/day03.txt", "r") as file:
    text = file.read()
    for i, c in enumerate(text):
        if i % 2 == 0:
            coord = mapping(c)
            if coord:
                santa_x, santa_y = santaPath[-1]
                santaPath.append((santa_x + coord[0], santa_y + coord[1]))
        else:
            coord = mapping(c)
            if coord:
                robo_x, robo_y = roboPath[-1]
                roboPath.append((robo_x + coord[0], robo_y + coord[1]))

uniqueHouses = len(set(santaPath + roboPath))

print (uniqueHouses)