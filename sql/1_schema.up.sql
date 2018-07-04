CREATE TABLE item_status (
  id SERIAL NOT NULL PRIMARY KEY,
  label TEXT
);

CREATE TABLE item (
  id SERIAL NOT NULL PRIMARY KEY,
  content TEXT,
  status_id INTEGER NOT NULL REFERENCES item_status(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE item_history (
  item_id INTEGER NOT NULL REFERENCES item(id),
  status_id INTEGER NOT NULL REFERENCES item_status(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY(item_id, status_id)
);

CREATE TABLE tag (
  tag_hash bytea NOT NULL PRIMARY KEY,
  label TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE item_tag (
  item_id INTEGER NOT NULL REFERENCES item(id),
  tag_hash bytea NOT NULL REFERENCES tag(tag_hash),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY(item_id, tag_hash)
);
