package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const PostJson = `{
	"id": 1,
	"post_type_id": "1",
	"creation_date": "2013-12-18T20:25:25.003",
	"score": "31",
	"view_count": "14902",
	"body": "Some post body",
	"owner_user_id": "3",
	"last_activity_date": "2013-12-19T04:18:10.800",
	"title": "How does the Kindle's reading rate algorithm work?",
	"tags": "<kindle><time-to-read><kindle-touch>",
	"answer_count": "1",
	"comment_count": "0",
	"favorite_count": "5"
}`

func TestPosts_UnmarshalJSON(t *testing.T)  {
	t.Run("Can unmarshal from expected JSON format", func(t *testing.T) {
		var post Post
		err := json.Unmarshal([]byte(PostJson), &post)
		assert.NoError(t, err)

		assert.Equal(t, 1, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, 3, post.UserId)
		assert.Equal(t, "Some post body", post.Body)
		assert.Equal(t, time.Date(2013, 12, 18, 20, 25, 25, 3000000, time.UTC), time.Time(post.CreationDate))
	})
}
