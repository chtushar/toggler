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
    color VARCHAR(255) DEFAULT NULL,
    org_id INT REFERENCES organizations(id),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Flags Group State table
CREATE TABLE flags_group_states (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    version INT,
    code TEXT DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW()
);
-- Flags Group table
CREATE TABLE flags_groups (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    org_id INT REFERENCES organizations(id),
    folder_id INT REFERENCES folders(id),
    current_version INT REFERENCES flags_group_states(id),
    environment_id INT REFERENCES environments(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(
        uuid,
        org_id,
        folder_id,
        environment_id
    )
);