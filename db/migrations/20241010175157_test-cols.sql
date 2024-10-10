-- migrate:up
CREATE TABLE test_cols (
  id SERIAL PRIMARY KEY,
  data JSONB
);	

-- migrate:down
DROP TABLE test_cols;

