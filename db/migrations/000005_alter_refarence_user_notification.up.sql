ALTER TABLE notifications
ADD CONSTRAINT fk_user_notification
FOREIGN KEY (user_id) REFERENCES users(user_id)
ON DELETE CASCADE;