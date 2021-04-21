FROM ubuntu:latest
RUN mkdir -p /netcat_test
RUN apt-get update && apt-get install netcat -y
COPY . .
CMD echo -n 'Ping' | nc server 12345 -w5