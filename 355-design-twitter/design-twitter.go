type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h TweetHeap) Less(i, j int) bool { return h[i].timestamp > h[j].timestamp }

func (h *TweetHeap) Push(x any) {
	if v, ok := x.(Tweet); ok {
		*h = append(*h, v)
	}
}

func (h *TweetHeap) Pop() any {
	tempHeap := *h
	size := len(tempHeap)
	last := tempHeap[size-1]
	*h = tempHeap[:size-1]
	return last
}

type UserId = int

type Tweet struct {
	id        int
	timestamp int
}

type Twitter struct {
	timestamp int
	tweets    map[UserId][]Tweet
	followers map[UserId]map[UserId]struct{}
}

func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int][]Tweet),
		followers: make(map[int]map[int]struct{}),
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.timestamp++

	t.tweets[userId] = append(t.tweets[userId], Tweet{
		id:        tweetId,
		timestamp: t.timestamp,
	})
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	tweetsHeap := new(TweetHeap)
	heap.Init(tweetsHeap)

	if userTweets, found := t.tweets[userId]; found {
		for _, tweet := range userTweets {
			heap.Push(tweetsHeap, tweet)
		}
	}

	if followees, found := t.followers[userId]; found {
		for followeeId := range followees {
			if userTweets, ok := t.tweets[followeeId]; ok {
				for _, tweet := range userTweets {
					heap.Push(tweetsHeap, tweet)
				}
			}
		}
	}

	result := make([]int, 0, 10)

	i := 0

	for i < 10 && tweetsHeap.Len() > 0 {
		tweet := heap.Pop(tweetsHeap).(Tweet)
		result = append(result, tweet.id)
		i++
	}

	return result
}

func (t *Twitter) Follow(followerId, followeeId int) {
	if _, found := t.followers[followerId]; !found {
		t.followers[followerId] = make(map[int]struct{})
	}

	t.followers[followerId][followeeId] = struct{}{}
}

func (t *Twitter) Unfollow(followerId, followeeId int) {
	if _, found := t.followers[followerId]; found {
		delete(t.followers[followerId], followeeId)
	}
}