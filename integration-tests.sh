#!/bin/bash
# Just something simple to show things work as advertised

# Search for hello
echo "Get hello"
curl -s http://localhost:8000/search/hello
echo -e "\n"

# Case sensitivity test
echo "Get Hello"
curl -s http://localhost:8000/search/hello
echo -e "\n"

# Search for nonexisting
echo "Get nonexisting"
curl -s http://localhost:8000/search/nonexisting
echo -e "\n"

# Update wordlist
response_code=$(curl -s -o /dev/null -i -w "%{http_code}" --header "Content-Type: application/json" \
  --request PUT \
  --data '{"word": ["put", "hello", "red", "white"]}' \
  http://localhost:8000/words)
echo "Received HTTP code (should be 200)"
echo $response_code
echo -e "\n"

# Top Five
echo "Top Five"
curl -s http://localhost:8000/words/top5
echo -e "\n"
