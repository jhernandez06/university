--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

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
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id uuid NOT NULL,
    codigo character varying(255) NOT NULL,
    nombre character varying(255) NOT NULL,
    creditos integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- Name: decanoes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.decanoes (
    id uuid NOT NULL,
    nombre character varying(255) NOT NULL,
    apellido character varying(255) NOT NULL,
    cedula character varying(255) NOT NULL,
    rol character varying(255) NOT NULL,
    celular character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.decanoes OWNER TO postgres;

--
-- Name: facultads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.facultads (
    id uuid NOT NULL,
    decano_id uuid NOT NULL,
    numero character varying NOT NULL,
    ubicacion character varying NOT NULL,
    nombre character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.facultads OWNER TO postgres;

--
-- Name: decano_facultads; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.decano_facultads AS
 SELECT d.nombre,
    d.apellido,
    f.nombre AS facultad_nombre,
    f.ubicacion,
    f.numero
   FROM (public.facultads f
     JOIN public.decanoes d ON ((d.id = f.decano_id)));


ALTER TABLE public.decano_facultads OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: teacher_courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.teacher_courses (
    id uuid NOT NULL,
    course_id uuid NOT NULL,
    teacher_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.teacher_courses OWNER TO postgres;

--
-- Name: teachers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.teachers (
    id uuid NOT NULL,
    facultad_id uuid NOT NULL,
    nombre character varying(255) NOT NULL,
    apellido character varying(255) NOT NULL,
    cedula character varying(255) NOT NULL,
    titulo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.teachers OWNER TO postgres;

--
-- Name: teachers_facultads; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.teachers_facultads AS
 SELECT f.nombre AS facultad,
    t.nombre,
    t.apellido,
    t.titulo
   FROM (public.teachers t
     JOIN public.facultads f ON ((f.id = t.facultad_id)));


ALTER TABLE public.teachers_facultads OWNER TO postgres;

--
-- Name: view_teacher_courses; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.view_teacher_courses AS
 SELECT f.nombre AS facultad,
    c.nombre AS course,
    t.nombre,
    t.apellido,
    t.titulo
   FROM (((public.teacher_courses tc
     JOIN public.courses c ON ((c.id = tc.course_id)))
     JOIN public.teachers t ON ((t.id = tc.teacher_id)))
     JOIN public.facultads f ON ((f.id = t.facultad_id)));


ALTER TABLE public.view_teacher_courses OWNER TO postgres;

--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: decanoes decanoes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.decanoes
    ADD CONSTRAINT decanoes_pkey PRIMARY KEY (id);


--
-- Name: facultads facultads_decano_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.facultads
    ADD CONSTRAINT facultads_decano_id_key UNIQUE (decano_id);


--
-- Name: facultads facultads_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.facultads
    ADD CONSTRAINT facultads_pkey PRIMARY KEY (id);


--
-- Name: teacher_courses teacher_courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teacher_courses
    ADD CONSTRAINT teacher_courses_pkey PRIMARY KEY (id);


--
-- Name: teachers teachers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT teachers_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: facultads fk_decano_facultad; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.facultads
    ADD CONSTRAINT fk_decano_facultad FOREIGN KEY (decano_id) REFERENCES public.decanoes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: teachers fk_facultad_teacher; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT fk_facultad_teacher FOREIGN KEY (facultad_id) REFERENCES public.facultads(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: teacher_courses teacher_courses_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teacher_courses
    ADD CONSTRAINT teacher_courses_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: teacher_courses teacher_courses_teacher_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teacher_courses
    ADD CONSTRAINT teacher_courses_teacher_id_fkey FOREIGN KEY (teacher_id) REFERENCES public.teachers(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

