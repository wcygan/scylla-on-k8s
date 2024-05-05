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