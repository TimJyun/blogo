package apis

import (
	"blogo/ent"
	"context"
	"fmt"
	"time"

	"blogo/config"

	"github.com/kataras/iris/v12"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Token   string `json:"token"`
}

func CreatePost(config config.Config, db *ent.Client) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		var request CreatePostRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			fmt.Println("CreatePost:read json error")
			ctx.Err()
			return
		}

		if config.Password != request.Token {
			fmt.Println("CreatePost:token is error")
			ctx.Err()
			return
		}

		err = db.Post.Create().
			SetTitle(request.Title).
			SetContent(request.Content).
			SetVersion(0).
			SetCreateTime(time.Now()).
			SetLastUpdate(time.Now()).
			Exec(context.Background())

		if err != nil {
			fmt.Println("CreatePost:create post error")
		}

		ctx.JSON(Result{Ok: err == nil})

	}
}
