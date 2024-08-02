CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcategories (
    id UUID PRIMARY KEY,
    category_id UUID REFERENCES categories(id),
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcategory_news (
    subcategory_id UUID REFERENCES subcategories(id),
    news_id UUID REFERENCES news(id) ON DELETE CASCADE,
    PRIMARY KEY (subcategory_id, news_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS news (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    links JSONB
);

CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY,
    username TEXT,
    password TEXT,
    avatar TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS superadmins (
    id UUID PRIMARY KEY,
    phone_number TEXT,
    password TEXT,
    avatar TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- INSERT INTO superadmins (id, username, password) VALUES('acc98ad0-43a1-4ac5-ba90-7dc1f1a34d1e', 'test', 'test');

-- Insert into categories table
-- Insert Categories
INSERT INTO categories (id, name) VALUES
-- Uzbekistan
('b1d7357d-1d92-4e77-a8d0-1d394c5b2ef2', 'O''zbekiston'),

('99d0724b-2d49-4140-8b60-0c41fe214b82', 'JONLI'),

-- World
('09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Dunyo'),

-- LIVE
('f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Ko''ngilochar'),

-- Sports
('9d57d482-6db2-43ec-8513-6478c066aa51', 'Sport'),

-- Auto
('23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3', 'Avto'),

-- Technology
('6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Texnologiya'),

-- Economy
('e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Iqtisod'),

-- Show Business
('72a5ec87-44ed-438c-b9d0-0e391ed7da4d', 'Shou-biznes'),

-- Military News
('b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69', 'Harbiy yangiliklar'),

-- Daily
('a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Kundalik'),

-- TOP 10
('7e27d7bb-258d-4df5-810d-9b5c3146a606', 'TOP 10');


-- Insert into subcategories table
-- Insert Subcategories
INSERT INTO subcategories (id, category_id, name) VALUES
-- Uzbekistan
('cfc38315-9a1e-4f38-8f4e-63196e5e4b4d', 'b1d7357d-1d92-4e77-a8d0-1d394c5b2ef2', 'O''zbekiston'),

-- World
('10c5f9f4-9b32-4e9a-9b91-ff4f38cf1d62', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Siyosat'),
('bfb0cda3-d527-4ea7-8b8b-08f36e8a84d4', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Jamiyat'),
('0ec7c541-3b4c-46a1-8e29-6c5014c7c84b', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Hodisalar'),
('74678315-cc88-453d-84aa-c56d3a9176fc', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Jinoyat'),
('d6d6d04a-07c6-4666-93b5-979d146dc7b8', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Nizolar'),

-- Entertainment
('27c3f98f-3ad7-4c8f-9f91-91b2dd9b2271', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Filmlar'),
('96cce354-ecb6-47c2-9c09-b7e2e7d185b1', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Teatr'),
('9cb6f8c6-4ed6-4b29-9e1f-9780122db0b1', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Restoranlar'),
('b7e0b4c1-bf7e-452c-9b89-d5424e02b122', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Konsertlar'),
('7c3d4b65-0f3e-4962-9d6b-cf2e1b53ff68', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Stand-up komediya'),
('99a4b6a8-d51e-4b83-92f6-246a2a2a4f60', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Parklar'),
('4cb2e539-cbbf-4b1f-a4d3-1f7d66330b1f', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Ko''rgazmalar'),
('f4c04dbe-6bfa-4e92-b07e-c24dce0ea5a4', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Kolleksiyalar'),
('ce0b3c85-27f6-4a4d-9b8b-c71720f8e724', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Bolalar'),
('b56b69b0-505d-47f6-bc7d-7b2f84fa0c15', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Bayramlar'),
('745dce4b-5cb2-4868-915e-b2b06cb5d31a', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Sport'),
('a55fd728-4d44-4c7e-8311-bf2c8e5b8e67', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Bazmlar'),
('ea4931d8-7d38-4c4f-83d7-5c5d3641f042', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Kino kitubxonasi'),
('9785b78b-9e9c-4ef5-9b6c-fbf64ae5f29b', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Yo''riqnomalar'),

-- LIVE
('8c550a80-817f-4e46-9f35-507f2233fa0b', '99d0724b-2d49-4140-8b60-0c41fe214b82', 'Instagram Reels'),

-- Sports
('32b066bc-6c61-46ae-9fa7-4e3a7e6d36d3', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Futbol'),
('7ac905e4-37b6-4c18-96d6-f9304c1c72e0', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Hokkey'),
('b9d7a6de-d7a8-4a54-b44e-0266b70204c5', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Boks/MMA'),
('1c2b5e12-5b5b-40f8-9eb7-d244f4473e8f', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Avtomobil sporti'),
('e290bd87-92d8-4c7e-905e-5015b535dbf2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Tennis'),
('4e00a79b-982a-407e-b83e-9e43e8aef3f2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Basketbol'),
('5079c95f-73d3-4c62-8b27-2045292efc4e', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Figurali uchish'),
('6a1f0234-fbc3-4b7c-8140-750ab49f71ec', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Kibersport'),
('6b9479d2-67d3-4f4a-b3f2-4a61ff5466c2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Shaxmat'),
('4c9e3912-56a2-47d2-9a2a-daf4349eac5f', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Yozgi sport'),
('33d43e0c-0452-4744-9c26-560fb60619c7', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Qishgi sport'),
('743a8c3d-7a7e-44ea-a6b0-36b76ad908b1', '9d57d482-6db2-43ec-8513-6478c066aa51', 'O''zbekiston'),
('02951c36-2ec5-4c46-bc65-3a344e7275a8', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Chempionat rasmlari'),

-- Auto
('2eeb5d0e-3487-4b96-873c-572273a6a50a', '23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3', 'Avto'),

-- Technology
('e949e339-d060-4d06-bb38-06b06970a3f8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Fan'),
('8c4d0875-91d3-4e46-8c36-d76d34d960b5', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Fazo'),
('a0581f69-51e1-48ef-b7fc-8238e7b7a4e8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Qurollar'),
('9f6d3b39-4c5c-4202-8923-cb3db238d8dc', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Tarix'),
('0308c01b-bc22-4a6f-8927-ea2b79289b51', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Sog''liq'),
('564deef4-d9b6-4c6f-a14d-79d76a78c9d8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Texnikalar'),
('aa935589-46b0-45e3-8062-5a7b0e52fc29', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Gadjetlar'),
('b2f52b78-e4ac-44d6-a4b2-3e372d2a205f', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'O''zbekiston'),

-- Economy
('efb450b5-9ac4-42e7-a7c0-1a8b2c1e12d0', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Iqtisod'),
('cbb7a1d5-593d-4f6e-a64d-bdd8db3e5e6d', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Kompaniyalar'),
('fa8c0e93-8a16-4e2b-bbab-cf2e0e1e81d3', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Shaxsiy hisob'),
('a1db5e43-56c5-4a3b-91a4-69959e4e7d6a', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Ko''chmas mulk'),
('bdad84c3-9e4f-47ec-8b5f-97ad9b5d7db4', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Import'),
('f7d7c9d5-5b71-4994-8b47-8d705e5e11b5', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Biznes muhiti'),
('fc88f84e-cfb5-45e0-bb1e-5b90b4a7df51', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'O''zbekiston'),

-- Show Business
('a0d44f37-8e54-47d7-82d0-51a7ec728f37', '72a5ec87-44ed-438c-b9d0-0e391ed7da4d', 'Shou biznes'),

-- Military News
('68a4c6d2-facb-40d2-8c6b-9c9b9f3c40c1', 'b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69', 'Harbiy yangiliklar'),

-- Daily
('0a281039-8711-4f0e-9a4f-e4c2a9de003f', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Oziq-ovqat'),
('6f1e61a8-8f4b-47f1-88c7-935e611ed8f0', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Psixologiya'),
('d97b0f91-1d61-464f-8d8b-c32a63e827f6', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Trendlar'),
('b2e095eb-bd51-4a61-89f0-d54d4a00c1d2', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Bolalar'),
('0a9b0690-7031-4860-b605-9900df22868f', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Uy va bog'''),
('3e8f3066-b36c-4a34-8d89-5e4b9c5b37d1', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Voqealar'),
('f0b7e501-72b1-4f2d-8fae-4986b9e59715', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Janjallar'),

-- TOP 10
('dfbdbf6e-9c93-44bc-a302-cfc1ac16e75b', '7e27d7bb-258d-4df5-810d-9b5c3146a606', 'TOP 10');















-- -- Update categories table with Uzbek translations
-- UPDATE categories SET name = 'Oʻzbekiston' WHERE id = '';
-- UPDATE categories SET name = 'Dunyo' WHERE id = '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d';
-- UPDATE categories SET name = 'Koʻngilochar' WHERE id = 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1';
-- UPDATE categories SET name = 'Sport' WHERE id = '9d57d482-6db2-43ec-8513-6478c066aa51';
-- UPDATE categories SET name = 'Avto' WHERE id = '23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3';
-- UPDATE categories SET name = 'Texnologiya' WHERE id = '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb';
-- UPDATE categories SET name = 'Iqtisodiyot' WHERE id = 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e';
-- UPDATE categories SET name = 'Shou-biznes' WHERE id = '72a5ec87-44ed-438c-b9d0-0e391ed7da4d';
-- UPDATE categories SET name = 'Harbiy yangiliklar' WHERE id = 'b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69';
-- UPDATE categories SET name = 'Kunlik' WHERE id = 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb';
-- UPDATE categories SET name = 'TOP 10' WHERE id = '7e27d7bb-258d-4df5-810d-9b5c3146a606';



-- -- Update subcategories table with Uzbek translations
-- UPDATE subcategories SET name = 'Oʻzbekiston' WHERE id = '';
-- UPDATE subcategories SET name = 'Siyosat' WHERE id = '10c5f9f4-9b32-4e9a-9b91-ff4f38cf1d62';
-- UPDATE subcategories SET name = 'Jamiyat' WHERE id = 'bfb0cda3-d527-4ea7-8b8b-08f36e8a84d4';
-- UPDATE subcategories SET name = 'Voqealar' WHERE id = '0ec7c541-3b4c-46a1-8e29-6c5014c7c84b';
-- UPDATE subcategories SET name = 'Jinoyat' WHERE id = '74678315-cc88-453d-84aa-c56d3a9176fc';
-- UPDATE subcategories SET name = 'Nizolar' WHERE id = 'd6d6d04a-07c6-4666-93b5-979d146dc7b8';
-- UPDATE subcategories SET name = 'Filmlar' WHERE id = '27c3f98f-3ad7-4c8f-9f91-91b2dd9b2271';
-- UPDATE subcategories SET name = 'Teatr' WHERE id = '96cce354-ecb6-47c2-9c09-b7e2e7d185b1';
-- UPDATE subcategories SET name = 'Restoranlar' WHERE id = '9cb6f8c6-4ed6-4b29-9e1f-9780122db0b1';
-- UPDATE subcategories SET name = 'Kontsertlar' WHERE id = 'b7e0b4c1-bf7e-452c-9b89-d5424e02b122';
-- UPDATE subcategories SET name = 'Stand-up' WHERE id = '7c3d4b65-0f3e-4962-9d6b-cf2e1b53ff68';
-- UPDATE subcategories SET name = 'Parklar' WHERE id = '99a4b6a8-d51e-4b83-92f6-246a2a2a4f60';
-- UPDATE subcategories SET name = 'Koʻrgazmalar' WHERE id = '4cb2e539-cbbf-4b1f-a4d3-1f7d66330b1f';
-- UPDATE subcategories SET name = 'Kolleksiyalar' WHERE id = 'f4c04dbe-6bfa-4e92-b07e-c24dce0ea5a4';
-- UPDATE subcategories SET name = 'Bolalar' WHERE id = 'ce0b3c85-27f6-4a4d-9b8b-c71720f8e724';
-- UPDATE subcategories SET name = 'Festival' WHERE id = 'b56b69b0-505d-47f6-bc7d-7b2f84fa0c15';
-- UPDATE subcategories SET name = 'Sport' WHERE id = '745dce4b-5cb2-4868-915e-b2b06cb5d31a';
-- UPDATE subcategories SET name = 'Partiyalar' WHERE id = 'a55fd728-4d44-4c7e-8311-bf2c8e5b8e67';
-- UPDATE subcategories SET name = 'Film kutubxonasi' WHERE id = 'ea4931d8-7d38-4c4f-83d7-5c5d3641f042';
-- UPDATE subcategories SET name = 'Qoʻllanmalar' WHERE id = '9785b78b-9e9c-4ef5-9b6c-fbf64ae5f29b';
-- UPDATE subcategories SET name = 'Instagram Reels' WHERE id = '8c550a80-817f-4e46-9f35-507f2233fa0b';
-- UPDATE subcategories SET name = 'Futbol' WHERE id = '32b066bc-6c61-46ae-9fa7-4e3a7e6d36d3';
-- UPDATE subcategories SET name = 'Xokkey' WHERE id = '7ac905e4-37b6-4c18-96d6-f9304c1c72e0';
-- UPDATE subcategories SET name = 'Boks/MMA' WHERE id = 'b9d7a6de-d7a8-4a54-b44e-0266b70204c5';
-- UPDATE subcategories SET name = 'Motorsport' WHERE id = '1c2b5e12-5b5b-40f8-9eb7-d244f4473e8f';
-- UPDATE subcategories SET name = 'Tennis' WHERE id = 'e290bd87-92d8-4c7e-905e-5015b535dbf2';
-- UPDATE subcategories SET name = 'Basketbol' WHERE id = '4e00a79b-982a-407e-b83e-9e43e8aef3f2';
-- UPDATE subcategories SET name = 'Figurali uchish' WHERE id = '5079c95f-73d3-4c62-8b27-2045292efc4e';
-- UPDATE subcategories SET name = 'Kibersport' WHERE id = '6a1f0234-fbc3-4b7c-8140-750ab49f71ec';
-- UPDATE subcategories SET name = 'Shaxmat' WHERE id = '6b9479d2-67d3-4f4a-b3f2-4a61ff5466c2';
-- UPDATE subcategories SET name = 'Yozgi sport' WHERE id = '4c9e3912-56a2-47d2-9a2a-daf4349eac5f';
-- UPDATE subcategories SET name = 'Qishki sport' WHERE id = '33d43e0c-0452-4744-9c26-560fb60619c7';
-- UPDATE subcategories SET name = 'Oʻzbekiston' WHERE id = '743a8c3d-7a7e-44ea-a6b0-36b76ad908b1';
-- UPDATE subcategories SET name = 'Chempionat tasvirlari' WHERE id = '02951c36-2ec5-4c46-bc65-3a344e7275a8';
-- UPDATE subcategories SET name = 'Avto' WHERE id = '2eeb5d0e-3487-4b96-873c-572273a6a50a';
-- UPDATE subcategories SET name = 'Fan' WHERE id = 'e949e339-d060-4d06-bb38-06b06970a3f8';
-- UPDATE subcategories SET name = 'Kosmos' WHERE id = '8c4d0875-91d3-4e46-8c36-d76d34d960b5';
-- UPDATE subcategories SET name = 'Qurollar' WHERE id = 'a0581f69-51e1-48ef-b7fc-8238e7b7a4e8';
-- UPDATE subcategories SET name = 'Tarix' WHERE id = '9f6d3b39-4c5c-4202-8923-cb3db238d8dc';
-- UPDATE subcategories SET name = 'Salomatlik' WHERE id = '0308c01b-bc22-4a6f-8927-ea2b79289b51';
-- UPDATE subcategories SET name = 'Texnikalar' WHERE id = '564deef4-d9b6-4c6f-a14d-79d76a78c9d8';
-- UPDATE subcategories SET name = 'Gadjetlar' WHERE id = 'aa935589-46b0-45e3-8062-5a7b0e52fc29';
-- UPDATE subcategories SET name = 'Oʻzbekiston' WHERE id = 'b2f52b78-e4ac-44d6-a4b2-3e372d2a205f';
-- UPDATE subcategories SET name = 'Iqtisodiyot' WHERE id = 'efb450b5-9ac4-42e7-a7c0-1a8b2c1e12d0';
-- UPDATE subcategories SET name = 'Kompaniyalar' WHERE id = 'cbb7a1d5-593d-4f6e-a64d-bdd8db3e5e6d';
-- UPDATE subcategories SET name = 'Shaxsiy hisob' WHERE id = 'fa8c0e93-8a16-4e2b-bbab-cf2e0e1e81d3';
-- UPDATE subcategories SET name = 'Koʻchmas mulk' WHERE id = 'a1db5e43-56c5-4a3b-91a4-69959e4e7d6a';
-- UPDATE subcategories SET name = 'Import oʻrnini bosish' WHERE id = 'bdad84c3-9e4f-47ec-8b5f-97ad9b5d7db4';
-- UPDATE subcategories SET name = 'Shahar muhiti' WHERE id = 'd01e30d1-48f2-47b7-a44c-8e0dba897a3e';
-- UPDATE subcategories SET name = 'Biznes iqlimi' WHERE id = 'f7d7c9d5-5b71-4994-8b47-8d705e5e11b5';
-- UPDATE subcategories SET name = 'Oʻzbekiston' WHERE id = 'fc88f84e-cfb5-45e0-bb1e-5b90b4a7df51';
-- UPDATE subcategories SET name = 'Shou-biznes' WHERE id = 'a0d44f37-8e54-47d7-82d0-51a7ec728f37';
-- UPDATE subcategories SET name = 'Harbiy yangiliklar' WHERE id = '68a4c6d2-facb-40d2-8c6b-9c9b9f3c40c1';
-- UPDATE subcategories SET name = 'Oziq-ovqat' WHERE id = '0a281039-8711-4f0e-9a4f-e4c2a9de003f';
-- UPDATE subcategories SET name = 'Psixologiya' WHERE id = '6f1e61a8-8f4b-47f1-88c7-935e611ed8f0';
-- UPDATE subcategories SET name = 'Trendlari' WHERE id = 'd97b0f91-1d61-464f-8d8b-c32a63e827f6';
-- UPDATE subcategories SET name = 'Bolalar' WHERE id = 'b2e095eb-bd51-4a61-89f0-d54d4a00c1d2';
-- UPDATE subcategories SET name = 'Uy va bogʻ' WHERE id = '0a9b0690-7031-4860-b605-9900df22868f';
-- UPDATE subcategories SET name = 'Salomatlik' WHERE id = '3e8f3066-b36c-4a34-8d89-5e4b9c5b37d1';
-- UPDATE subcategories SET name = 'Karyera' WHERE id = 'f0b7e501-72b1-4f2d-8fae-4986b9e59715';
-- UPDATE subcategories SET name = 'Sayohat' WHERE id = 'fc5e11d2-64b6-42ba-8f0a-b9a9c8e66c42';
-- UPDATE subcategories SET name = 'Iqtisodiyot' WHERE id = '6b3b11b1-4b16-452b-a03d-1f7e9a61969b';
-- UPDATE subcategories SET name = 'Mashinalar' WHERE id = '4fc0479f-fd48-4e83-b3ba-b3df4d91c6da';
-- UPDATE subcategories SET name = 'Sport' WHERE id = 'c38327f6-5492-4d56-8d69-efb6ab93ae32';
-- UPDATE subcategories SET name = 'Oʻyinlar' WHERE id = 'cfd2e8e4-3bfb-4f16-8554-f54ea865a2dc';
-- UPDATE subcategories SET name = 'Koʻngilochar' WHERE id = 'e45888f6-d0d7-48a4-9184-d27f67d0b7a4';
-- UPDATE subcategories SET name = 'TOP 10' WHERE id = 'dfbdbf6e-9c93-44bc-a302-cfc1ac16e75b';
