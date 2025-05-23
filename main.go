package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var idPost int
	fmt.Println("Ingrese un id de Post para obtenerlo")
	fmt.Scanf("%d", &idPost)

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", idPost)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la peticion", err)
		return

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error al hacer la peticion", err)
		return
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return
	}

	fmt.Println("Title: ", post.Title)
	fmt.Println("Contenido: ", post.Body)
}
