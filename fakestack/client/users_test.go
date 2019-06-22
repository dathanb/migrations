package client

import (
	"encoding/xml"
	"github.com/dathanb/migrations/fakestack/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

const UserXml = "<row Id=\"-1\" Reputation=\"1\" CreationDate=\"2016-08-02T00:14:10.580\" DisplayName=\"Community\" LastAccessDate=\"2016-08-02T00:14:10.580\" Location=\"on the server farm\" AboutMe=\"&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;\" Views=\"0\" UpVotes=\"3\" DownVotes=\"661\" AccountId=\"-1\" />"
const UserXmlWithLeadingCharData = "\n<row Id=\"-1\" Reputation=\"1\" CreationDate=\"2016-08-02T00:14:10.580\" DisplayName=\"Community\" LastAccessDate=\"2016-08-02T00:14:10.580\" Location=\"on the server farm\" AboutMe=\"&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;\" Views=\"0\" UpVotes=\"3\" DownVotes=\"661\" AccountId=\"-1\" />"
const TwoUsersXml = `"<row Id="-1" Reputation="1" CreationDate="2016-08-02T00:14:10.580" DisplayName="FirstUser" LastAccessDate="2016-08-02T00:14:10.580" Location="on the server farm" AboutMe="&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;" Views="0" UpVotes="3" DownVotes="661" AccountId="-1" />"
"<row Id="1" Reputation="1" CreationDate="2016-08-02T00:14:10.580" DisplayName="SecondUser" LastAccessDate="2016-08-02T00:14:10.580" Location="on the server farm" AboutMe="&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;" Views="0" UpVotes="3" DownVotes="661" AccountId="-1" />"`
const UsersXml = `<?xml version="1.0" encoding="utf-8"?>
<users>
  <row Id="-1" Reputation="1" CreationDate="2016-08-02T00:14:10.580" DisplayName="Community" LastAccessDate="2016-08-02T00:14:10.580" Location="on the server farm" AboutMe="&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;" Views="0" UpVotes="3" DownVotes="661" AccountId="-1" />
  <row Id="1" Reputation="101" CreationDate="2016-08-02T15:36:45.333" DisplayName="Adam Lear" LastAccessDate="2019-04-05T19:06:40.560" Location="New York, NY" AboutMe="&#xA;&#xA;&lt;p&gt;Developer at Stack Overflow focusing on public Q&amp;amp;A products. Canadian working in the American idiom.&lt;/p&gt;&#xA;&#xA;&lt;p&gt;Once upon a time:&lt;/p&gt;&#xA;&#xA;&lt;ul&gt;&#xA;&lt;li&gt;community manager at Stack Overflow&lt;/li&gt;&#xA;&lt;li&gt;elected moderator on Stack Overflow and Software Engineering&lt;/li&gt;&#xA;&lt;li&gt;desktop software developer ¯\_(ツ)_/¯ &lt;/li&gt;&#xA;&lt;/ul&gt;&#xA;&#xA;&lt;p&gt;Email me a link to your favorite Wikipedia article: &lt;code&gt;adam@stackoverflow.com&lt;/code&gt;.&lt;/p&gt;&#xA;" Views="103" UpVotes="0" DownVotes="0" ProfileImageUrl="https://i.stack.imgur.com/SMEGn.jpg?s=128&amp;g=1" AccountId="37099" />
</users>`

func Test_ReadUser(t *testing.T) {
	t.Run("Can read a user from XML", func(t *testing.T) {
		reader := strings.NewReader(UserXml)
		decoder := xml.NewDecoder(reader)
		user, err := readUser(decoder)

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T00:14:10.580")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, -1, user.Id)
		assert.Equal(t, "Community", user.DisplayName)
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)
	})
	t.Run("Can read a user from XML with leading CharData", func(t *testing.T) {
		reader := strings.NewReader(UserXmlWithLeadingCharData)
		decoder := xml.NewDecoder(reader)
		user, err := readUser(decoder)

		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T00:14:10.580")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, -1, user.Id)
		assert.Equal(t, "Community", user.DisplayName)
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)
	})
	t.Run("Can read two users in a row from XML", func(t *testing.T) {
		expectedTime, err := time.Parse(models.TimeFormat, "2016-08-02T00:14:10.580")

		reader := strings.NewReader(TwoUsersXml)
		decoder := xml.NewDecoder(reader)
		user, err := readUser(decoder)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, -1, user.Id)
		assert.Equal(t, "FirstUser", user.DisplayName)
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)

		user, err = readUser(decoder)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.Id)
		assert.Equal(t, "SecondUser", user.DisplayName)
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)
	})
}

func Test_ReadUsers(t *testing.T) {
	t.Run("Can read all users from XML", func(t *testing.T) {
		reader := strings.NewReader(UsersXml)

		users := make(chan models.User)
		go readUsers(reader, users)

		user, _ := <- users
		assert.Equal(t, -1, user.Id)
		assert.Equal(t, "Community", user.DisplayName)
		expectedTime, _ := time.Parse(models.TimeFormat, "2016-08-02T00:14:10.580")
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)

		user, _ = <- users
		assert.Equal(t, 1, user.Id)
		assert.Equal(t, "Adam Lear", user.DisplayName)
		expectedTime, _ = time.Parse(models.TimeFormat, "2016-08-02T15:36:45.333")
		assert.Equal(t, models.Time(expectedTime), user.CreationDate)

		_, closed := <- users
		assert.Equal(t, true, closed)
	})
}
