FROM golang:alpine as builder
# install git, ca-certificates 
RUN apk update && apk add git && apk add ca-certificates
# create githubstatususer to run the binary
RUN adduser -D -g '' githubstatususer
# copy source
COPY . $GOPATH/src/github-status/
WORKDIR $GOPATH/src/github-status/
# set ENV to build binary for scratch
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64 
# test and build
RUN go test -v
RUN go build -a -installsuffix cgo -o /go/bin/github-status

# STEP 2 build a small image
FROM scratch
# copy certs, user and binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/github-status /go/bin/github-status
# define user
USER githubstatususer
# run binary
ENTRYPOINT ["/go/bin/github-status"]
