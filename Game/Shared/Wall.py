from Shared import Item
from Shared.GameConstants import *

class Wall(Item, object):
    def __init__(self, position, sprite, game):
        self.__game = game

        super(Wall, self).__init__(position, GameConstants.TILE_SIZE, sprite)

    def render(self):
        self.getGame().screen.blit(self.getSprite(), self.getPosition())

    def getGame(self):
        return self.__game
