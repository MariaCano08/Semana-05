package main

import (
	"fmt"

	"./Multimedia"
)

func menu() {
	var id int
	detener:= false
	m := new(Multimedia.ContenidoWeb)
	var t, f, c string
	var d int
	var fr int64

	for !detener {
		fmt.Println("Que vas a capturar\n1)Imagen\n2)Audio\n3)Video\n4)Mostrar\n5)Salir")
		fmt.Scan(&id)
		if id == 1 {
			fmt.Println("Captura Imagen")
			a := new(Multimedia.Imagen)
			fmt.Println("Titulo:")
			fmt.Scan(&t)
			fmt.Println("Formato:")
			fmt.Scan(&f)
			fmt.Println("Canales:")
			fmt.Scan(&c)
			a.Titulo = t
			a.Formato = f
			a.Canales = c
			m.Multimedia = append(m.Multimedia, a)

		} else {
			if id == 2 {
				b := new(Multimedia.Audio)
				fmt.Println("Titulo:")
				fmt.Scan(&t)
				fmt.Println("Formato:")
				fmt.Scan(&f)
				fmt.Println("Duracion:")
				fmt.Scan(&d)
				b.Titulo = t
				b.Formato = f
				b.Duracion = d
				m.Multimedia = append(m.Multimedia, b)
			} else {
				if id == 3 {
					fmt.Println("Captura Video")
					v := new(Multimedia.Video)
					fmt.Println("Titulo:")
					fmt.Scan(&t)
					fmt.Println("Formato:")
					fmt.Scan(&f)
					fmt.Println("Freames:")
					fmt.Scan(&fr)
					v.Titulo = t
					v.Formato = f
					v.Frames = fr
					m.Multimedia = append(m.Multimedia, v)

				} else {
					if id == 4 {
						fmt.Print(m.Mostrar())
					} else {
						if id == 5 {
							detener = true
						}
					}
				}
			}
		}
	}
}

func main() {
	menu()
}