DELETE FROM erpv1.emails WHERE user_id = (SELECT id FROM erpv1.users WHERE username='admin');
DELETE FROM erpv1.users WHERE `username`='admin';