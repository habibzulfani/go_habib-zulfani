package main

// Kendaraan adalah tipe data untuk kendaraan umum.
type Kendaraan struct {
    TotalRoda        int
    KecepatanPerJam  int
}

// Mobil adalah tipe data untuk mobil yang merupakan kendaraan.
type Mobil struct {
    Kendaraan
}

// Berjalan meningkatkan kecepatan mobil sebanyak kecepatan baru.
func (m *Mobil) Berjalan() {
    m.TambahKecepatan(10)
}

// TambahKecepatan menambahkan kecepatan mobil dengan kecepatan baru.
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
    m.KecepatanPerJam += kecepatanBaru
}

func main() {
    // MobilCepat adalah sebuah objek mobil.
    mobilCepat := Mobil{}
    mobilCepat.Berjalan()
    mobilCepat.Berjalan()
    mobilCepat.Berjalan()

    // MobilLamban adalah sebuah objek mobil lainnya.
    mobilLamban := Mobil{}
    mobilLamban.Berjalan()
}
