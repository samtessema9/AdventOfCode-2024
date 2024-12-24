#!/bin/bash

# Ensure we have the correct number of arguments
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <day> <output-file-path>"
    exit 1
fi

# Read CLI arguments
day=$1
output_file=$2

# Ensure 'day' is an integer
if ! [[ "$day" =~ ^[0-9]+$ ]]; then
    echo "Error: <day> must be an integer."
    exit 1
fi

# Locate the .env file
env_file="./.env"

if [ ! -f "$env_file" ]; then
    echo "Error: .env file not found in the parent directory."
    exit 1
fi

# Read the session ID from the .env file
session_id=$(grep '^SESSION_ID=' "$env_file" | cut -d '=' -f2)
echo "session_id: $session_id"

if [ -z "$session_id" ]; then
    echo "Error: SESSION_ID not found in .env file."
    exit 1
fi

# Make the request to fetch the input data and write it to the specified output file
url="https://adventofcode.com/2024/day/$day/input"
response=$(curl -o $output_file -w "%{http_code}" -H "Cookie: session=$session_id" $url)
http_code=${response: -3}

# Check if the request was successful
if [ "$http_code" -ne 200 ]; then
    echo "Error: Failed to fetch input data. HTTP code $http_code."
    rm -f $output_file
    exit 1
fi

echo "Input data successfully written to $output_file."
