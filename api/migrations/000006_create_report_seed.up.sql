INSERT INTO categories (name)
VALUES
('フレンチ'),
('中華'),
('和食'),
('イタリアン'),
('韓国料理');



INSERT INTO recipes (id, title, thumbnail_url, recipe, category_id, ingredient, created_at, updated_at)
VALUES
('01J6CXNF4GYJFTNTTH9TZ3MSJG', 'トマトサラダ', 'https://video.com/production/videos/b331637e-1067-4471-acf3-/.jpg', 'トマトとバジルのシンプルなサラダ。', 1, '[{"ingredient": "トマト", "amount": "100g"}, {"ingredient": "バジル", "amount": "10枚"}]', '2024-08-01 10:00:00', '2024-08-01 10:00:00'),
('01J6CXP7VZ8DMDNRWAKY5CNPBK', '麻婆豆腐', 'https://video..com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', '辛い中華風の豆腐料理。', 2, '[{"ingredient": "豆腐", "amount": "300g"}, {"ingredient": "ひき肉", "amount": "100g"}]', '2024-08-05 12:30:00', '2024-08-05 12:30:00'),
('01J6CXPHN9ZDYGQSDSKY0ERVBY', '天ぷら', 'https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'サクサクの揚げ物。', 3, '[{"ingredient": "エビ", "amount": "3尾"}, {"ingredient": "小麦粉", "amount": "50g"}]', '2024-08-10 08:45:00', '2024-08-10 08:45:00'),
('01J6CXPTXMX9CYFNQPHAM6KY3H', 'カルボナーラ', 'https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'クリーミーなパスタ料理。', 4, '[{"ingredient": "スパゲッティ", "amount": "200g"}, {"ingredient": "卵", "amount": "2個"}]', '2024-08-15 09:15:00', '2024-08-15 09:15:00'),
('01J6CXQ7ZVQ5SDFSV2SBJSR8G0', 'キムチチゲ', 'https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'ピリ辛の韓国風鍋料理。', 5, '[{"ingredient": "キムチ", "amount": "200g"}, {"ingredient": "豚肉", "amount": "150g"}]', '2024-08-20 14:50:00', '2024-08-20 14:50:00');

INSERT INTO reports (id, comment, thumbnail_url, author_id, recipe_id, created_at, updated_at)
VALUES
('1e8b6d70-9c2f-11eb-a8b3-0242ac130003', 'トマトサラダのレビュー。簡単で美味しかったです！', 'https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'a4e8b5d4-9c1f-11eb-a8b3-0242ac130003', '01J6CXNF4GYJFTNTTH9TZ3MSJG', '2024-09-01 09:00:00', '2024-09-01 09:00:00'),
('2f8c7e80-9c2f-11eb-a8b3-0242ac130003', '麻婆豆腐のレビュー。辛さがちょうど良く、ご飯が進みました。', 'https://..com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'b5f2c3d2-9c1f-11eb-a8b3-0242ac130003', '01J6CXP7VZ8DMDNRWAKY5CNPBK', '2024-09-02 14:15:00', '2024-09-02 14:15:00'),
('3a9d8f90-9c2f-11eb-a8b3-0242ac130003', 'カルボナーラのレビュー。濃厚でクリーミー、家族にも好評でした。', 'https://.kurashiru.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg', 'c6d3f6e8-9c1f-11eb-a8b3-0242ac130003', '01J6CXPTXMX9CYFNQPHAM6KY3H', '2024-09-03 18:30:00', '2024-09-03 18:30:00');
