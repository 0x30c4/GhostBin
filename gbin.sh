#!/bin/bash

# Usage: ./ghostbin.sh [-f filename] [-e expire_seconds] [-r max_reads] [-d deepurl_length] [-s secret]

# Default values
EXPIRE="18446744073709551615"  # No expiration by default
READS="0"                      # No read limit by default
DEEPURL="0"                    # Default URL length
SECRET=""                      # No secret by default
FILENAME="-"                   # Default to pipe input

# Function to display usage
usage() {
  echo "Usage: $0 [-f filename] [-e expire_seconds] [-r max_reads] [-d deepurl_length] [-s secret]"
  echo "  -f: The filename to upload. Use '-' to pipe output of a command."
  echo "  -e: Expire time in seconds (default: no expiration)."
  echo "  -r: Maximum number of reads (default: unlimited)."
  echo "  -d: Length of the URL (default: random URL length)."
  echo "  -s: Secret for deletion (default: none)."
  exit 1
}

# Parse command-line arguments
while getopts "f:e:r:d:s:" opt; do
  case ${opt} in
    f ) FILENAME=$OPTARG ;;
    e ) EXPIRE=$OPTARG ;;
    r ) READS=$OPTARG ;;
    d ) DEEPURL=$OPTARG ;;
    s ) SECRET=$OPTARG ;;
    * ) usage ;;
  esac
done

# Check if filename is provided or input is piped
if [ -z "$FILENAME" ] && [ -t 0 ]; then
  usage
fi

# Build the curl command
CURL_CMD="curl -F \"f=@${FILENAME}\""
[ "$EXPIRE" != "18446744073709551615" ] && CURL_CMD+=" -F \"expire=${EXPIRE}\""
[ "$READS" != "0" ] && CURL_CMD+=" -F \"read=${READS}\""
[ "$DEEPURL" != "0" ] && CURL_CMD+=" -F \"deepurl=${DEEPURL}\""
[ -n "$SECRET" ] && CURL_CMD+=" -F \"secret=${SECRET}\""

# Append the GhostBin URL
CURL_CMD+=" gbin.me"

# Execute the curl command
if [ "$FILENAME" == "-" ]; then
  # Handle piped input
  cat | eval $CURL_CMD
else
  # Handle file input
  eval $CURL_CMD
fi
