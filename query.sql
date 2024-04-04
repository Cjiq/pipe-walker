-- name: GetLines :many
SELECT 
  lines.id as line_id, 
  lines.name as line_name, 
  GROUP_CONCAT(l2.name) as neighbors
FROM lines
  JOIN nodes ON lines.n_id = nodes.n_id
  JOIN nodes as n2 ON nodes.node_id = n2.node_id and nodes.n_id != n2.n_id
  JOIN lines as l2 ON n2.n_id = l2.n_id
GROUP BY 
  lines.id, 
  lines.name;

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

