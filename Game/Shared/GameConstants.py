import os

class GameConstants:    
    SCREEN_SIZE = [1200,800]
    TILE_SIZE = [42,42]

    SPRITE_TREE = os.path.join("Assets", "tree2.png")
    SPRITE_TARGET = os.path.join("Assets", "target.png")
    SPRITE_BOULDER = os.path.join("Assets", "boulder2.png")
    SPRITE_PLAYER = os.path.join("Assets", "player.png")
    SPRITE_GROUND = os.path.join("Assets", "ground.png")

    PLAYING_SCENE = 0
    GAMEOVER_SCENE = 1
    HIGHSCORE_SCENE = 2
    MENU_SCENE = 3