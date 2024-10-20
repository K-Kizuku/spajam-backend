ALTER TABLE buildings
ADD CONSTRAINT fk_user_building
FOREIGN KEY (building_id) REFERENCES users(user_id)
ON DELETE CASCADE;