

-- DROP DATABASE IF EXISTS fsdisk;

-- CREATE DATABASE fsdisk;

-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS fsdevice
(
  id                BIGINT,
  customerid        BIGINT,
  name              TEXT,
  blocksize         INT DEFAULT 4096,
  sizegb            BIGINT,
  created           TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY(customerid, id)
);

CREATE TABLE IF NOT EXISTS fsnode
(
  parent      BIGINT,
  inode       BIGINT,
  name        TEXT,
  attr        JSONB,
  created     TIMESTAMP,
  fsdevice_id BIGINT,
  PRIMARY KEY (parent, inode)
);

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO ubuntu;
