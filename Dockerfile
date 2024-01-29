FROM --platform=Linux/amd64 golang:1.21.6

RUN mkdir /impact
RUN mkdir /scripts
RUN mkdir /var/log/impact

# Copy the app files to the container
COPY go.sum /impact
COPY go.mod /impact
COPY main.go /impact
COPY ./src /impact/src
COPY ./scripts /scripts

WORKDIR /impact

CMD [ "/scripts/entrypoint.sh" ]