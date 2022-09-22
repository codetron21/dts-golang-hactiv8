package resources

type Student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var students []Student = make([]Student, 0)

func AddStudents(arguments ...Student) {
	students = append(students, arguments...)
}

func GetStudents() []Student {
	return students
}

func FindStudent(index int) (student Student, founded bool) {
	// supaya absen hanya menerima bilangan positif
	if index <= 0 || index > len(students) {
		return
	}

	student = students[index-1]
	studentPtr := &student

	if studentPtr != nil {
		founded = true
	}

	//if(student != Student{})

	// for i, v := range students {
	// 	if i+1 == index {
	// 		founded = true
	// 		student = v
	// 		return
	// 	}
	// }

	return
}
