package main

/* 	Data Kelompok */
/*	Judul							: Aplikasi Daring 1
		Anggota Kelompok	:	1301194204 Mar'i Fauzan Rambe
												1301194183 Sabrina Adinda Sari
*/

import  (
	"fmt"
	"os"
  "os/exec"
  "runtime"
)

const N = 99999

type data struct {
	id int
	sekul string
  kls int
  pel string
  smt int
  silb string
  mat string
}
var tab [N]data
var tabCadangan [N]data
var i , tot, id int

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() {
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

var key string

// Menu Utama
func main() {
		var x string

		CallClear()
    fmt.Println(".::Selamat Datang di aplikasi DARING 1::.")
		fmt.Println("==========================================\n")
		fmt.Println(".:::::::::::::::::::::::::::::::::::::::::.")
		fmt.Print("(1. Insert)  (2. View)  (0. Exit)\n")
		fmt.Print(":::::::::::::::::::::::::::::::::::::::::::\nPilih : ")
		fmt.Scan(&x)
		CallClear()
	  if x == "1" {
	  	insert()
	    main()
	  } else if x == "2" {
			menuView()
	  } else if x == "0" {
			fmt.Println("Thank you, See You Again")
		} else {
			main()
		}
}

// Menu View
func menuView() {
	var x string

	fmt.Print(".::::::::::::::::::::::::::::::::::::::::::::::::::::::::.\n")
	fmt.Print("(1. Edit)  (2. Delete)  (3. Cari)  (4. Sorting)  (0. Exit)\n")
	fmt.Print("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::\nPilih : ")
	fmt.Scan(&x)
	CallClear()
	if x == "1" {
		menuEdit()
	} else if x == "2" {
		menuDelete()
	} else if x == "3" {
		menuCari()
	} else if x == "4" {
		menuSort()
	} else if x == "0" {
		main()
	} else {
		menuView()
	}
}

// Insert
func insert() {
  var confirm string

	i = tot
	fmt.Print("::::::::::::::::::::::::::::::::::::::::::::\n")
  fmt.Print("Apakah Ingin Input : ")
  fmt.Scan(&confirm)
	fmt.Print("::::::::::::::::::::::::::::::::::::::::::::\n")
  for (confirm != "no") && (i < N) {
		id++
		tab[i].id = id
		fmt.Print("Sekolah		: ")
	  fmt.Scan(&tab[i].sekul)
	  fmt.Print("Kelas		: ")
	  fmt.Scan(&tab[i].kls)
	  fmt.Print("Pelajaran	: ")
	  fmt.Scan(&tab[i].pel)
	  fmt.Print("Semester	: ")
	  fmt.Scan(&tab[i].smt)
	  fmt.Print("Silabus		: ")
	  fmt.Scan(&tab[i].silb)
	  fmt.Print("Materi		: ")
	  fmt.Scan(&tab[i].mat)
	  CallClear()
		i++
    tot++
		fmt.Print("::::::::::::::::::::::::::::::::::::::::::::\n")
    fmt.Print("Apakah anda ingin input lagi : ")
    fmt.Scan(&confirm)
		fmt.Print("::::::::::::::::::::::::::::::::::::::::::::\n")
	}
  CallClear()
}

// Selection Sort - Sekolah
func sortS() {
	var pass,i,idx_min int
	var temp data

	tabCadangan = tab
	i = 0
	pass = 0
	for pass < tot - 1 {
		idx_min = pass
		i = pass + 1
		for i < tot {
			if tabCadangan[idx_min].sekul > tabCadangan[i].sekul {
				idx_min = i
			}
			i++
		}
		temp = tabCadangan[idx_min]
		tabCadangan[idx_min] = tabCadangan[pass]
		tabCadangan[pass] = temp
		pass++
	}
}

// Selection Sort - Materi
func sortK() {
	var pass,i,idx_min int
	var temp data

	i = 0
	tabCadangan = tab
	pass = 0
	for pass < tot - 1 {
		idx_min = pass
		i = pass + 1
		for i < tot {
			if tabCadangan[idx_min].mat > tabCadangan[i].mat {
				idx_min = i
			}
			i++
		}
		temp = tabCadangan[idx_min]
		tabCadangan[idx_min] = tabCadangan[pass]
		tabCadangan[pass] = temp
		pass++
	}
}

// Binary Search - Sekolah
func binaryS(x *string) {
	var btm,top,mid int
	var found bool

	sortS()
	btm = 0
	top = tot
	found = false
	for top >= btm && found == false {
		mid = (btm + top) / 2
		if *x < tabCadangan[mid].sekul {
			top = mid - 1
		} else if *x > tabCadangan[mid].sekul {
			btm = mid + 1
		} else {
			found = true
			if tabCadangan[mid].id != 0 {
				fmt.Println(".::::::::::::::::::::::::::::::::::::.")
				fmt.Println("ID		: ",tabCadangan[mid].id)
	      fmt.Println("Sekolah		: ",tabCadangan[mid].sekul)
				fmt.Println("Kelas		: ",tabCadangan[mid].kls)
				fmt.Println("Pelajaran	: ",tabCadangan[mid].pel)
				fmt.Println("Semester	: ",tabCadangan[mid].smt)
				fmt.Println("Silabus		: ",tabCadangan[mid].silb)
				fmt.Println("Materi		: ",tabCadangan[mid].mat)
				fmt.Println("::::::::::::::::::::::::::::::::::::::")
			}
		}
	}
	if found == false {
		fmt.Println("Data Tidak Ada")
	}
}

// Binary Search - Materi
func binaryM(x *string) {
	var btm,top,mid int
	var found bool

	sortK()
	btm = 0
	top = tot
	found = false
	for top >= btm && found == false {
		mid = (btm + top) / 2
		if *x < tabCadangan[mid].mat {
			top = mid - 1
		} else if *x > tabCadangan[mid].mat {
			btm = mid + 1
		} else {
			if tabCadangan[mid].id != 0 {
				found = true
				fmt.Println(".::::::::::::::::::::::::::::::::::::.")
				fmt.Println("ID		: ",tabCadangan[mid].id)
	      fmt.Println("Sekolah		: ",tabCadangan[mid].sekul)
				fmt.Println("Kelas		: ",tabCadangan[mid].kls)
				fmt.Println("Pelajaran	: ",tabCadangan[mid].pel)
				fmt.Println("Semester	: ",tabCadangan[mid].smt)
				fmt.Println("Silabus		: ",tabCadangan[mid].silb)
				fmt.Println("Materi		: ",tabCadangan[mid].mat)
				fmt.Println("::::::::::::::::::::::::::::::::::::::")
			}
		}
	}
	if found == false {
		fmt.Println("Data Tidak Ada")
	}
}

// Sequential Search
func sequential(x,y *string) {
	var sum int
	var s,m string

	s = *x
	m = *y
	i = 0
	for (i < tot) {
  	if (s == tab[i].sekul || m == tab[i].mat) {
			if tab[i].id != 0 {
				fmt.Println(".::::::::::::::::::::::::::::::::::::.")
				fmt.Println("ID		: ",tab[i].id)
	      fmt.Println("Sekolah		: ",tab[i].sekul)
				fmt.Println("Kelas		: ",tab[i].kls)
				fmt.Println("Pelajaran	: ",tab[i].pel)
				fmt.Println("Semester	: ",tab[i].smt)
				fmt.Println("Silabus		: ",tab[i].silb)
				fmt.Println("Materi		: ",tab[i].mat)
				fmt.Println("::::::::::::::::::::::::::::::::::::::")
				sum++
			}
		}
    i++
  }
	if sum == 0 {
		fmt.Println("Data Tidak Ditemukan")
	}
}

// Menu Search
func menuCari() {
  var x int
	var s,m string

  i = 0
	fmt.Println(".::::::::::::::::::::::::::::::::::.")
	fmt.Print("(1. Sequential)  (2. Binary)  (0. Exit)\n")
	fmt.Print("====================================\nPilih Metode Pencarian : ")
	fmt.Scan(&x)
	if (x == 1) {  					// Sequential Search
		fmt.Println(".::::::::::::::::::::::::::::::::::.")
		fmt.Print("(1. Sekolah)  (2. Materi)  (0. Exit)\n")
		fmt.Print("====================================\nCari Berdasarkan : ")
		fmt.Scan(&x)
		if (x == 1) {
			fmt.Print("Sekolah	: ")
			fmt.Scan(&s)
		} else if (x == 2) {
			fmt.Print("Materi	: ")
			fmt.Scan(&m)
		} else if (x == 0) {
			CallClear()
			main()
		} else {
			menuCari()
		}
		sequential(&s,&m)
		fmt.Print("Ketik '0' jika ingin Exit : ")
		fmt.Scan(&x)
		if x == 0 {
			CallClear()
			main()
		} else {
			menuCari()
		}
	} else if (x == 2) {    // Binary Search
		fmt.Println(".::::::::::::::::::::::::::::::::::.")
		fmt.Print("(1. Sekolah)  (2. Materi)  (0. Exit)\n")
		fmt.Print("====================================\nCari Berdasarkan : ")
		fmt.Scan(&x)
		if (x == 1) {
			fmt.Print("Sekolah	: ")
			fmt.Scan(&s)
			binaryS(&s)
		} else if (x == 2) {
			fmt.Print("Materi	: ")
			fmt.Scan(&m)
			binaryM(&m)
		} else if (x == 0) {
			CallClear()
			main()
		} else {
			menuCari()
		}
	} else if (x == 0) {
		CallClear()
		menuView()
	} else {
		menuCari()
	}
	fmt.Print("Ketik '0' jika ingin Exit : ")
	fmt.Scan(&x)
	if x == 0 {
		CallClear()
		main()
	} else {
		menuCari()
	}
}

// Delete
func menuDelete() {
	var sum int
	var x string

	i = 0
	fmt.Print("Input Materi yang ingin di hapus : ")
	fmt.Scan(&x)

	for (i < tot) {
		if tab[i].mat == x {
			//found = true
			tab[i].id = 0
			tab[i].sekul = ""
		  tab[i].kls = 0
		  tab[i].pel = ""
		  tab[i].smt = 0
		  tab[i].silb = ""
		  tab[i].mat = ""
			sum++
		}
		i++
	}
	if sum > 0 {
		fmt.Println("Data Berhasil Dihapus")
	} else {
		fmt.Println("Data Tidak Ada")
	}

	fmt.Print("Ketik '0' jika ingin Exit : ")
	fmt.Scan(&x)
	if x == "0" {
		CallClear()
		menuView()
	} else {
		menuDelete()
	}
}

// Edit
func menuEdit(){
	var x,sum int
	var c string

	i = 0
	c = "no"

	for c != "yes" {
		fmt.Print("Input ID data yang ingin di Edit : ")
		fmt.Scan(&x)

		for (i < tot) {
			if tab[i].id == x && tab[i].id != 0 {
				fmt.Println(".::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::.")
				fmt.Print("(1. Sekolah)  (2. Kelas) (3. Pelajaran)  (4. Semester)  (5. Silabus)  (6. Materi)  (0. Exit)\n")
				fmt.Print("Pilih yang akan di Edit : ")
				fmt.Scan(&x)
				fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

				switch x {
			    case 1:
						fmt.Print("Input Sekolah : ")
						fmt.Scan(&tab[i].sekul)
			    case 2:
						fmt.Print("Input Kelas : ")
						fmt.Scan(&tab[i].kls)
			    case 3:
						fmt.Print("Input Pelajaran : ")
						fmt.Scan(&tab[i].pel)
					case 4:
						fmt.Print("Input Semester : ")
						fmt.Scan(&tab[i].smt)
					case 5:
						fmt.Print("Input Silabus : ")
						fmt.Scan(&tab[i].silb)
					case 6:
						fmt.Print("Input Materi : ")
						fmt.Scan(&tab[i].mat)
					default :
						menuView()
			    }
					sum++

			}

			i++
		}

		if sum == 0 {
			fmt.Println("Data Tidak Ada")
		}

		i = 0
		sum = 0
		fmt.Print("Sudah Selesai Edit : ")
		fmt.Scan(&c)
		CallClear()
	}
	fmt.Println("Data Berhasil Diedit")
	fmt.Print("Ketik '0' jika ingin Exit : ")
	fmt.Scan(&x)
	if x == 0 {
		CallClear()
		menuView()
	} else {
		menuEdit()
	}
}

// Menu Sort
func menuSort() {
	var a,b,x int
	var pass,idx_min int
	var temp data

	i = 0
	CallClear()
	tabCadangan = tab
	fmt.Print(".:::::::::::::::::::::::::::::.\n")
	fmt.Print("(1. Ascending)  (2. Descending)\n")
	fmt.Print("Pilih urutan sorting : ")
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print("(1. Sekolah)  (2. Materi)\n")
		fmt.Print("Pilih Berdasarkan : ")
	  fmt.Scan(&b)
	  if b == 1 {
  		// Selection Sort - Ascending
			// Berdasarkan Sekolah
			pass = 0
			for pass < tot - 1 {
				idx_min = pass
				i = pass + 1
				for i < tot {
					if tabCadangan[idx_min].sekul > tabCadangan[i].sekul {
						idx_min = i
					}
					i++
				}
				temp = tabCadangan[idx_min]
				tabCadangan[idx_min] = tabCadangan[pass]
				tabCadangan[pass] = temp
				pass++
			}
	    hasilSort()
	  } else if b == 2 {
			// Berdasarkan Materi
			pass = 0
			for pass < tot - 1 {
				idx_min = pass
				i = pass + 1
				for i < tot {
					if tabCadangan[idx_min].mat > tabCadangan[i].mat {
						idx_min = i
					}
					i++
				}
				temp = tabCadangan[idx_min]
				tabCadangan[idx_min] = tabCadangan[pass]
				tabCadangan[pass] = temp
				pass++
			}
	  	hasilSort()
	  }
	} else if a == 2 {
		fmt.Print("(1. Sekolah)  (2. Materi)\n")
		fmt.Print("Pilih Berdasarkan : ")
	  fmt.Scan(&b)
	  if b == 1 {
    	//Insertion Sort - Descending
			pass = -1
			for pass <= tot - 1 {
				i = pass + 1
				temp = tabCadangan[i]
				for i > 0 && temp.sekul > tabCadangan[i-1].sekul {
					tabCadangan[i] = tabCadangan[i-1]
					i--
				}
				tabCadangan[i] = temp
				pass++
			}
	    hasilSort()
	  } else if b == 2 {
			pass = -1
			for pass <= tot - 1 {
				i = pass + 1
				temp = tabCadangan[i]
				for i > 0 && temp.mat > tabCadangan[i-1].mat {
					tabCadangan[i] = tabCadangan[i-1]
					i = i - 1
				}
				tabCadangan[i] = temp
				pass = pass + 1
			}
	    hasilSort()
	  }
	}

	fmt.Print("Ketik '0' jika ingin Exit : ")
	fmt.Scan(&x)
	if x == 0 {
		CallClear()
		menuView()
	} else {
		menuSort()
	}

}

// Cetak Hasil Sort (Hasil Sort Tidak Permanen)
func hasilSort() {
	for i := 0; i < tot; i++ {
		if tabCadangan[i].id != 0 {
			fmt.Println("::::::::::::::::::::::::::::::::::::::")
			fmt.Println("ID		: ",tabCadangan[i].id)
			fmt.Println("Sekolah		: ",tabCadangan[i].sekul)
			fmt.Println("Kelas		: ",tabCadangan[i].kls)
			fmt.Println("Pelajaran	: ",tabCadangan[i].pel)
			fmt.Println("Semester	: ",tabCadangan[i].smt)
			fmt.Println("Silabus		: ",tabCadangan[i].silb)
			fmt.Println("Materi		: ",tabCadangan[i].mat)
		}
	}
	fmt.Println()
}
