# Dockerfile

FROM golang:latest as builder

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

COPY /bin /app

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app /app
USER xouser:xouser

ENTRYPOINT [ "./app/xo" ]
