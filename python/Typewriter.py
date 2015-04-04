#!/usr/bin/python
import os, sys, time, string
class TypeWriter:
    def Write(self, line, delay = 0.9):
        for a in line:
            sys.stdout.write(a)
            sys.stdout.flush()
            time.sleep(delay)
        sys.stdout.write("\n")

T = TypeWriter()
T.Write(string.join(sys.argv[1::]))
