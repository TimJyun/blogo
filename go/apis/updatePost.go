package apis

import (
	"blogo/config"
	"blogo/ent"
	"blogo/ent/post"
	"context"
	"time"

	"github.com/kataras/iris/v12"
)

type UpdatePostRequest struct {
	Version uint32 `json:"version"`
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Token   string `json:"token"`
}

func UpdatePost(config config.Config, db *ent.Client) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		var request UpdatePostRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			ctx.Err()
			println("UpdatePost:read json err")
			return
		}

		if config.Password != request.Token {
			ctx.Err()
			println("UpdatePost:token is err")
			return
		}

		println("request.ID:", request.ID)
		println("request.Version:", request.Version)

		err = db.Post.UpdateOneID(request.ID).
			SetTitle(request.Title).
			SetContent(request.Content).
			SetLastUpdate(time.Now()).
			Where(post.Version(request.Version)).
			Exec(context.Background())

		if err != nil {
			println(err.Error())

		}

		ctx.JSON(Result{Ok: err == nil})
	}
}
