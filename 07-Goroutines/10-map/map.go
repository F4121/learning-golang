package map_test

/*
Map
- Golang memiliki sebuah struct bernama sync.Map
- Map ini mirip golang map, namun yang membedakan, Map ini aman untuk menggunakan concurrent menggunakan goroutine
- Ada beberapa function yang bisa kita gunakan di Map:
	- Store(key, value) untuk menyimpan data ke Map
	- Load(key) untuk menggambil data dari Map menggunakan key
	- Delete(key) untuk menghapus data dari Map menggunakan key
	- Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map
**/
