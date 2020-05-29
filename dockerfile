FROM golang:latest

ENV APP_NAME "gotest"
ENV CODE_PATH "github.com/pravinhp/${APP_NAME}"

# Create work env
RUN mkdir -p  $GOPATH/src/$CODE_PATH/
ADD .  $GOPATH/src/$CODE_PATH/
WORKDIR $GOPATH/src/$CODE_PATH/

RUN chmod u+x $GOPATH/src/$CODE_PATH/crud/crud

CMD cd $GOPATH/src/$CODE_PATH/
CMD ./crud/crud