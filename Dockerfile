FROM jcorioland/azure-terratest:0.12.3

RUN mkdir /go/src/ossparis-demo
COPY . /go/src/ossparis-demo
WORKDIR /go/src/ossparis-demo

RUN chmod +x run-tests.sh

ENTRYPOINT [ "./run-tests.sh" ]