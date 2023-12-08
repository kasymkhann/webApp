package repository

func GetComments() ([]string, error) {
	return Client.LRange("comment", 0, 10).Result()
}

func PostComments(comment string) error {
	return Client.LPush("comment", comment).Err()
}
