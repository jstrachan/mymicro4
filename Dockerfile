FROM alpine:3.2
ADD mymicro4-srv /mymicro4-srv
ENTRYPOINT [ "/mymicro4-srv" ]
