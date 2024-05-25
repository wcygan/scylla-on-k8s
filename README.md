Using [cqlsh]()

```
kubectl exec -it POD_NAME -- cqlsh 
```

Run some cql commands:

```sql
CREATE KEYSPACE mykeyspace WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
USE mykeyspace;
CREATE TABLE mytable (mykey text PRIMARY KEY, mycolumn text);
INSERT INTO mytable (mykey, mycolumn) VALUES ('mykey', 'myvalue');
SELECT * FROM mytable;
```

We will build a gRPC server that accepts "ping" messages, and simply writes the ping timestamp to a ScyllaDB table.

Upon receiving a ping, the server will generate a uuid and record the timestamp at which the ping was received.

## Ping Server

```bash
grpcurl -plaintext -proto proto/ping/v1/ping.proto  localhost:50051 ping.v1.PingService/Ping
```