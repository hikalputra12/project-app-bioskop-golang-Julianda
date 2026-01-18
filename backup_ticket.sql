--
-- PostgreSQL database dump
--

\restrict 0OOAfjS2YW5mRQHfdGKo2Nmjre7tubkCKf6PmpRX3Dol56weZvqtHTrS6FePLnP

-- Dumped from database version 16.11
-- Dumped by pg_dump version 16.11

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: booking_seat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.booking_seat (
    id integer NOT NULL,
    user_id integer,
    showtime_id integer,
    seat_id integer,
    payment_method_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    payment_details jsonb,
    status character varying(20)
);


ALTER TABLE public.booking_seat OWNER TO postgres;

--
-- Name: booking_seat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.booking_seat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.booking_seat_id_seq OWNER TO postgres;

--
-- Name: booking_seat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.booking_seat_id_seq OWNED BY public.booking_seat.id;


--
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id integer NOT NULL,
    name character varying(255),
    location text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_cinema_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_cinema_id_seq OWNER TO postgres;

--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_cinema_id_seq OWNED BY public.cinemas.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255),
    description text,
    duration integer,
    release_date date,
    genre character varying(100),
    rating double precision,
    poster_url text,
    director_info jsonb,
    cast_info jsonb,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- Name: movies_movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_movies_id_seq OWNER TO postgres;

--
-- Name: movies_movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_movies_id_seq OWNED BY public.movies.id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id integer NOT NULL,
    name character varying(50),
    logo_url text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- Name: payment_methods_payment_method_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_payment_method_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_payment_method_id_seq OWNER TO postgres;

--
-- Name: payment_methods_payment_method_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_payment_method_id_seq OWNED BY public.payment_methods.id;


--
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    studio_id integer,
    seat_number character varying(10),
    is_available boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- Name: seats_seat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_seat_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_seat_id_seq OWNER TO postgres;

--
-- Name: seats_seat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_seat_id_seq OWNED BY public.seats.id;


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id integer,
    expired_at timestamp without time zone,
    revoked_at timestamp without time zone,
    last_active timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- Name: showtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.showtimes (
    id integer NOT NULL,
    movie_id integer,
    studio_id integer,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    price numeric(10,2),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.showtimes OWNER TO postgres;

--
-- Name: showtimes_showtime_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.showtimes_showtime_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.showtimes_showtime_id_seq OWNER TO postgres;

--
-- Name: showtimes_showtime_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.showtimes_showtime_id_seq OWNED BY public.showtimes.id;


--
-- Name: studios; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.studios (
    id integer NOT NULL,
    cinema_id integer,
    name character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.studios OWNER TO postgres;

--
-- Name: studios_studio_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.studios_studio_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.studios_studio_id_seq OWNER TO postgres;

--
-- Name: studios_studio_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.studios_studio_id_seq OWNED BY public.studios.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    email text NOT NULL,
    password_hash character varying(255) NOT NULL,
    phone_number character varying(20),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    is_verified boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.id;


--
-- Name: booking_seat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat ALTER COLUMN id SET DEFAULT nextval('public.booking_seat_id_seq'::regclass);


--
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_cinema_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_movies_id_seq'::regclass);


--
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_payment_method_id_seq'::regclass);


--
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_seat_id_seq'::regclass);


--
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_showtime_id_seq'::regclass);


--
-- Name: studios id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios ALTER COLUMN id SET DEFAULT nextval('public.studios_studio_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: booking_seat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking_seat (id, user_id, showtime_id, seat_id, payment_method_id, created_at, updated_at, deleted_at, payment_details, status) FROM stdin;
31	17	1	2	2	2026-01-18 02:01:24.372098	\N	\N	\N	PENDING
30	17	1	1	2	2026-01-18 02:01:24.369631	\N	\N	{"CVV": "123", "CardNumber": "1234567812345678", "ExpiryDate": "01/06"}	PAID
18	\N	1	3	1	2026-01-16 03:46:11.787092	\N	\N	{"CVV": "123", "CardNumber": "1234567812345678", "ExpiryDate": "01/06"}	PAID
29	\N	1	5	2	2026-01-17 23:37:07.092814	\N	\N	{"CVV": "123", "CardNumber": "1234567812345678", "ExpiryDate": "01/06"}	PAID
28	\N	1	4	2	2026-01-17 23:37:07.087838	\N	\N	{"CVV": "123", "CardNumber": "1234567812345678", "ExpiryDate": "01/06"}	PAID
\.


--
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cinemas (id, name, location, created_at, updated_at, deleted_at) FROM stdin;
1	Grand Indonesia XXI	Jl. M.H. Thamrin No.1	2026-01-13 17:38:31.683841	\N	\N
2	PIM 3 XXI	Pondok Indah Mall 3	2026-01-13 17:38:31.683841	\N	\N
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, title, description, duration, release_date, genre, rating, poster_url, director_info, cast_info, created_at, updated_at, deleted_at) FROM stdin;
1	Avengers: Infinity War	The Avengers unite to battle their most powerful enemy yet -- Thanos.	149	2018-04-27	Action	8.4	https://poster.url/avengers.jpg	{"name": "Anthony Russo", "photo_url": "https://img.com/russo.jpg"}	[{"name": "Robert Downey Jr.", "role": "Iron Man", "photo_url": "https://img.com/rdj.jpg"}, {"name": "Chris Hemsworth", "role": "Thor", "photo_url": "https://img.com/chris.jpg"}]	2026-01-13 17:38:31.683841	\N	\N
2	Inception	A thief who steals corporate secrets through the use of dream-sharing technology.	148	2010-07-16	Sci-Fi	8.8	https://poster.url/inception.jpg	{"name": "Christopher Nolan", "photo_url": "https://img.com/nolan.jpg"}	[{"name": "Leonardo DiCaprio", "role": "Cobb", "photo_url": "https://img.com/leo.jpg"}]	2026-01-13 17:38:31.683841	\N	\N
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_methods (id, name, logo_url, created_at, updated_at, deleted_at) FROM stdin;
1	Credit Card	https://img.com/cc.png	2026-01-13 17:38:31.683841	\N	\N
2	Gopay	https://img.com/gopay.png	2026-01-13 17:38:31.683841	\N	\N
3	Virtual Account	https://img.com/va.png	2026-01-13 17:38:31.683841	\N	\N
\.


--
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (id, studio_id, seat_number, is_available, created_at, updated_at, deleted_at) FROM stdin;
6	1	B3	t	2026-01-13 17:38:31.683841	\N	\N
3	1	A3	f	2026-01-13 17:38:31.683841	2026-01-16 03:46:11.793538	\N
4	1	B1	f	2026-01-13 17:38:31.683841	2026-01-18 01:57:14.216324	\N
5	1	B2	f	2026-01-13 17:38:31.683841	2026-01-18 01:57:14.218822	\N
1	1	A1	f	2026-01-13 17:38:31.683841	2026-01-18 02:02:10.019722	\N
2	1	A2	f	2026-01-13 17:38:31.683841	2026-01-18 02:02:10.021491	\N
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, user_id, expired_at, revoked_at, last_active, created_at) FROM stdin;
3e5e7c22-2dca-4422-a19e-5c1449eeeb60	17	2026-01-19 01:48:49.987946	2026-01-18 01:49:03.232485	2026-01-18 01:48:49.987946	2026-01-18 01:48:49.987946
ffa6cc7e-c81f-448b-b225-397713613945	17	2026-01-19 02:05:42.614499	2026-01-18 17:47:41.09151	2026-01-18 02:05:42.615018	2026-01-18 01:49:24.199588
6ba60f3f-19ec-4291-8f71-7d51473a5188	26	2026-01-19 21:04:17.377514	2026-01-18 21:13:00.859344	2026-01-18 21:04:17.377514	2026-01-18 21:04:17.377514
710c21a5-1bc5-4669-aa3e-911223d62ac3	17	2026-01-19 21:13:36.955703	2026-01-18 21:17:34.440789	2026-01-18 21:13:36.955702	2026-01-18 21:13:36.955702
\.


--
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.showtimes (id, movie_id, studio_id, start_time, end_time, price, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	2026-01-14 17:38:31.683841	2026-01-14 20:08:31.683841	50000.00	2026-01-13 17:38:31.683841	\N	\N
2	2	2	2026-01-15 17:38:31.683841	2026-01-15 19:38:31.683841	45000.00	2026-01-13 17:38:31.683841	\N	\N
\.


--
-- Data for Name: studios; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.studios (id, cinema_id, name, created_at, updated_at, deleted_at) FROM stdin;
1	1	Studio 1	2026-01-13 17:38:31.683841	\N	\N
2	1	Studio 2	2026-01-13 17:38:31.683841	\N	\N
3	2	Studio 1	2026-01-13 17:38:31.683841	\N	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password_hash, phone_number, created_at, updated_at, deleted_at, is_verified) FROM stdin;
17	rania	rania@gmail.com	$2a$14$LZNGxCHjTBmRFB.9DNaBdeA6Gb0BqxufQ/qPgxDuTFZI8mJwe/16u	0823239932424	2026-01-18 01:35:26.78029	2026-01-18 01:35:26.78029	\N	f
26	haikal	haikal.putra1210@gmail.com	$2a$14$tdF0MQlEQxi0HUMKy5jkvOaUGVNMM9M3A9NdEHbS6scjSOUeFr9Fa	0812345678\n	2026-01-18 21:02:21.582347	2026-01-18 21:02:21.582347	\N	t
\.


--
-- Name: booking_seat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.booking_seat_id_seq', 31, true);


--
-- Name: cinemas_cinema_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_cinema_id_seq', 2, true);


--
-- Name: movies_movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_movies_id_seq', 2, true);


--
-- Name: payment_methods_payment_method_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_payment_method_id_seq', 3, true);


--
-- Name: seats_seat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_seat_id_seq', 6, true);


--
-- Name: showtimes_showtime_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.showtimes_showtime_id_seq', 2, true);


--
-- Name: studios_studio_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.studios_studio_id_seq', 3, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 26, true);


--
-- Name: booking_seat booking_seat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT booking_seat_pkey PRIMARY KEY (id);


--
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- Name: studios studios_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios
    ADD CONSTRAINT studios_pkey PRIMARY KEY (id);


--
-- Name: booking_seat unique_seat_booking; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT unique_seat_booking UNIQUE (showtime_id, seat_id, deleted_at);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: booking_seat booking_seat_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT booking_seat_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id);


--
-- Name: booking_seat booking_seat_seat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT booking_seat_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id);


--
-- Name: booking_seat booking_seat_showtime_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT booking_seat_showtime_id_fkey FOREIGN KEY (showtime_id) REFERENCES public.showtimes(id);


--
-- Name: booking_seat booking_seat_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seat
    ADD CONSTRAINT booking_seat_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: seats seats_studio_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_studio_id_fkey FOREIGN KEY (studio_id) REFERENCES public.studios(id);


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: showtimes showtimes_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- Name: showtimes showtimes_studio_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_studio_id_fkey FOREIGN KEY (studio_id) REFERENCES public.studios(id);


--
-- Name: studios studios_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios
    ADD CONSTRAINT studios_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);


--
-- PostgreSQL database dump complete
--

\unrestrict 0OOAfjS2YW5mRQHfdGKo2Nmjre7tubkCKf6PmpRX3Dol56weZvqtHTrS6FePLnP

