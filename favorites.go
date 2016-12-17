package main

func getUserFavorites(username string) []string {
	cached := serverCache.GetCache("favorites", username)
	if cached == nil {
		initUserFavorites(username)
	}
	favMap := serverCache.GetCache("favorites", username).(map[string]bool)
	favIDs := []string{}
	for k, v := range favMap {
		if v {
			favIDs = append(favIDs, k)
		}
	}
	// Return the IDs of favorite recipes
	return favIDs
}

func addToFavorites(username string, recipeID string) {
	cached := serverCache.GetCache("favorites", username)
	if cached == nil {
		initUserFavorites(username)
	}
	favMap := serverCache.GetCache("favorites", username).(map[string]bool)
	favMap[recipeID] = true
}

func removeFromFavorites(username string, recipeID string) {
	cached := serverCache.GetCache("favorites", username)
	if cached == nil {
		initUserFavorites(username)
	}
	favMap := serverCache.GetCache("favorites", username).(map[string]bool)
	favMap[recipeID] = false
}

func favoriteExists(username string, recipeID string) bool {
	cached := serverCache.GetCache("favorites", username)
	if cached == nil {
		initUserFavorites(username)
	}
	favMap := serverCache.GetCache("favorites", username).(map[string]bool)
	return favMap[recipeID]
}

func initUserFavorites(username string) {
	serverCache.SetCache("favorites", username, map[string]bool{})
}
