BEGIN;

CREATE TABLE system (
  name  VARCHAR(80) PRIMARY KEY NOT NULL ,
  value text NOT NULL
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name varchar(255),
  email varchar(255) NOT NULL
);

CREATE TABLE devices (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  user_id INT NOT NULL,

  CONSTRAINT devices_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE device_metrics (
  id SERIAL PRIMARY KEY,
  device_id INT NOT NULL,
  metric_1 INT,
  metric_2 INT,
  metric_3 INT,
  metric_4 INT,
  metric_5 INT,
  local_time TIMESTAMP,
  server_time TIMESTAMP DEFAULT NOW(),

  CONSTRAINT device_metrics_device_id_fk FOREIGN KEY (device_id) REFERENCES devices (id) ON DELETE CASCADE
);

CREATE TABLE device_alerts (
  id SERIAL PRIMARY KEY,
  device_id INT,
  message TEXT
);

insert into users (id, name, email) values (1, 'Test User', 'mcdia@bk.ru');
insert into devices (id, name, user_id) values (1, 'Test Device', 1);

END;