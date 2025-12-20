-- lists the names of the songs that feature other artists
SELECT s.name
FROM songs s
JOIN artists a ON a.id = s.artist_id
WHERE s.name LIKE '%feat.%';