package client

import (
	"encoding/xml"
	"github.com/dathanb/migrations/fakestack/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

var UserXml = "<row Id=\"-1\" Reputation=\"1\" CreationDate=\"2016-08-02T00:14:10.580\" DisplayName=\"Community\" LastAccessDate=\"2016-08-02T00:14:10.580\" Location=\"on the server farm\" AboutMe=\"&lt;p&gt;Hi, I'm not really a person.&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I'm a background process that helps keep this site clean!&lt;/p&gt;&#xD;&#xA;&lt;p&gt;I do things like&lt;/p&gt;&#xD;&#xA;&lt;ul&gt;&#xD;&#xA;&lt;li&gt;Randomly poke old unanswered questions every hour so they get some attention&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own community questions and answers so nobody gets unnecessary reputation from them&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own downvotes on spam/evil posts that get permanently deleted&lt;/li&gt;&#xD;&#xA;&lt;li&gt;Own suggested edits from anonymous users&lt;/li&gt;&#xD;&#xA;&lt;li&gt;&lt;a href=&quot;http://meta.stackoverflow.com/a/92006&quot;&gt;Remove abandoned questions&lt;/a&gt;&lt;/li&gt;&#xD;&#xA;&lt;/ul&gt;\" Views=\"0\" UpVotes=\"3\" DownVotes=\"661\" AccountId=\"-1\" />"

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
}
