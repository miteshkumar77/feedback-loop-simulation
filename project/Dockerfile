FROM golang

ADD . /project/

WORKDIR /project

RUN make srv

ENTRYPOINT [ "./bin/server" ]

EXPOSE 5300