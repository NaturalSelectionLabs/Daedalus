#!/bin/sh -l

yaml=$(/daedalus argocd)
echo yaml
echo "app=$yaml" >> $GITHUB