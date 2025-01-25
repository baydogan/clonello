CREATE TABLE users (
                       username VARCHAR PRIMARY KEY,
                       hashed_password VARCHAR NOT NULL,
                       full_name VARCHAR NOT NULL,
                       email VARCHAR UNIQUE NOT NULL,
                       password_changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);