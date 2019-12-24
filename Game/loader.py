import sys
import os


def loadMap():
    with open("level.txt") as f:
        data = f.readlines()
    f.close()    
    return data

def loadMaps():
    with open(os.path.join("Levels","maps.txt")) as f:
        data = f.readlines()
    f.close()    
    return data  

def loadLevels():
    levels = []
    lines = loadMaps()
    level = []
    lidx = 0
    for idx, line in enumerate(lines):
        if line.find("Maze") >= 0:
            if len(level) > 0:
                level.pop()
                levels.append(level)
                level = []
            lidx = 0
        if lidx > 6:
            level.append(line)
        lidx += 1
    return levels
    


