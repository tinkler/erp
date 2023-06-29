INSERT INTO erpv1.users(username, phone_number)
	VALUES ('admin', '18176386025');
INSERT INTO erpv1.emails (user_id, address) SELECT u.ID, 'tinkler@163.com' FROM erpv1.users AS u WHERE username = 'admin';