FROM ghcr.io/betagorobot/betagomermaidserver-golang:master as builder

COPY betago.zip /data/

WORKDIR /data/

RUN unzip betago.zip

RUN CGO_ENABLED=0 go build -mod vendor -ldflags="-w -s" -o betaGo-source ./*.go &&\
    upx -9 -o betaGo betaGo-source


FROM ghcr.io/betagorobot/betagomermaidserver-chrome:master as runner
# FROM scratch as runner


COPY --from=builder  /data/betaGo /betaGo

WORKDIR /

CMD ["./betaGo"]
