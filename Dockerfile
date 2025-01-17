FROM golang:1.16-buster as build
WORKDIR /docs-cpi-pandemia
COPY go.* ./
COPY main.go .
COPY bar/ ./bar/
COPY cache/ ./cache/
COPY cli/ ./cli/
COPY downloader/ ./downloader/
COPY filesystem/ ./filesystem/
COPY text/ ./text/
COPY zip/ ./zip/
RUN go get && go build -o /usr/bin/docs-cpi-pandemia

FROM debian:buster-slim
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=build /usr/bin/docs-cpi-pandemia /usr/bin/docs-cpi-pandemia
CMD ["docs-cpi-pandemia"]
