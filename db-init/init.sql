-- Adminer 5.2.1 PostgreSQL 16.6 dump

DROP TABLE IF EXISTS "projects";
DROP SEQUENCE IF EXISTS project_id_seq;
CREATE SEQUENCE project_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."projects" (
    "id" integer DEFAULT nextval('project_id_seq') NOT NULL,
    "name" character varying(255) NOT NULL,
    "description" text NOT NULL,
    "created_by" integer NOT NULL,
    CONSTRAINT "project_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "tasks";
DROP SEQUENCE IF EXISTS tasks_id_seq;
CREATE SEQUENCE tasks_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."tasks" (
    "id" integer DEFAULT nextval('tasks_id_seq') NOT NULL,
    "project_id" integer NOT NULL,
    "title" character varying(255) NOT NULL,
    "description" text NOT NULL,
    "status" character varying(255) NOT NULL,
    "deadline" timestamp NOT NULL
) WITH (oids = false);


DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "name" character varying(255) NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE UNIQUE INDEX users_email_unique ON public.users USING btree (email);


ALTER TABLE ONLY "public"."projects" ADD CONSTRAINT "projects_id_fkey" FOREIGN KEY (created_by) REFERENCES users(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."tasks" ADD CONSTRAINT "tasks_id_fkey" FOREIGN KEY (project_id) REFERENCES projects(id) NOT DEFERRABLE;

-- 2025-07-15 19:21:06 UTC
