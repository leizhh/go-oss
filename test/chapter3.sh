#!/bin/bash

curl -v 127.0.0.1:30000/objects/test3 -XPUT -d"this is object test3"

echo -n "this is object test3" | openssl dgst -sha256 -binary | base64
curl -v 127.0.0.1:30000/objects/test3 -XPUT -d"this is object test3" -H "Digest: SHA-256=GYqqAdFPt+CScnUDc0/Gcu3kwcWmOADKNYpiZtdbgsM="

curl 127.0.0.1:30000/objects/test3
echo

echo -n "this is object test3 version 2" | openssl dgst -sha256 -binary | base64
curl -v 127.0.0.1:30000/objects/test3 -XPUT -d"this is object test3 version 2" -H "Digest: SHA-256=cAPvsxZe1PR54zIESQy0BaxC1pYJIvaHSF3qEOZYYIo="

curl -v 127.0.0.1:30000/objects/test3 -XPUT -d"this is object test3 version 3" -H "Digest: SHA-256=v8EJIZMsSfWXGdrlV2dFe4wUkinaWN6f1ql6cOu1KWA="

curl 127.0.0.1:30000/objects/test3
echo

curl 127.0.0.1:30000/objects/test3
echo
curl 127.0.0.1:30000/locate/GYqqAdFPt+CScnUDc0%2FGcu3kwcWmOADKNYpiZtdbgsM=
echo
curl 127.0.0.1:30000/locate/cAPvsxZe1PR54zIESQy0BaxC1pYJIvaHSF3qEOZYYIo=
echo
curl 127.0.0.1:30000/versions/test3
echo
curl 127.0.0.1:30000/objects/test3?version=1
echo
curl -v 127.0.0.1:30000/objects/test3 -XDELETE

curl -v 127.0.0.1:30000/objects/test3
echo

curl 127.0.0.1:30000/versions/test3
echo
curl 127.0.0.1:30000/objects/test3?version=1
echo
curl 127.0.0.1:30000/objects/test3?version=2
echo
