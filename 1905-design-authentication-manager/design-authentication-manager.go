type AuthenticationManager struct {
	timeToLive time.Duration
	tokens     map[string]time.Time
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: time.Duration(timeToLive) * time.Second,
		tokens:     make(map[string]time.Time),
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	expirationDate := time.Unix(int64(currentTime), 0).Add(am.timeToLive)
	am.tokens[tokenId] = expirationDate
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	current := time.Unix(int64(currentTime), 0)
	if expirationDate, found := am.tokens[tokenId]; found && expirationDate.After(current) {
		am.tokens[tokenId] = current.Add(am.timeToLive)
	}
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	current := time.Unix(int64(currentTime), 0)
	count := 0

	for _, expirationDate := range am.tokens {
		if expirationDate.After(current) {
			count++
		}
	}

	return count
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */