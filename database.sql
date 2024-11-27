CREATE TABLE Animals (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255),
    jenis VARCHAR(50),
    usia INTEGER,
    jenis_kelamin VARCHAR(10),
    deskripsi TEXT,
    status VARCHAR(50),
    lokasi VARCHAR(255)
);

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    nomor_telepon VARCHAR(15),
    alamat TEXT
);

CREATE TABLE Adoptions (
    id SERIAL PRIMARY KEY,
    hewan_id INT REFERENCES Animals(id),
    pengguna_id INT REFERENCES Users(id),
    tanggal_adopsi DATE
);

CREATE TABLE Release_Adoptions (
    id SERIAL PRIMARY KEY,
    hewan_id INT REFERENCES Animals(id),
    pengguna_id INT REFERENCES Users(id),
    tanggal_lepas DATE
);
