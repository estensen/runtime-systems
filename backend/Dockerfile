FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/estensen/runtime-systems/backend
WORKDIR /go/src/github.com/estensen/runtime-systems/backend

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
        then \
        app; \
        else \
        # Hot reload code when it changes
        go get github.com/pilu/fresh && \
        fresh; \
        fi

EXPOSE 8080
