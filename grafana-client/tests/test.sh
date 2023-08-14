#!/bin/bash

### GET
curl -i -X GET -H "Authorization: Bearer glsa_zhl48pD4kccNbIHCia8TCHVllHpN36oc_dc5c0e32" 'http://10.10.10.85/api/annotations'

### POST
curl -i -X POST -H "Authorization: Bearer glsa_zhl48pD4kccNbIHCia8TCHVllHpN36oc_dc5c0e32" -H "Content-Type: application/json" -d '@annotation-sample.json' 'http://10.10.10.85/api/annotations'
