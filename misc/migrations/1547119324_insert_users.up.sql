INSERT INTO users SELECT generate_series(1,100000) AS id, md5(random()::text) AS Name, md5(random()::text) AS Phone, md5(random()::text) AS Date, md5(random()::text) AS City, md5(random()::text) AS Country, md5(random()::text) AS Email, md5(random()::text) AS Coordinates;