--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.5

-- Started on 2019-09-03 01:46:37

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

--
-- TOC entry 198 (class 1255 OID 24584)
-- Name: auto_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.auto_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$;


ALTER FUNCTION public.auto_updated_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 24576)
-- Name: approval; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.approval (
    id text NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    service_rule integer NOT NULL,
    comment text,
    status integer NOT NULL,
    deadline timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.approval OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 24586)
-- Name: servicerule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.servicerule (
    id integer NOT NULL,
    name text NOT NULL,
    apikey text NOT NULL,
    url text NOT NULL
);


ALTER TABLE public.servicerule OWNER TO postgres;

--
-- TOC entry 2818 (class 0 OID 24576)
-- Dependencies: 196
-- Data for Name: approval; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.approval (id, title, description, service_rule, comment, status, deadline, created_at, updated_at) FROM stdin;
5834ddba-7a1b-4c49-b010-1b7924ac7679	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:16:51.88905+08	2019-09-01 23:26:44.475694+08
193b8ce8-cbfd-4388-aa50-7730d60c81d0	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:16:54.431546+08	2019-09-01 23:26:44.475694+08
0fde2ae0-9a87-4e59-90be-536996fabf2b	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:16:55.286292+08	2019-09-01 23:26:44.475694+08
a2c0850d-4ab3-448d-a583-671785186be5	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:16:56.653294+08	2019-09-01 23:26:44.475694+08
6dd58623-613a-4a77-aa0e-ba32e6488b5e	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:22:24.79507+08	2019-09-01 23:26:44.475694+08
79a71571-d5a2-49d5-b647-0cde66beda12	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 13:39:17.384026+08	2019-09-01 23:26:44.475694+08
abd9e336-943b-4cfb-8f8a-a9f0d0a69564	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 14:38:35.374373+08	2019-09-01 23:26:44.475694+08
e42d09f1-eddb-4d9e-86c2-19c749ddbeec	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 14:38:49.323318+08	2019-09-01 23:26:44.475694+08
10e2bb30-c058-4375-92e7-eff9d51e6ea6	ABCzzzD	desc	1	Commenzt	2	2019-09-22 21:42:31+08	2019-09-01 23:47:33.673342+08	2019-09-01 23:47:33.674341+08
23041b85-bb9d-4ad6-bbde-ba2d787db935	ABCzzzD1	desc	1	Commenzt	2	2019-09-22 21:42:31+08	2019-09-01 23:47:41.223169+08	2019-09-01 23:47:41.224169+08
e07406d4-1b35-47af-9532-e5f850b45963	ABCz11zzD	desc	1	TESTTESTTzEzSTz	3	2019-09-22 20:42:31+08	2019-09-01 13:16:55.98238+08	2019-09-02 01:52:26.485927+08
56fce895-84f4-465b-b203-c00bfe90cfc5	ABCzzzD1z	desc	1	TESTTESTTzEzSTz	2	2019-09-22 21:42:31+08	2019-09-01 23:48:22.274383+08	2019-09-02 01:32:13.953997+08
e27d1583-ecaf-4457-af65-e66be666de73	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 20:42:31+08	2019-09-01 17:35:32.606913+08	2019-09-01 23:26:44.475694+08
e976a3d1-a712-4366-afa9-31100ddd59b3	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 19:42:31+08	2019-09-01 18:11:59.709972+08	2019-09-01 23:26:44.475694+08
18256fea-4b59-41cb-8c9c-0ddbfa1e32d6	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 21:42:31+08	2019-09-01 18:12:06.440198+08	2019-09-01 23:26:44.475694+08
14618e16-55d2-4123-a7b5-5e0a7d64c833	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 21:42:31+08	2019-09-01 19:05:35.205064+08	2019-09-01 23:26:44.475694+08
1eebb31e-1923-425b-89e5-df3be06262d2	ABCz11zzD	desc	1	TESTTESTTzEST	2	2019-09-22 21:42:31+08	2019-09-01 19:05:37.430317+08	2019-09-01 23:26:44.475694+08
c802f0d7-5e15-439e-a8e0-c3df8c560136	ABCzzzD1	desc	1	Commenzt	2	2019-09-22 21:42:31+08	2019-09-01 23:48:19.454834+08	2019-09-01 23:48:19.455833+08
438114d3-aa8b-4e6d-a512-fc9ae2c1605e	ABCz11zzD	desc	2	TESTTESTTzEzSTz	2	2019-09-22 20:42:31+08	2019-09-01 13:05:49.583452+08	2019-09-02 00:52:22.817149+08
\.


--
-- TOC entry 2819 (class 0 OID 24586)
-- Dependencies: 197
-- Data for Name: servicerule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.servicerule (id, name, apikey, url) FROM stdin;
2	test23	adsfasdf	fsadadsfasdf
1	UserSvc	usersvcapikey	http://localhost:8080/approval
\.


--
-- TOC entry 2691 (class 2606 OID 24583)
-- Name: approval approval_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.approval
    ADD CONSTRAINT approval_pkey PRIMARY KEY (id);


--
-- TOC entry 2694 (class 2606 OID 24593)
-- Name: servicerule servicerule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.servicerule
    ADD CONSTRAINT servicerule_pkey PRIMARY KEY (id);


--
-- TOC entry 2692 (class 1259 OID 24599)
-- Name: fki_ServiceRuleFK; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "fki_ServiceRuleFK" ON public.approval USING btree (service_rule);


--
-- TOC entry 2696 (class 2620 OID 24585)
-- Name: approval auto_field_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER auto_field_updated_at BEFORE UPDATE ON public.approval FOR EACH ROW EXECUTE PROCEDURE public.auto_updated_at_column();


--
-- TOC entry 2695 (class 2606 OID 24594)
-- Name: approval ServiceRuleFK; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.approval
    ADD CONSTRAINT "ServiceRuleFK" FOREIGN KEY (service_rule) REFERENCES public.servicerule(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2019-09-03 01:46:37

--
-- PostgreSQL database dump complete
--

