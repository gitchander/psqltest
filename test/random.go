package main

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/gitchander/psqltest/random"
)

var (
	timestampMin = dateToTime(2000, time.January, 1).UnixNano()
	timestampMax = dateToTime(2100, time.January, 1).UnixNano()
)

func dateToTime(year int, month time.Month, day int) time.Time {
	// time
	const (
		hour = 0
		min  = 0
		sec  = 0

		nsec = 0 // nanoseconds
	)
	return time.Date(
		year, month, day,
		hour, min, sec, nsec,
		time.Local)
}

func randInt64(r *rand.Rand, min, max int64) int64 {
	return min + r.Int63n(max-min)
}

func randTimestamp(r *rand.Rand) int64 {
	return randInt64(r, timestampMin, timestampMax)
}

func randTask(r *rand.Rand) *Task {
	t := &Task{
		// ID
		Alias:     randAlias(r),
		Timestamp: time.Unix(0, randTimestamp(r)),
		Groups:    randGroups(r),
		Number:    randNumber(r),
		FieldU64:  randFieldU64(r),
	}
	return t
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")

func randAlias(r *rand.Rand) string {
	return randWord(r)
}

func randGroup(r *rand.Rand) string {
	return randWord(r)
}

func randGroups(r *rand.Rand) string {
	groups := make([]string, random.IntMinMax(r, 2, 8))
	for i := range groups {
		groups[i] = randGroup(r)
	}
	return "/" + strings.Join(groups, "/")
}

func randWord(r *rand.Rand) string {
	n := random.IntMinMax(r, 3, 10)
	rs := randRunesByCorpus(r, n, alphabet)
	return string(rs)
}

func randRunesByCorpus(r *rand.Rand, n int, corpus []rune) []rune {
	rs := make([]rune, n)
	for i := range rs {
		rs[i] = randRuneByCorpus(r, corpus)
	}
	return rs
}

func randRuneByCorpus(r *rand.Rand, corpus []rune) rune {
	return corpus[r.Intn(len(corpus))]
}

func randByCorpus[T any](r *rand.Rand, corpus []T) T {
	return corpus[r.Intn(len(corpus))]
}

var digits = []rune("0123456789")

func randNumber(r *rand.Rand) string {
	n := random.IntMinMax(r, 20, 30)
	rs := randRunesByCorpus(r, n, digits)
	return string(rs)
}

func randFieldU64(r *rand.Rand) uint64 {
	switch n := r.Intn(4); n {
	case 0:
		return 0
	case 1:
		return math.MaxInt32
	case 2:
		return math.MaxInt64
	default:
		return uint64(r.Int63()) >> (r.Intn(64))
	}
}
