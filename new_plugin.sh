#!/bin/bash
# Get PLUGIN_NAME from command line arguments
PLUGIN_NAME=$1

# Check if PLUGIN_NAME is provided
if [ -z "$PLUGIN_NAME" ]; then
  echo "Please provide PLUGIN_NAME as an argument"
  exit 1
fi
# Copy the contents of the 'boilerplate' directory to the 'PLUGIN_NAME' directory
cp -r boilerplate/ protoc-gen-$PLUGIN_NAME

# Replace the string 'boilerplate' with '$PLUGIN_NAME' in the Makefile
if [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' "s/boilerplate/$PLUGIN_NAME/" protoc-gen-$PLUGIN_NAME/Makefile
else
  sed -i "s/boilerplate/$PLUGIN_NAME/" protoc-gen-$PLUGIN_NAME/Makefile
fi
