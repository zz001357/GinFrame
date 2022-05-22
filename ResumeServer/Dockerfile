FROM alpine:last

MAINTAINER ZhangZe "407102799@qq.com"

# 在容器根目录 创建一个 GinFrame 目录和子目录ResumeServer
WORKDIR /GinFrame
WORKDIR /ResumeServer

#COPY <相对Dockerfile的文件路径> <文件放置位置>
COPY ./resume .

EXPOSE 8006

# 运行golang程序的命令
ENTRYPOINT ["/GinFrame/ResumeServer/resume"]

