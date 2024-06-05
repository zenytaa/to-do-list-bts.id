# to-do-list-bts.id

### How To Run

```bash
#move to directory
$ cd to-do-list-bts.id
#create .env
$ cp env.example .env
#prepare database
$ cd to-do-list-bts.id/sql
#go into postgres psql
$ sudo -u postgres psql
#run init.sql
\i init.sql

```