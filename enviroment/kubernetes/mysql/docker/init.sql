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
  mediumintc mediumint not null,
  intc int not null,
  bigintc bigint not null,

  floatc float(1, 1) not null,
  doublec double(1, 1) not null,
  decimalc decimal(1, 1) not null,

  charc char(1) not null,
  varcharc varchar(1) not null,

  tinytextc tinytext not null,
  textc text not null,
  mediumtextc mediumtext not null,
  longtextc longtext not null,

  datec date not null,
  datetimec datetime(0) not null,
  timestampc timestamp(1) not null,
  timec time(6) not null,
  yearc year not null,

  bitc bit(1) not null,
  binaryc binary(255) not null,
  varbinaryc varbinary(255) not null,
  tinyblobc tinyblob not null,
  blobc blob not null,
  mediumblobc mediumblob not null,
  longblobc longblob not null
);
-- auto_inc
-- nullable
-- default value
-- primary
-- uniq
-- compound primary key

drop database if exists test_db;
create database if not exists test_db;
use test_db;

create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);

drop database if exists test_connection;
create database if not exists test_connection;
use test_connection;

create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);
