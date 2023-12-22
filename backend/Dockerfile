# Use an appropriate base image for your Go application
# For example:
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app/backend

# Copy necessary files into the container
COPY . .

# Build your Go application
RUN go mod download
RUN go build -o backend ./cmd/backend

EXPOSE 8080/tcp

# Set the entry point for the container
CMD ["./backend"]
