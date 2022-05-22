#使用多阶段构建，解决可能因为系统架构不一，造成的手工打包不兼容的问题

#第一阶段构建golang的应用程序
FROM golang:latest as build-env
MAINTAINER ZhangZe "407102799@qq.com"

ENV APP_PATH=$GOPATH/GinFrame
# 在容器根目录 操作
WORKDIR $APP_PATH
#COPY <相对Dockerfile的文件路径 电脑位置> <docker位置 文件放置位置>
COPY ResumeServer ./ResumeServer
COPY proto ./proto
COPY go.sum .
COPY go.mod .

#设置go的一些常用环境
RUN go env -w GOPROXY=https://goproxy.cn,direct  \
    && go env -w GO111MODULE=on \
    && go mod download \
    && cd ResumeServer \
    && go build
RUN cd ResumeServer && ls

#第二阶段构建运行
FROM scratch
WORKDIR /app/
COPY --from=0 $GOPATH/GinFrame/ResumeServer /resumeServer
EXPOSE 8006
# 运行golang程序的命令
ENTRYPOINT ./resumeServer






#FROM alpine:latest
##alpine 补充基础包
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#
#MAINTAINER ZhangZe "407102799@qq.com"
#
## 在容器根目录 创建一个 GinFrame 目录和子目录ResumeServer
#WORKDIR resume_app
#
#
##COPY <相对Dockerfile的文件路径 电脑位置> <docker位置 文件放置位置>
#COPY ./resume/resume ./
#COPY ./config_resume.yaml ./
#
#EXPOSE 8006
#
## 运行golang程序的命令
#ENTRYPOINT /resume_app/resume
