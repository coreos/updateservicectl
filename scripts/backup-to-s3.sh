#!/bin/bash -e

bucket=$1
if [ -z $bucket ]; then
    echo "usage: $0 <bucket-name>"
    exit 1
fi

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
tmpfile=`mktemp -t "roller-backup"`
$dir/../bin/updatectl database backup $tmpfile

resource="/${bucket}/backup.tar.gz"
contentType="application/x-compressed-tar"
dateValue=`date -ju +"%a, %d %b %Y %H:%M:%S GMT"`
stringToSign="PUT\n\n${contentType}\n${dateValue}\n${resource}"

signature=`echo -en ${stringToSign} | openssl sha1 -hmac ${S3_SECRET_KEY} -binary | base64`

curl --upload-file ${tmpfile} \
    -H "Date: ${dateValue}" \
    -H "Content-type: ${contentType}" \
    -H "Authorization: AWS ${S3_ACCESS_KEY}:${signature}" \
    http://s3.amazonaws.com${resource}
