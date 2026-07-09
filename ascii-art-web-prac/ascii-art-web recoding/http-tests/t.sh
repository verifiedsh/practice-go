#!bin/bash
curl -i http://localhost:8080/ok
curl -i http://localhost:8080/badrequest
curl -i http://localhost:8080/notfound
curl -i http://localhost:8080/error
curl -i http://localhost:8080/