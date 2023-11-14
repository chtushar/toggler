CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Organization table
CREATE TABLE organizations (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- User table
CREATE TABLE users (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    active BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
-- Organization Members table
CREATE TABLE organization_members (
    user_id INT REFERENCES users(id),
    org_id INT REFERENCES organizations(id),
    UNIQUE(user_id, org_id)
);
-- API Keys table
CREATE TABLE api_keys (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    api_key TEXT NOT NULL,
    allowed_domains VARCHAR [] NOT NULL,
    org_id INT REFERENCES organizations(id),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Folder table
CREATE TABLE folders (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    org_id INT REFERENCES organizations(id),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Environment table
CREATE TABLE environments (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    color CHAR(7) DEFAULT NULL,
    org_id INT REFERENCES organizations(id),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Flags Group table
CREATE TABLE flags_groups (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    org_id INT REFERENCES organizations(id),
    folder_id INT REFERENCES folders(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(
        uuid,
        org_id,
        folder_id
    )
);
-- Flags Group State table
CREATE TABLE flags_group_states (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    flags_group_id INT REFERENCES flags_groups(id),
    js TEXT DEFAULT NULL,
    environment_id INT REFERENCES environments(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(
        flags_group_id,
        environment_id
    )
);