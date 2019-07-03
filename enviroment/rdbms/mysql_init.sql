drop database if exists test_micro_test;
create database if not exists test_micro_test;
use test_micro_test;

drop table if exists test;
create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);

drop table if exists record_completion_all_type;
create table if not exists record_completion_all_type (
  dummy char(30),

  tinyintc tinyint not null,
  smallintc smallint not null,
  intc int not null,

  datec date not null
);
-- auto_inc
-- nullable
-- default value
-- primary
-- uniq
-- compound primary key

drop database if exists test_connection_repository;
create database if not exists test_connection_repository;

drop database if exists test_connection;
create database if not exists test_connection;
use test_connection;

create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);
