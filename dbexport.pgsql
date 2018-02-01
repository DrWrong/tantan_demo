--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.3
-- Dumped by pg_dump version 9.5.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: relationship; Type: TABLE; Schema: public; Owner: drwrong
--

CREATE TABLE relationship (
    id integer NOT NULL,
    uid1 integer,
    uid2 integer,
    state smallint
);


ALTER TABLE relationship OWNER TO drwrong;

--
-- Name: relationship_id_seq; Type: SEQUENCE; Schema: public; Owner: drwrong
--

CREATE SEQUENCE relationship_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE relationship_id_seq OWNER TO drwrong;

--
-- Name: relationship_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: drwrong
--

ALTER SEQUENCE relationship_id_seq OWNED BY relationship.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: drwrong
--

CREATE TABLE "user" (
    id integer NOT NULL,
    name character varying(255)
);


ALTER TABLE "user" OWNER TO drwrong;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: drwrong
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO drwrong;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: drwrong
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: drwrong
--

ALTER TABLE ONLY relationship ALTER COLUMN id SET DEFAULT nextval('relationship_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: drwrong
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- Data for Name: relationship; Type: TABLE DATA; Schema: public; Owner: drwrong
--

COPY relationship (id, uid1, uid2, state) FROM stdin;
1	1	2	2
2	2	1	1
\.


--
-- Name: relationship_id_seq; Type: SEQUENCE SET; Schema: public; Owner: drwrong
--

SELECT pg_catalog.setval('relationship_id_seq', 3, true);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: drwrong
--

COPY "user" (id, name) FROM stdin;
1	Alice
2	Alice2
3	DrWrong
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: drwrong
--

SELECT pg_catalog.setval('user_id_seq', 3, true);


--
-- Name: relationship_pkey; Type: CONSTRAINT; Schema: public; Owner: drwrong
--

ALTER TABLE ONLY relationship
    ADD CONSTRAINT relationship_pkey PRIMARY KEY (id);


--
-- Name: relationship_uid1_uid2_key; Type: CONSTRAINT; Schema: public; Owner: drwrong
--

ALTER TABLE ONLY relationship
    ADD CONSTRAINT relationship_uid1_uid2_key UNIQUE (uid1, uid2);


--
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: drwrong
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: drwrong
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM drwrong;
GRANT ALL ON SCHEMA public TO drwrong;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

