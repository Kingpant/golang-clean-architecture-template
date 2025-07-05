CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT now ()
    );

--bun:split

CREATE TYPE user_log_action AS ENUM (
    'CREATE',
    'UPDATE'
);

CREATE TABLE
    user_logs (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        action user_log_action NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now ()
    );