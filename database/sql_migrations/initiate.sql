-- +migrate Up
-- +migrate StatementBegin

create table user_level (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    access          TEXT[] NOT NULL,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_user_level()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_user_level_time BEFORE INSERT ON user_level FOR EACH ROW EXECUTE PROCEDURE update_created_column_user_level();

CREATE OR REPLACE FUNCTION update_modified_column_user_level()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_user_level_time BEFORE UPDATE ON user_level FOR EACH ROW EXECUTE PROCEDURE update_modified_column_user_level();

create table users (
    id              serial PRIMARY KEY,
    username        varchar(256) NOT NULL UNIQUE,
    password        varchar(20) NOT NULL,
    level           INT NOT NULL references user_level(id),
    token           TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_users()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_users_time BEFORE INSERT ON users FOR EACH ROW EXECUTE PROCEDURE update_created_column_users();

CREATE OR REPLACE FUNCTION update_modified_column_users()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_users_time BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column_users();


create table seller (
    id              serial PRIMARY KEY,
    email           varchar(256) NOT NULL UNIQUE,
    name            varchar(256) NOT NULL UNIQUE,
    phone_number    varchar(20) NOT NULL,
    username        varchar(256) NOT NULL UNIQUE references users(username),
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_seller()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_seller_time BEFORE INSERT ON seller FOR EACH ROW EXECUTE PROCEDURE update_created_column_seller();

CREATE OR REPLACE FUNCTION update_modified_column_seller()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_seller_time BEFORE UPDATE ON seller FOR EACH ROW EXECUTE PROCEDURE update_modified_column_seller();

create table brand (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_brand()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_brand_time BEFORE INSERT ON brand FOR EACH ROW EXECUTE PROCEDURE update_created_column_brand();

CREATE OR REPLACE FUNCTION update_modified_column_brand()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_brand_time BEFORE UPDATE ON brand FOR EACH ROW EXECUTE PROCEDURE update_modified_column_brand();

create table treatment_type (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_treatment_type()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_treatment_type_time BEFORE INSERT ON treatment_type FOR EACH ROW EXECUTE PROCEDURE update_created_column_treatment_type();

CREATE OR REPLACE FUNCTION update_modified_column_treatment_type()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_treatment_type_time BEFORE UPDATE ON treatment_type FOR EACH ROW EXECUTE PROCEDURE update_modified_column_treatment_type();

create table treatment (
    id              serial PRIMARY KEY,
    car             INT NOT NULL,
    type            INT NOT NULL references treatment_type(id),
    date            TIMESTAMP,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_treatment()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_treatment_time BEFORE INSERT ON treatment FOR EACH ROW EXECUTE PROCEDURE update_created_column_treatment();

CREATE OR REPLACE FUNCTION update_modified_column_treatment()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_treatment_time BEFORE UPDATE ON treatment FOR EACH ROW EXECUTE PROCEDURE update_modified_column_treatment();


create table casualty_level (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_casualty_level()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_casualty_level_time BEFORE INSERT ON casualty_level FOR EACH ROW EXECUTE PROCEDURE update_created_column_casualty_level();

CREATE OR REPLACE FUNCTION update_modified_column_casualty_level()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_casualty_level_time BEFORE UPDATE ON casualty_level FOR EACH ROW EXECUTE PROCEDURE update_modified_column_casualty_level();

create table casualty_type (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    level           INT NOT NULL references casualty_level(id),
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_casualty_type()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_casualty_type_time BEFORE INSERT ON casualty_type FOR EACH ROW EXECUTE PROCEDURE update_created_column_casualty_type();

CREATE OR REPLACE FUNCTION update_modified_column_casualty_type()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_casualty_type_time BEFORE UPDATE ON casualty_type FOR EACH ROW EXECUTE PROCEDURE update_modified_column_casualty_type();

create table casualty (
    id              serial PRIMARY KEY,
    car             INT NOT NULL,
    type            INT NOT NULL references casualty_type(id),
    date            TIMESTAMP,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_casualty()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_casualty_time BEFORE INSERT ON casualty FOR EACH ROW EXECUTE PROCEDURE update_created_column_casualty();

CREATE OR REPLACE FUNCTION update_modified_column_casualty()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_casualty BEFORE UPDATE ON casualty FOR EACH ROW EXECUTE PROCEDURE update_modified_column_casualty();

create table registration_status (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_registration_status()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_registration_status_time BEFORE INSERT ON registration_status FOR EACH ROW EXECUTE PROCEDURE update_created_column_registration_status();

CREATE OR REPLACE FUNCTION update_modified_column_registration_status()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_registration_status BEFORE UPDATE ON registration_status FOR EACH ROW EXECUTE PROCEDURE update_modified_column_registration_status();

create table transmission (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_transmission()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_transmission_time BEFORE INSERT ON transmission FOR EACH ROW EXECUTE PROCEDURE update_created_column_transmission();

CREATE OR REPLACE FUNCTION update_modified_column_transmission()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_transmission BEFORE UPDATE ON transmission FOR EACH ROW EXECUTE PROCEDURE update_modified_column_transmission();

create table condition (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_condition()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_condition_time BEFORE INSERT ON condition FOR EACH ROW EXECUTE PROCEDURE update_created_column_condition();

CREATE OR REPLACE FUNCTION update_modified_column_condition()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_condition BEFORE UPDATE ON condition FOR EACH ROW EXECUTE PROCEDURE update_modified_column_condition();

create table status (
    id              serial PRIMARY KEY,
    name            varchar(50) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_status()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_status_time BEFORE INSERT ON status FOR EACH ROW EXECUTE PROCEDURE update_created_column_status();

CREATE OR REPLACE FUNCTION update_modified_column_status()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_status BEFORE UPDATE ON status FOR EACH ROW EXECUTE PROCEDURE update_modified_column_status();

create table car (
    id              serial PRIMARY KEY,
    car_type        varchar(256) NOT NULL,
    brand           INT NOT NULL references brand(id),
    flat_number     varchar(10) NOT NULL,
    transmission    INT NOT NULL references transmission(id),
    colour          varchar(100),
    year            INT,
    condition       INT NOT NULL references condition(id),
    registration_status INT NOT NULL references registration_status(id),
    seller          BIGINT NOT NULL references users(id),
    casualty        INT[],
    treatment       INT[],
    status          INT NOT NULL references status(id),
    created_at      TIMESTAMP,
    created_by      BIGINT,
    modified_at     TIMESTAMP,
    modified_by     BIGINT
);

CREATE OR REPLACE FUNCTION update_created_column_car()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_created_column_car_time BEFORE INSERT ON car FOR EACH ROW EXECUTE PROCEDURE update_created_column_car();

CREATE OR REPLACE FUNCTION update_modified_column_car()
RETURNS TRIGGER AS $$
BEGIN
NEW.modified_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_modified_column_car BEFORE UPDATE ON car FOR EACH ROW EXECUTE PROCEDURE update_modified_column_car();

-- +migrate StatementEnd
