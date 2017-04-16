FROM scratch
MAINTAINER embano1@live.com
ADD jsonapi /
ENTRYPOINT ["/jsonapi"]
EXPOSE 8080

