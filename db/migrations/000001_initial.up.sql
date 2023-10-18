CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Organization table
CREATE TABLE organizations (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- User table
CREATE TABLE users (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
-- Organization Members table
CREATE TABLE organization_members (
    user_uuid UUID REFERENCES users(uuid),
    org_uuid UUID REFERENCES organizations(uuid),
    UNIQUE(user_uuid, org_uuid)
);
-- Folder table
CREATE TABLE folders (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organizations(uuid),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Environment table 
CREATE TABLE environments (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organizations(uuid),
    created_at TIMESTAMP DEFAULT NOW()
);
-- Flags Group State table
CREATE TABLE flags_group_states (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    version INT,
    code TEXT DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW()
);
-- Flags Group table
CREATE TABLE flags_groups (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    org_uuid UUID REFERENCES organizations(uuid),
    folder_uuid UUID REFERENCES folders(uuid),
    current_version UUID REFERENCES flags_group_states(uuid),
    environment_uuid UUID REFERENCES environments(uuid),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(
        uuid,
        org_uuid,
        folder_uuid,
        environment_uuid
    )
);