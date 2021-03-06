DROP TABLE IF EXISTS logs;

CREATE TABLE logs
(
  id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
  sha VARCHAR(255),
  date_test DATETIME,
  name_group VARCHAR(255),
  name_test VARCHAR(255),
  result VARCHAR(255),
  bench_time DECIMAL(18,4),
  output TEXT
);
