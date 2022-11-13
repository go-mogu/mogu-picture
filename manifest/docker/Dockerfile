FROM alpine:3.16

###############################################################################
#                                INSTALLATION
###############################################################################
ENV APP_NAME                                mogu-picture
ENV WORKDIR                                 /app
ADD resource                                $WORKDIR/
ADD ./$APP_NAME   $WORKDIR/$APP_NAME


RUN chmod +x $WORKDIR/$APP_NAME

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./$APP_NAME
