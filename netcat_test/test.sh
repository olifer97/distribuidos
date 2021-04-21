#!/bin/bash
docker build -f ./Dockerfile -t netcat_test:latest .
docker run --network distribuidos_testing_net netcat_test:latest > netcat_test.out
cmp netcat_test.out response_expected.in && echo "✅ Send correctly" || echo "❌ Send Error"
rm netcat_test.out