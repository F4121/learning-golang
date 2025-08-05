package tabletests

/*
Table Test
- Sebelumnya kita sudah belajar tentang sub test
- Jika ddiperhatikan sebeneranya dengan sub test, kita bisa membbuat test secara dinamis
- Dan fitur sub test ini biasanya digunakan oleh programmer golang untuk membbuat test dengan konsep table test
- Table test yaitu dimana kita menyediakan data berupa slice yang berisi paramenter dan ekspektasi hasil dari unit test
- Lalu slice tersebut kita iterasi menggunakan sub test
**/

func HelloWorld(name string) string {
	return "Hello, " + name
}
