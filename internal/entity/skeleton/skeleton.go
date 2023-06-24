package skeleton

// Skeleton model
type Skeleton struct {
	SkeletonID int `db:"skeleton_id" json:"skeleton_id"`
	//SkeletonName string `db:"skeleton_name" json:"skeleton_name"`
	SkeletonType string `db:"skeleton_type" json:"skeleton_type"`
}

type DataSiswa struct {
	NamaSiswa  string  `db:"nama" json:"nama"`
	Kelas      string  `db:"kelas" json:"kelas"`
	NilaiSmst1 float64 `db:"nilai_smst_1" json:"nilai_smst_1"`
	NilaiSmst2 float64 `db:"nilai_smst_2" json:"nilai_smst_2"`
	NilaiSmst3 float64 `db:"nilai_smst_3" json:"nilai_smst_3"`
	NilaiSmst4 float64 `db:"nilai_smst_4" json:"nilai_smst_4"`
	NilaiSmst5 float64 `db:"nilai_smst_5" json:"nilai_smst_5"`
}
