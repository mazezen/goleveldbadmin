FROM alpine

RUN mkdir -p /goleveldbadmin

COPY ./goleveldbadmin /goleveldbadmin
COPY ./config.yaml /goleveldbadmin

# 设置工作目录
WORKDIR /goleveldbadmin

# 运行 Go 项目
CMD ["./goleveldbadmin"]
