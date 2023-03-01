# STEP 1 build a base image
FROM golang:1.19.3-alpine AS builder

### create appuser ###
RUN adduser -D -g '' elf

# create workspace
WORKDIR /opt/app/
COPY go.mod go.sum ./

# fetch dependancies
RUN go mod download && go mod verify

# copy the source code
COPY . .

### build binary ###
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/user-service ./cmd


# STEP 2 build a small image
FROM alpine:3.17.0
LABEL language="golang"

# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd

COPY . .

### copy the static executable ###

#to have permissions to execute the binary.
COPY --from=builder --chown=elf:1000 /go/bin/user-service /user-service

# use a non-root user
USER elf

# run app
ENTRYPOINT ["./user-service"]
