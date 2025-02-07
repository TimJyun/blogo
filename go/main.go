package main

import (
	"context"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"blogo/apis"
	"blogo/config"
	"blogo/ent"

	"github.com/iris-contrib/middleware/cors"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var config config.Config = config.GetOrInitConfig()

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	}))

	app.HandleDir("/", iris.Dir("./dist"), iris.DirOptions{IndexName: "index.html", SPA: true})

	app.Post("/api/getPostList", apis.GetPostList(config, client))

	app.Post("/api/getPost", apis.GetPost(config, client))

	app.Post("/api/deletePost", apis.DeletePost(config, client))

	app.Post("/api/createPost", apis.CreatePost(config, client))

	app.Post("/api/updatePost", apis.UpdatePost(config, client))

	app.Run(iris.Addr(":80"))

}
