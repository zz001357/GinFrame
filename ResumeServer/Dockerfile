FROM golang:alpine3.15

MAINTAINER zhangze

RUN mkdir "GinFrmae"

WORKDIR /GinFrame

#添加go.mod go.sum
ADD . go.mod
ADD . go.sum

#运行go mod
#RUN go mod download

RUN mkdir "ResumeServer"
WORKDIR /ResumeServer

COPY . .

EXPOSE 8006

ENTRYPOINT ["./server.go"]



