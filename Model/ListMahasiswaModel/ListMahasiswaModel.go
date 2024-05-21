package ListMahasiswaModel

type ListMahasiswa struct {
	// KTPID      string `json:"ktpId" form:"ktpId"`
	// NamaTamu   string `json:"namaTamu" form:"namaTamu"`
	// Tujuan     string `json:"tujuan" form:"tujuan"`
	// Keterangan string `json:"keterangan" form:"keterangan"`
	ID   int    `json:"id" form:"id"`
	NIM  string `json:"nim" form:"nim"`
	Nama string `json:"nama" form:"nama"`
}

// TableName memberikan nama tabel yang eksplisit
func (ListMahasiswa) TableName() string {
	return "tb_mahasiswa"
}
