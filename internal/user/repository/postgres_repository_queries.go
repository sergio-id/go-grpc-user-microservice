package repository

// here we store all the queries for the postgresql repository
const (
	createUserCommand = `INSERT INTO "user" (email, password, first_name, last_name, about, phone_number, gender, status, last_ip, last_device, avatar_url) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
					RETURNING *`

	updateUserCommand = `UPDATE "user" 
					SET first_name = COALESCE(NULLIF($1, ''), first_name), 
					    last_name = COALESCE(NULLIF($2, ''), last_name),
						about = COALESCE(NULLIF($3, ''), about),
						phone_number = COALESCE(NULLIF($4, ''), phone_number),
						gender = $5,
						status = $6,
						last_ip = COALESCE(NULLIF($7, ''), last_ip),
						last_device = COALESCE(NULLIF($8, ''), last_device),
						avatar_url = COALESCE(NULLIF($9, ''), avatar_url),
					    updated_at = now()
					WHERE id = $10 
					RETURNING *`

	deleteByIDCommand = `DELETE FROM "user" WHERE id = $1`

	getByIDQuery = `SELECT * FROM "user" WHERE id = $1`

	getByEmailQuery = `SELECT * FROM "user" WHERE email = $1`
)
