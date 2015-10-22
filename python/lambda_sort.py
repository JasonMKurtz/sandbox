#!/usr/bin/env python2.7
""" 
    Requires python2.7 
    Example: 

    $ ./lambda_sort.py
    3 digits: 384
    2 digits: 48
    1 digits: 6
""" 

import operator
i = 1000
input = [ x for x in xrange(0, i) ]
n = filter(lambda x: (x % 5) != 0 and (x % 4) != 0 and len(set(str(x))) == len(str(x)), input)
count = { i: 0 for i in set([ len(str(x)) for x in n ]) }
for i in n: count[len(str(i))] += 1
                                             
for i in sorted(count.items(), key=operator.itemgetter(1))[::-1]: 
    print "%d digits: %d" % (i[0], i[1])
