# syntax=docker/dockerfile:1

# build frontend
FROM node:24-alpine as frontend
WORKDIR /app
COPY web/package.json .
RUN npm install

COPY web .
RUN npm run build

# bulding backend
FROM golang:1.26-alpine AS backend
WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN go build -o server ./cmd

# final image
FROM scratch
WORKDIR /app
COPY --from=frontend /app/dist ./static
COPY --from=backend /app/server .

ENV STATIC_DIR="./static"
CMD ["./server"]