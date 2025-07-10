-- name: PingDB :one
SELECT 1;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ,
    deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    card_id VARCHAR(100) UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_id INTEGER NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    active BOOLEAN DEFAULT TRUE,
    type VARCHAR(32) NOT NULL DEFAULT 'balance',
    assigned_at TIMESTAMPTZ DEFAULT now(),
     deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE balances (
    id SERIAL PRIMARY KEY,
    user_id INTEGER  REFERENCES users(id) ON DELETE CASCADE,
    card_id INTEGER UNIQUE REFERENCES cards(id) NOT NULL,
    ride_cost NUMERIC(10,2) NOT NULL,
    balance NUMERIC(10,2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT unique_user_card UNIQUE (user_id, card_id)
);

CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    card_id INTEGER REFERENCES cards(id) ON DELETE SET NULL,
    device_id VARCHAR(100) NOT NULL,
    floor INTEGER,
    timestamp TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE charges (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    amount NUMERIC(10,2) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('topup', 'ride')),
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    device_id varchar(255) UNIQUE NOT NULL,
    location VARCHAR(255),
    installed_at TIMESTAMPTZ DEFAULT now(),
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS card_activations (
    id SERIAL PRIMARY KEY,
    card_id INTEGER NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    activation_start DATE NOT NULL,
    activation_end DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- AUTH USERS TABLE FOR AUTHENTICATION
CREATE TABLE IF NOT EXISTS auth_users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'customer')),
    device_id varchar(255) REFERENCES devices(device_id) ON DELETE CASCADE, 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE INDEX idx_cards_card_id ON cards(card_id);
CREATE INDEX idx_trips_user_id ON trips(user_id);
CREATE INDEX idx_charges_user_id ON charges(user_id);
