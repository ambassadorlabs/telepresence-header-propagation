FROM --platform=linux/amd64 golang:1.20.4-alpine3.18

ENV PORT=8080

RUN apk add --no-cache libc6-compat

WORKDIR /usr/src/finalupper/

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY ./main.go ./main.go

RUN go build main.go

EXPOSE ${8080}
ENTRYPOINT [ "./main" ]
