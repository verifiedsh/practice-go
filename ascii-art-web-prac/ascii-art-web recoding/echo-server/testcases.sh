#!bin/bash
curl -i http://localhost:8080/
curl -i http://localhost:8080/echo -d text="hello world"
curl -i http://localhost:8080/echo -d text=""
curl -i http://localhost:8080/wrong