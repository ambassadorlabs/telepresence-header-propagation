## Running locally

This project was built using Gradle 8.1.1 and OpenJDK 17. I can't guarantee compatibility with any other versions.
Happy to accept help on working on compatibility.

* `$ gradle assemble`
* `$ java -jar ./build/libs/finalupper.jar`

## Building container

* `$ docker buildx build -o type=docker --platform=linux/amd64 --tag thedevelopnik/tp-headers-operator-java:1.0 .`
