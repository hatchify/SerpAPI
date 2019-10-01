package serpapi

// LocalResult represents a local result entry
type LocalResult struct {
	Position    int    `json:"position"`
	Title       string `json:"title"`
	Reviews     int    `json:"reviews"`
	Price       string `json:"price"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Hours       string `json:"hours"`
	Thumbnail   string `json:"thumbnail"`
}

// LocalMap represents the local map for the current search
type LocalMap struct {
	Link           string         `json:"link"`
	GPSCoordinates GPSCoordinates `json:"gps_coordinates"`
	Image          string         `json:"image"`
}

// GPSCoordinates represents GPS coordinates for a local map
type GPSCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  int     `json:"altitude"`
}
