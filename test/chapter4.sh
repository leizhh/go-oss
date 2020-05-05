#!/bin/bash

echo -n "this object will have only 1 instance" | openssl dgst -sha256 -binary | base64

curl -v 127.0.0.1:30000/objects/test4_1 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=incorrecthash"

curl -v 127.0.0.1:30000/objects/-9-XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="

curl -v 127.0.0.1:30000/objects/test4_2 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="

ls -ltr ./storage/?/objects
echo
curl 127.0.0.1:30000/objects/test4_1
echo
curl 127.0.0.1:30000/objects/test4_2
echo
curl 127.0.0.1:30000/locate/aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY=
echo
curl -v 127.0.0.1:30000/versions/test4_1
echo
curl 127.0.0.1:30000/versions/test4_2