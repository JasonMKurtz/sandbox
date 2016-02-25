#!/usr/bin/python
"""
Given a list of items and their price, print an alphabetical-sorted list of the items and total prices. 

Input: 
20
ONION 100
BANANA FRIES 12
ONION 100
MANGO 60
BANANA FRIES 12
MANGO 60
TOMATO 80
POTATO CHIPS 30
POTATO CHIPS 30
MANGO 60
POTATO CHIPS 30
TOMATO 80
APPLE JUICE 10
TOMATO 80
CANDY 5
POTATO CHIPS 30
BANANA FRIES 12
ONION 100
TOMATO 80
POTATO CHIPS 30

Output: 
APPLE JUICE 10
BANANA FRIES 36
CANDY 5
MANGO 180
ONION 300
POTATO CHIPS 150
TOMATO 320

"""
import operator
ic = int(raw_input())
stuff = dict()
for i in xrange(ic): 
    r = raw_input().split()
    p = r[len(r) - 1]
    r.pop()
    n = '-'.join(r)
    if n in stuff: 
        stuff[n] += int(p)
    else: 
        stuff[n] = int(p)
        
s = sorted(stuff.items(), key=operator.itemgetter(0))
for k, p in s: 
    print k.replace("-", " "), p
