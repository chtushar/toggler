CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Organization table
CREATE TABLE IF NOT EXISTS organization (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
);
-- User table
CREATE TABLE IF NOT EXISTS user (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    org_uuid UUID REFERENCES organization(uuid),
);
-- Folder table
CREATE TABLE IF NOT EXISTS folder (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organization(uuid),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
);
-- Environment table
CREATE TABLE IF NOT EXISTS environment (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organization(uuid),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
);
-- Flags Group State table
CREATE TABLE IF NOT EXISTS flags_group_state (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    version INT,
    code TEXT DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
);
-- Flags Group table
CREATE TABLE IF NOT EXISTS flags_group (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organization(uuid),
    folder_uuid UUID REFERENCES folder(uuid),
    current_version UUID REFERENCES flags_group_state(uuid),
    environment_uuid UUID REFERENCES environment(uuid),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(
        uuid,
        org_uuid,
        folder_uuid,
        current_version,
        environment_uuid
    )
);