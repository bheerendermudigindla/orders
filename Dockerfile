#Building the binary.
FROM amd64/golang as build

ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOOS=linux   \
    GOPATH= \
    GOARCH="amd64" \
    TZ=Asia/Kolkata
    
###
RUN mkdir -p /usr/local/go/src/sellerapp

COPY go.mod .
COPY go.sum .
 
#Download Dependency
RUN go mod download
 
COPY . /usr/local/go/src/sellerapp
WORKDIR /usr/local/go/src/sellerapp
 
#Running the go service
RUN go build -o sellerapp_order sellerapp_order.go

RUN cp -r /usr/local/go/src/sellerapp/sellerapp_order /sellerapp_order

RUN mkdir /deployer/
RUN cp -r /usr/local/go/src/webapp-eventdatabuilder/deployer/git_details.txt /deployer/git_details.txt
 
CMD ["sellerapp_order"]