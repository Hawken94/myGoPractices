package zerolog

import (
	"testing"

	"myGoPractices/thirdPackage/zhuozhuo"

	"github.com/rs/zerolog/log"
)

func TestZerolog(t *testing.T) {
	log.Logger = log.With().Caller().Logger()

	log.Info().Msgf("hello,%v", "world")
	str, i := zhuozhuo.GetLongestSubstrings("aabcceddabc")
	log.Error().Msgf("str:%s i:%d", str, i)
	t.Error()
}
