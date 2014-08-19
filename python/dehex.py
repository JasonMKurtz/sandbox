#!/usr/bin/python
import sys
if len(sys.argv) > 0:
    print "%s" % (sys.argv[1].decode("hex"))
else:
    print "./dehex.py <hex string>"
