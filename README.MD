# github-status

This application is something I'm using to wrap my head around golang. This simply sets up a few routes and tells us if Github is available or not.

### Building the Application

You can build the application quickly by running the following.

```
make build
```

Yep. That's it. You should see the following output:

```
Sending build context to Docker daemon  522.2kB
Step 1/13 : FROM golang:alpine as builder
 ---> 95ec94706ff6
Step 2/13 : RUN apk update && apk add git && apk add ca-certificates
 ---> Using cache
 ---> 8be54cc12270
Step 3/13 : RUN adduser -D -g '' githubstatususer
 ---> Using cache
 ---> ecb16db01859
Step 4/13 : COPY . $GOPATH/src/tonyganga/github-status/
 ---> Using cache
 ---> c27a6fcd9ebf
Step 5/13 : WORKDIR $GOPATH/src/tonyganga/github-status/
 ---> Using cache
 ---> e8e95fcf766b
Step 6/13 : RUN go get -d -v
 ---> Using cache
 ---> 9b64b7691ba0
Step 7/13 : RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/github-status
 ---> Using cache
 ---> acedf86099f0
Step 8/13 : FROM scratch
 ---> 
Step 9/13 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> 39406aaaeefe
Step 10/13 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> d3744320a95d
Step 11/13 : COPY --from=builder /go/bin/github-status /go/bin/github-status
 ---> Using cache
 ---> 2a62773107d6
Step 12/13 : USER githubstatususer
 ---> Using cache
 ---> 35c3a1bd0157
Step 13/13 : ENTRYPOINT ["/go/bin/github-status"]
 ---> Using cache
 ---> c9f99f149d01
Successfully built c9f99f149d01
Successfully tagged github-status:latest

```

You can then view the image by running the following.

```
docker images

REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
github-status       latest              c9f99f149d01        2 minutes ago       7.48MB
<none>              <none>              acedf86099f0        2 minutes ago       356MB

```

You can then run the application by doing the following.

```
docker run -d -p 8000:8000 github-status 
```

You should then be able to access the api, try a simple curl.

```
$ curl 172.17.0.2:8000/status
* TCP_NODELAY set
* Connected to 172.17.0.2 (172.17.0.2) port 8000 (#0)
> GET /status HTTP/1.1
> Host: 172.17.0.2:8000
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Thu, 18 Oct 2018 02:58:54 GMT
< Content-Length: 110
< 
* Connection #0 to host 172.17.0.2 left intact

{"status":"good","created_on":"2018-10-08T22:39:11Z","body":"Everything operating normally.","Available":true}%  

```
