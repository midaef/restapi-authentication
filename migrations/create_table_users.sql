CREATE TABLE public.users
(
    id serial,
    user character varying(32),
    password character varying(64),
    token character varying(128),
    words text,
    premium character varying(10),
    book_marks text
);

ALTER TABLE public.users
    OWNER to postgres;