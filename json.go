package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title string  `json:"title"`
	NumeroVideos int `json:"numero_videos"`
}

type Cursos [] Curso

func main() {
	
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 5},
			Curso{"Curso de Ruby", 10},
			Curso{"Curso de Python", 25},
		}
		json.NewEncoder(writer).Encode(cursos)
	})

	http.ListenAndServe(":8000", nil)
}