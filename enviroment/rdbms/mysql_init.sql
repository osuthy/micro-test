create database if not exists test_micro_test;
use test_micro_test;

create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);

create database if not exists test_connection_repository;

create database if not exists test_connection;
use test_connection;

create table if not exists test (
  column1 char(30),
  column2 char(30),
  primary key (column1)
);
