CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY NOT NULL,
  username VARCHAR(255) NOT NULL DEFAULT '',
  first_name VARCHAR(255) NOT NULL CHECK (first_name <> ''),
  last_name VARCHAR(255) DEFAULT '',
  admin BOOLEAN DEFAULT FALSE,
  phone VARCHAR(12) CHECK (CHAR_LENGTH(phone) = 12),
  action VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);