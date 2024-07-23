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

-- INSERT INTO superadmins (id, username, password) VALUES('acc98ad0-43a1-4ac5-ba90-7dc1f1a34d1e', 'test', 'test');

-- Insert into categories table
INSERT INTO categories(id, name)
VALUES
    ('a5b7d3b6-81f7-4d63-8d3e-5e4d6e5d8b7d', 'Uzbekistan'),
    ('c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'In the world'),
    ('d7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Entertainment'),
    ('e8d7c6b5-4a3d-4d6e-89f6-4e5d6e7d8c9e', 'LIVE'),
    ('f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Sport'),
    ('a9d8e7c6-8e4d-4a6b-9f2e-4d6e7c8d9e7f', 'Auto'),
    ('b7e8c6d5-9d4a-4b6e-87f9-4e5d6e8c9d7f', 'Technologies'),
    ('c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Economics'),
    ('d5e8c6b7-9d4a-4e6f-87e9-4d6e7c8d9f7e', 'Show businesses'),
    ('e6d8c7b9-4e8a-4d6e-89f7-4d5e8c7d9e7f', 'Military news'),
    ('f7e8c9d5-4a6d-4e8b-89f7-4d6e8c7d9e7f', 'Daily'),
    ('a8d9c6e7-4e8b-4d6e-89f7-4d5e8c7d9e7f', 'TOP10');


-- Insert into subcategories table
INSERT INTO subcategories(id, category_id, name)
VALUES
    ('b9e7c8d6-4a5e-4d6f-8e7d-4d5e6c8d9f7a', 'c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'Politics'),
    ('c6d5e8b7-4a6f-4d8e-9e7d-4d5e6c8d9f7b', 'c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'Community'),
    ('d7e8c6b5-4a8d-4e9e-9f7d-4d5e6c8d9f7c', 'c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'Incidents'),
    ('e8d7c6b5-4a9e-4d6f-8e7d-4d5e6c8d9f7d', 'c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'Crimes'),
    ('f9e8d7c6-4a6b-4d9e-8f7d-4d5e6c8d9f7e', 'c68d6e8e-2a6d-4a5d-918e-4e5d6e6d7a8e', 'Conflicts'),
    ('a9d8e7c6-4b9e-4d8f-8e7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'MOVIE'),
    ('b7e8c6d5-4a9e-4d8e-8f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Restaurants'),
    ('c6d7e8b9-4d8e-4a6f-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Concerts'),
    ('d5e8c6b7-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'StandUp'),
    ('e6d8c7b9-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Parks'),
    ('f7e8c9d5-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Exhibitions'),
    ('a8d9c6e7-4b9e-4d8e-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Collections'),
    ('b9e7c8d6-4a5e-4d6f-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Children'),
    ('c6d5e8b7-4a6f-4d8e-8f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Festivals'),
    ('d7e8c6b5-4a9e-4d8f-9f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Parties'),
    ('e8d7c6b5-4a9e-4d8e-8f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Film library'),
    ('f9e8d7c6-4a9e-4d8e-8f7d-4d5e6c8d9f7f', 'd7e8c7f1-7c9e-4b3a-91b2-4d6f7e8c9a9d', 'Guide'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Football'),
    ('b7e8c6d5-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Hockey'),
    ('c6d7e8b9-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Box, MMA'),
    ('d5e8c6b7-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Autosport'),
    ('e6d8c7b9-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Tennis'),
    ('f7e8c9d5-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Basketball'),
    ('a8d9c6e7-4b9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Figure skating'),
    ('b9e7c8d6-4a5e-4d6f-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Cybersport'),
    ('c6d5e8b7-4a6f-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Chess'),
    ('d7e8c6b5-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Summer sports'),
    ('e8d7c6b5-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Winter sports'),
    ('f9e8d7c6-4a9e-4d8e-9f7d-4d5e6c8d9f7f', 'f9e8d7c6-9b2e-4a5d-87e4-4e5d6e8d9c7f', 'Uzbekistan'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9f7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Science'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Cosmos'),
    ('c6d7e8b9-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Weapon'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'History'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Health'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Technique'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9f7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Gadgets'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4e5d6e8c9d7f', 'c6d7e8b9-4e8a-4b6d-89f7-4d5e6c7e8f9a', 'Uzbekistan'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c7e8f9a', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Economics'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Companies'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Personal account'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Real estate'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Import substitution'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Business climate'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Accounting'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Management'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Finance'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Stock exchange'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Mergers and acquisitions'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Forbes'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Rising stars'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Investments'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Real estate'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Product'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'IT'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Software'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Services'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Gadgets'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Digital tools'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'PC components'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Mobile phones'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Mobile apps'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Devices'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Car accessories'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Automobile news'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Cars'),
    ('a9d8e7c6-4b9e-4d8e-9f7d-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'News'),
    ('b7e8c6d5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Hobbies'),
    ('c6d7e8b9-4e8a-4b6d-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Electronics'),
    ('d7e8c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Cars'),
    ('e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Fashion'),
    ('f9e8d7c6-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'e8d7c6b5-4a6d-4b6e-87f9-4d5e6c8d9d7f', 'Beauty')
ON CONFLICT (article_id, tag_id) DO NOTHING;
