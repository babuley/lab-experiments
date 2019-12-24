import uuid

class Item:

    def __init__(self, position, size, sprite):
        self.__position = position
        self.__size = size
        self.__sprite = sprite
        self.__id = uuid.uuid4()

    def setPosition(self, position):
        self.__position = position

    def getPosition(self):
        return self.__position

    def getSize(self):
        return self.__size

    def getSprite(self):
        return self.__sprite

    def intersects(self, other):
        pass

    def getId(self):
        return self.__id

    def setId(self, id):
        self.__id = id