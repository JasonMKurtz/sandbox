#!/usr/bin/env python
import re, sys
regex = "(?P<url>(?P<proto>https?)?:?\/?\/?(?P<hostname>[^\/:]+\.?))?:?(?P<port>[0-9]+)?(?P<path>[^?#]+)?#?(?P<anchor>[^\?]+)?\??(?P<params>[^\n]+)?"
#string = 'https://nyc-wdevtool-01:8080/jenkins/job/ecommerce-platform-master-appc/1679/'
string = ' '.join(sys.argv[1:])

m = re.search(regex, string)
print string
print m.groupdict()
