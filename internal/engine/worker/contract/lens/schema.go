package lens

type PublicationV1 struct {
	Version          string      `json:"version"`
	MetadataID       string      `json:"metadata_id"`
	Content          string      `json:"content"`
	ExternalURL      string      `json:"external_url"`
	Image            interface{} `json:"image"`
	ImageMimeType    interface{} `json:"imageMimeType"`
	Name             string      `json:"name"`
	Tags             []string    `json:"tags"`
	AnimationURL     interface{} `json:"animation_url"`
	MainContentFocus string      `json:"mainContentFocus"`
	ContentWarning   interface{} `json:"contentWarning"`
	Attributes       []struct {
		TraitType   string `json:"traitType"`
		DisplayType string `json:"displayType"`
		Value       any    `json:"value"`
	} `json:"attributes"`
	Media  []PublicationMedia `json:"media"`
	Locale string             `json:"locale"`
	AppID  string             `json:"appId"`
}

type PublicationMedia struct {
	Item string `json:"item"`
	Type string `json:"type"`
}

type PublicationV2 struct {
	Schema string `json:"$schema"`
	Lens   struct {
		ID               string `json:"id"`
		AppID            string `json:"appId"`
		Locale           string `json:"locale"`
		MainContentFocus string `json:"mainContentFocus"`
		Content          string `json:"content"`
		Image            struct {
			Item string `json:"item"`
			Type string `json:"type"`
		} `json:"image"`
		Tags  []string `json:"tags"`
		Video struct {
			Item       string `json:"item"`
			Attributes []struct {
				Type  string `json:"type"`
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"attributes"`
			Type     string `json:"type"`
			AltTag   string `json:"altTag"`
			Cover    string `json:"cover"`
			Duration int    `json:"duration"`
			License  string `json:"license"`
		} `json:"video"`
	} `json:"lens"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	AnimationURL string `json:"animation_url"`
	ExternalURL  string `json:"external_url"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}
