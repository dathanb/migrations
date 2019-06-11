package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const UserJson = `{
    "id": -1,
    "reputation": "1",
    "creation_date": "2013-12-18T19:59:50.927",
    "display_name": "Community",
    "last_access_date": "2013-12-18T19:59:50.927",
    "website_url": "http://meta.stackexchange.com/",
    "location": "on the server farm",
    "about_me": "<p>Hi, I'm not really a person.</p>\n\n<p>I'm a background process that helps keep this site clean!</p>\n\n<p>I do things like</p>\n\n<ul>\n<li>Randomly poke old unanswered questions every hour so they get some attention</li>\n<li>Own community questions and answers so nobody gets unnecessary reputation from them</li>\n<li>Own downvotes on spam/evil posts that get permanently deleted</li>\n<li>Own suggested edits from anonymous users</li>\n<li><a href=\"http://meta.stackexchange.com/a/92006\">Remove abandoned questions</a></li>\n</ul>\n",
    "views": "0",
    "up_votes": "12",
    "down_votes": "432",
    "account_id": "-1"
}`

func TestUser_UnmarshalJSON(t *testing.T) {
	t.Run("Can unmarshal User from expected JSON", func(t *testing.T) {
		var user User
		err := json.Unmarshal([]byte(UserJson), &user)
		assert.NoError(t, err)
		assert.Equal(t, -1, user.Id)
		assert.Equal(t, "Community", user.DisplayName)
		assert.Equal(t, time.Date(2013, 12, 18, 19, 59, 50, 927000000, time.UTC), time.Time(user.CreationDate))
	})
}
