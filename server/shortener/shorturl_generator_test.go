package shortener

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const userId = "e0dba740-fc4b-4977-872c-d360239e6b1a"
//tests from https://github.com/eddywm/go-shortener-wm/blob/master/shortener/shorturl_generator_test.go
func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1 := GenerateShortLink(initialLink_1, userId)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := GenerateShortLink(initialLink_2, userId)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(initialLink_3, userId)


	assert.Equal(t, shortLink_1, "jTa4L57P")
	assert.Equal(t, shortLink_2, "d66yfx7N")
	assert.Equal(t, shortLink_3, "dhZTayYQ")
}