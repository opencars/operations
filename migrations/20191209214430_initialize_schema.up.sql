BEGIN;

CREATE TABLE IF NOT EXISTS resources(
  "id"            SERIAL      PRIMARY KEY,
  "uid"           VARCHAR(36) NOT NULL UNIQUE,
  "name"          TEXT        NOT NULL,
  "last_modified" TIMESTAMP   NOT NULL,
  "url"           TEXT        NOT NULL,
  "created_at"    TIMESTAMP   NOT NULL DEFAULT NOW()
);
CREATE INDEX resources_uid_idx ON resources("uid");

CREATE TABLE IF NOT EXISTS operations(
  "resource_id"  INT         NOT NULL REFERENCES resources("id"),
  "person"       TEXT        NOT NULL,
  "reg_address"  TEXT,
  "code"         SMALLINT    NOT NULL,
  "name"         TEXT        NOT NULL,
  "reg_date"     DATE        NOT NULL,
  "office_id"    INT         NOT NULL,
  "office_name"  TEXT        NOT NULL,
  "make"         TEXT        NOT NULL,
  "model"        TEXT        NOT NULL,
  "year"         SMALLINT    NOT NULL,
  "color"        TEXT        NOT NULL,
  "kind"         TEXT        NOT NULL,
  "body"         TEXT        NOT NULL,
  "purpose"      TEXT        NOT NULL,
  "fuel"         TEXT,
  "capacity"     INT,
  "own_weight"   REAL,
  "total_weight" REAL,
  "number"       TEXT        NOT NULL
);

CREATE INDEX operations_number_idx ON operations("number");

END;

