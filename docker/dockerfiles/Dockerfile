#指定编译环境为golang
FROM golang:latest As build
#维护者名称
MAINTAINER welcome112@qq.com
#指定代码目录
ENV ROOT=/welcom112s-go
#拷贝当前目录所有文件到$ROOT/testrunner目录下
COPY ../.. $ROOT/testrunner
#切换到$ROOT/gogostudy 目录
WORKDIR $ROOT/testrunner
#构建产物到系统的/bin/testrunner下
RUN  go build -o /bin/testrunner
#分层构建
FROM ubuntu:latest
#拷贝上一层的产物到当前层
COPY --from=build /bin/testrunner  /bin/testrunner
#拷贝执行脚本到docker目录
COPY docker-entrypoint.sh /app/
#切换工作目录
WORKDIR /app/
#赋予脚本执行权限
RUN chmod -R 777 docker-entrypoint.sh
#指定容器执行指令
ENTRYPOINT ["./docker-entrypoint.sh"]
#为脚本指定参数
#指定执行命令
CMD ["release"]
