CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS PAYMENTS (
  id uuid  DEFAULT uuid_generate_v4 (),
  status VARCHAR NOT NULL,
  amount decimal(2) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  PRIMARY KEY (id)
);