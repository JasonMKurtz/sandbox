#!/usr/bin/env python
import re
regex = "(?P<url>(?P<proto>[^:]+):\/\/?(?P<hostname>[^\/:]+)):(?P<port>[^\/]+)(?P<path>.+)"
string = 'https://nyc-wdevtool-01:8080/jenkins/job/ecommerce-platform-master-appc/1679/'

m = re.search(regex, string)
print string
print "URL: %s, protocol: %s, hostname: %s, port: %s, path: %s" % (m.group('url'), m.group('proto'), m.group('hostname'), m.group('port'), m.group('path'))
