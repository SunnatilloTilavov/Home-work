package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"

	// Faylni ochish yoki mavjud bo'lsa ochib, yo'qligicha yaratish
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fayl ochishda xatolik:", err)
		return
	}
	defer file.Close()

	// Fayldan o'qish
	readData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Fayldan o'qishda xatolik:", err)
		return
	}

	// Agar faylda ma'lumot bo'lsa, uni chiqarib berish
	if len(readData) > 0 {
		data := []byte("Salom, Golang!\n")
		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Faylga yozishda xatolik:", err)
			return
		}

		readData, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Fayldan o'qishda xatolik:", err)
			return
		}
		fmt.Println("Fayldan o'qilgan ma'lumot:", string(readData))
		
		
	} else {
		// Fayl bo'sh bo'lsa yangi ma'lumotni yozish
		data := []byte("Salom, Golang!\n")
		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Faylga yozishda xatolik:", err)
			return
		}
		fmt.Println("Yangi ma'lumot faylga yozildi.")
		fmt.Println("Fayldan o'qilgan ma'lumot:", string(data))
	}
}
