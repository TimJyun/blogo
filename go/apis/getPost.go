package apis

import (
	"blogo/config"
	"blogo/ent"
	"blogo/ent/post"
	"context"
	"fmt"

	"github.com/kataras/iris/v12"
)

type GetPostRequest struct {
	ID uint32 `json:"id"`
}

func GetPost(config config.Config, db *ent.Client) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		var request GetPostRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			fmt.Println("GetPost:read json error")
			ctx.Err()
			return
		}

		posts, err := db.Post.Query().Where(post.ID(request.ID)).All(context.Background())
		if err != nil || len(posts) != 1 {
			fmt.Println("GetPost:query post error")
			ctx.Err()
			return
		}

		ctx.JSON(&posts[0])
	}
}
