package model

import "time"

type PekerjaanAlumni struct {
	ID                int       `json:"id"`
	AlumniID          int       `json:"alumni_id"`
	NamaPerusahaan    string    `json:"nama_perusahaan"`
	PosisiJabatan     string    `json:"posisi_jabatan"`
	BidangIndustri    string    `json:"bidang_industri"`
	LokasiKerja       string    `json:"lokasi_kerja"`
	GajiRange         string    `json:"gaji_range"`
	TanggalMulaiKerja time.Time `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja"`
	StatusPekerjaan   string    `json:"status_pekerjaan"`
	DeskripsiPekerjaan string    `json:"deskripsi_pekerjaan"`
    IsDeleted           bool      `json:"is_deleted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Request untuk CREATE pekerjaan
type CreatePekerjaanRequest struct {
    AlumniID            int     `json:"alumni_id"`
    NamaPerusahaan      string  `json:"nama_perusahaan"`
    PosisiJabatan       string  `json:"posisi_jabatan"`
    BidangIndustri      string  `json:"bidang_industri"`
    LokasiKerja         string  `json:"lokasi_kerja"`
    GajiRange           string  `json:"gaji_range"`
    TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`   // string → biar gampang diparse
    TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"` // nullable
    StatusPekerjaan     string  `json:"status_pekerjaan"`
    DeskripsiPekerjaan  string  `json:"deskripsi_pekerjaan"`
}

// Request untuk UPDATE pekerjaan
type UpdatePekerjaanRequest struct {
    NamaPerusahaan      string  `json:"nama_perusahaan"`
    PosisiJabatan       string  `json:"posisi_jabatan"`
    BidangIndustri      string  `json:"bidang_industri"`
    LokasiKerja         string  `json:"lokasi_kerja"`
    GajiRange           string  `json:"gaji_range"`
    TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`
    TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"`
    StatusPekerjaan     string  `json:"status_pekerjaan"`
    DeskripsiPekerjaan  string  `json:"deskripsi_pekerjaan"`
}
