FROM alpine:latest

LABEL maintainer="mail@task.media"

RUN echo "hi new feature" > /helloWorld

ENV SOME_BUG=FIX

