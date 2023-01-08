#!/bin/sh
curl --data '{"id":"100", "seq":"2", "caller":"hoge", "callee":"fuga"}' -X POST localhost:8090/api
