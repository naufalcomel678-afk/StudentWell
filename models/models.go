package models

import "time"

type Role struct {
	IDRole   uint   `gorm:"primaryKey;column:role_id"`
	NamaRole string `gorm:"column:nama_role"`
}

func (Role) TableName() string { return "role" }

type Jurusan struct {
	IDJurusan   uint   `gorm:"primaryKey;column:id_jurusan"`
	NamaJurusan string `gorm:"column:nama_jurusan"`
}

func (Jurusan) TableName() string { return "jurusan" }

type User struct {
	IDUser   uint   `gorm:"primaryKey;column:id_user"`
	Nama     string `gorm:"column:nama"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	RoleID   uint   `gorm:"column:role_id;type:int"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

func (User) TableName() string { return "user" }

type Mahasiswa struct {
	IDUser        uint      `gorm:"primaryKey;column:id_user;autoIncrement:false"` // Matikan autoIncrement
	IDJurusan     uint      `gorm:"column:id_jurusan"`
	Angkatan      int       `gorm:"column:angkatan"`
	TanggalDaftar time.Time `gorm:"column:tanggal_daftar"`
	User          User      `gorm:"foreignKey:IDUser"`
	Jurusan       Jurusan   `gorm:"foreignKey:IDJurusan"`
}

func (Mahasiswa) TableName() string { return "mahasiswa" }

type Kategori struct {
	IDKategori   uint   `gorm:"primaryKey;column:id_kategori"` // Harus id_kategori
	NamaKategori string `gorm:"column:nama_kategori"`          // Harus nama_kategori
}

func (Kategori) TableName() string { return "kategori" }

type RangeStres struct {
	IDRange      uint   `gorm:"primaryKey;column:id_range"`
	SkorMin      int    `gorm:"column:skor_min"`
	SkorMax      int    `gorm:"column:skor_max"`
	TingkatStres string `gorm:"column:tingkat_stres"`
}

func (RangeStres) TableName() string { return "range_stres" }

type Pertanyaan struct {
	IDPertanyaan   uint     `gorm:"primaryKey;column:id_pertanyaan"`
	IDKategori     uint     `gorm:"column:id_kategori"`
	TeksPertanyaan string   `gorm:"column:teks_pertanyaan"`
	Bobot          int      `gorm:"column:bobot"`
	Kategori       Kategori `gorm:"foreignKey:IDKategori"`
}

func (Pertanyaan) TableName() string { return "pertanyaan" }

type Kuesioner struct {
	IDKuesioner uint      `gorm:"primaryKey;column:id_kuesioner"`
	TanggalIsi  time.Time `gorm:"column:tanggal_isi;type:date"`
	IDUser      uint      `gorm:"column:id_user"`
	User        User      `gorm:"foreignKey:IDUser"`
}

func (Kuesioner) TableName() string { return "kuesioner" }

type DetailJawaban struct {
	IDDetail     uint       `gorm:"primaryKey;column:id_detail"`
	IDKuesioner  uint       `gorm:"column:id_kuesioner"`
	IDPertanyaan uint       `gorm:"column:id_pertanyaan"`
	NilaiJawaban int        `gorm:"column:nilai_jawaban"`
	Kuesioner    Kuesioner  `gorm:"foreignKey:IDKuesioner"`
	Pertanyaan   Pertanyaan `gorm:"foreignKey:IDPertanyaan"`
}

func (DetailJawaban) TableName() string { return "detail_jawaban" }

type HasilStres struct {
	IDHasil      uint       `gorm:"primaryKey;column:id_hasil"`
	IDKuesioner  uint       `gorm:"column:id_kuesioner"`
	IDRange      uint       `gorm:"column:id_range"`
	TotalSkor    int        `gorm:"column:total_skor"`
	Rekomendasi  string     `gorm:"column:rekomendasi"`
	TanggalHasil time.Time  `gorm:"column:tanggal_hasil;type:date"`
	Kuesioner    Kuesioner  `gorm:"foreignKey:IDKuesioner"`
	RangeStres   RangeStres `gorm:"foreignKey:IDRange"`
}

func (HasilStres) TableName() string { return "hasil_stres" }
