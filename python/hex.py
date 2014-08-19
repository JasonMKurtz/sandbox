#!/usr/bin/python
import sys
if len(sys.argv) > 0:
    print "%s" % (sys.argv[1].encode("hex"))
else:
    print "./hex.py <string>"
