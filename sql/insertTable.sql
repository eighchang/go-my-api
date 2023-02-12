# articles

INSERT INTO public.articles
(title, contents, username, nice, created_at)
VALUES('firstPost', 'This is my first blog', 'eiji', 2, now());
INSERT INTO public.articles
(title, contents, username, nice, created_at)
VALUES('2nd', '2nd blog post', 'eiji', 2, now());

# comments
INSERT INTO public."comments"
(article_id, message, create_at)
VALUES(1, '1st comment yeah', now());

INSERT INTO public."comments"
(article_id, message, create_at)
VALUES(1, 'welcome', now());