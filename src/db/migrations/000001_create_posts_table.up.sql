BEGIN;
CREATE TABLE if NOT EXISTS "Posts"
(
	id serial
		constraint "Posts_pkey"
			primary key,
	category varchar(255) default 'JOB'::character varying,
	type varchar(255) [],
	content text,
	unit varchar(255),
	price integer,
	"issueDate" timestamp with time zone,
	"workingDays" integer,
	location varchar(255) [],
	area numeric(19,10),
	title varchar(255),
	"createdAt" timestamp with time zone not null,
	"updatedAt" timestamp with time zone not null
);

CREATE INDEX if NOT EXISTS posts_category
	on "Posts" (category);

COMMIT;
