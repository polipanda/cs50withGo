-- list all movies released in 2010 and their ratings, in descending order by rating.
-- For movies with the same rating, order them alphabetically by title.
SELECT m.title, r.rating
FROM movies m
JOIN ratings r ON r.movie_id = m.id
WHERE year = 2010
ORDER BY r.rating DESC, m.title;