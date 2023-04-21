# Use the official Golang image to create a build
FROM golang:latest AS build

# Set the working directory to the project directory
WORKDIR /app

# Copy all the files except the settings/config.json file
COPY . .

RUN rm -rf settings/config.json

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a lightweight Alpine Linux image
FROM alpine:latest

# Copy the built Go binary from the previous stage
COPY --from=build /app/app .

# Set the binary as the entry point of the container
ENTRYPOINT ["./app"]
