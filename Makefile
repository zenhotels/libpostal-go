all:
	cgogen postal.yml

clean:
	rm -f postal/cgo_helpers.go postal/cgo_helpers.h postal/doc.go postal/types.go
	rm -f postal/postal.go

test:
	cd postal && go build
	