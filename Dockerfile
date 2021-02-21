FROM alpine:latest

LABEL maintainer="mail@task.media"

RUN echo "hi there release drafter" > /helloWorld

ENV SOME_BUG=FIX

