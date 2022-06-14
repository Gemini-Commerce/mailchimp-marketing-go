CUR_UID := $(shell id -u)
CUR_GID := $(shell id -g)

generate:
	docker run --rm -it --user ${CUR_UID}:${CUR_GID} -e GOPATH=$(go env GOPATH):/go -e DEBUG=true -v `pwd`:/go/src/client quay.io/goswagger/swagger \
	 generate client --skip-validation --with-expand -f /go/src/client/swagger.json -A mailchimp-marketing --default-scheme https -t /go/src/client