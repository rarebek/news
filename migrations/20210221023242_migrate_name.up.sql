CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcategories (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id),
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS news (
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcategory_news (
    subcategory_id INT REFERENCES subcategories(id),
    news_id UUID REFERENCES news(id),
    PRIMARY KEY (subcategory_id, news_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY,
    phone_number TEXT,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS superadmins (
    id UUID PRIMARY KEY,
    phone_number TEXT,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INSERT INTO superadmins (id, phone_number, password) VALUES('acc98ad0-43a1-4ac5-ba90-7dc1f1a34d1e', 'test', 'test');

INSERT INTO categories(name)
VALUES
    ('Uzbekistan'),
    ('In the world'),
    ('Entertainment'),
    ('LIVE'),
    ('Sport'),
    ('Auto'),
    ('Technologies'),
    ('Economics'),
    ('Show businesses'),
    ('Military news'),
    ('Daily'),
    ('TOP10');

INSERT INTO subcategories(category_id, name)
VALUES
    (2, 'Politics'),
    (2, 'Community'),
    (2, 'Incidents'),
    (2, 'Crimes'),
    (2, 'Conflicts'),
    (3, 'MOVIE'),
    (3, 'Restaurants'),
    (3, 'Concerts'), 
    (3, 'StandUp'),
    (3, 'Parks'),
    (3, 'Exhibitions'),
    (3, 'Collections'),
    (3, 'Children'),
    (3, 'Festivals'),
    (3, 'Parties'),
    (3, 'Film library'),
    (3, 'Guide'),
    (5, 'Football'),
    (5, 'Hockey'),
    (5, 'Box, MMA'),
    (5, 'Autosport'),
    (5, 'Tennis'),
    (5, 'Basketball'),
    (5, 'Figure skating'),
    (5, 'Cybersport'),
    (5, 'Chess'),
    (5, 'Summer sports'),
    (5, 'Winter sports'),
    (5, 'Uzbekistan'),
    (7, 'Science'),
    (7, 'Cosmos'),
    (7, 'Weapon'),
    (7, 'History'),
    (7, 'Health'),
    (7, 'Technique'),
    (7, 'Gadgets'),
    (7, 'Uzbekistan'),
    (8, 'Economics'),
    (8, 'Companies'), 
    (8, 'Personal account'),
    (8, 'Real estate'),
    (8, 'Import substitution'),
    (8, 'Business climate'),
    (8, 'Uzbekistan'),
    (11, 'Food'),
    (11, 'Psychology'),
    (11, 'Trends'),
    (11, 'Children'),
    (11, 'Home and garden'),
    (11, 'Events'),
    (11, 'Scandals');
