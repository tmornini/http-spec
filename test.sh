#!/bin/bash

go install .

PREFIX=https://api.subledger.com

http-spec -prefix $PREFIX example-file-not-found.htsf          && exit 1
http-spec -prefix $PREFIX example-request-only.htsf            && exit 1
http-spec -prefix $PREFIX example-post.htsf                    || exit 1
http-spec -prefix $PREFIX example-regexp-and-substitution.htsf || exit 1

echo Tests ran successfully!
