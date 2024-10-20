CREATE TABLE IF NOT EXISTS  buildings (
  building_id   text PRIMARY KEY,
  zip integer      NOT NULL,
  building_type text      NOT NULL,
  old_level integer      NOT NULL
);

