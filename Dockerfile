################
# BUILD BINARY #
################

FROM golang:alpine3.17 as builder

WORKDIR /app

COPY . .

RUN apk add --no-cache nodejs npm

RUN npm install --prefix web

RUN npm run --prefix web build

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" .

#####################
# MAKE SMALL BINARY #
#####################
FROM scratch

# Copy the executable.
WORKDIR /app

COPY --from=builder /app/theitem /app
COPY --from=builder /app/config.prod.json /app

# ENTRYPOINT ["theitem", "appitem"]