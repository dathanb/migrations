package client

import (
	"encoding/xml"
	"github.com/dathanb/migrations/fakestack/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

const PostXml = `<row Id="1" PostTypeId="1" AcceptedAnswerId="3" CreationDate="2016-08-02T15:39:14.947" Score="8" ViewCount="384" Body="Question 1 body" OwnerUserId="8" LastEditorUserId="10135" LastEditDate="2018-10-18T10:45:18.660" LastActivityDate="2018-10-18T10:45:18.660" Title="What is &quot;backprop&quot;?" Tags="&lt;neural-networks&gt;&lt;backpropagation&gt;&lt;terminology&gt;&lt;definitions&gt;" AnswerCount="3" CommentCount="3" FavoriteCount="1" />`
const PostXmlWithLeadingCharData = `\n<row Id="1" PostTypeId="1" AcceptedAnswerId="3" CreationDate="2016-08-02T15:39:14.947" Score="8" ViewCount="384" Body="Question 1 body" OwnerUserId="8" LastEditorUserId="10135" LastEditDate="2018-10-18T10:45:18.660" LastActivityDate="2018-10-18T10:45:18.660" Title="What is &quot;backprop&quot;?" Tags="&lt;neural-networks&gt;&lt;backpropagation&gt;&lt;terminology&gt;&lt;definitions&gt;" AnswerCount="3" CommentCount="3" FavoriteCount="1" />`
const TwoPostsXml = `<row Id="1" PostTypeId="1" AcceptedAnswerId="3" CreationDate="2016-08-02T15:39:14.947" Score="8" ViewCount="384" Body="Question 1 body" OwnerUserId="8" LastEditorUserId="10135" LastEditDate="2018-10-18T10:45:18.660" LastActivityDate="2018-10-18T10:45:18.660" Title="What is &quot;backprop&quot;?" Tags="&lt;neural-networks&gt;&lt;backpropagation&gt;&lt;terminology&gt;&lt;definitions&gt;" AnswerCount="3" CommentCount="3" FavoriteCount="1" />
  <row Id="2" PostTypeId="1" AcceptedAnswerId="9" CreationDate="2016-08-02T15:40:20.623" Score="10" ViewCount="404" Body="Question 2 body" OwnerUserId="8" LastEditorUserId="2444" LastEditDate="2019-02-23T22:36:19.090" LastActivityDate="2019-02-23T22:36:37.133" Title="How does noise affect generalization?" Tags="&lt;neural-networks&gt;&lt;machine-learning&gt;&lt;statistical-ai&gt;&lt;generalization&gt;" AnswerCount="3" CommentCount="0" FavoriteCount="1" />`
const PostsXml = `<?xml version="1.0" encoding="utf-8"?>
<posts>
  <row Id="1" PostTypeId="1" AcceptedAnswerId="3" CreationDate="2016-08-02T15:39:14.947" Score="8" ViewCount="384" Body="Question 1 body" OwnerUserId="8" LastEditorUserId="10135" LastEditDate="2018-10-18T10:45:18.660" LastActivityDate="2018-10-18T10:45:18.660" Title="What is &quot;backprop&quot;?" Tags="&lt;neural-networks&gt;&lt;backpropagation&gt;&lt;terminology&gt;&lt;definitions&gt;" AnswerCount="3" CommentCount="3" FavoriteCount="1" />
  <row Id="2" PostTypeId="1" AcceptedAnswerId="9" CreationDate="2016-08-02T15:40:20.623" Score="10" ViewCount="404" Body="Question 2 body" OwnerUserId="8" LastEditorUserId="2444" LastEditDate="2019-02-23T22:36:19.090" LastActivityDate="2019-02-23T22:36:37.133" Title="How does noise affect generalization?" Tags="&lt;neural-networks&gt;&lt;machine-learning&gt;&lt;statistical-ai&gt;&lt;generalization&gt;" AnswerCount="3" CommentCount="0" FavoriteCount="1" />
</posts>`


func Test_ReadPost(t *testing.T) {
	t.Run("Can read a post from XML", func(t *testing.T) {
		reader := strings.NewReader(PostXml)
		decoder := xml.NewDecoder(reader)
		post, err := readPost(decoder)

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T15:39:14.947")

		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 1, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 1 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)
	})
	t.Run("Can read a post from XML with leading CharData", func(t *testing.T) {
		reader := strings.NewReader(PostXmlWithLeadingCharData)
		decoder := xml.NewDecoder(reader)
		post, err := readPost(decoder)

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T15:39:14.947")

		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 1, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 1 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)
	})
	t.Run("Can read two users in a row from XML", func(t *testing.T) {
		reader := strings.NewReader(TwoPostsXml)
		decoder := xml.NewDecoder(reader)
		post, err := readPost(decoder)

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T15:39:14.947")

		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 1, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 1 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)

		post, err = readPost(decoder)

		expectedTime, err = time.Parse(models.TimeFormat, "2016-08-02T15:40:20.623")

		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 2, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 2 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)
	})
}

func Test_ReadPosts(t *testing.T) {
	t.Run("Can read all users from XML", func(t *testing.T) {
		reader := strings.NewReader(PostsXml)

		posts := make(chan models.Post)
		go readPosts(reader, posts)

		post := <- posts

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T15:39:14.947")
		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 1, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 1 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)

		post = <- posts

		expectedTime, err = time.Parse(models.TimeFormat, "2016-08-02T15:40:20.623")

		assert.Nil(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, 2, post.Id)
		assert.Equal(t, 1, post.PostType)
		assert.Equal(t, "Question 2 body", post.Body)
		assert.Equal(t, models.Time(expectedTime), post.CreationDate)

		_, closed := <- posts
		assert.Equal(t, true, closed)
	})
}
