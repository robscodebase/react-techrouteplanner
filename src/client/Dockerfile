FROM node:9.8
RUN apt-get update
WORKDIR /client
COPY ./entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
