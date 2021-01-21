# Neulhan Commerce Server with Go
docker-compose lambda postgresql golang gorm gin

- ~~set HotReload feature~~
- set Logging Middleware
- ~~set Go Server to run in Docker Container~~

## Run Project
```
docker-compose build
docker-compose up
```

## Import, UnImport packages
```bash
go mod tidy
```

## Verify packages
```bash
go mod verify
```

## Connect To Postgresql 
```
docker exec -it nc_postgres /bin/bash
...
root@49d68bd0cacd:/# psql -U neulhan NC
psql (13.1 (Debian 13.1-1.pgdg100+1))
Type "help" for help.
...
NC=# \c
You are now connected to database "NC" as user "neulhan".
```

```
NC=# \dt
          List of relations
 Schema |   Name    | Type  |  Owner  
--------+-----------+-------+---------
 public | customers | table | neulhan
 public | orders    | table | neulhan
 public | products  | table | neulhan
(3 rows)
...

NC=# SELECT * FROM customers LIMIT 1;
 id |          created_at           |          updated_at           | deleted_at |  name | email | pass | logged_in 
----+-------------------------------+-------------------------------+------------+-------+-------+------+-----------
  1 | 2021-01-19 00:14:24.580872+09 | 2021-01-19 00:14:24.580872+09 |            |       |       |      | t
(1 row)
...
```