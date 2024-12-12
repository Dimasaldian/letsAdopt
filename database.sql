-- Active: 1732702326143@@127.0.0.1@5432@letsadoptdb
-- Membuat tabel Users
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT 'user' CHECK (role IN ('admin', 'user')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel Pets (Tabel untuk hewan peliharaan)
CREATE TABLE Pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('dog', 'cat', 'bird', 'other')),
    breed VARCHAR(100),
    age INT,
    description TEXT,
    vaccinated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel Adoptions
CREATE TABLE Adoptions (
    id_adopt SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    pet_id INT NOT NULL,
    reason TEXT NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
    notification_sent BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (pet_id) REFERENCES Pets(id) ON DELETE CASCADE
);

-- Membuat tabel Admins
CREATE TABLE Admins (
    admin_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    privileges TEXT,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- Membuat tabel Notifications
CREATE TABLE Notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    adoption_id INT NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (adoption_id) REFERENCES Adoptions(id_adopt) ON DELETE CASCADE
);

-- Update tabel Pets dengan kolom alamat
ALTER TABLE Pets
ADD COLUMN addresses VARCHAR(255);

-- Membuat tabel Addresses
CREATE TABLE Addresses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,
    pet_id INT REFERENCES Pets(id) ON DELETE CASCADE,
    street VARCHAR(100),
    city VARCHAR(50),
    state VARCHAR(50),
    postal_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel pet_images
CREATE TABLE pet_images (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES Pets(id) ON DELETE CASCADE
);

-- Membuat index untuk pet_id di tabel pet_images
CREATE INDEX idx_pet_id ON pet_images(pet_id);


-- Insert dummy data into Users
INSERT INTO Users (name, email, password, role) VALUES
('Alice Smith', 'alice@example.com', 'password123', 'user'),
('Bob Johnson', 'bob@example.com', 'password456', 'admin'),
('Charlie Brown', 'charlie@example.com', 'password789', 'user');

-- Insert dummy data into Pets
INSERT INTO Pets (name, type, breed, age, description, vaccinated) VALUES
('Dam', 'cat', 'Golden Retriever', 5, 'Friendly and energetic cat.', TRUE);

-- Insert dummy data into Adoptions
INSERT INTO Adoptions (user_id, pet_id, reason, status) VALUES
(1, 1, 'Looking for a companion.', 'approved'),
(2, 2, 'I want a low-maintenance pet.', 'pending'),
(3, 3, 'I love birds and want to take care of them.', 'approved');

-- Insert dummy data into Admins
INSERT INTO Admins (user_id, privileges) VALUES
(2, 'Manage users, approve adoptions');

-- Insert dummy data into Notifications
INSERT INTO Notifications (user_id, adoption_id, message) VALUES
(1, 1, 'Your adoption request for Rex has been approved.'),
(2, 2, 'Your adoption request for Whiskers is pending.'),
(3, 3, 'Your adoption request for Tweety has been approved.');

-- Insert dummy data into Addresses
INSERT INTO Addresses (user_id, pet_id, street, city, state, postal_code) VALUES
(1, 1, '123 Maple St', 'Springfield', 'Illinois', '62701'),
(2, 2, '456 Oak Ave', 'Chicago', 'Illinois', '60601'),
(3, 3, '789 Pine Rd', 'Madison', 'Wisconsin', '53703');

-- Insert dummy data into pet_images
INSERT INTO pet_images (pet_id, url) VALUES
(1, 'https://example.com/images/rex.jpg'),
(2, 'https://example.com/images/whiskers.jpg'),
(3, 'https://example.com/images/tweety.jpg');

ALTER TABLE notifications ALTER COLUMN adoption_id TYPE INT USING adoption_id::integer;

ALTER TABLE notifications DROP CONSTRAINT notifications_adoption_id_fkey;
ALTER TABLE notifications ADD CONSTRAINT notifications_adoption_id_fkey FOREIGN KEY (adoption_id) REFERENCES adoptions(id_adopt) ON DELETE CASCADE;