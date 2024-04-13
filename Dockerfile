# Use the official RabbitMQ image as the base image
FROM rabbitmq:latest

# Optionally, you can customize the RabbitMQ configuration by providing a custom configuration file
# COPY rabbitmq.conf /etc/rabbitmq/rabbitmq.conf

# Expose the default RabbitMQ port
EXPOSE 5672

# Optionally, expose the RabbitMQ management plugin port
EXPOSE 15672

# You can customize the RabbitMQ environment variables if needed
# ENV RABBITMQ_CONFIG_FILE=/etc/rabbitmq/rabbitmq.conf

# Start RabbitMQ server
CMD ["rabbitmq-server"]
