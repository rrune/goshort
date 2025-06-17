#!/bin/bash

# Simple "script" to convert the README.md file in the root to the landing page of the website
# uses pandoc because I like pandoc and it works

echo -e "$(cat components/header.html) \n $(pandoc ../README.md) \n $(cat components/footer.html)" > public/index.html