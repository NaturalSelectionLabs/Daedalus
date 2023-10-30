#!/bin/sh -l

yaml=$(/daedalus argocd)
echo $yaml
echo "app=$yaml" >> $GITHUB_OUTPUT
#echo $GITHUB_OUTPUT