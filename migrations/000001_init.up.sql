CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE subscription_status AS ENUM ('active', 'expired');
CREATE TYPE payment_status AS ENUM ('success', 'failed');

CREATE TABLE subscriptions (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    service_name TEXT        NOT NULL,
    price        INTEGER     NOT NULL CHECK (price > 0),
    user_id      UUID        NOT NULL,
    start_date   TIMESTAMPTZ NOT NULL,
    end_date     TIMESTAMPTZ NOT NULL,
    auto_renew   BOOLEAN     NOT NULL DEFAULT false,
    status       subscription_status NOT NULL DEFAULT 'active',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions (user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions (status);
CREATE INDEX idx_subscriptions_end_date ON subscriptions (end_date);

CREATE TABLE payments (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    subscription_id UUID           NOT NULL REFERENCES subscriptions (id) ON DELETE CASCADE,
    amount          INTEGER        NOT NULL CHECK (amount > 0),
    status          payment_status NOT NULL DEFAULT 'success',
    paid_at         TIMESTAMPTZ    NOT NULL DEFAULT now()
);

CREATE INDEX idx_payments_subscription_id ON payments (subscription_id);
