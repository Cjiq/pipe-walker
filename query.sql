-- name: GetLines :many
SELECT 
  l.id AS line_id, 
  l.name AS line_name, 
  GROUP_CONCAT(l2.name) as neighbors
FROM lines AS l
  JOIN nodes AS n  ON l.n_id = n.n_id
  JOIN nodes AS n2 ON n.node_id = n2.node_id and n.n_id != n2.n_id
  JOIN lines AS l2 ON n2.n_id = l2.n_id
GROUP BY 
  l.id, 
  l.name
ORDER BY 
  n.node_id;

-- name: GetLine :one
SELECT * FROM lines
JOIN nodes ON lines.n_id = nodes.n_id
WHERE lines.id = ?;

-- name: CreateLine :one
INSERT INTO lines (name, type, n_id) VALUES (?, ?, ?)
RETURNING *;

-- name: CreateNode :one
INSERT INTO nodes (name, node_id, n_id) VALUES (?, ?, ?)
RETURNING *;

