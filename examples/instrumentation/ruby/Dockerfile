FROM --platform=linux/amd64 ruby:3.2.2-alpine3.17

ENV PORT=4567
ENV GEM_HOME="/usr/local/bundle"
ENV PATH $GEM_HOME/bin:$GEM_HOME/gems/bin:$PATH

RUN apk add --no-cache ruby-bundler make g++ bash

WORKDIR /usr/src/uppercase/

COPY ./ ./

RUN bundle install

EXPOSE ${PORT}
ENTRYPOINT [ "bash", "./start.sh" ]
