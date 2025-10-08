package repository

import (
	"backendgo/app/model"
	"backendgo/database"
	"database/sql"
	"fmt"
	"log"
)

// Get All
func GetAllPekerjaan() ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, is_deleted, created_at, updated_at
		FROM pekerjaan_alumni ORDER BY created_at DESC
		WHERE is_deleted = FALSE
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		var ts sql.NullTime
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &ts,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if ts.Valid {
			t := ts.Time
			p.TanggalSelesaiKerja = &t
		}
		list = append(list, p)
	}
	return list, nil
}

// Get by ID
func GetPekerjaanByID(id int) (model.PekerjaanAlumni, error) {
	var p model.PekerjaanAlumni
	var ts sql.NullTime
	err := database.DB.QueryRow(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, is_deleted, created_at, updated_at
		FROM pekerjaan_alumni WHERE id=$1
	`, id).Scan(
		&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
		&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &ts,
		&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt,
	)
	if ts.Valid {
		t := ts.Time
		p.TanggalSelesaiKerja = &t
	}
	return p, err
}

// Get by Alumni
func GetPekerjaanByAlumniID(alumniID int) ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, is_deleted, created_at, updated_at
		FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY created_at DESC
	`, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		var ts sql.NullTime
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &ts,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if ts.Valid {
			t := ts.Time
			p.TanggalSelesaiKerja = &t
		}
		list = append(list, p)
	}
	return list, nil
}

// Create
func CreatePekerjaan(p model.PekerjaanAlumni) (model.PekerjaanAlumni, error) {
	err := database.DB.QueryRow(`
		INSERT INTO pekerjaan_alumni (
			alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
			tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan,
			created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,&13)
		RETURNING id
	`, p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
		p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan,
		p.DeskripsiPekerjaan, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		log.Println("DB error CreatePekerjaan:", err)
	}
	return p, err
}

// Update
func UpdatePekerjaan(p model.PekerjaanAlumni) (model.PekerjaanAlumni, error) {
	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni
		SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, gaji_range=$5,
		    tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=$10
		WHERE id=$11
	`, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja, p.GajiRange,
		p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan, p.UpdatedAt, p.ID)
	return p, err
}

// Delete
func DeletePekerjaan(id int) error {
	_, err := database.DB.Exec("DELETE FROM pekerjaan_alumni WHERE id=$1", id)
	return err
}

// Pagination
func GetAllPekerjaanWithPagination(search, sortBy, order string, limit, offset int) ([]model.PekerjaanAlumni, error) {
	query := fmt.Sprintf(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja,
		       gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan,
		       deskripsi_pekerjaan, created_at, updated_at
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortBy, order)

	rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		var ts sql.NullTime
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &ts,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if ts.Valid {
			t := ts.Time
			p.TanggalSelesaiKerja = &t
		}
		list = append(list, p)
	}
	return list, nil
}

// Count
func CountPekerjaan(search string) (int, error) {
	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1
	`, "%"+search+"%").Scan(&count)
	return count, err
}

func SoftDeletePekerjaan(id int) error {
	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni
		SET is_deleted = TRUE, updated_at = NOW()
		WHERE id = $1
	`, id)
	return err
}

func CheckPekerjaanOwnedByAlumni(pekerjaanID int, alumniID int) (bool, error) {
	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM pekerjaan_alumni
		WHERE id = $1 AND alumni_id = $2
	`, pekerjaanID, alumniID).Scan(&count)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}


func GetTrashedPekerjaan() ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, is_deleted, created_at, updated_at
		FROM pekerjaan_alumni
		WHERE is_deleted = TRUE
		ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		var ts sql.NullTime
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &ts,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if ts.Valid {
			t := ts.Time
			p.TanggalSelesaiKerja = &t
		}
		list = append(list, p)
	}
	return list, nil
}

func RestorePekerjaan(id int) error {
	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni
		SET is_deleted = FALSE, updated_at = NOW()
		WHERE id = $1
	`, id)

	return err
}

func HardDeletePekerjaan(id int) error {
	result, err := database.DB.Exec(`
		DELETE FROM pekerjaan_alumni
		WHERE id = $1 AND is_deleted = TRUE
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("data tidak ditemukan atau belum dihapus (soft delete)")
	}

	return nil
}

