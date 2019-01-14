DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
  id SERIAL PRIMARY KEY,
  Name varchar(255) default NULL,
  Phone varchar(100) default NULL,
  Date varchar(255),
  City varchar(255),
  Country varchar(100) default NULL,
  Email varchar(255) default NULL,
  Coordinates varchar(40) default NULL
);