## Queueing-API

This API is responsible for receiving and processing messages from RabbitMQ. It dials RabbitMQ through AMPQ and listens to messages. Once a message is received, it is sent to the subscriber. For example:

Customer Service here subscribes to receive any messages that are sent to a email_confirmation queue. A message arrives, queueing-api will send it to the customer service which will handle it accordingly. 

