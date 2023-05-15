FROM --platform=linux/amd64 gradle:8.1.1-jdk17-alpine

ENV SERVER_PORT=8081

WORKDIR /usr/src/finalupper/

COPY ./ ./

RUN gradle assemble

EXPOSE ${SERVER_PORT}
ENTRYPOINT [ "java", "-jar", "./build/libs/finalupper.jar" ]
