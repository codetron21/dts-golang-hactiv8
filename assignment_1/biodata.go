package main

import (
	"assignment_1/resources"
	"fmt"
	"os"
	"strconv"
)

type Student = resources.Student

func init() {
	resources.AddStudents(createResources()...)
}

func main() {
	absentNumber, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Nomor absen tidak dikenali")
		return
	}

	result, isFounded := resources.FindStudent(absentNumber)

	if isFounded {
		printStudent(result)
	} else {
		fmt.Printf("Teman kelas dengan nomor absen %d tidak ditemukan", absentNumber)
	}
}

func printStudent(student Student) {
	fmt.Printf(
		`
Nama: %s
Alamat: %s
Pekerjaan: %s
Alasan: %s
		`,
		student.Nama, student.Alamat, student.Pekerjaan, student.Alasan,
	)
}

func createResources() []Student {
	return []Student{
		{
			Nama:      "Adadua karunia putera",
			Alamat:    "Depok",
			Pekerjaan: "-",
			Alasan:    "Karena menarik",
		},
		{
			Nama:      "Ramli",
			Alamat:    "Jakarta",
			Pekerjaan: "Programmer",
			Alasan:    "Karena terpaksa",
		},
		{
			Nama:      "Budi",
			Alamat:    "Bandung",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Karena bagus",
		},
	}
}
