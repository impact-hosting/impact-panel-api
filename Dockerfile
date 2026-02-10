FROM --platform=Linux/amd64 golang:1.21.6

RUN apt update
RUN apt install mariadb-client -y

RUN mkdir /impact
RUN mkdir /scripts
RUN mkdir /var/log/impact

# Copy the app files to the container
COPY ./src /impact
COPY ./scripts /scripts

WORKDIR /impact

CMD [ "/scripts/entrypoint.sh" ]