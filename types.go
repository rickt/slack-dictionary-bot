package slackdictionarybot

// struct for basic JSON request to be sent to Oxford Dictionary API
type OxfordRequest struct {
	Type    string `json:"Accept"`
	App_ID  string `json:"app_id"`
	App_Key string `json:"app_key"`
}

// struct for Oxford Dictionary API reply
type OxfordReply struct {
	Metadata struct {
		Provider string `json:"provider"`
	} `json:"metadata"`
	Results []struct {
		ID             string `json:"id"`
		Language       string `json:"language"`
		LexicalEntries []struct {
			Entries []struct {
				Etymologies         []string `json:"etymologies"`
				GrammaticalFeatures []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"grammaticalFeatures"`
				HomographNumber string `json:"homographNumber"`
				Senses          []struct {
					Definitions []string `json:"definitions"`
					Domains     []string `json:"domains"`
					Examples    []struct {
						Text string `json:"text"`
					} `json:"examples"`
					ID               string   `json:"id"`
					ShortDefinitions []string `json:"short_definitions"`
					Subsenses        []struct {
						Definitions []string `json:"definitions"`
						Examples    []struct {
							Text string `json:"text"`
						} `json:"examples"`
						ID               string   `json:"id"`
						ShortDefinitions []string `json:"short_definitions"`
						ThesaurusLinks   []struct {
							EntryID string `json:"entry_id"`
							SenseID string `json:"sense_id"`
						} `json:"thesaurusLinks"`
						Registers []string `json:"registers,omitempty"`
					} `json:"subsenses"`
					ThesaurusLinks []struct {
						EntryID string `json:"entry_id"`
						SenseID string `json:"sense_id"`
					} `json:"thesaurusLinks"`
				} `json:"senses"`
			} `json:"entries"`
			Language        string `json:"language"`
			LexicalCategory string `json:"lexicalCategory"`
			Pronunciations  []struct {
				AudioFile        string   `json:"audioFile"`
				Dialects         []string `json:"dialects"`
				PhoneticNotation string   `json:"phoneticNotation"`
				PhoneticSpelling string   `json:"phoneticSpelling"`
			} `json:"pronunciations"`
			Text string `json:"text"`
		} `json:"lexicalEntries"`
		Type string `json:"type"`
		Word string `json:"word"`
	} `json:"results"`
}

// struct for runtime environment variables
type envVars struct {
	APIBaseURL  string
	APPID       string
	APPKey      string
	APIPath     string
	VerifyToken string
}

// EOF
