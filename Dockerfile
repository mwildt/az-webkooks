FROM scratch

MAINTAINER mail@maltewildt.de

COPY main /main

EXPOSE 80

CMD ["/main"]
