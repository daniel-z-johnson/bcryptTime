FROM golang:1.15.2 AS build
WORKDIR /src
COPY . .
RUN go build -o bbcrypt
FROM scratch
COPY --from=build /src/bbcrypt / 
CMD ["/bbcrypt"]
