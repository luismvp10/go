package main

import (
	"net/http"

)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//http.ServeFile(writer,request,"./index.html")

		//fmt.Println(request.URL.Path)
		//http.ServeFile(writer,request, request.URL.Path[1:])
		fileServer := http.FileServer(http.Dir("public"))
		http.Handle("/", http.StripPrefix("/", fileServer))

	})
	http.ListenAndServe(":8000", nil)

}