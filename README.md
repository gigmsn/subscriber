# GIGMSN Subscriber [![Build Status](https://travis-ci.org/gigmsn/subscriber.svg?branch=master)](https://travis-ci.org/gigmsn/subscriber)

A gopher RabbitMQ consumer

### HOWTO:

1. Start consumer container for reading messages from queue:

    make consumer/up

2. Stop service and remove container:

    make consumer/stop
