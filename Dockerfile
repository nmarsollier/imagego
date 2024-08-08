# Docker para desarrollo
FROM golang:1.22.6-bullseye

WORKDIR /go/src/github.com/nmarsollier/imagego

ENV REDIS_URL=host.docker.internal:6379
ENV RABBIT_URL=amqp://host.docker.internal
ENV AUTH_SERVICE_URL=http://host.docker.internal:3000

# Puerto de Image Service y debug
EXPOSE 3001

CMD ["go" , "run" , "/go/src/github.com/nmarsollier/imagego"]