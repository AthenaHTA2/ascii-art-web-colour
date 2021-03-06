# syntax=docker/dockerfile:1
FROM golang:1.17.5
RUN mkdir /colour
WORKDIR /colour
# Download necessary Go modules
COPY go.mod ./
# COPY go.sum ./
# download all packages in mod file
RUN go mod download
# upload the entire 'Helena_Showcase_24May2022' application
ADD . /colour
RUN go mod tidy
# log-in to Git before cloning the 'Helena_Showcase_24May2022' repository
# RUN git config --global user.email "sedmakh2@gmail.com"
# RUN git config --global user.name "AthenaHTA2"
# RUN git config --global --add url."git@learn.01founders.co:".insteadOf "https://git.learn.01founders.co/"
# RUN go get git.learn.01founders.co/AthenaHTA2/Helena_Showcase_24May2022
# RUN cd /audit && git clone https://learn.01founders.co/git/AthenaHTA2/Helena_Showcase_24May2022.git
# Dockerfile is in the 'Helena_Showcase_24May2022' folder
# RUN cd /colour && /Helena_Showcase_24May2022
RUN cd /colour
# Next build a static application binary named 'binarycolour'
RUN go build -o binarycolour
# The port that connects to docker daemon
LABEL version="1.0"
LABEL description="Ascii-art-colour showcase project by Helena Sedmak, London, May 2022"
# Tell Docker to execute the 'binarycolour' command when this image is used to start a container.
ENTRYPOINT [ "/colour/binarycolour" ]