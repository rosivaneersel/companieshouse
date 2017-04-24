#!/bin/bash

if [[ $# -eq 0 ]] ; then
    echo 'No arguments provided'
    exit 1
fi

git fetch $1
git merge $1/master