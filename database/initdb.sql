

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
  PRIMARY KEY(customerid, id)
);

CREATE TABLE IF NOT EXISTS fsnode
(
  parent      BIGINT,
  inode       BIGINT,
  name        TEXT,
  attr        JSONB,
  fsdevice_id BIGINT      REFERENCES fsdevice(id)
  PRIMARY KEY (parent, inode)
);
