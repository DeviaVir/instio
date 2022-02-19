# Base our image on an official, minimal image of our preferred golang
FROM golang:1.16

# Note: The default golang docker image, already has the GOPATH env variable set.
# GOPATH is located at /go
ENV GO_SRC $GOPATH/src
ENV INSTIO_GITHUB github.com/DeviaVir/instio
ENV INSTIO_ROOT $GO_SRC/$INSTIO_GITHUB

# Consider updating package in the future. For instance ca-certificates etc.
# RUN apt-get update -qq && apt-get install -y build-essential

# Make the instio root directory
RUN mkdir -p $INSTIO_ROOT

# All commands will be run inside of instio root
WORKDIR $INSTIO_ROOT

# Copy the instio source into instio root.
COPY . .

# the following runs the code inside of the $GO_SRC/$INSTIO_GITHUB directory
RUN go get -u $INSTIO_GITHUB...

# Define the scripts we want run once the container boots
# CMD [ "" ]