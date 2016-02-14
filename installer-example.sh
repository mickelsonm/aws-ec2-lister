#!/bin/bash

LISTER_PATH=$HOME"/Downloads/darwin-ec2-lister"

# handle optional arguments
for arg in "$@"; do
  case $arg in
    "--latest" )
    echo "Grabbing the latest of the lister";
    if [ -f "$LISTER_PATH" ]; then
      rm $LISTER_PATH;
    fi
  esac
done

# we can use the lister if it already is where it should be
if [ -f "$LISTER_PATH" ]; then
  aws ec2 describe-instances | $LISTER_PATH;
else # we configure it
  echo "Configuring the lister...";
  aws s3 cp s3://my-executables/darwin-ec2-lister $LISTER_PATH;
  chmod +x $LISTER_PATH;
  aws ec2 describe-instances | $LISTER_PATH;
fi
