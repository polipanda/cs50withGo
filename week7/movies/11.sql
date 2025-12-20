-- list the titles of the five highest rated movies (in order) that
-- Chadwick Boseman starred in, starting with the highest rated
-- SELECT title, ratings.rating
SELECT *
from movies
JOIN stars ON stars.movie_id = movies.id
LEFT JOIN people ON people.id = stars.person_id
LEFT JOIN ratings ON ratings.movie_id = movies.id
WHERE people.name = 'Chadwick Boseman'
-- there's multiple records for actors in the same movie ;/
GROUP BY title
ORDER BY ratings.rating DESC
LIMIT 5;

SELECT * FROM stars where person_id = 1569276 and movie_id = 4154756;