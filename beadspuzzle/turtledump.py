from turtle import *
from beads import *

#NOTE: That just uses the other script to get the number of combinations 
#and then uses turtle to draw the beads in green

tt = Turtle()
tt.speed(10)
def drawSquare(t, gap, fill):
    t.color('green')
    if (fill == 1):
        t.begin_fill()
    for i in range(4):        
        t.forward(10)
        t.right(90)
    if (fill == 1):
        t.end_fill()      


mySet = getCombinations(12)
print(mySet)

for s in range(len(mySet)):
    subset = mySet[s]
    t = s * 30
    for i in range(len(subset)):
        o = subset[i]
        tt.penup()
        tt.goto(i*30, t)
        tt.pendown()
        drawSquare(tt, 10, o)
    
done()