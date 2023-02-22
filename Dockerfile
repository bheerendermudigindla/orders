#Building the binary.
FROM amd64/golang as build

ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOOS=linux   \
    GOPATH= \
    GOARCH="amd64" \
    TZ=Asia/Kolkata
    
###
RUN mkdir -p /usr/local/go/src/orders

COPY go.mod .
COPY go.sum .
 
#Download Dependency
RUN go mod download
 
COPY . /usr/local/go/src/orders
WORKDIR /usr/local/go/src/orders
 
#Running the go service
RUN go build -o main main.go

RUN cp -r /usr/local/go/src/orders/main /main

RUN mkdir /deployer/
RUN cp -r /usr/local/go/src/webapp-eventdatabuilder/deployer/git_details.txt /deployer/git_details.txt
 
CMD ["main"]