BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --


CREATE TABLE public.category
(
    id   SERIAL PRIMARY KEY,
    name_ru TEXT NOT NULL,
    api_name TEXT NOT NULL UNIQUE
);

CREATE TABLE public.product
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name_ru       TEXT NOT NULL,
    api_name      TEXT NOT NULL UNIQUE,
    description   TEXT NOT NULL,
    image_path    TEXT NOT NULL,
    price         BIGINT NOT NULL,
    category_id   SERIAL references public.category,
    CONSTRAINT positive_price CHECK (price > 0)
);

CREATE TABLE public.user
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username    TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL
);

-- DATA --

INSERT INTO public.user (id, username, password) VALUES (
    DEFAULT,
   'admin',
   crypt('admin', gen_salt('bf'))
);

INSERT INTO public.category (name_ru, api_name) VALUES ('Молочные продукты', 'milk_products');
INSERT INTO public.category (name_ru, api_name) VALUES ('Канцелярия', 'office_products');

INSERT INTO public.product (id, name_ru, api_name, description, image_path, price, category_id)
VALUES (
       DEFAULT,
       'Яйцо',
       'egg',
       '10 яиц в коробке',
       '/static/egg.jpg',
       100,
       1
);

INSERT INTO public.product (id, name_ru, api_name, description, image_path, price, category_id)
VALUES (
        DEFAULT,
       'Молоко',
       'milk',
       'Молоко с 3% процентами жирности',
       '/static/milk.jpg',
       119,
       1
);


COMMIT;