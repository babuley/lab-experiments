import pygame
from Scenes.Scene import Scene
from Shared import *
from Shared import Player
from Shared import Boulder

import random


class PlayingGameScene(Scene, object):

    def __resetBlackBox(self):
        self.__blackBox = []

    def __init__(self, game):
        super(PlayingGameScene, self).__init__(game)
        self.__resetBlackBox()

    def render(self):
        super(PlayingGameScene, self).render()

        player = self.getGame().getPlayer()
        player.render()

        self.clearText()

        self.addText("Sokoban ", x=0,
                     y=GameConstants.SCREEN_SIZE[1] - 60, size=30)

        self.addText("Level: " + str(self.getGame().getLevel()),
                     x=0, y=GameConstants.SCREEN_SIZE[1] - 30, size=30)

    def isPositionOccupied(self, pos):
        def taken(w): return w.getPosition()[
            0] == pos[0] and w.getPosition()[1] == pos[1]
        occupied = len(list(filter(taken, self.getGame().getAssets()))) != 0
        return occupied or len(self.getBoulderAtPosition(pos)) != 0

    def isSceneCompleted(self):
        boulders = self.getGame().getBoulders()
        targets = self.getGame().getTargets()
        c = 0
        def matchItemToTarget(b, t): return b.getPosition()[0] == t.getPosition(
        )[0] and b.getPosition()[1] == t.getPosition()[1]
        for b in boulders:
            for t in targets:
                if matchItemToTarget(b, t):
                    c += 1
        return c == len(boulders)

    def getBoulderAtPosition(self, pos):
        def cond(b): return b.getPosition()[
            0] == pos[0] and b.getPosition()[1] == pos[1]
        return list(filter(cond, self.getGame().getBoulders()))

    def newPositionRight(pos):
        return (pos[0] + GameConstants.TILE_SIZE[0], pos[1])

    def newPositionLeft(pos):
        return (pos[0] - GameConstants.TILE_SIZE[0], pos[1])

    def newPositionUp(pos):
        return (pos[0], pos[1] - GameConstants.TILE_SIZE[0])

    def newPositionDown(pos):
        return (pos[0], pos[1] + GameConstants.TILE_SIZE[0])

    __directions = {"RIGHT":  newPositionRight, "LEFT": newPositionLeft,
                    "UP": newPositionUp, "DOWN": newPositionDown}

    def findBoulder(self, cand):
        return list(filter(lambda r: r.getId() == cand.getId(), self.getGame().getBoulders()))

    def moveBack(self):
        if len(self.__blackBox) > 0:
            o = self.__blackBox.pop()
            player = self.getGame().getPlayer()
            player.setPosition(o[0].getPosition())
            if o[1]:
                b = self.findBoulder(o[1])
                if len(b) != 0:
                    b[0].setPosition(o[1].getPosition())

    def moveTo(self, dir):
        player = self.getGame().getPlayer()
        c = player.getPosition()
        newPosition = self.__directions[dir](c)
        b = self.getBoulderAtPosition(newPosition)
        bb = Boulder(newPosition, None, None)
        p = Player(c, None, None)

        if len(b) != 0:
            bb = Boulder(newPosition, None, None)
            bb.setId(b[0].getId())
        self.__blackBox.append([p, bb])

        if len(b) != 0:
            boulderNewPosition = self.__directions[dir](newPosition)
            if not self.isPositionOccupied(boulderNewPosition):
                b[0].setPosition(boulderNewPosition)
        o = self.isPositionOccupied(newPosition)
        if not o:
            player.setPosition(newPosition)
            r = random.randint(0, 13)
            if r % 5 == 0:
                self.getGame().playSound()
        else:
            self.__blackBox.pop()

    def onGameComplete(self):
        self.__resetBlackBox()
        self.getGame().goToNextLevel()

    def handleEvents(self, events):
        super(PlayingGameScene, self).handleEvents(events)
        for event in events:
            if event.type == pygame.QUIT:
                exit()

            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_BACKSPACE:
                    self.moveBack()
                if event.key == pygame.K_RIGHT:
                    self.moveTo("RIGHT")
                if event.key == pygame.K_LEFT:
                    self.moveTo("LEFT")
                if event.key == pygame.K_UP:
                    self.moveTo("UP")
                if event.key == pygame.K_DOWN:
                    self.moveTo("DOWN")

                if self.isSceneCompleted():
                    self.onGameComplete()
