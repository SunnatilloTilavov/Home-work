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

	// Agar faylda ma'lumot bo'lsa, uni davomidan qo'shish
	if len(readData) > 0 {
		// Faylni oxirgi joyidan boshlab o'qib chiqish
		_, err := file.Seek(-1, os.SEEK_END)
		if err != nil {
			fmt.Println("Faylni oxirgi joyidan o'qib chiqishda xatolik:", err)
			return
		}

		for i := 1; i <= 100; i++ {
			newData := fmt.Sprintf("\n %v", i)
			_, err = file.Write([]byte(newData))
			if err != nil {
				fmt.Println("Faylga yozishda xatolik:", err)
				return
			}
	}

		// O'zgartirilgan ma'lumotni qayta o'qish
		updatedData, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Fayldan o'qishda xatolik:", err)
			return
		}
		fmt.Println("O'zgartirilgan ma'lumot:", string(updatedData))
	} else {
		// Fayl bo'sh bo'lsa yangi ma'lumotni yozish
		data := []byte("0\n")
		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Faylga yozishda xatolik:", err)
			return
		}
		fmt.Println("Yangi ma'lumot faylga yozildi.")
		fmt.Println("Fayldan o'qilgan ma'lumot:", string(data))
	}
}
