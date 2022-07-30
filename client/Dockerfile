FROM golang:1.18 AS build

COPY . /src
WORKDIR /src
RUN go get .
RUN CGO_ENABLED=0 go build -v

FROM scratch
COPY --from=build /src/client /app/client
WORKDIR /app
CMD [ "/app/client" ]
