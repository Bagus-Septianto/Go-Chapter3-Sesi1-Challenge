## Perintah Challenge :

Buatlah sebuah service untuk melakukan POST data dalam format JSON dan secara acak setiap **15 detik** dengan angka random antara **1-100** untuk value `water` dan `wind`. Gunakan url post berikut untuk menjalankan simulasi post request :
https://jsonplaceholder.typicode.com/posts

Kemudian tampilkan pada terminal **hasil post**nya. Selain itu kalian harus menentukan status `water` dan `wind` tersebut.
Dengan ketentuan :
- jika `water` dibawah 5 maka status aman
- jika `water` antara 6 - 8 maka status siaga
- jika `water` diatas 8 maka status bahaya
- jika `wind` dibawah 6 maka status aman
- jika `wind` antara 7 - 15 maka status siaga
- jika `wind` diatas 15 maka status bahaya
- value `water` dalam satuan meter
- value `wind` dalam satuan meter per detik

### Note: 
- di perintah diminta untuk angka random 1-100 tapi di ketentuan yang masuk kondisi hanya 1-15/15% (?), sisanya akan menjadi "bahaya" *tetap random 1-100 dan 85% keluaran `status bahaya`*
- hasil post dari [jsonplaceholder](https://jsonplaceholder.typicode.com/posts) menambahkan property baru yaitu `id` tapi di gambar challenge hanya ada `water` dan `wind`. *id masih dibiarkan ada/tidak dihapus dari response*
- ketentuan "value ... dalam satuan ..." tidak tahu mau digunakan dimana. di gambar contoh tidak dilihatkan ada tambahan "meter" dan "meter per detik". *belum diimplementasikan*