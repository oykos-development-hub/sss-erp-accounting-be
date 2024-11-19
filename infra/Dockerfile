FROM golang:1.22 as builder
ARG GOPRIVATE
ARG CELERITAS_RAT

RUN git config --global url."https://nonblank:${CELERITAS_RAT}@${GOPRIVATE}/erp/celeritas".insteadOf "https://${GOPRIVATE}/erp/celeritas"

ENV CGO_ENABLED=0

RUN go env -w GOPRIVATE="${GOPRIVATE}" 
RUN go env -w GOINSECURE="${GOPRIVATE}"
RUN git config --global http.sslVerify false

COPY . /
WORKDIR /
RUN go build -o cmd/accounting .

FROM scratch
COPY --from=builder /cmd/accounting /accounting
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /
CMD ["./accounting"] 