# melogale
Melogale is yet another database research project. 

This project is primarily focused on querying planning and how to convert a potentially complex
SQL query into an agnostic query plan that is unaware of the database completely. The query plan
is only made aware of data in the data base when the query plan is actually executed. Thus the plan
must be created in a pessimistic fashion that does not rely on items being present in the database.

An example of an agnostic plan is the follow explain output:

```
LVL  ACTN  NAME                      DESC                                                                             KEY
========================================================================================================================================================================================================================
[-1] NONE  query                     CREATE TABLE users (id bigint primary key, account_id bigint references accounts (account_id), email text unique, password text, first_name text, last_name text);
[00] GET   table header              table with name [users] must not exist                                           74057573657273
[02] ID    new object                new object id for TABLE - users                                                  730000057573657273
[02] SET   table header              create table header: users                                                       74057573657273
[03] SET   column header             create column header: users -> id                                                63ffffffffffffffff026964
[03] SET   column header             create column header: users -> account_id                                        63ffffffffffffffff0a6163636f756e745f6964
[00] GET   table header              table with name [accounts] must exist                                            74086163636f756e7473
[01] GET   column header             column with name [account_id] must exist on table [accounts]                     63ffffffffffffffff0a6163636f756e745f6964
[04] SET   foreign key constraint    create foreign key constraint header: fk_users_account_id_accounts_account_id    66ffffffffffffffff27666b5f75736572735f6163636f756e745f69645f6163636f756e74735f6163636f756e745f6964
[03] SET   column header             create column header: users -> email                                             63ffffffffffffffff05656d61696c
[04] SET   unique constraint         create unique constraint header: uq_users_email                                  78ffffffffffffffff0e75715f75736572735f656d61696c
[03] SET   column header             create column header: users -> password                                          63ffffffffffffffff0870617373776f7264
[03] SET   column header             create column header: users -> first_name                                        63ffffffffffffffff0a66697273745f6e616d65
[03] SET   column header             create column header: users -> last_name                                         63ffffffffffffffff096c6173745f6e616d65
```

The plan would be executed in the order here, and certain query plan items would cache data
for future query plan items. This means that each query plan node only depends on the previous
plan nodes. The only assumption that is made is the format the data is encoded in inside the
key-value store.


I do still need to figure out how to do more dynamic plans. The plan above to create a table is
relatively simple compared to some of the query plans that a SQL database might need to deal with.

For example; in a select query, the overall query plan would change significantly if the table being
queried has an index that would satisfy the filtered columns.
The query plan could change even further if that index happens to also store the resulting columns
needed for the result; the table would not even need to be scanned. An index only scan could satisfy
this query completely.


The above query plan could also be optimized a bit more in the future. The steps in the plan could
be sorted by LVL. This would move things like verifying that the foreign table and its columns
actually exist before we even start to create columns in the database. Allowing the query plan to
fail much sooner if there was a problem.

The optimal plan would be sorted like this:

```
LVL  ACTN  NAME                      DESC                                                                             KEY
========================================================================================================================================================================================================================
[-1] NONE  query                     CREATE TABLE users (id bigint primary key, account_id bigint references accounts (account_id), email text unique, password text, first_name text, last_name text);
[00] GET   table header              table with name [users] must not exist                                           74057573657273
[00] GET   table header              table with name [accounts] must exist                                            74086163636f756e7473
[01] GET   column header             column with name [account_id] must exist on table [accounts]                     63ffffffffffffffff0a6163636f756e745f6964
[02] ID    new object                new object id for TABLE - users                                                  730000057573657273
[02] SET   table header              create table header: users                                                       74057573657273
[03] SET   column header             create column header: users -> id                                                63ffffffffffffffff026964
[03] SET   column header             create column header: users -> account_id                                        63ffffffffffffffff0a6163636f756e745f6964
[03] SET   column header             create column header: users -> email                                             63ffffffffffffffff05656d61696c
[03] SET   column header             create column header: users -> password                                          63ffffffffffffffff0870617373776f7264
[03] SET   column header             create column header: users -> first_name                                        63ffffffffffffffff0a66697273745f6e616d65
[03] SET   column header             create column header: users -> last_name                                         63ffffffffffffffff096c6173745f6e616d65
[04] SET   foreign key constraint    create foreign key constraint header: fk_users_account_id_accounts_account_id    66ffffffffffffffff27666b5f75736572735f6163636f756e745f69645f6163636f756e74735f6163636f756e745f6964
[04] SET   unique constraint         create unique constraint header: uq_users_email                                  78ffffffffffffffff0e75715f75736572735f656d61696c
```



An example of an insert plan (WORK IN PROGRESS):
```
LVL  ACTN  NAME                      DESC                                                                             KEY
==============================================================================================================================================================
[-1] NONE  query                     INSERT INTO users (account_id, email, first_name, last_name) VALUES(123, 'me@me.com', 'Elliot', 'Courant');
[00] GET   table header              table with name [users] must exist                                               74057573657273
[01] GET   column header             column with name [account_id] must exist on table [users]                        63ffffffffffffffff0a6163636f756e745f6964
[01] GET   column header             column with name [email] must exist on table [users]                             63ffffffffffffffff05656d61696c
[01] GET   column header             column with name [first_name] must exist on table [users]                        63ffffffffffffffff0a66697273745f6e616d65
[01] GET   column header             column with name [last_name] must exist on table [users]                         63ffffffffffffffff096c6173745f6e616d65
[03] NONE  insert                    insert 1 record(s) into: users -> account_id, email, first_name, last_name
```