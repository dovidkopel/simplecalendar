FROM frolvlad/alpine-python2

RUN pip install aws-sam-cli
RUN mkdir /srv

WORKDIR /srv

EXPOSE 3000
# nohup /usr/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 &