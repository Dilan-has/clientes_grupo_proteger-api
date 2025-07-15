USE grupo_proteger;
CREATE TABLE legal_rep (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    cc VARCHAR(20) UNIQUE
);
CREATE TABLE clients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    nit VARCHAR(20) UNIQUE,
    address VARCHAR(150),
    phone VARCHAR(20),
    email VARCHAR(100),
    id_rep INT,
    arl VARCHAR(100),
    FOREIGN KEY (id_rep) REFERENCES legal_rep(id)
);
CREATE TABLE affiliates (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    cc VARCHAR(20) UNIQUE,
    eps VARCHAR(100),
    status ENUM('ACTIVE', 'RETIRED') DEFAULT 'ACTIVE',
    id_client INT,
    pension VARCHAR(100),
    risk VARCHAR(100),
    birthdate DATE,
    caja VARCHAR(100),
    income DECIMAL(10, 2),
    FOREIGN KEY (id_client) REFERENCES clients(id)
);
CREATE TABLE credentials (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_client INT,
    organization VARCHAR(100),
    user VARCHAR(50),
    pass VARCHAR(255),
    FOREIGN KEY (id_client) REFERENCES clients(id)
);
CREATE TABLE user (
    user VARCHAR(50) PRIMARY KEY,
    password VARCHAR(255) -- Asumiendo que se almacenan contraseñas en hash
);

