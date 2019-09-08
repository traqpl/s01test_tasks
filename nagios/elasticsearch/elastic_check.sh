#!/bin/bash

FOUNDCOUNTER=`curl -s -X GET "localhost:9200/index/message/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query": {
        "match_phrase" : {
            "error" : {
                "query" : "Handbill not printed"
            }
        }
    }
}
' | grep value | cut -f2 -d: | cut -f1 -d,`
if [ $FOUNDCOUNTER == 3 ]; then
    echo "CRITICAL"
    exit 2
else
    echo "OK"
    exit 0
fi