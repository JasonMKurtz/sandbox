#!/usr/bin/env python
import re, sys
regex = "(?P<url>:?(?P<proto>https?)?:?\/?\/?(?P<hostname>[^\/]+))\/?:(?P<port>[0-9]+)?(?P<path>\/.+)?"
#string = 'https://nyc-wdevtool-01:8080/jenkins/job/ecommerce-platform-master-appc/1679/'
string = ' '.join(sys.argv[1:])

m = re.search(regex, string)
print string
print "URL: %s, protocol: %s, hostname: %s, port: %s, path: %s" % (m.group('url'), m.group('proto'), m.group('hostname'), "80" if m.group('port') is None else m.group('port'), m.group('path'))
