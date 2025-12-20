-- list the names of all people who starred in a movie in which Kevin Bacon also starred
SELECT p2.name
FROM stars s1
JOIN stars s2 ON s1.movie_id = s2.movie_id
JOIN movies m ON m.id = s1.movie_id
JOIN people p1 ON p1.id = s1.person_id
JOIN people p2 ON p2.id = s2.person_id
WHERE p1.name = 'Kevin Bacon'
AND p1.birth = 1958
AND p1.id != p2.id
GROUP BY p2.name;