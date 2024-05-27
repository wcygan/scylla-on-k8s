## Quickstart

Start the application

```bash
skaffold dev
```

Ping the server a few times to create some data

```
grpcurl -plaintext -proto proto/ping/v1/ping.proto  localhost:50051 ping.v1.PingService/Ping
```

## Scylla

You can check the data that was created using [cqlsh](https://opensource.docs.scylladb.com/stable/cql/cqlsh.html)

```
kubectl exec -it $(kubectl get pods -l app=scylla-db -o jsonpath='{.items[0].metadata.name}') -- cqlsh
```

Run some cql commands:

```cql
USE events;
DESCRIBE TABLES;
select * from pings;
```

## Port Forwarding

Sometimes, if a port is not available, kubectl will choose a different port for the service. Keep an eye on the skaffold logs for this:

```
Port forwarding statefulset/scylla-db in namespace default, remote port 9042 -> http://127.0.0.1:9042
Port forwarding deployment/server in namespace default, remote port 50051 -> http://127.0.0.1:50052
```

Now if I wanted to send a gRPC request, it would be to `localhost:50052` instead of `localhost:50051`.

```
grpcurl -plaintext -proto proto/ping/v1/ping.proto  localhost:50052 ping.v1.PingService/Ping
```