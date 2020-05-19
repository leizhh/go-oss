#!/bin/bash

curl -v 127.0.0.1:30000/objects/test -XPUT -d"this is object test"

echo -n "this is object test" | openssl dgst -sha256 -binary | base64
curl -v 127.0.0.1:30000/objects/test -XPUT -d"this is object test" -H "Digest: SHA-256=Fbv+3RiNfmoPVV0NoaSgcn5lA+3PbAsaV5s2OE3c1/M="

curl 127.0.0.1:30000/objects/test
echo

echo -n "this is object test version 2" | openssl dgst -sha256 -binary | base64
curl -v 127.0.0.1:30000/objects/test -XPUT -d"this is object test version 2" -H "Digest: SHA-256=F0S3Z92DbOgD55sf233cpfA/AwMnBRX82h6fYt0AZfk="

curl 127.0.0.1:30000/objects/test
echo

curl 127.0.0.1:30000/objects/test
echo
curl 127.0.0.1:30000/locate/F0S3Z92DbOgD55sf233cpfA%2FAwMnBRX82h6fYt0AZfk=
echo
curl 127.0.0.1:30000/locate/Fbv+3RiNfmoPVV0NoaSgcn5lA+3PbAsaV5s2OE3c1%2FM=
echo
curl 127.0.0.1:30000/versions/test
echo
curl 127.0.0.1:30000/objects/test?version=1
echo
curl -v 127.0.0.1:30000/objects/test -XDELETE

curl -v 127.0.0.1:30000/objects/test
echo

curl 127.0.0.1:30000/versions/test
echo
curl 127.0.0.1:30000/objects/test?version=1
echo
curl 127.0.0.1:30000/objects/test?version=2
echo
