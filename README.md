# 💬 Komentar Analyzer

Program ini adalah aplikasi berbasis CLI (Command Line Interface) yang dirancang untuk **menyimpan, mengelola, dan menganalisis komentar** berdasarkan kata-kata yang ditentukan pengguna. Tujuan utamanya adalah untuk membantu pengguna memahami kecenderungan positif, netral, atau negatif dari komentar yang masuk secara sederhana dan terstruktur.

---

## ⚙️ Cara Kerja

Komentar yang dimasukkan ke dalam program akan **dibandingkan dengan sample kata** yang bisa diinput sendiri oleh pengguna. Terdapat 3 jenis sample yang dianalisis:

- ✅ **Good**: Berisi kata-kata yang umum digunakan dalam komentar positif.
- ❌ **Bad**: Berisi kata-kata yang umum digunakan dalam komentar negatif.
- 🔁 **Multiplier**: Berisi kata-kata penguat seperti "sangat", "paling", atau "sering", yang dapat meningkatkan bobot skor.

Setiap komentar akan diproses kata per kata, dibandingkan dengan sample yang tersedia, dan menghasilkan **skor numerik** yang merepresentasikan sentimen komentar:
- Skor diatas 1 → positif
- Skor dibawah 0.5 → negatif
- Skor diantara 1 dan 0.5 → netral

---

## ✨ Fitur

- 📝 Memasukkan komentar
- ✏️ Mengedit komentar
- ❌ Menghapus komentar
- 📥 Memasukkan sample kata (Good, Bad, Multiplier)
- 🛠️ Mengedit sample kata
- 🗑️ Menghapus sample kata
- 📊 Analisis skor komentar
- 🔢 Sorting komentar berdasarkan skor menggunakan **Insertion Sort**
- 🔤 Sorting komentar berdasarkan abjad menggunakan **Selection Sort**
- 📏 Sorting komentar berdasarkan panjang kalimat menggunakan **Selection Sort**
- 🔍 Pencarian komentar berdasarkan keyword menggunakan **Sequential Search**
- 🔍 Pencarian komentar berdasarkan kategori menggunakan **Binary Search**

---

## 🧪 Instalasi

Cukup **salin (copy)** seluruh folder proyek ini ke dalam komputer kamu, atau buka terminal di folder tujuan dan execute command dibawah :
```bash
git clone https://github.com/kieldinata/Tubes
```
---

## ▶️ Penggunaan

Masuk ke folder tempat file `main.go` berada, lalu jalankan perintah berikut di terminal:

```bash
go run main.go
```
