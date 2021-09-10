#!/usr/bin/env python3

import sys
import json

with open(sys.argv[1]) as f:
  data = json.load(f)

if "feVersion" in data.keys():
    print("flogo_fe")
else:
    print("flogo_oss")