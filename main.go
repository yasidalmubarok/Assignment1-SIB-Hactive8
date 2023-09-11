package main

import (
	"fmt"
	"os"
)

type namaTemanKelas struct {
	ID         uint8
	Name       string
	Address    string
	Occupation string
	Reason     string
}

var temanKelas []namaTemanKelas = []namaTemanKelas{
	{
		ID:         1,
		Name:       "Arga",
		Address:    "Semarang",
		Occupation: "Cyber Security",
		Reason:     "Karena bahasa Go memiliki Concurency",
	},
	{
		ID:         2,
		Name:       "Doanta",
		Address:    "Medan",
		Occupation: "Android Dev",
		Reason:     "Gak tau pengen aja belajar Golang",
	},
	{
		ID:         3,
		Name:       "Fanny",
		Address:    "Pekalongan",
		Occupation: "Game Dev",
		Reason:     "Bahasanya mudah cuy gk nyusahin(katanya)",
	},
	{
		ID:         4,
		Name:       "Saddam",
		Address:    "Jambi",
		Occupation: "Backend Dev",
		Reason:     "Ada Concurency sama Garbage Collection yang membuat saya tertarik mempelajari bahasa Go",
	},
}

func cariNamaTemanKelas(name string) *namaTemanKelas {
	for i, teman := range temanKelas {
		if teman.Name == name {
			return &temanKelas[i]
		}
	}
	return nil
}
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <namaTeman>")
		return
	}

	namaTeman := args[1]
	teman := cariNamaTemanKelas(namaTeman)

	if teman != nil {
		fmt.Println("ID: ", teman.ID)
		fmt.Println("Nama: ", teman.Name)
		fmt.Println("Alamat: ", teman.Address)
		fmt.Println("Pekerjaan: ", teman.Occupation)
		fmt.Println("Alasan: ", teman.Reason)
	}else {
		fmt.Printf("Teman yang bernama %s tidak ada pada daftar kelas", namaTeman)
	}
}
