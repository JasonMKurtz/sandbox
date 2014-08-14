#!/usr/bin/python
import os, sys, time
class TypeWriter:
    def Write(self, line, speed):
        for a in line:
            sys.stdout.write(a)
            sys.stdout.flush()
            time.sleep(speed)
        sys.stdout.write("\n")

T = TypeWriter()
T.Write("I'm a typewriter.", 0.3)
