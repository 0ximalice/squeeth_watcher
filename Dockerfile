# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Golang application source code to the container
COPY ./aggregator /app

# Build the Golang application
RUN go build -o aggregator .

# Switch to a smaller base image for the final image
FROM python:3.9-slim

# Set the working directory for the final image
WORKDIR /app

# Install supervisord
RUN apt-get update && apt-get install -y supervisor

# Copy the built Golang binary from the previous stage
COPY --from=0 /app/aggregator /app/aggregator

# Install Streamlit using pip
RUN pip install streamlit

# Copy the Streamlit app code to the container
COPY ./simulator /app

# Copy the supervisord configuration file
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Expose the port on which Streamlit will run
EXPOSE 8501

# Command to run both Golang app and Streamlit app
CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
