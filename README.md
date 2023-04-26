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



