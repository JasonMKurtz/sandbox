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

    def wholeHyp(self): 
        hyps = [ ]
        for i in range(1, 100): 
            for j in range(2, 100):
                if self.doHyp(i, j).is_integer() and i < j: 
                    hyps.append([i, j, self.doHyp(i, j)])

        for h in hyps: 
            print "%f, %f = %f" % h[0], h[1], h[2]

        
P = Pythag(3, 4)
print P.getHyp()
print P.wholeHyp()


 
