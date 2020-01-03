import random
import math
from functools import reduce

# Symetrical side
totalBeads = 8
symbols = ['O', 'X']

# Randomly generate candidates of arrays of particular combination size


def getMeCandidateArray(sizeOfTarget):
    cand = []
    for i in range(sizeOfTarget):
        o = random.randint(0, 1)
        cand.append(o)
    return cand

# Check whether 2 arrays are the same


def isTheSameArray(a, b):
    c = 0
    for i in range(len(a)):
        for j in range(len(b)):
            if i == j and a[i] == b[j]:
                c = c + 1
    return c == len(a)

# Check whether we already have the given combination in the stash


def alreadyIn(all, cand):
    return len(list(filter(lambda c: isTheSameArray(c, cand), all))) != 0

# Calculate the sum of the elements in the array


def getSum(r):
    return reduce(lambda x, y: (x + y), r)

# Just to decorate 0,1 to nicer symbols. like O and X, etc, or A,B or + -


def decorateSymbols(r):
    return list(map(lambda a: symbols[a]))


def getTarget(b):
    return math.factorial(b) / ((math.factorial(b/2) * math.factorial(b - b/2)))

# Find all the combinations that meet the symmetry condition


def getCombinations(totalBeads):
    beadsOnSide = int(totalBeads / 2)
    targetSet = []
    while len(targetSet) < getTarget(beadsOnSide):
        r = getMeCandidateArray(beadsOnSide)
        if (not alreadyIn(targetSet, r) and getSum(r) == beadsOnSide/2):
            targetSet.append(r)
    return targetSet

# Decorate the final result


def decorateArray(cand):
    return list(map(lambda i: list(map(lambda a: symbols[a], i)), cand))


print("Number of total beads:", totalBeads)
print("Number of beads on side if symmetrical:", int(totalBeads / 2))
print("Target number of combinations:", getTarget(int(totalBeads / 2)))
print(decorateArray(getCombinations(totalBeads)))
