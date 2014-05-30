FROM ubuntu
MAINTAINER Mark McGranaghan "mmcgrana@gmail.com"

RUN apt-get update
RUN apt-get install -y wget ca-certificates build-essential git mercurial bzr

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /usr/local/go/

RUN wget --no-verbose https://storage.googleapis.com/golang/go1.2.2.src.tar.gz
RUN tar -xz -f go1.2.2.src.tar.gz -C /usr/local 
RUN rm go1.2.2.src.tar.gz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1

RUN go get github.com/samuel/go-zookeeper/zk
RUN go get github.com/mmcgrana/zk

CMD ["/bin/bash"]
