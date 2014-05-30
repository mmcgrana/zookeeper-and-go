FROM ubuntu
MAINTAINER Mark McGranaghan "mmcgrana@gmail.com"

RUN apt-get update
RUN apt-get install -y openjdk-7-jdk wget

Run wget -q -O /opt/zookeeper-3.4.6.tar.gz http://apache.mirrors.pair.com/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz 
RUN tar -xz -f /opt/zookeeper-3.4.6.tar.gz -C /opt
RUN rm /opt/zookeeper-3.4.6.tar.gz

ADD run.sh /opt/zookeeper-3.4.6/run.sh
RUN chmod +x /opt/zookeeper-3.4.6/run.sh

RUN mkdir /var/zookeeper
VOLUME ["/var/zookeeper"]

CMD ["/opt/zookeeper-3.4.6/run.sh"]
