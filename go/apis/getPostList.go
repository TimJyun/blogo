package apis

import (
	"blogo/config"
	"blogo/ent"
	"context"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

type PostPreview struct {
	ID         uint32    `json:"id"`
	Title      string    `json:"title"`
	LastUpdate time.Time `json:"lastUpdate"`
}

func GetPostList(config config.Config, db *ent.Client) func(ctx iris.Context) {

	return func(ctx iris.Context) {
		posts, err := db.Post.Query().All(context.Background())
		if err != nil {
			fmt.Println("GetPostList:query failed")
			ctx.Err()
			return
		}

		var postList []PostPreview = make([]PostPreview, len(posts))

		for idx, post := range posts {
			postList[idx] = PostPreview{
				ID:         post.ID,
				Title:      post.Title,
				LastUpdate: post.LastUpdate,
			}
		}

		fmt.Println("postList-len:", len(postList))

		ctx.JSON(&postList)

	}
}
