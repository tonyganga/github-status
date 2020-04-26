# github-status

This application is something I'm using to wrap my head around golang. This simply sets up a few routes and tells us if Github is available or not.

### Building the Application

You can build the application quickly by running the following.

```
make build
```

Yep. That's it. You should see the following output:

```
docker build -t github-status:latest .
Sending build context to Docker daemon  243.7kB
Step 1/16 : FROM golang:alpine as builder
 ---> ccda0e5ccbfc
Step 2/16 : RUN apk update && apk add git && apk add ca-certificates
 ---> Using cache
 ---> 4cbc84ad9250
Step 3/16 : RUN adduser -D -g '' githubstatususer
 ---> Using cache
 ---> 285441a29844
Step 4/16 : COPY . $GOPATH/src/github-status/
 ---> adb249c25fdc
Step 5/16 : WORKDIR $GOPATH/src/github-status/
 ---> Running in 84146b64ddf8
Removing intermediate container 84146b64ddf8
 ---> f4292452a3df
Step 6/16 : ENV CGO_ENABLED 0
 ---> Running in 018449773e71
Removing intermediate container 018449773e71
 ---> bbea7ada25cd
Step 7/16 : ENV GOOS linux
 ---> Running in fbbd1e57cc75
Removing intermediate container fbbd1e57cc75
 ---> a83d7acda0d9
Step 8/16 : ENV GOARCH amd64
 ---> Running in 1c3a143aaa5f
Removing intermediate container 1c3a143aaa5f
 ---> e6e89fb9e93f
Step 9/16 : RUN go test -v
 ---> Running in 41404e0b10cf
go: downloading github.com/gorilla/mux v1.6.2
go: downloading github.com/stretchr/testify v1.2.2
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/davecgh/go-spew v1.1.1
=== RUN   TestHealthCheckHandler
--- PASS: TestHealthCheckHandler (0.00s)
=== RUN   TestGithubStatusHandler
2020/04/26 17:05:07 GET /api/last-message.json
2020/04/26 17:05:07 Closing body of request
2020/04/26 17:05:07 Unmarshaling last-message.json
2020/04/26 17:05:07 Github returned an outage state.
--- PASS: TestGithubStatusHandler (1.19s)
PASS
ok      github.com/tonyganga/github-status      1.190s
Removing intermediate container 41404e0b10cf
 ---> 23973d40ac28
Step 10/16 : RUN go build -a -installsuffix cgo -o /go/bin/github-status
 ---> Running in e939335544c0
Removing intermediate container e939335544c0
 ---> 9d894ed491a6
Step 11/16 : FROM scratch
 ---> 
Step 12/16 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> 97433d6c14b4
Step 13/16 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> 11be22532be1
Step 14/16 : COPY --from=builder /go/bin/github-status /go/bin/github-status
 ---> Using cache
 ---> f9727be7f6da
Step 15/16 : USER githubstatususer
 ---> Using cache
 ---> 817e4c2fec22
Step 16/16 : ENTRYPOINT ["/go/bin/github-status"]
 ---> Using cache
 ---> d0d1a1012c02
Successfully built d0d1a1012c02
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
