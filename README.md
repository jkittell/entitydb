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
BenchmarkNewEntity-16          	  568731	      1913 ns/op
BenchmarkEntityDB_Search-16    	     262	   4996220 ns/op
BenchmarkEntityDB_Insert-16    	     358	   3374957 ns/op
BenchmarkEntityDB_Lookup-16    	     217	   5523728 ns/op
BenchmarkEntityDB_Delete-16    	     142	   8357930 ns/op
```



