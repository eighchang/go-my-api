-- public.articles definition
-- Drop table
-- DROP TABLE public.articles;
CREATE TABLE public.articles (
  article_id serial4 NOT NULL,
  title varchar(100) NOT NULL,
  contents text NOT NULL,
  username varchar(100) NOT NULL,
  nice int4 NOT NULL,
  created_at date NOT NULL,
  CONSTRAINT articles_pkey PRIMARY KEY (article_id)
);
-- public."comments" definition
-- Drop table
-- DROP TABLE public."comments";
CREATE TABLE public."comments" (
  comment_id serial4 NOT NULL,
  article_id int4 NOT NULL,
  message text NOT NULL,
  create_at date NOT NULL,
  CONSTRAINT comments_pkey PRIMARY KEY (comment_id)
);
-- public."comments" foreign keys
ALTER TABLE public."comments"
ADD CONSTRAINT comments_fk FOREIGN KEY (article_id) REFERENCES public.articles(article_id) ON DELETE CASCADE ON UPDATE CASCADE;