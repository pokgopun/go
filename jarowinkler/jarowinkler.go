/* https://theweeklychallenge.org/blog/perl-weekly-challenge-010/
Challenge #2
Write a script to find Jaro-Winkler distance between two strings. For more information check wikipedia page.
*/
package jarowinkler

import (
	"math"
	"regexp"
)

const scalingFactor = 0.1

type jarowinkler struct {
	str1, str2                                           string
	l1, l2                                               int
	m                                                    map[byte][]int
	maxDistance, matchCount, transposeCount, prefixCount float64
	similarity0, similarity, distance0, distance         float64
}

func New(str1, str2 string) (jw jarowinkler) {
	jw = jarowinkler{
		str1: str1,
		str2: str2,
		l1:   len(str1),
		l2:   len(str2),
	}
	jw.maxDistance = math.Max(float64(jw.l1), float64(jw.l2))/2 - 1
	jw.m = make(map[byte][]int)
	var (
		chrDistance    float64
		prevMatchIndex int
	)
	matchAndCount := func(i int, b byte) {
		for j, v := range jw.m[b] {
			chrDistance = math.Abs(float64(v - i))
			if chrDistance <= jw.maxDistance {
				jw.matchCount++
				if v < prevMatchIndex {
					jw.transposeCount++
				}
				prevMatchIndex = v
				jw.m[b] = append(jw.m[b][:j], jw.m[b][j+1:]...)
				break
			}
		}
	}
	for i, b := range []byte(jw.str1) {
		if jw.str1[:(i+1)%(jw.l1+1)] == jw.str2[:(i+1)%(jw.l2+1)] {
			jw.prefixCount++
		}
		_, ok := jw.m[b]
		if ok {
			if len(jw.m[b]) > 0 {
				matchAndCount(i, b)
			}
			continue
		}
		s := regexp.MustCompile(string(b)).FindAllIndex([]byte(jw.str2), -1)
		if len(s) > 0 {
			jw.m[b] = []int{}
			for _, v := range s {
				jw.m[b] = append(jw.m[b], v[0])
			}
			matchAndCount(i, b)
		}
	}
	switch {
	case jw.str1 == jw.str2:
		jw.similarity0 = 1
		jw.similarity = 1
	case jw.str1 == "" || jw.str2 == "" || jw.matchCount == 0:
		jw.similarity0 = 0
		jw.similarity = 0
	default:
		jw.similarity0 = (jw.matchCount/float64(jw.l1) + jw.matchCount/float64(jw.l2) + (jw.matchCount-jw.transposeCount)/jw.matchCount) / 3
	}
	jw.similarity = jw.similarity0 + scalingFactor*jw.prefixCount*(1-jw.similarity0)
	jw.distance0 = 1 - jw.similarity0
	jw.distance = 1 - jw.similarity
	return jw
}
func (jw *jarowinkler) Distance(wink bool) (r float64) {
	if wink == true {
		r = jw.distance
	} else {
		r = jw.distance0
	}
	return math.Round(r*1000) / 1000
}
func (jw *jarowinkler) Similarity(wink bool) (r float64) {
	if wink == true {
		r = jw.similarity
	} else {
		r = jw.similarity0
	}
	return math.Round(r*1000) / 1000
}
