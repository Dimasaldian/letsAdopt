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
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE -- Sesuaikan dengan nama kolom di tabel Users
);

-- Membuat tabel Notifications
CREATE TABLE Notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    adoption_id INT NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE, -- Sesuaikan dengan nama kolom di tabel Users
    FOREIGN KEY (adoption_id) REFERENCES Adoptions(id) ON DELETE CASCADE -- Sesuaikan dengan nama kolom di tabel Adoptions
);


-- Index untuk efisiensi query
CREATE INDEX idx_adoptions_status ON Adoptions(status);
CREATE INDEX idx_adoptions_user_id ON Adoptions(user_id);
CREATE INDEX idx_adoptions_pet_id ON Adoptions(pet_id);
