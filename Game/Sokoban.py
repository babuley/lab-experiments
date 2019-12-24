from Shared import GameConstants
from Shared import Wall, Target, Player, Boulder, Floor
from loader import *
from Scenes import *
import pygame, sys
import os
import random

class Sokoban:

    def __initSets(self):
        self.__assets = []
        self.__targets = []
        self.__boulders = []
        self.__floor = []
        self.__player = None

    def __initSound(self):
        self.__sounds = []
        self.__currentSound = None
        for file in os.listdir(os.path.join("Assets","Sound")):                  
            if file.endswith(".wav"):
                self.__sounds.append(pygame.mixer.Sound(os.path.join("Assets","Sound",file)))

    def __initWalls(self):
        self.__wallTilesCookieCutter = []
        for w in os.listdir(os.path.join("Assets","Walls")):                  
            if w.endswith(".png"):
                self.__wallTilesCookieCutter.append(pygame.image.load(os.path.join("Assets","Walls",w)))

    def __initBoulders(self):
        self.__itemsCookieCutter = []
        for it in os.listdir(os.path.join("Assets","Item")):                  
            if it.endswith(".png"):
                self.__itemsCookieCutter.append(pygame.image.load(os.path.join("Assets","Item",it)))                


    def __init__(self):
        pygame.init()
        pygame.mixer.init()
        pygame.display.set_caption("Sokoban")

        self.__clock = pygame.time.Clock()
        self.screen = pygame.display.set_mode(GameConstants.SCREEN_SIZE, pygame.DOUBLEBUF, 32)
        
        self.__maps = loadLevels()
        self.__currentLevel = 0
        self.map = self.__maps[self.__currentLevel]

        self.__initSets()
        self.__scenes = (
            PlayingGameScene(self),         
            MenuScene(self)
        )

        self.__currentScene = 0        
        self.__initSound()  
        self.__initWalls()
        self.__initBoulders()

        pygame.mouse.set_visible(0)


    def __pickItemFrom(self, bag):
        if len(bag) == 0:
            return None
        r = random.randint(0, len(bag) - 1)
        return bag[r]

    def __loadAssets(self):              
        s = 42
        i, j = 0, 0        
        wallTile = self.__pickItemFrom(self.__wallTilesCookieCutter)
        itemTile = self.__pickItemFrom(self.__itemsCookieCutter)
        for m in self.map:                        
            i+=1
            j = 0
            for n in m:
                j+=1
                if n == 'X':
                    t = Wall((s * j, s * i), wallTile, self)
                    self.__assets.append(t)                    
                if n == '.':                     
                    t = Target((s * j, s * i), pygame.image.load(GameConstants.SPRITE_TARGET), self)
                    self.__targets.append(t)                    
                if n == '@':                       
                    t = Player((s * j, s * i), pygame.image.load(GameConstants.SPRITE_PLAYER), self)
                    self.__player = t
                if n == '*':                                      
                    t = Boulder((s * j, s * i), itemTile, self)
                    self.__boulders.append(t)    
                if n != '\n':
                    t = Floor((s * j, s * i), pygame.image.load(GameConstants.SPRITE_GROUND), self)
                    self.__floor.append(t)   


    def start(self):
        self.__loadAssets()   

        while 1:
            self.__clock.tick(100)
            self.screen.fill((0, 0, 0))           

            toRender = [self.__floor, self.__assets, self.__targets, self.__boulders]

            for set in toRender:
                for a in set:
                    a.render()     

            currentScene = self.__scenes[self.__currentScene]
            currentScene.handleEvents(pygame.event.get())
            currentScene.render()         
            pygame.display.update()

    def getPlayer(self):
        return self.__player

    def getAssets(self):
        return self.__assets

    def getBoulders(self):
        return self.__boulders        

    def getTargets(self):
        return self.__targets    

    def goToNextLevel(self):
        if self.__currentLevel < len(self.__maps):
            self.__currentLevel += 1
            self.map = self.__maps[self.__currentLevel]
            self.__initSets()
            self.start()

    def changeScene(self, scene):
        pass

    def getLevel(self):
        return self.__currentLevel

    def playSound(self, soundClip = None):     
        if len(self.__sounds) == 0:
            return                      
        r = random.randint(0, len(self.__sounds) - 1)
        if self.__currentSound:
            self.__currentSound.stop()
        self.__currentSound = self.__sounds[r]        
        self.__currentSound.play()

    def reset(self):
        self.__currentLevel = 0

if __name__ == "__main__":
	Sokoban().start()