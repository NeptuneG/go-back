ALTER TABLE user_points ADD tx_id varchar;
CREATE UNIQUE INDEX ON user_points(tx_id);
