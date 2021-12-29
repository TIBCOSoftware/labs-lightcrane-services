#!/usr/bin/env python3

import sys
import io
import json
from urllib.parse import urlparse

raw_data = sys.stdin.read()

end_points = {}
for line in raw_data.splitlines():
    if 'GET -' in line:
        url = urlparse(line[line.index('https'):])
        print('scheme  :', url.scheme)
        print('netloc  :', url.netloc)
        print('path    :', url.path)
        print('params  :', url.params)
        print('query   :', url.query)
        print('fragment:', url.fragment)
        print('username:', url.username)
        print('password:', url.password)
        print('hostname:', url.hostname, '(netloc in lower case)')
        print('port    :', url.port)
        end_points['URL'] = '{}://{}'.format(url.scheme, url.netloc)
        end_points['Properties'] = {}
        end_points['Properties']['URI'] = {}
        end_points['Properties']['URI']['Service'] = url.path

if 'URL' in end_points :
    f = open("endpoint.json", "w")
    f.write(json.dumps(end_points))
    f.close()