FROM golang:1.20.13-bookworm

RUN apt update
RUN apt install -y libdlib-dev libopenblas-dev
RUN apt install -y libjpeg62-turbo-dev


# https://askubuntu.com/questions/623578/installing-blas-and-lapack-packages
RUN apt install -y libatlas-base-dev
