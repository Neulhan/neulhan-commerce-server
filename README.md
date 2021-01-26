# Neulhan Commerce Server with Go
![](static/ONAIR.png)
elastic-beanstalk docker postgresql golang gorm gin

- set Rolling deploy
- set Logging Middleware
- set SSL https connection
- create unit test
- set Validator for each model

## Run Project Development
```bash
source alias.sh
up
down
```

## Deploy Project
```bash
eb deploy
```

## Tidy packages
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
