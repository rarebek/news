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
('b1d7357d-1d92-4e77-a8d0-1d394c5b2ef2', 'Uzbekistan'),

-- World
('09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'World'),

-- LIVE
('f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Entertainment'),

-- Sports
('9d57d482-6db2-43ec-8513-6478c066aa51', 'Sports'),

-- Auto
('23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3', 'Auto'),

-- Technology
('6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Technology'),

-- Economy
('e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Economy'),

-- Show Business
('72a5ec87-44ed-438c-b9d0-0e391ed7da4d', 'Show Business'),

-- Military News
('b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69', 'Military News'),

-- Daily
('a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Daily'),

-- TOP 10
('7e27d7bb-258d-4df5-810d-9b5c3146a606', 'TOP 10');


-- Insert into subcategories table
-- Insert Subcategories
INSERT INTO subcategories (id, category_id, name) VALUES
-- Uzbekistan
('cfc38315-9a1e-4f38-8f4e-63196e5e4b4d', 'b1d7357d-1d92-4e77-a8d0-1d394c5b2ef2', 'Uzbekistan'),

-- World
('10c5f9f4-9b32-4e9a-9b91-ff4f38cf1d62', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Politics'),
('bfb0cda3-d527-4ea7-8b8b-08f36e8a84d4', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Society'),
('0ec7c541-3b4c-46a1-8e29-6c5014c7c84b', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Incidents'),
('74678315-cc88-453d-84aa-c56d3a9176fc', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Crime'),
('d6d6d04a-07c6-4666-93b5-979d146dc7b8', '09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d', 'Conflicts'),

-- Entertainment
('27c3f98f-3ad7-4c8f-9f91-91b2dd9b2271', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Movies'),
('96cce354-ecb6-47c2-9c09-b7e2e7d185b1', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Theater'),
('9cb6f8c6-4ed6-4b29-9e1f-9780122db0b1', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Restaurants'),
('b7e0b4c1-bf7e-452c-9b89-d5424e02b122', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Concerts'),
('7c3d4b65-0f3e-4962-9d6b-cf2e1b53ff68', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Stand-up'),
('99a4b6a8-d51e-4b83-92f6-246a2a2a4f60', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Parks'),
('4cb2e539-cbbf-4b1f-a4d3-1f7d66330b1f', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Exhibitions'),
('f4c04dbe-6bfa-4e92-b07e-c24dce0ea5a4', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Collections'),
('ce0b3c85-27f6-4a4d-9b8b-c71720f8e724', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Kids'),
('b56b69b0-505d-47f6-bc7d-7b2f84fa0c15', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Festivals'),
('745dce4b-5cb2-4868-915e-b2b06cb5d31a', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Sports'),
('a55fd728-4d44-4c7e-8311-bf2c8e5b8e67', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Parties'),
('ea4931d8-7d38-4c4f-83d7-5c5d3641f042', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Film Library'),
('9785b78b-9e9c-4ef5-9b6c-fbf64ae5f29b', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Guides'),

-- LIVE
('8c550a80-817f-4e46-9f35-507f2233fa0b', 'f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1', 'Instagram Reels'),

-- Sports
('32b066bc-6c61-46ae-9fa7-4e3a7e6d36d3', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Football'),
('7ac905e4-37b6-4c18-96d6-f9304c1c72e0', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Hockey'),
('b9d7a6de-d7a8-4a54-b44e-0266b70204c5', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Boxing/MMA'),
('1c2b5e12-5b5b-40f8-9eb7-d244f4473e8f', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Motorsport'),
('e290bd87-92d8-4c7e-905e-5015b535dbf2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Tennis'),
('4e00a79b-982a-407e-b83e-9e43e8aef3f2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Basketball'),
('5079c95f-73d3-4c62-8b27-2045292efc4e', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Figure Skating'),
('6a1f0234-fbc3-4b7c-8140-750ab49f71ec', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Esports'),
('6b9479d2-67d3-4f4a-b3f2-4a61ff5466c2', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Chess'),
('4c9e3912-56a2-47d2-9a2a-daf4349eac5f', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Summer Sports'),
('33d43e0c-0452-4744-9c26-560fb60619c7', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Winter Sports'),
('743a8c3d-7a7e-44ea-a6b0-36b76ad908b1', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Uzbekistan'),
('02951c36-2ec5-4c46-bc65-3a344e7275a8', '9d57d482-6db2-43ec-8513-6478c066aa51', 'Championship Images'),

-- Auto
('2eeb5d0e-3487-4b96-873c-572273a6a50a', '23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3', 'Auto'),

-- Technology
('e949e339-d060-4d06-bb38-06b06970a3f8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Science'),
('8c4d0875-91d3-4e46-8c36-d76d34d960b5', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Space'),
('a0581f69-51e1-48ef-b7fc-8238e7b7a4e8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Weapons'),
('9f6d3b39-4c5c-4202-8923-cb3db238d8dc', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'History'),
('0308c01b-bc22-4a6f-8927-ea2b79289b51', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Health'),
('564deef4-d9b6-4c6f-a14d-79d76a78c9d8', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Techniques'),
('aa935589-46b0-45e3-8062-5a7b0e52fc29', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Gadgets'),
('b2f52b78-e4ac-44d6-a4b2-3e372d2a205f', '6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb', 'Uzbekistan'),

-- Economy
('efb450b5-9ac4-42e7-a7c0-1a8b2c1e12d0', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Economy'),
('cbb7a1d5-593d-4f6e-a64d-bdd8db3e5e6d', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Companies'),
('fa8c0e93-8a16-4e2b-bbab-cf2e0e1e81d3', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Personal Account'),
('a1db5e43-56c5-4a3b-91a4-69959e4e7d6a', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Real Estate'),
('bdad84c3-9e4f-47ec-8b5f-97ad9b5d7db4', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Import Substitution'),
('d01e30d1-48f2-47b7-a44c-8e0dba897a3e', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Urban Environment'),
('f7d7c9d5-5b71-4994-8b47-8d705e5e11b5', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Business Climate'),
('fc88f84e-cfb5-45e0-bb1e-5b90b4a7df51', 'e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e', 'Uzbekistan'),

-- Show Business
('a0d44f37-8e54-47d7-82d0-51a7ec728f37', '72a5ec87-44ed-438c-b9d0-0e391ed7da4d', 'Show Business'),

-- Military News
('68a4c6d2-facb-40d2-8c6b-9c9b9f3c40c1', 'b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69', 'Military News'),

-- Daily
('0a281039-8711-4f0e-9a4f-e4c2a9de003f', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Food'),
('6f1e61a8-8f4b-47f1-88c7-935e611ed8f0', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Psychology'),
('d97b0f91-1d61-464f-8d8b-c32a63e827f6', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Trends'),
('b2e095eb-bd51-4a61-89f0-d54d4a00c1d2', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Kids'),
('0a9b0690-7031-4860-b605-9900df22868f', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Home and Garden'),
('3e8f3066-b36c-4a34-8d89-5e4b9c5b37d1', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Health'),
('f0b7e501-72b1-4f2d-8fae-4986b9e59715', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Career'),
('fc5e11d2-64b6-42ba-8f0a-b9a9c8e66c42', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Travel'),
('6b3b11b1-4b16-452b-a03d-1f7e9a61969b', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Economy'),
('4fc0479f-fd48-4e83-b3ba-b3df4d91c6da', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Cars'),
('c38327f6-5492-4d56-8d69-efb6ab93ae32', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Sport'),
('cfd2e8e4-3bfb-4f16-8554-f54ea865a2dc', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Games'),
('e45888f6-d0d7-48a4-9184-d27f67d0b7a4', 'a1e67972-b9b0-4eec-a51d-8fa8d08e51bb', 'Entertainment'),

-- TOP 10
('dfbdbf6e-9c93-44bc-a302-cfc1ac16e75b', '7e27d7bb-258d-4df5-810d-9b5c3146a606', 'TOP 10');
