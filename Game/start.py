import pygame, sys
import random
import os 
pygame.init()

width = 800
height = 600
windowSize = (width, height)
screen = pygame.display.set_mode(windowSize)

#myriadProFont = pygame.font.SysFont("Myriad Pro", 48)
#myriadProFont.render(" Shark ", 1, (255,0,0), (255,255,255))
a = os.path.join("Assets", "julien.jpeg")
hello = pygame.image.load(a)
pygame.mouse.set_visible(0)



helloSize= hello.get_size()
helloWidth = helloSize[0]
helloHeight = helloSize[1]
directionX = 1
directionY = 1

x,y = 0,0

clock = pygame.time.Clock()
while 1:
    clock.tick(40)
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            sys.exit()

        if event.type == pygame.KEYDOWN:
            if event.key == pygame.K_RIGHT:
                x += 10
            if event.key == pygame.K_LEFT:
                x -= 10
            if event.key == pygame.K_DOWN:
                y += 10
            if event.key == pygame.K_UP:
                y -= 10

        if event.type == pygame.MOUSEMOTION:
            mousePos = pygame.mouse.get_pos()
            x, y = mousePos
  
    screen.fill( (255,255,255) )

    
    




    screen.blit(hello, (x,y))


    pygame.display.update()


