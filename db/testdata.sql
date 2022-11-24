DO $$
DECLARE
  user_id  INTEGER := 12345;
BEGIN

insert into public.user_spendings (user_id, category, amount, date)
values
(user_id, 'Еда', 10, '2022-05-02'),
(user_id, 'Лекарства', 20, '2022-05-21'),
(user_id, 'Еда', 20, '2022-06-21'),
(user_id, 'Транспорт', 20, '2022-08-11'),
(user_id, 'Хозтовары', 20, '2022-08-15'),
(user_id, 'Развлечения', 20, '2022-09-01'),
(user_id, 'Хозтовары', 20, '2022-09-02'),
(user_id, 'Еда', 20, '2022-09-11'),
(user_id, 'Транспорт', 20, '2022-09-19'),
(user_id, 'Хозтовары', 20, '2022-09-24'),
(user_id, 'Еда', 20, '2022-09-28'),
(user_id, 'Развлечения', 20, '2022-10-01'),
(user_id, 'Бытовая техника', 20, '2022-10-02'),
(user_id, 'Лекарства', 20, '2022-10-05'),
(user_id, 'Развлечения', 20, '2022-10-10'),
(user_id, 'Еда', 20, '2022-10-11'),
(user_id, 'Услуги', 20, '2022-10-12'),
(user_id, 'Транспорт', 20, '2022-10-18'),
(user_id, 'Развлечения', 20, '2022-10-19'),
(user_id, 'Еда', 20, '2022-10-20'),
(user_id, 'Электроника', 1234, '2022-10-20'),
(user_id, 'Услуги', 456, '2022-10-21'),
(user_id, 'Развлечения', 100, '2022-10-22');


-------------------------------------------------------------------------------------------------


insert into public.currency_rates (currency, timestamp, rate)
values
('CNY', '2022-10-18 17:05:14.676684', 8.4591),
('EUR', '2022-10-18 17:05:14.774525', 60.625999),
('USD', '2022-10-18 17:05:14.939161', 60.875),
('USD', '2022-10-18 17:06:14.613743', 60.875),
('CNY', '2022-10-18 17:06:14.710270', 8.4591),
('EUR', '2022-10-18 17:06:14.870589', 60.625999),
('USD', '2022-10-18 17:07:14.539522', 60.875),
('EUR', '2022-10-18 17:07:14.636414', 60.625999),
('CNY', '2022-10-18 17:07:14.736490', 8.4591),
('USD', '2022-10-18 17:08:14.595390', 60.875),
('EUR', '2022-10-18 17:08:14.806353', 60.625999),
('CNY', '2022-10-18 17:08:14.912332', 8.4591),
('USD', '2022-10-18 17:09:14.534527', 60.875),
('CNY', '2022-10-18 17:09:14.630859', 8.4591),
('EUR', '2022-10-18 17:09:14.731248', 60.625999),
('USD', '2022-10-18 17:10:14.547602', 60.875),
('CNY', '2022-10-18 17:10:14.666055', 8.4591),
('EUR', '2022-10-18 17:10:14.822736', 60.625999),
('USD', '2022-10-18 17:11:14.621875', 60.875),
('EUR', '2022-10-18 17:11:14.737637', 60.625999),
('CNY', '2022-10-18 17:11:14.848950', 8.4591),
('CNY', '2022-10-18 17:12:14.542922', 8.4591),
('EUR', '2022-10-18 17:12:14.643048', 60.625999),
('USD', '2022-10-18 17:12:14.749886', 60.875),
('CNY', '2022-10-18 17:12:56.853153', 8.4591),
('EUR', '2022-10-18 17:12:56.954186', 60.625999),
('USD', '2022-10-18 17:12:57.109880', 60.875),
('EUR', '2022-10-18 17:13:56.714203', 60.625999),
('USD', '2022-10-18 17:13:56.818520', 60.875),
('CNY', '2022-10-18 17:13:56.915827', 8.4591),
('CNY', '2022-10-18 17:14:42.431681', 8.4591),
('EUR', '2022-10-18 17:14:42.538507', 60.625999),
('USD', '2022-10-18 17:14:42.694295', 60.875),
('CNY', '2022-10-18 17:22:39.054778', 8.4591),
('EUR', '2022-10-18 17:22:39.165048', 60.625999),
('USD', '2022-10-18 17:22:39.264278', 60.875),
('CNY', '2022-10-18 17:23:38.875871', 8.4591),
('EUR', '2022-10-18 17:23:38.975529', 60.625999),
('USD', '2022-10-18 17:23:39.077890', 60.875),
('EUR', '2022-10-18 17:24:38.896135', 60.625999),
('USD', '2022-10-18 17:24:38.984218', 60.875),
('CNY', '2022-10-18 17:24:39.086646', 8.4591),
('EUR', '2022-10-18 17:25:38.863770', 60.625999),
('USD', '2022-10-18 17:25:38.978601', 60.875),
('CNY', '2022-10-18 17:25:39.079746', 8.4591),
('USD', '2022-10-18 17:26:38.923893', 60.875),
('CNY', '2022-10-18 17:26:39.024413', 8.4591),
('EUR', '2022-10-18 17:26:39.122212', 60.625999),
('EUR', '2022-10-18 17:27:38.868505', 60.625999),
('USD', '2022-10-18 17:27:39.021419', 60.875),
('CNY', '2022-10-18 17:27:39.118078', 8.4591),
('USD', '2022-10-18 17:28:38.908716', 60.875),
('CNY', '2022-10-18 17:28:39.013498', 8.4591),
('EUR', '2022-10-18 17:28:39.114224', 60.625999),
('EUR', '2022-10-18 17:29:38.859486', 60.625999),
('USD', '2022-10-18 17:29:38.958629', 60.875),
('CNY', '2022-10-18 17:29:39.069297', 8.4591),
('EUR', '2022-10-18 17:30:38.877282', 60.625999),
('USD', '2022-10-18 17:30:38.994022', 60.875),
('CNY', '2022-10-18 17:30:39.093743', 8.4591),
('USD', '2022-10-18 17:31:38.869980', 60.875),
('CNY', '2022-10-18 17:31:38.984704', 8.4591),
('EUR', '2022-10-18 17:31:39.102599', 60.625999),
('EUR', '2022-10-18 17:32:38.924503', 60.625999),
('USD', '2022-10-18 17:32:39.029961', 60.875),
('CNY', '2022-10-18 17:32:39.131973', 8.4591),
('EUR', '2022-10-18 17:33:38.874239', 60.625999),
('USD', '2022-10-18 17:33:38.989179', 60.875),
('CNY', '2022-10-18 17:33:39.090697', 8.4591),
('EUR', '2022-10-18 17:34:38.872932', 60.625999),
('USD', '2022-10-18 17:34:38.988642', 60.875),
('CNY', '2022-10-18 17:34:39.090229', 8.4591),
('USD', '2022-10-18 17:35:38.883524', 60.875),
('CNY', '2022-10-18 17:35:38.981845', 8.4591),
('EUR', '2022-10-18 17:35:39.138694', 60.625999),
('EUR', '2022-10-18 17:36:38.921523', 60.625999),
('USD', '2022-10-18 17:36:39.086680', 60.875),
('CNY', '2022-10-18 17:36:39.187072', 8.4591),
('USD', '2022-10-18 17:37:38.868227', 60.875),
('EUR', '2022-10-18 17:37:38.966955', 60.625999),
('CNY', '2022-10-18 17:37:39.068442', 8.4591),
('EUR', '2022-10-18 17:38:38.861025', 60.625999),
('USD', '2022-10-18 17:38:38.960297', 60.875),
('CNY', '2022-10-18 17:38:39.064442', 8.4591),
('EUR', '2022-10-18 17:39:38.863731', 60.625999),
('USD', '2022-10-18 17:39:38.963196', 60.875),
('CNY', '2022-10-18 17:39:39.074687', 8.4591),
('USD', '2022-10-18 17:40:38.886345', 60.875),
('CNY', '2022-10-18 17:40:38.992930', 8.4591),
('EUR', '2022-10-18 17:40:39.095004', 60.625999),
('EUR', '2022-10-18 17:41:38.925644', 60.625999),
('USD', '2022-10-18 17:41:39.029084', 60.875),
('CNY', '2022-10-18 17:41:39.182409', 8.4591),
('USD', '2022-10-18 17:42:38.874382', 60.875),
('EUR', '2022-10-18 17:42:38.975000', 60.625999),
('CNY', '2022-10-18 17:42:39.080254', 8.4591),
('USD', '2022-10-18 17:43:38.875052', 60.875),
('CNY', '2022-10-18 17:43:38.980841', 8.4591),
('EUR', '2022-10-18 17:43:39.082628', 60.625999),
('EUR', '2022-10-18 17:44:38.934768', 60.625999),
('USD', '2022-10-18 17:44:39.070014', 60.875),
('CNY', '2022-10-18 17:44:39.211673', 8.4591),
('EUR', '2022-10-18 17:45:38.896559', 60.625999),
('USD', '2022-10-18 17:45:39.071017', 60.875),
('CNY', '2022-10-18 17:45:39.179361', 8.4591),
('EUR', '2022-10-18 17:54:41.680828', 60.625999),
('USD', '2022-10-18 17:54:41.781148', 60.875),
('CNY', '2022-10-18 17:54:41.887489', 8.4591),
('CNY', '2022-10-18 17:55:41.477034', 8.4591),
('EUR', '2022-10-18 17:55:41.642724', 60.625999),
('USD', '2022-10-18 17:55:41.798982', 60.875),
('CNY', '2022-10-18 17:56:41.460757', 8.4591),
('EUR', '2022-10-18 17:56:41.562493', 60.625999),
('USD', '2022-10-18 17:56:41.662241', 60.875),
('CNY', '2022-10-18 17:57:41.443782', 8.4591),
('EUR', '2022-10-18 17:57:41.540224', 60.625999),
('USD', '2022-10-18 17:57:41.638609', 60.875),
('CNY', '2022-10-18 17:58:41.458554', 8.4591),
('USD', '2022-10-18 17:58:41.666292', 60.875),
('EUR', '2022-10-18 17:58:41.573558', 60.625999),
('CNY', '2022-10-18 17:59:41.448886', 8.4591),
('USD', '2022-10-18 17:59:41.547364', 60.875),
('EUR', '2022-10-18 17:59:41.649251', 60.625999),
('CNY', '2022-10-18 18:00:41.467648', 8.4591),
('EUR', '2022-10-18 18:00:41.577530', 60.625999),
('USD', '2022-10-18 18:00:41.679781', 60.875),
('CNY', '2022-10-18 18:01:41.468664', 8.4591),
('EUR', '2022-10-18 18:01:41.953383', 60.625999),
('USD', '2022-10-18 18:01:42.148875', 60.875),
('CNY', '2022-10-18 18:02:41.456218', 8.4591),
('EUR', '2022-10-18 18:02:41.557332', 60.625999),
('USD', '2022-10-18 18:02:41.657462', 60.875),
('CNY', '2022-10-18 18:03:41.455876', 8.4591),
('EUR', '2022-10-18 18:03:41.605845', 60.625999),
('USD', '2022-10-18 18:03:41.722848', 60.875),
('USD', '2022-10-18 18:04:10.695510', 60.875),
('CNY', '2022-10-18 18:04:10.803629', 8.4591),
('EUR', '2022-10-18 18:04:10.593477', 60.625999),
('EUR', '2022-10-18 18:05:10.427390', 60.625999),
('CNY', '2022-10-18 18:05:10.526368', 8.4591),
('USD', '2022-10-18 18:05:10.634531', 60.875),
('EUR', '2022-10-18 18:06:10.437422', 60.625999),
('CNY', '2022-10-18 18:06:10.536208', 8.4591),
('USD', '2022-10-18 18:06:10.701740', 60.875),
('EUR', '2022-10-18 18:07:10.426515', 60.625999),
('USD', '2022-10-18 18:07:10.539588', 60.875),
('CNY', '2022-10-18 18:07:10.656991', 8.4591),
('EUR', '2022-10-18 18:08:10.440828', 60.625999),
('CNY', '2022-10-18 18:08:10.546430', 8.4591),
('USD', '2022-10-18 18:08:10.672686', 60.875),
('EUR', '2022-10-18 18:09:10.434419', 60.625999),
('USD', '2022-10-18 18:09:10.535081', 60.875),
('CNY', '2022-10-18 18:09:10.632773', 8.4591),
('CNY', '2022-10-18 18:10:10.432463', 8.4591),
('USD', '2022-10-18 18:10:10.545695', 60.875),
('EUR', '2022-10-18 18:10:10.658622', 60.625999),
('EUR', '2022-10-18 18:11:10.439660', 60.625999),
('USD', '2022-10-18 18:11:10.982231', 60.875),
('CNY', '2022-10-18 18:11:11.138782', 8.4591),
('EUR', '2022-10-18 18:12:10.437078', 60.625999),
('CNY', '2022-10-18 18:12:10.589102', 8.4591),
('USD', '2022-10-18 18:12:10.743046', 60.875),
('USD', '2022-10-18 18:13:10.435748', 60.875),
('EUR', '2022-10-18 18:13:10.536207', 60.625999),
('CNY', '2022-10-18 18:13:10.640015', 8.4591),
('CNY', '2022-10-18 18:14:10.487908', 8.4591),
('USD', '2022-10-18 18:14:10.588254', 60.875),
('EUR', '2022-10-18 18:14:10.686252', 60.625999),
('CNY', '2022-10-18 18:15:10.445971', 8.4591),
('USD', '2022-10-18 18:15:10.601025', 60.875),
('EUR', '2022-10-18 18:15:10.703114', 60.625999),
('CNY', '2022-10-18 18:16:10.430247', 8.4591),
('EUR', '2022-10-18 18:16:10.531876', 60.625999),
('USD', '2022-10-18 18:16:10.647800', 60.875),
('USD', '2022-10-18 18:17:10.493818', 60.875),
('EUR', '2022-10-18 18:17:10.594422', 60.625999),
('CNY', '2022-10-18 18:17:10.695520', 8.4591),
('EUR', '2022-10-18 18:18:10.440782', 60.625999),
('CNY', '2022-10-18 18:18:10.543521', 8.4591),
('USD', '2022-10-18 18:18:10.645074', 60.875),
('USD', '2022-10-18 18:19:10.441944', 60.875),
('CNY', '2022-10-18 18:19:10.589682', 8.4591),
('EUR', '2022-10-18 18:19:10.693162', 60.625999),
('EUR', '2022-10-18 18:20:10.440227', 60.625999),
('CNY', '2022-10-18 18:20:10.540701', 8.4591),
('USD', '2022-10-18 18:20:10.698748', 60.875),
('EUR', '2022-10-18 18:21:10.488508', 60.625999),
('USD', '2022-10-18 18:21:10.594774', 60.875),
('CNY', '2022-10-18 18:21:10.696809', 8.4591),
('EUR', '2022-10-18 18:22:10.444873', 60.625999),
('CNY', '2022-10-18 18:22:10.547629', 8.4591),
('USD', '2022-10-18 18:22:10.651882', 60.875),
('EUR', '2022-10-18 18:23:31.842334', 60.625999),
('USD', '2022-10-18 18:23:31.940170', 60.875),
('CNY', '2022-10-18 18:23:32.089658', 8.4591),
('USD', '2022-10-18 18:24:31.853333', 60.875),
('EUR', '2022-10-18 18:24:32.001568', 60.625999),
('CNY', '2022-10-18 18:24:31.752596', 8.4591),
('EUR', '2022-10-18 18:25:31.755398', 60.625999),
('USD', '2022-10-18 18:25:32.097306', 60.875),
('CNY', '2022-10-18 18:25:31.949376', 8.4591),
('EUR', '2022-10-18 18:26:31.714313', 60.625999),
('CNY', '2022-10-18 18:26:31.817498', 8.4591),
('USD', '2022-10-18 18:26:31.911363', 60.875),
('EUR', '2022-10-18 18:27:31.813154', 60.625999),
('CNY', '2022-10-18 18:27:31.954633', 8.4591),
('USD', '2022-10-18 18:27:32.054200', 60.875),
('EUR', '2022-10-18 18:28:31.713962', 60.625999),
('CNY', '2022-10-18 18:28:31.812859', 8.4591),
('USD', '2022-10-18 18:28:31.909135', 60.875),
('EUR', '2022-10-18 18:29:31.710909', 60.625999),
('CNY', '2022-10-18 18:29:31.810894', 8.4591),
('USD', '2022-10-18 18:29:31.939660', 60.875),
('EUR', '2022-10-18 18:30:31.698691', 60.625999),
('USD', '2022-10-18 18:30:31.818176', 60.875),
('CNY', '2022-10-18 18:30:31.914328', 8.4591),
('EUR', '2022-10-18 18:31:31.790847', 60.625999),
('CNY', '2022-10-18 18:31:31.892903', 8.4591),
('USD', '2022-10-18 18:31:32.005066', 60.875),
('EUR', '2022-10-18 18:32:31.721682', 60.625999),
('CNY', '2022-10-18 18:32:31.836349', 8.4591),
('USD', '2022-10-18 18:32:31.931357', 60.875),
('CNY', '2022-10-18 18:33:31.702447', 8.4591),
('USD', '2022-10-18 18:33:31.802788', 60.875),
('EUR', '2022-10-18 18:33:31.952048', 60.625999),
('EUR', '2022-10-18 18:34:31.713585', 60.625999),
('CNY', '2022-10-18 18:34:31.807852', 8.4591),
('USD', '2022-10-18 18:34:31.905215', 60.875),
('EUR', '2022-10-18 18:35:31.702025', 60.625999),
('CNY', '2022-10-18 18:35:31.800723', 8.4591),
('USD', '2022-10-18 18:35:31.899199', 60.875),
('EUR', '2022-10-18 18:36:31.704733', 60.625999),
('CNY', '2022-10-18 18:36:31.803533', 8.4591),
('USD', '2022-10-18 18:36:31.900637', 60.875),
('EUR', '2022-10-18 18:37:31.755728', 60.625999),
('CNY', '2022-10-18 18:37:31.851585', 8.4591),
('USD', '2022-10-18 18:37:31.949782', 60.875),
('EUR', '2022-10-18 18:38:31.712359', 60.625999),
('CNY', '2022-10-18 18:38:31.962875', 8.4591),
('USD', '2022-10-18 18:38:32.020481', 60.875),
('CNY', '2022-10-18 18:39:31.705777', 8.4591),
('EUR', '2022-10-18 18:39:31.810110', 60.625999),
('USD', '2022-10-18 18:39:31.909844', 60.875),
('EUR', '2022-10-18 18:40:31.694892', 60.625999),
('CNY', '2022-10-18 18:40:31.792296', 8.4591),
('USD', '2022-10-18 18:40:31.900106', 60.875),
('CNY', '2022-10-18 18:41:31.703576', 8.4591),
('USD', '2022-10-18 18:41:31.806029', 60.875),
('EUR', '2022-10-18 18:41:31.901334', 60.625999),
('CNY', '2022-10-18 18:42:31.701136', 8.4591),
('USD', '2022-10-18 18:42:31.801083', 60.875),
('EUR', '2022-10-18 18:42:31.953263', 60.625999),
('EUR', '2022-10-18 18:43:31.703535', 60.625999),
('CNY', '2022-10-18 18:43:31.812558', 8.4591),
('USD', '2022-10-18 18:43:31.907695', 60.875),
('CNY', '2022-10-18 18:44:31.764440', 8.4591),
('USD', '2022-10-18 18:44:31.864971', 60.875),
('EUR', '2022-10-18 18:44:31.961619', 60.625999),
('USD', '2022-10-18 18:45:31.713880', 60.875),
('EUR', '2022-10-18 18:45:31.898840', 60.625999),
('CNY', '2022-10-18 18:45:32.002301', 8.4591),
('EUR', '2022-10-18 18:46:31.774427', 60.625999),
('CNY', '2022-10-18 18:46:31.876736', 8.4591),
('USD', '2022-10-18 18:46:31.986058', 60.875),
('EUR', '2022-10-18 18:47:31.702253', 60.625999),
('CNY', '2022-10-18 18:47:31.807550', 8.4591),
('USD', '2022-10-18 18:47:31.906046', 60.875),
('EUR', '2022-10-18 18:48:31.704281', 60.625999),
('CNY', '2022-10-18 18:48:31.804157', 8.4591),
('USD', '2022-10-18 18:48:31.906432', 60.875),
('USD', '2022-10-18 18:49:31.693919', 60.875),
('CNY', '2022-10-18 18:49:31.800353', 8.4591),
('EUR', '2022-10-18 18:49:31.901606', 60.625999),
('EUR', '2022-10-18 18:51:06.251073', 60.625999),
('USD', '2022-10-18 18:51:06.352406', 60.875),
('CNY', '2022-10-18 18:51:06.455255', 8.4591),
('EUR', '2022-10-18 18:52:06.105741', 60.625999),
('CNY', '2022-10-18 18:52:06.205834', 8.4591),
('USD', '2022-10-18 18:52:06.309120', 60.875),
('EUR', '2022-10-18 18:53:06.092108', 60.625999),
('USD', '2022-10-18 18:53:06.194882', 60.875),
('CNY', '2022-10-18 18:53:06.358660', 8.4591),
('EUR', '2022-10-18 18:54:06.096558', 60.625999),
('CNY', '2022-10-18 18:54:06.198217', 8.4591),
('USD', '2022-10-18 18:54:06.297412', 60.875),
('USD', '2022-10-18 18:55:06.085595', 60.875),
('CNY', '2022-10-18 18:55:06.185873', 8.4591),
('EUR', '2022-10-18 18:55:06.286806', 60.625999),
('EUR', '2022-10-18 18:56:06.264728', 60.625999),
('USD', '2022-10-18 18:56:06.516124', 60.875),
('CNY', '2022-10-18 18:56:06.363226', 8.4591),
('EUR', '2022-10-18 18:57:06.094869', 60.625999),
('USD', '2022-10-18 18:57:06.194244', 60.875),
('CNY', '2022-10-18 18:57:06.316545', 8.4591),
('CNY', '2022-10-18 18:58:23.414057', 8.4591),
('USD', '2022-10-18 18:58:24.981942', 60.875),
('EUR', '2022-10-18 18:58:25.081716', 60.625999),
('EUR', '2022-10-18 18:59:06.096833', 60.625999),
('USD', '2022-10-18 18:59:06.201709', 60.875),
('CNY', '2022-10-18 18:59:06.302020', 8.4591),
('USD', '2022-10-18 19:00:06.232964', 60.875),
('EUR', '2022-10-18 19:00:06.333831', 60.625999),
('CNY', '2022-10-18 19:00:06.090236', 8.4591),
('USD', '2022-10-18 19:01:09.874687', 60.875),
('EUR', '2022-10-18 19:01:10.080327', 60.625999),
('CNY', '2022-10-18 19:01:09.986124', 8.4591),
('CNY', '2022-10-18 19:02:06.094025', 8.4591),
('EUR', '2022-10-18 19:02:06.195808', 60.625999),
('USD', '2022-10-18 19:02:06.303036', 60.875),
('CNY', '2022-10-18 19:03:06.089811', 8.4591),
('USD', '2022-10-18 19:03:06.189391', 60.875),
('EUR', '2022-10-18 19:03:06.299288', 60.625999),
('USD', '2022-10-18 19:04:06.093103', 60.875),
('CNY', '2022-10-18 19:04:06.192841', 8.4591),
('EUR', '2022-10-18 19:04:06.291176', 60.625999),
('USD', '2022-10-18 19:05:06.102149', 60.875),
('EUR', '2022-10-18 19:05:06.306875', 60.625999),
('CNY', '2022-10-18 19:05:06.203406', 8.4591),
('CNY', '2022-10-18 19:06:07.587593', 8.4591),
('EUR', '2022-10-18 19:06:07.690789', 60.625999),
('USD', '2022-10-18 19:06:07.788341', 60.875),
('EUR', '2022-10-18 19:07:06.108536', 60.625999),
('USD', '2022-10-18 19:07:06.226061', 60.875),
('CNY', '2022-10-18 19:07:06.373495', 8.4591),
('EUR', '2022-10-18 23:07:18.489014', 60.625999),
('USD', '2022-10-18 23:07:18.585476', 60.875),
('CNY', '2022-10-18 23:07:18.738129', 8.4591),
('CNY', '2022-10-18 23:08:18.421713', 8.4591),
('USD', '2022-10-18 23:08:18.585185', 60.875),
('EUR', '2022-10-18 23:08:18.689853', 60.625999),
('USD', '2022-10-18 23:09:18.299270', 60.875),
('CNY', '2022-10-18 23:09:18.400778', 8.4591),
('EUR', '2022-10-18 23:09:18.498460', 60.625999),
('CNY', '2022-10-18 23:10:18.300776', 8.4591),
('USD', '2022-10-18 23:10:18.399888', 60.875),
('EUR', '2022-10-18 23:10:18.499766', 60.625999),
('USD', '2022-10-18 23:11:18.306068', 60.875),
('CNY', '2022-10-18 23:11:18.408411', 8.4591),
('EUR', '2022-10-18 23:11:18.557903', 60.625999),
('CNY', '2022-10-18 23:12:18.304385', 8.4591),
('USD', '2022-10-18 23:12:18.402187', 60.875),
('EUR', '2022-10-18 23:12:18.498898', 60.625999),
('USD', '2022-10-18 23:13:18.301659', 60.875),
('CNY', '2022-10-18 23:13:18.398676', 8.4591),
('EUR', '2022-10-18 23:13:18.499249', 60.625999),
('EUR', '2022-10-18 23:14:18.299599', 60.625999),
('CNY', '2022-10-18 23:14:18.459374', 8.4591),
('USD', '2022-10-18 23:14:18.555308', 60.875),
('EUR', '2022-10-18 23:15:18.304634', 60.625999),
('USD', '2022-10-18 23:15:18.404226', 60.875),
('CNY', '2022-10-18 23:15:18.501775', 8.4591),
('CNY', '2022-10-18 23:16:18.299255', 8.4591),
('USD', '2022-10-18 23:16:18.425851', 60.875),
('EUR', '2022-10-18 23:16:18.527657', 60.625999),
('USD', '2022-10-18 23:17:18.306020', 60.875),
('CNY', '2022-10-18 23:17:18.406116', 8.4591),
('EUR', '2022-10-18 23:17:18.509462', 60.625999),
('CNY', '2022-10-18 23:18:18.307265', 8.4591),
('EUR', '2022-10-18 23:18:18.430629', 60.625999),
('USD', '2022-10-18 23:18:18.579959', 60.875),
('EUR', '2022-10-18 23:19:18.315116', 60.625999),
('USD', '2022-10-18 23:19:18.417926', 60.875),
('CNY', '2022-10-18 23:19:18.527431', 8.4591),
('CNY', '2022-10-18 23:31:55.019457', 8.4591),
('EUR', '2022-10-18 23:31:55.134047', 60.625999),
('USD', '2022-10-18 23:31:55.233380', 60.875),
('CNY', '2022-10-18 23:32:54.867598', 8.4591),
('USD', '2022-10-18 23:32:54.971582', 60.875),
('EUR', '2022-10-18 23:32:55.072768', 60.625999),
('EUR', '2022-10-18 23:33:54.874283', 60.625999),
('USD', '2022-10-18 23:33:54.982658', 60.875),
('CNY', '2022-10-18 23:33:55.085636', 8.4591),
('CNY', '2022-10-20 22:40:27.276384', 8.4229),
('EUR', '2022-10-20 22:40:27.381373', 59.549),
('USD', '2022-10-20 22:40:27.487306', 61),
('USD', '2022-10-20 22:41:27.052012', 61),
('CNY', '2022-10-20 22:41:27.148672', 8.4229),
('EUR', '2022-10-20 22:41:27.247278', 59.549),
('USD', '2022-10-20 22:42:27.049565', 61),
('CNY', '2022-10-20 22:42:27.149987', 8.4229),
('EUR', '2022-10-20 22:42:27.246595', 59.549),
('CNY', '2022-10-20 22:43:27.096545', 8.4229),
('EUR', '2022-10-20 22:43:27.196117', 59.549),
('USD', '2022-10-20 22:43:27.299053', 61),
('USD', '2022-10-20 22:44:27.055868', 61),
('CNY', '2022-10-20 22:44:27.152196', 8.4229),
('EUR', '2022-10-20 22:44:27.253188', 59.549),
('USD', '2022-10-20 22:45:27.044001', 61),
('CNY', '2022-10-20 22:45:27.139734', 8.4229),
('EUR', '2022-10-20 22:45:27.236461', 59.549),
('EUR', '2022-10-20 22:46:27.059626', 59.549),
('USD', '2022-10-20 22:46:27.155016', 61),
('CNY', '2022-10-20 22:46:27.251061', 8.4229),
('USD', '2022-10-20 22:47:27.063308', 61),
('CNY', '2022-10-20 22:47:27.167581', 8.4229),
('EUR', '2022-10-20 22:47:27.264776', 59.549),
('USD', '2022-10-20 22:48:27.119681', 61),
('EUR', '2022-10-20 22:48:27.224205', 59.549),
('CNY', '2022-10-20 22:48:27.366660', 8.4229),
('EUR', '2022-10-20 22:49:27.042618', 59.549),
('USD', '2022-10-20 22:49:27.139466', 61),
('CNY', '2022-10-20 22:49:27.236457', 8.4229),
('CNY', '2022-10-20 22:50:32.694793', 8.4229),
('EUR', '2022-10-20 22:50:32.793904', 59.549),
('USD', '2022-10-20 22:51:27.189702', 61),
('EUR', '2022-10-20 22:51:27.331026', 59.549),
('CNY', '2022-10-20 22:51:27.497388', 8.4229),
('EUR', '2022-10-20 22:52:30.799887', 59.549),
('USD', '2022-10-20 22:52:32.157525', 61),
('CNY', '2022-10-20 22:52:32.604108', 8.4229),
('EUR', '2022-10-20 22:53:27.052111', 59.549),
('USD', '2022-10-20 22:53:27.151176', 61),
('CNY', '2022-10-20 22:53:27.251881', 8.4229),
('EUR', '2022-10-20 22:54:27.046205', 59.549),
('USD', '2022-10-20 22:54:27.182364', 61),
('CNY', '2022-10-20 22:54:27.282819', 8.4229),
('USD', '2022-10-20 22:55:27.056460', 61),
('CNY', '2022-10-20 22:55:27.153930', 8.4229),
('EUR', '2022-10-20 22:55:27.308643', 59.549),
('EUR', '2022-10-20 22:56:27.055670', 59.549),
('USD', '2022-10-20 22:56:27.166604', 61),
('CNY', '2022-10-20 22:56:27.275656', 8.4229),
('USD', '2022-10-20 22:57:27.050890', 61),
('CNY', '2022-10-20 22:57:27.149126', 8.4229),
('EUR', '2022-10-20 22:57:27.246900', 59.549),
('EUR', '2022-10-20 22:58:27.056704', 59.549),
('USD', '2022-10-20 22:58:27.156537', 61),
('CNY', '2022-10-20 22:58:27.313194', 8.4229),
('EUR', '2022-10-20 22:59:27.052123', 59.549),
('USD', '2022-10-20 22:59:27.149511', 61),
('CNY', '2022-10-20 22:59:27.249291', 8.4229),
('USD', '2022-10-20 23:00:27.056077', 61),
('CNY', '2022-10-20 23:00:27.152843', 8.4229),
('EUR', '2022-10-20 23:00:27.261902', 59.549),
('EUR', '2022-10-20 23:01:27.055249', 59.549),
('USD', '2022-10-20 23:01:27.159934', 61),
('CNY', '2022-10-20 23:01:27.260530', 8.4229),
('USD', '2022-10-20 23:02:27.059015', 61),
('CNY', '2022-10-20 23:02:27.155435', 8.4229),
('EUR', '2022-10-20 23:02:27.258812', 59.549),
('EUR', '2022-10-20 23:03:27.053755', 59.549),
('USD', '2022-10-20 23:03:27.161875', 61),
('CNY', '2022-10-20 23:03:27.263554', 8.4229),
('USD', '2022-10-20 23:04:27.053557', 61),
('CNY', '2022-10-20 23:04:27.150121', 8.4229),
('EUR', '2022-10-20 23:04:27.251394', 59.549),
('USD', '2022-10-20 23:05:27.056227', 61),
('CNY', '2022-10-20 23:05:27.167794', 8.4229),
('EUR', '2022-10-20 23:05:27.267366', 59.549),
('USD', '2022-10-20 23:06:27.060594', 61),
('CNY', '2022-10-20 23:06:27.158744', 8.4229),
('EUR', '2022-10-20 23:06:27.320578', 59.549),
('CNY', '2022-10-20 23:07:27.047731', 8.4229),
('EUR', '2022-10-20 23:07:27.150427', 59.549),
('USD', '2022-10-20 23:07:27.250651', 61),
('USD', '2022-10-20 23:08:27.056894', 61),
('EUR', '2022-10-20 23:08:27.212265', 59.549),
('CNY', '2022-10-20 23:08:27.312351', 8.4229),
('USD', '2022-10-20 23:09:27.046990', 61),
('EUR', '2022-10-20 23:09:27.145172', 59.549),
('CNY', '2022-10-20 23:09:27.249492', 8.4229),
('USD', '2022-10-20 23:10:27.060033', 61),
('CNY', '2022-10-20 23:10:27.161395', 8.4229),
('EUR', '2022-10-20 23:10:27.264552', 59.549),
('USD', '2022-10-20 23:11:27.046978', 61),
('EUR', '2022-10-20 23:11:27.148495', 59.549),
('CNY', '2022-10-20 23:11:27.247129', 8.4229),
('EUR', '2022-10-20 23:12:27.056443', 59.549),
('USD', '2022-10-20 23:12:27.157359', 61),
('CNY', '2022-10-20 23:12:27.260186', 8.4229),
('USD', '2022-10-20 23:13:27.057718', 61),
('CNY', '2022-10-20 23:13:27.158405', 8.4229),
('EUR', '2022-10-20 23:13:27.256133', 59.549),
('USD', '2022-10-20 23:14:27.046451', 61),
('CNY', '2022-10-20 23:14:27.147495', 8.4229),
('EUR', '2022-10-20 23:14:27.275611', 59.549),
('USD', '2022-10-20 23:15:27.055841', 61),
('CNY', '2022-10-20 23:15:27.177489', 8.4229),
('EUR', '2022-10-20 23:15:27.342502', 59.549),
('CNY', '2022-10-20 23:27:31.741211', 8.4229),
('EUR', '2022-10-20 23:27:31.849238', 59.549),
('USD', '2022-10-20 23:27:31.945410', 61),
('USD', '2022-10-20 23:28:31.600034', 61),
('EUR', '2022-10-20 23:28:31.698900', 59.549),
('CNY', '2022-10-20 23:28:31.796948', 8.4229),
('CNY', '2022-10-20 23:29:31.593314', 8.4229),
('USD', '2022-10-20 23:29:31.688991', 61),
('EUR', '2022-10-20 23:29:31.785910', 59.549),
('CNY', '2022-10-21 00:02:41.749735', 8.4229),
('EUR', '2022-10-21 00:02:41.849093', 59.549),
('USD', '2022-10-21 00:02:41.946955', 61),
('CNY', '2022-10-21 00:03:41.540606', 8.4229),
('USD', '2022-10-21 00:03:41.704338', 61),
('EUR', '2022-10-21 00:03:41.779476', 59.549),
('CNY', '2022-10-21 00:04:41.488818', 8.4229),
('USD', '2022-10-21 00:04:41.591532', 61),
('EUR', '2022-10-21 00:04:41.693425', 59.549),
('CNY', '2022-10-21 00:05:41.489998', 8.4229),
('USD', '2022-10-21 00:05:41.587070', 61),
('EUR', '2022-10-21 00:05:41.684500', 59.549),
('CNY', '2022-10-21 00:06:41.489034', 8.4229),
('USD', '2022-10-21 00:06:41.588594', 61),
('EUR', '2022-10-21 00:06:41.687557', 59.549),
('CNY', '2022-10-21 00:07:41.494456', 8.4229),
('USD', '2022-10-21 00:07:41.646819', 61),
('EUR', '2022-10-21 00:07:41.744800', 59.549),
('CNY', '2022-10-21 00:08:41.500258', 8.4229),
('USD', '2022-10-21 00:08:41.604652', 61),
('EUR', '2022-10-21 00:08:41.703269', 59.549),
('USD', '2022-10-21 00:09:41.501958', 61),
('EUR', '2022-10-21 00:09:41.604700', 59.549),
('CNY', '2022-10-21 00:09:41.711506', 8.4229),
('CNY', '2022-10-21 00:10:41.492553', 8.4229),
('USD', '2022-10-21 00:10:41.649570', 61),
('EUR', '2022-10-21 00:10:41.804819', 59.549),
('CNY', '2022-10-21 00:11:41.496145', 8.4229),
('USD', '2022-10-21 00:11:41.596019', 61),
('EUR', '2022-10-21 00:11:41.695398', 59.549),
('USD', '2022-10-21 00:12:41.487289', 61),
('EUR', '2022-10-21 00:12:41.590132', 59.549),
('CNY', '2022-10-21 00:12:41.746078', 8.4229),
('USD', '2022-10-21 00:13:41.494158', 61),
('EUR', '2022-10-21 00:13:41.596419', 59.549),
('CNY', '2022-10-21 00:13:41.706209', 8.4229),
('USD', '2022-10-21 00:14:33.657952', 61),
('CNY', '2022-10-21 00:14:33.759787', 8.4229),
('EUR', '2022-10-21 00:14:33.859342', 59.549),
('EUR', '2022-10-21 00:16:04.244358', 59.549),
('USD', '2022-10-21 00:16:04.349585', 61),
('CNY', '2022-10-21 00:16:40.659275', 8.4229),
('EUR', '2022-10-21 00:16:40.761667', 59.397701),
('USD', '2022-10-21 00:16:40.859407', 60.965),
('EUR', '2022-10-21 00:20:20.696659', 59.397701),
('USD', '2022-10-21 00:20:20.798599', 60.965),
('CNY', '2022-10-21 00:20:20.897212', 8.4229),
('CNY', '2022-10-21 00:21:25.180224', 8.4229),
('EUR', '2022-10-21 00:21:25.336313', 59.397701),
('USD', '2022-10-21 00:21:25.439369', 60.965),
('CNY', '2022-10-21 00:22:25.039652', 8.4229),
('USD', '2022-10-21 00:22:25.140100', 60.965),
('EUR', '2022-10-21 00:22:25.291528', 59.397701),
('USD', '2022-10-21 01:13:31.149585', 60.965),
('CNY', '2022-10-21 01:13:31.258150', 8.4229),
('EUR', '2022-10-21 01:13:31.418142', 59.397701),
('USD', '2022-10-21 01:14:30.855677', 60.965),
('EUR', '2022-10-21 01:14:30.961219', 59.397701),
('CNY', '2022-10-21 01:14:31.069937', 8.4229),
('USD', '2022-10-21 01:15:30.853148', 60.965),
('CNY', '2022-10-21 01:15:30.963711', 8.4229),
('EUR', '2022-10-21 01:15:31.072301', 59.397701),
('USD', '2022-10-21 01:16:30.856095', 60.965),
('EUR', '2022-10-21 01:16:30.975819', 59.397701),
('CNY', '2022-10-21 01:16:31.083877', 8.4229),
('CNY', '2022-10-21 01:17:30.851334', 8.4229),
('EUR', '2022-10-21 01:17:30.961524', 59.397701),
('USD', '2022-10-21 01:17:31.066241', 60.965),
('USD', '2022-10-21 01:18:30.862323', 60.965),
('EUR', '2022-10-21 01:18:30.971262', 59.397701),
('CNY', '2022-10-21 01:18:31.077166', 8.4229),
('USD', '2022-10-21 01:19:30.904780', 60.965),
('CNY', '2022-10-21 01:19:31.023096', 8.4229),
('EUR', '2022-10-21 01:19:31.124810', 59.397701),
('USD', '2022-10-21 01:20:30.861848', 60.965),
('EUR', '2022-10-21 01:20:30.965781', 59.397701),
('CNY', '2022-10-21 01:20:31.072503', 8.4229),
('CNY', '2022-10-21 01:21:30.855995', 8.4229),
('EUR', '2022-10-21 01:21:30.965306', 59.397701),
('USD', '2022-10-21 01:21:31.073700', 60.965),
('EUR', '2022-10-21 01:22:30.853983', 59.397701),
('USD', '2022-10-21 01:22:30.964104', 60.965),
('CNY', '2022-10-21 01:22:31.071423', 8.4229),
('CNY', '2022-10-21 01:44:00.908616', 8.4229),
('EUR', '2022-10-21 01:44:01.019719', 59.397701),
('USD', '2022-10-21 01:44:01.130493', 60.965),
('CNY', '2022-10-21 01:45:00.727995', 8.4229),
('EUR', '2022-10-21 01:45:00.942457', 59.397701),
('USD', '2022-10-21 01:45:00.835616', 60.965),
('CNY', '2022-10-21 01:46:00.731200', 8.4229),
('EUR', '2022-10-21 01:46:00.851617', 59.397701),
('USD', '2022-10-21 01:46:00.978349', 60.965),
('CNY', '2022-10-21 01:47:00.726157', 8.4229),
('USD', '2022-10-21 01:47:00.832731', 60.965),
('EUR', '2022-10-21 01:47:00.943899', 59.397701),
('CNY', '2022-10-21 01:53:47.216710', 8.4229),
('EUR', '2022-10-21 01:53:47.327895', 59.397701),
('USD', '2022-10-21 01:53:47.434845', 60.965),
('CNY', '2022-10-21 18:29:26.066893', 8.4229),
('EUR', '2022-10-21 18:29:26.169347', 59.549),
('USD', '2022-10-21 18:29:26.270623', 60.75),
('CNY', '2022-10-21 19:27:57.031231', 8.4229),
('EUR', '2022-10-21 19:27:57.132948', 59.549),
('USD', '2022-10-21 19:27:57.236714', 60.75),
('CNY', '2022-10-21 19:48:35.710649', 8.4229),
('EUR', '2022-10-21 19:48:35.813842', 59.549),
('USD', '2022-10-21 19:48:35.912097', 60.75),
('EUR', '2022-10-22 00:43:34.755955', 59.549),
('CNY', '2022-10-22 00:43:34.958112', 8.4229),
('USD', '2022-10-22 00:43:34.854527', 60.75),
('EUR', '2022-10-22 00:44:34.489269', 59.549),
('CNY', '2022-10-22 00:44:34.590568', 8.4229),
('USD', '2022-10-22 00:44:34.692213', 60.75),
('USD', '2022-10-22 02:40:54.184646', 61.615002),
('CNY', '2022-10-22 02:40:54.283829', 8.4229),
('EUR', '2022-10-22 02:40:54.391016', 59.313099);

END $$;