from microbit import *
import random
import music

display.show(Image.HEART)

while True:
    if accelerometer.was_gesture('shake'):
        display.show(random.randint(1, 6))
        sleep(2000)
        display.show("Feed me Seemuir")
        sleep(2000)
        display.show("")
    if button_a.was_pressed():
        display.scroll(temperature())
    if button_b.was_pressed():
        for x in range(2):
            music.play(["C4:4", "D4", "E4", "C4"])
        for x in range(2):
            music.play(["E4:4", "F4", "G4:8"])      
