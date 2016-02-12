#!/bin/bash

LISTER_PATH=$HOME"/Downloads/darwin-ec2-lister"

if [ -f "$LISTER_PATH" ]; then
  aws ec2 describe-instances | $LISTER_PATH;
else
  echo "Configuring the lister...";
  aws s3 cp s3://my-executables/darwin-ec2-lister $LISTER_PATH;
  chmod +x $LISTER_PATH;
  aws ec2 describe-instances | $LISTER_PATH;
fi
