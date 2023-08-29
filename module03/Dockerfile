FROM alpine

WORKDIR /app

ENV TZ=Asia/Shanghai

RUN apk add tzdata

COPY bin/* /usr/bin/

ENTRYPOINT ["server"]
