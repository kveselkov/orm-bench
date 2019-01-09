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

INSERT INTO users SELECT generate_series(1,100000) AS id, md5(random()::text) AS Name, md5(random()::text) AS Phone, md5(random()::text) AS Date, md5(random()::text) AS City, md5(random()::text) AS Country, md5(random()::text) AS Email, md5(random()::text) AS Coordinates;