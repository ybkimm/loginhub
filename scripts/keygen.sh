#!/bin/sh
dd if=/dev/random bs=64 count=1 2> /dev/null | sha1sum | cut -b 1-40 | base64 | cut -b 1-48
