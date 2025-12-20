-- list the names of all people who starred in Toy Story
SELECT p.name
FROM movies m
JOIN stars ON m.id = stars.movie_id
JOIN people p ON p.id = stars.person_id
WHERE title = 'Toy Story';