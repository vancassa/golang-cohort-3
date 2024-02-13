package main

import (
	"fmt"
	"os"
	"strings"
)

type Student struct {
	id					int
	name				string
	address			string
	occupation	string
	reason			string
}

func main() {
	students := []Student{
		{ id: 1, name: "Abdul", address: "Jl. Satu", occupation: "Arsitek", reason: "Alasan Abdul" },
		{ id: 2, name: "Budi", address: "Jl. Dua", occupation: "Budayawan", reason: "Alasan Budi" },
		{ id: 3, name: "Chandra", address: "Jl. Tiga", occupation: "Camat", reason: "Alasan Chandra" },
		{ id: 4, name: "Deni", address: "Jl. Empat", occupation: "Damkar", reason: "Alasan Deni" },
		{ id: 5, name: "Emon", address: "Jl. Lima", occupation: "Engineer", reason: "Alasan Emon" },
	}

	if (len(os.Args) < 2) {
		fmt.Println("Tolong masukan nama atau nomor absen")	
		fmt.Println("Contoh: 'go run main.go Fitri' atau 'go run main.go 2'")
		return
	}

	cliValues := os.Args[1:]
	found := false

	for _, cliVal := range cliValues {
		for _, student := range students {
			if (cliVal == fmt.Sprint(student.id) || strings.ToLower(cliVal) == strings.ToLower(student.name)) {
				if (found) {
					fmt.Println(strings.Repeat("=", 20))
				}

				printInfo(student)
				found = true
			}
		}
	}

	if (!found) {
		fmt.Println("Data tidak ditemukan")
	}
}

func printInfo (student Student) {
	fmt.Println("ID:", student.id)
	fmt.Println("nama:", student.name)
	fmt.Println("alamat:", student.address)
	fmt.Println("pekerjaan:", student.occupation)
	fmt.Println("alasan:", student.reason)
}