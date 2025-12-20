-- list the names of all people who starred in
-- a movie released in 2004, ordered by birth year
-- f a person appeared in more than one movie in 2004, they should only appear in your results once.
SELECT DISTINCT(p.name)
FROM people p
JOIN stars ON p.id = stars.person_id
JOIN movies m ON m.id = stars.movie_id
WHERE m.year = 2004
ORDER BY p.name;