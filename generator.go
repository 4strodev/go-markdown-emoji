package emoji

func getEmoji(name string) string {
	emoji, exists := emojiMap[name]
	if !exists {
		return ""
	}
        return emoji
}
