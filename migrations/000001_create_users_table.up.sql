CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
