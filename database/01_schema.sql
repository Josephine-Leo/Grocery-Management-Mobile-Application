DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS points_transactions;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS reward_claims;
DROP TABLE IF EXISTS rewards;
DROP TABLE IF EXISTS order_lines;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS fridge_entries;
DROP TABLE IF EXISTS payments;

enum delivery_option {
    delivery
    pickup
};

enum order_status {
    placed
    paid
    packed
    out_for_delivery
    ready_for_pickup
    completed
    cancelled
};

enum fridge_status {
  in_fridge
  used
  expired
  discarded
};

enum points_reason {
  earned
  redeemed
  adjustment
};

enum reward_status {
  claimed
  used
  expired
};

enum payment_status {
    pending
    succeeded
    failed
    refunded
};

CREATE TABLE users (
    user_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    phone_number varchar(20) NOT NULL,
    address varchar(255) NOT NULL,
    auth_uid varchar(255) NOT NULL,
    points integer NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items (
    item_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    sub_cat varchar(255),
    price decimal(10, 2) NOT NULL,
    discount varchar(255),
    image_url varchar(255),
    rating varchar(10),
    title varchar(255) NOT NULL,
    feature varchar(255),
    description text,
);

CREATE TABLE cart_items (
    cart_item_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    item_id integer NOT NULL REFERENCES items(item_id) ON DELETE CASCADE,
    quantity integer NOT NULL,
    added_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    remark varchar(255)
);

CREATE TABLE orders (
    order_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    total_cost decimal(10, 2) NOT NULL,
    status order_status NOT NULL,
    order_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delivery_address varchar(255),
    delivery_option delivery_option NOT NULL,
    delivery_fee decimal(10, 2) NOT NULL DEFAULT 0.00,
    points_earned integer NOT NULL DEFAULT 0,
    completed_at timestamp
);

CREATE TABLE order_lines (
    order_line_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_id integer NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    item_id integer NOT NULL REFERENCES items(item_id) ON DELETE CASCADE,
    quantity integer NOT NULL,
    unit_price decimal(10, 2) NOT NULL,
    remark varchar(255)
);

CREATE TABLE fridge_entries (
    entry_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_line_id integer NOT NULL REFERENCES order_lines(order_line_id) ON DELETE CASCADE,
    quantity_remaining decimal(10,2) NOT NULL,
    expires_at timestamp,
    status fridge_status NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
    payment_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_id integer NOT NULL REFEERENCES orders(order_id) ON DELETE CASCADE,
    provider varchar(255) NOT NULL,
    provider_payment_id varchar(255) NOT NULL,
    amount decimal(10, 2) NOT NULL,
    status payment_status NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    paid_at timestamp
);

CREATE TABLE rewards (
    reward_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(255) NOT NULL,
    description varchar(255),
    points_cost integer NOT NULL,
    valid_until timestamp
);

CREATE TABLE reward_claims (
    claim_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    reward_id integer NOT NULL REFERENCES rewards(reward_id) ON DELETE CASCADE,
    status reward_status NOT NULL,
    claimed_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    used_at timestamp,
    expires_at timestamp,
    order_id integer REFERENCES orders(order_id) ON DELETE SET NULL
);

CREATE TABLE points_transactions (
    txn_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    order_id integer REFERENCES orders(order_id) ON DELETE CASCADE,
    claim_id integer REFERENCES reward_claims(claim_id) ON DELETE CASCADE,
    change integer NOT NULL,
    reason points_reason NOT NULL,
    created_at timestamp NOT NULL
);