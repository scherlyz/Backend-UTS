package repository

import (
	"backendgo/app/model"
	"backendgo/database"
	
	"fmt"
)

// Get All Alumni
func GetAllAlumni() ([]model.Alumni, error) {
	rows, err := database.DB.Query(`
		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
		FROM alumni ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alumniList []model.Alumni
	for rows.Next() {
		var a model.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
			&a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}
	return alumniList, nil
}

// Get Alumni by ID
func GetAlumniByID(id int) (model.Alumni, error) {
	var a model.Alumni
	err := database.DB.QueryRow(`
		SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
		FROM alumni WHERE id=$1
	`, id).Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
		&a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
	return a, err
}

// Create Alumni
func CreateAlumni(a model.Alumni) (model.Alumni, error) {
	err := database.DB.QueryRow(`
		INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id
	`, a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat, a.CreatedAt, a.UpdatedAt).Scan(&a.ID)
	return a, err
}

// Update Alumni
func UpdateAlumni(a model.Alumni) (model.Alumni, error) {
	_, err := database.DB.Exec(`
		UPDATE alumni SET nim=$1, nama=$2, jurusan=$3, angkatan=$4, tahun_lulus=$5,
		    email=$6, no_telepon=$7, alamat=$8, updated_at=$9 WHERE id=$10
	`, a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat, a.UpdatedAt, a.ID)
	return a, err
}

// Delete Alumni
func DeleteAlumni(id int) error {
	_, err := database.DB.Exec("DELETE FROM alumni WHERE id=$1", id)
	return err
}

// Update Status Kematian
func UpdateStatusKematian(id int, status bool) error {
	_, err := database.DB.Exec(`
        UPDATE alumni SET status_kematian=$1, updated_at=NOW() WHERE id=$2
    `, status, id)
	return err
}

// Pagination
func GetAlumniRepo(search, sortBy, order string, limit, offset int) ([]model.Alumni, error) {
	query := fmt.Sprintf(`
        SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, status_kematian, created_at, updated_at
        FROM alumni
        WHERE nama ILIKE $1 OR email ILIKE $1 OR jurusan ILIKE $1
        ORDER BY %s %s
        LIMIT $2 OFFSET $3
    `, sortBy, order)

	rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Alumni
	for rows.Next() {
		var a model.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.StatusKematian,
			&a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func CountAlumniRepo(search string) (int, error) {
	var total int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM alumni WHERE nama ILIKE $1 OR email ILIKE $1 OR jurusan ILIKE $1
	`, "%"+search+"%").Scan(&total)
	return total, err
}
