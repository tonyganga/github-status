# github-status

This application is something I'm using to wrap my head around golang. This simply sets up a few routes and tells us if Github is available or not.

### Building the Application

You can build the application quickly by running the following.

```
docker build -t github-status:latest .

```

Yep. That's it. You should see the following output:

```
$ docker build -t github-status:latest .

Sending build context to Docker daemon  518.7kB
Step 1/10 : FROM golang:1.11
1.11: Pulling from library/golang
bc9ab73e5b14: Pull complete 
193a6306c92a: Pull complete 
e5c3f8c317dc: Pull complete 
a587a86c9dcb: Pull complete 
1bc310ac474b: Pull complete 
997731689cfb: Pull complete 
202760eb4a00: Pull complete 
Digest: sha256:63ec0e29aeba39c0fe2fc6551c9ca7fa16ddf95394d77ccee75bc7062526a96c
Status: Downloaded newer image for golang:1.11
 ---> 45e48f60e268
Step 2/10 : ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
Downloading [==================================================>]  8.499MB/8.499MB
 ---> fd2fbfa4a481
Step 3/10 : RUN chmod +x /usr/bin/dep
 ---> Running in 1eae5467fcb1
Removing intermediate container 1eae5467fcb1
 ---> 70bd22184ea3
Step 4/10 : WORKDIR $GOPATH/src/github.com/tonyganga/github-status
 ---> Running in a18467cbd9ab
Removing intermediate container a18467cbd9ab
 ---> dfda15f5d3c1
Step 5/10 : COPY Gopkg.toml Gopkg.lock ./
 ---> 034effa8b0d3
Step 6/10 : RUN dep ensure --vendor-only
 ---> Running in 25d3f9d28646
Removing intermediate container 25d3f9d28646
 ---> f84cfc7d1dc9
Step 7/10 : COPY . ./
 ---> 52ecb6f11816
Step 8/10 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .
 ---> Running in f035c14d94c4
Removing intermediate container f035c14d94c4
 ---> cbda58700fae
Step 9/10 : EXPOSE 8000
 ---> Running in 4597eb6034ce
Removing intermediate container 4597eb6034ce
 ---> b7177c5a5938
Step 10/10 : ENTRYPOINT ["/app"]
 ---> Running in 3046cfb84eca
Removing intermediate container 3046cfb84eca
 ---> 5b90316e89dd
Successfully built 5b90316e89dd
Successfully tagged github-status:latest

```

You can then run the application by doing the following.

```
docker run -d github-status 
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
