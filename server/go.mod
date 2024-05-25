module github.com/wcygan/scylla-on-k8s/server

go 1.22.0

require (
	github.com/google/uuid v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/wcygan/scylla-on-k8s/generated/go v0.0.0
	google.golang.org/grpc v1.63.2
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/wcygan/scylla-on-k8s/generated/go => ../generated/go

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.13.0
