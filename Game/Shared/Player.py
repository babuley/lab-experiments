from Shared import Item
from Shared.GameConstants import *


class Player(Item):
    def __init__(self, position, sprite, game):
        self.__game = game

        super(Player, self).__init__(position, GameConstants.TILE_SIZE, sprite)

    def render(self):
        self.getGame().screen.blit(self.getSprite(), self.getPosition())

    def getGame(self):
        return self.__game

    def setPosition(self, pos):
        newPosition = pos
        size = self.getSize()
        
        if newPosition[0] + size[0] > GameConstants.SCREEN_SIZE[0]:
            newPosition[0] = GameConstants.SCREEN_SIZE[0] - size[0]

        super(Player, self).setPosition(newPosition)
        