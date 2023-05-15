## Running locally

* `$ bundle install`
* `$ ruby uppercase.rb`

## Building container

* `$ docker buildx build -o type=docker --platform=linux/amd64 --tag thedevelopnik/tp-headers-instrumentation-ruby:1.0 .`
