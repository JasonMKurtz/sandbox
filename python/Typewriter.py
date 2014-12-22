#!/usr/bin/python
import os, sys, time, string
class TypeWriter:
    def Write(self, line, speed):
        for a in line:
            sys.stdout.write(a)
            sys.stdout.flush()
            time.sleep(speed)
        sys.stdout.write("\n")

T = TypeWriter()
T.Write(string.join(sys.argv[1::]), 0.1)
