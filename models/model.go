package model

import "time"

type PerhitunganInvetasiRequest struct {
	JenisKelamin  string  `json:"jenis_kelamin"`
	Usia          int     `json:"usia"`
	Perokok       string  `json:"perokok"`
	Nominal       float64 `json:"nominal"`
	LamaInvestasi int     `json:"lama_investasi"`
}

type Transaction struct {
	ID              int       `json:"id"`
	TglTransaksi    time.Time `json:"tgl_transaksi"`
	NoTransaction   string    `json:"no_transaction"`
	Nama            string    `json:"nama"`
	JenisKelamin    string    `json:"jenis_kelamin"`
	Usia            int       `json:"usia"`
	Email           string    `json:"email"`
	Perokok         string    `json:"perokok"`
	Nominal         int       `json:"nominal"`
	LamaInvestasi   int       `json:"lama_investasi"`
	PeriodeBayar    string    `json:"periode_pembayaran"`
	MetodeBayar     string    `json:"metode_bayar"`
	PembayaranBulan int       `json:"pembayaran_bulan"`
	TotalBayar      int       `json:"total_bayar"`
}

type InvestasiInput struct {
	Nama          string `json:"nama"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Usia          int    `json:"usia"`
	Email         string `json:"email"`
	Perokok       string `json:"perokok"`
	Nominal       int    `json:"nominal"`
	LamaInvestasi int    `json:"lama_investasi"`
	PeriodeBayar  string `json:"periode_pembayaran"`
	MetodeBayar   string `json:"metode_bayar"`
}

type InvestasiOutput struct {
	NoTransaction string `json:"no_transaction"`
	TglTransaksi  string `json:"tgl_transaksi"`
	Nama          string `json:"nama"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Usia          string `json:"usia"`
	Nominal       string `json:"nominal"`
	LamaInvestasi string `json:"lama_investasi"`
	PeriodeBayar  string `json:"periode_pembayaran"`
	MetodeBayar   string `json:"metode_bayar"`
	TotalBayar    string `json:"total_bayar"`
}

type InvestasiOutputAll struct {
	ID            int    `json:"id"`
	NoTransaction string `json:"no_transaction"`
	TglTransaksi  string `json:"tgl_transaksi"`
	Nama          string `json:"nama"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Usia          string `json:"usia"`
	Email         string `json:"email"`
	Perokok       string `json:"perokok"`
	Nominal       string `json:"nominal"`
	LamaInvestasi string `json:"lama_investasi"`
	PeriodeBayar  string `json:"periode_pembayaran"`
	MetodeBayar   string `json:"metode_bayar"`
	TotalBayar    string `json:"total_bayar"`
	Status        string `json:"status"`
}

type EditDataInput struct {
	NoTransaction string `json:"no_transaction"`
	Status        string `json:"status"`
	Supervisor    string `json:"supervisor"`
	Nama          string `json:"nama"`
	Usia          int    `json:"usia"`
}
