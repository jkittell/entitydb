# entitydb

## Store unstructured data in postgres

### Create the table in the postgres database

```
CREATE TABLE entity ( 
    id SERIAL PRIMARY KEY, 
    name TEXT,
    description TEXT,
    properties JSONB );
```

### Create new entitydb instance

```
db := NewEntityDB(host, port, user, password, dbname)
```

### Functions

```
Search
Insert
Lookup
Delete

```

### Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/jkittell/entitydb
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNewEntity-16           	  602362	      1909 ns/op
BenchmarkEntityDB_Search-16     	      39	  29223563 ns/op
BenchmarkEntityDB_Search2-16    	      40	  30150220 ns/op
BenchmarkEntityDB_Insert-16     	     274	   3681259 ns/op
BenchmarkEntityDB_Lookup-16     	     223	   5354542 ns/op
BenchmarkEntityDB_Delete-16     	     139	   8285716 ns/op
```



