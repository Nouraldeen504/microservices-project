# go_service/Dockerfile
FROM golang:1.17

# Set the working directory
WORKDIR /app

# Copy the Go file
COPY main.go ./

# Expose the port
EXPOSE 8080

# Command to run the app
CMD ["go", "run", "main.go"]