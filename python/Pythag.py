#!/usr/bin/python 
import math
class Pythag: 
    def __init__ (self, width, height): 
        self.width = width
        self.height = height

    def getHyp(self): 
        return math.sqrt(self.width ** 2 + self.height ** 2)

    def doHyp(self, width, height): 
        return math.sqrt(width ** 2 + height ** 2)

        
P = Pythag(3, 4)
print P.getHyp()


 
