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
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    creditos integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- Name: deans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.deans (
    id uuid NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    identification_card character varying(255) NOT NULL,
    rol character varying(255) NOT NULL,
    cell_phone_number character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.deans OWNER TO postgres;

--
-- Name: faculties; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.faculties (
    id uuid NOT NULL,
    dean_id uuid NOT NULL,
    number character varying NOT NULL,
    location character varying NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.faculties OWNER TO postgres;

--
-- Name: dean_faculties; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.dean_faculties AS
 SELECT d.first_name,
    d.last_name,
    f.name AS faculty_name,
    f.location,
    f.number
   FROM (public.faculties f
     JOIN public.deans d ON ((d.id = f.dean_id)));


ALTER TABLE public.dean_faculties OWNER TO postgres;

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
    faculty_id uuid NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    identification_card character varying(255) NOT NULL,
    job_title character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.teachers OWNER TO postgres;

--
-- Name: teacher_faculties; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.teacher_faculties AS
 SELECT f.name AS faculty,
    t.first_name,
    t.last_name,
    t.job_title
   FROM (public.teachers t
     JOIN public.faculties f ON ((f.id = t.faculty_id)));


ALTER TABLE public.teacher_faculties OWNER TO postgres;

--
-- Name: view_teacher_courses; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.view_teacher_courses AS
 SELECT f.name AS faculty,
    c.name AS course,
    t.first_name,
    t.last_name,
    t.job_title
   FROM (((public.teacher_courses tc
     JOIN public.courses c ON ((c.id = tc.course_id)))
     JOIN public.teachers t ON ((t.id = tc.teacher_id)))
     JOIN public.faculties f ON ((f.id = t.faculty_id)));


ALTER TABLE public.view_teacher_courses OWNER TO postgres;

--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: deans deans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deans
    ADD CONSTRAINT deans_pkey PRIMARY KEY (id);


--
-- Name: faculties faculties_dean_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.faculties
    ADD CONSTRAINT faculties_dean_id_key UNIQUE (dean_id);


--
-- Name: faculties faculties_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.faculties
    ADD CONSTRAINT faculties_pkey PRIMARY KEY (id);


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
-- Name: faculties fk_dean_faculty; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.faculties
    ADD CONSTRAINT fk_dean_faculty FOREIGN KEY (dean_id) REFERENCES public.deans(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: teachers fk_faculty_teacher; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT fk_faculty_teacher FOREIGN KEY (faculty_id) REFERENCES public.faculties(id) ON UPDATE CASCADE ON DELETE CASCADE;


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

