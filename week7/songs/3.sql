--list the names of the top 5 longest songs, in descending order of length
SELECT name from songs ORDER BY duration_ms DESC limit 5;