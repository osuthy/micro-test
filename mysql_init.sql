create database if not exists test_micro_test;

use test_micro_test;

create table if not exists users (
  id int not null auto_increment,
  name char(30),
  mail char(30),
  age int,
  primary key (id)
);
