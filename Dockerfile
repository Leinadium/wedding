# syntax=docker/dockerfile:1

# bulding backend
FROM golang:1.26-alpine AS backend
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal cmd ./
RUN go build -o server ./cmd

# final image
FROM scratch
WORKDIR /app
COPY --from=backend /app/server .
CMD ["./server"]