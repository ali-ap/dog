FROM golang:1.18.4-bullseye AS builder

LABEL maintainer="Ali Alipour"

WORKDIR /app

COPY ./*.go ./

COPY . .

RUN go build -o output/server cmd/main.go


FROM golang:1.18.4-bullseye AS final
WORKDIR /app
COPY --from=builder app/output/server .
COPY --from=builder app/configs/app_config.yml ./configs/app_config.yml
EXPOSE 80
EXPOSE 443
#RUN groupadd -r neusta && useradd --no-log-init -r -g neusta neusta
#RUN chown -R neusta:neusta  /app
#USER neusta
ENV GIN_MODE=release
CMD ["./server"]