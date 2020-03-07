# Dockerfile

FROM golang@sha256:1ff752199f17b70e5f4dc2ad7f3e7843c456eb7e1407ed158ed8c237dbf1476a as builder
ENV USER=xouser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY ./bin .
WORKDIR ./bin
#FROM scratch


#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/group /etc/group
#USER xouser:xouser

ENTRYPOINT [ "./xo" ]
