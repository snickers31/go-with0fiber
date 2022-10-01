

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    username VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    account_status BOOLEAN  NOT NULL DEFAULT(true),
    isDeleted BOOLEAN  NOT NULL DEFAULT(false),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

-- Add indexes
CREATE INDEX index_user ON users (id) WHERE isDeleted = false;