#Stage 1
#Builder stage
#Imports the full go compiler and builds the image
FROM golang:1.26 AS first-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web-app

#Stage 2
#Final stage
#Starts with a lightweight image (gcr.io/distroless/base-debian11) and keeps just the executable file
FROM gcr.io/distroless/base-debian11

COPY --from=first-stage /web-app /web-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/web-app"]
