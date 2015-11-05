#!/usr/bin/python
""" 
    Given: 
        5 3
        89 90 78 93 80
        90 91 85 88 86  
        91 92 83 89 90.5
    where 5 is the number of students, and 3 in the number of scores each student has, compute each student's average
"""
numStudents, numScores = [ int(i) for i in raw_input().split() ] # 5, 3  
scores = zip(*[ [ float(j) for j in raw_input().split() ] for i in xrange(numScores) ])
for st in scores: 
    total = 0 
    for s in st: 
        total += s 
    print total / numScores
