CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- or SERIAL if you want int64
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

