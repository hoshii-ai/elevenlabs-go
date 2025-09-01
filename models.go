package elevenlabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Language struct {
	LanguageId string `json:"language_id"`
	Name       string `json:"name"`
}

type Model struct {
	CanBeFineTuned                     bool       `json:"can_be_finetuned"`
	CanDoTextToSpeech                  bool       `json:"can_do_text_to_speech"`
	CanDoVoiceConversion               bool       `json:"can_do_voice_conversion"`
	CanUseSpeakerBoost                 bool       `json:"can_use_speaker_boost"`
	CanUseStyle                        bool       `json:"can_use_style"`
	Description                        string     `json:"description"`
	Languages                          []Language `json:"languages"`
	MaxCharactersRequestFreeUser       int        `json:"max_characters_request_free_user"`
	MaxCharactersRequestSubscribedUser int        `json:"max_characters_request_subscribed_user"`
	ModelId                            string     `json:"model_id"`
	Name                               string     `json:"name"`
	RequiresAlphaAccess                bool       `json:"requires_alpha_access"`
	ServesProVoices                    bool       `json:"serves_pro_voices"`
	TokenCostFactor                    float32    `json:"token_cost_factor"`
}

type TextToSpeechRequest struct {
	Text          string         `json:"text"`
	ModelID       string         `json:"model_id,omitempty"`
	VoiceSettings *VoiceSettings `json:"voice_settings,omitempty"`
}

type GetVoicesResponse struct {
	Voices []Voice `json:"voices"`
}

type AddVoiceResponse struct {
	VoiceId string `json:"voice_id"`
}

type Voice struct {
	AvailableForTiers       []string          `json:"available_for_tiers"`
	Category                string            `json:"category"`
	Description             string            `json:"description"`
	FineTuning              FineTuning        `json:"fine_tuning"`
	HighQualityBaseModelIds []string          `json:"high_quality_base_model_ids"`
	Labels                  map[string]string `json:"labels"`
	Name                    string            `json:"name"`
	PreviewUrl              string            `json:"preview_url"`
	Samples                 []VoiceSample     `json:"samples"`
	Settings                VoiceSettings     `json:"settings,omitempty"`
	Sharing                 VoiceSharing      `json:"sharing"`
	VoiceId                 string            `json:"voice_id"`
}

type VoiceSettings struct {
	SimilarityBoost float32 `json:"similarity_boost"`
	Stability       float32 `json:"stability"`
	Style           float32 `json:"style,omitempty"`
	SpeakerBoost    bool    `json:"use_speaker_boost,omitempty"`
}

type VoiceSharing struct {
	ClonedByCount          int               `json:"cloned_by_count"`
	DateUnix               int               `json:"date_unix"`
	Description            string            `json:"description"`
	DisableAtUnix          bool              `json:"disable_at_unix"`
	EnabledInLibrary       bool              `json:"enabled_in_library"`
	FinancialRewardEnabled bool              `json:"financial_reward_enabled"`
	FreeUsersAllowed       bool              `json:"free_users_allowed"`
	HistoryItemSampleId    string            `json:"history_item_sample_id"`
	Labels                 map[string]string `json:"labels"`
	LikedByCount           int               `json:"liked_by_count"`
	LiveModerationEnabled  bool              `json:"live_moderation_enabled"`
	Name                   string            `json:"name"`
	NoticePeriod           int               `json:"notice_period"`
	OriginalVoiceId        string            `json:"original_voice_id"`
	PublicOwnerId          string            `json:"public_owner_id"`
	Rate                   float32           `json:"rate"`
	ReviewMessage          string            `json:"review_message"`
	ReviewStatus           string            `json:"review_status"`
	Status                 string            `json:"status"`
	VoiceMixingAllowed     bool              `json:"voice_mixing_allowed"`
	WhitelistedEmails      []string          `json:"whitelisted_emails"`
}

type VoiceSample struct {
	FileName  string `json:"file_name"`
	Hash      string `json:"hash"`
	MimeType  string `json:"mime_type"`
	SampleId  string `json:"sample_id"`
	SizeBytes int    `json:"size_bytes"`
}

type FineTuning struct {
	FineTuningRequested         bool                  `json:"fine_tuning_requested"`
	FineTuningState             string                `json:"finetuning_state"`
	IsAllowedToFineTune         bool                  `json:"is_allowed_to_fine_tune"`
	Language                    string                `json:"language"`
	ManualVerification          ManualVerification    `json:"manual_verification"`
	ManualVerificationRequested bool                  `json:"manual_verification_requested"`
	SliceIds                    []string              `json:"slice_ids"`
	VerificationAttempts        []VerificationAttempt `json:"verification_attempts"`
	VerificationAttemptsCount   int                   `json:"verification_attempts_count"`
	VerificationFailures        []string              `json:"verification_failures"`
}

type ManualVerification struct {
	ExtraText       string `json:"extra_text"`
	Files           []File `json:"files"`
	RequestTimeUnix int    `json:"request_time_unix"`
}

type File struct {
	FileId         string `json:"file_id"`
	FileName       string `json:"file_name"`
	MimeType       string `json:"mime_type"`
	SizeBytes      int    `json:"size_bytes"`
	UploadDateUnix int    `json:"upload_date_unix"`
}

type VerificationAttempt struct {
	Accepted            bool      `json:"accepted"`
	DateUnix            int       `json:"date_unix"`
	LevenshteinDistance float32   `json:"levenshtein_distance"`
	Recording           Recording `json:"recording"`
	Similarity          float32   `json:"similarity"`
	Text                string    `json:"text"`
}

type Recording struct {
	MimeType       string `json:"mime_type"`
	RecordingId    string `json:"recording_id"`
	SizeBytes      int    `json:"size_bytes"`
	Transcription  string `json:"transcription"`
	UploadDateUnix int    `json:"upload_date_unix"`
}

type DownloadHistoryRequest struct {
	HistoryItemIds []string `json:"history_item_ids"`
}

type GetHistoryResponse struct {
	History           []HistoryItem `json:"history"`
	LastHistoryItemId string        `json:"last_history_item_id"`
	HasMore           bool          `json:"has_more"`
}

type HistoryItem struct {
	CharacterCountChangeFrom int           `json:"character_count_change_from"`
	CharacterCountChangeTo   int           `json:"character_count_change_to"`
	ContentType              string        `json:"content_type"`
	DateUnix                 int           `json:"date_unix"`
	Feedback                 Feedback      `json:"feedback"`
	HistoryItemId            string        `json:"history_item_id"`
	ModelId                  string        `json:"model_id"`
	RequestId                string        `json:"request_id"`
	Settings                 VoiceSettings `json:"settings"`
	ShareLinkId              string        `json:"share_link_id"`
	State                    string        `json:"state"`
	Text                     string        `json:"text"`
	VoiceCategory            string        `json:"voice_category"`
	VoiceId                  string        `json:"voice_id"`
	VoiceName                string        `json:"voice_name"`
}

type Feedback struct {
	AudioQuality    bool    `json:"audio_quality"`
	Emotions        bool    `json:"emotions"`
	Feedback        string  `json:"feedback"`
	Glitches        bool    `json:"glitches"`
	InaccurateClone bool    `json:"inaccurate_clone"`
	Other           bool    `json:"other"`
	ReviewStatus    *string `json:"review_status,omitempty"`
	ThumbsUp        bool    `json:"thumbs_up"`
}

type Subscription struct {
	AllowedToExtendCharacterLimit  bool    `json:"allowed_to_extend_character_limit"`
	CanExtendCharacterLimit        bool    `json:"can_extend_character_limit"`
	CanExtendVoiceLimit            bool    `json:"can_extend_voice_limit"`
	CanUseInstantVoiceCloning      bool    `json:"can_use_instant_voice_cloning"`
	CanUseProfessionalVoiceCloning bool    `json:"can_use_professional_voice_cloning"`
	CharacterCount                 int     `json:"character_count"`
	CharacterLimit                 int     `json:"character_limit"`
	Currency                       string  `json:"currency"`
	NextCharacterCountResetUnix    int     `json:"next_character_count_reset_unix"`
	VoiceLimit                     int     `json:"voice_limit"`
	ProfessionalVoiceLimit         int     `json:"professional_voice_limit"`
	Status                         string  `json:"status"`
	Tier                           string  `json:"tier"`
	MaxVoiceAddEdits               int     `json:"max_voice_add_edits"`
	VoiceAddEditCounter            int     `json:"voice_add_edit_counter"`
	HasOpenInvoices                bool    `json:"has_open_invoices"`
	NextInvoice                    Invoice `json:"next_invoice"`
	withInvoicingDetails           bool
}

type Invoice struct {
	AmountDueCents         int `json:"amount_due_cents"`
	NextPaymentAttemptUnix int `json:"next_payment_attempt_unix"`
}

type User struct {
	Subscription                Subscription `json:"subscription"`
	FirstName                   string       `json:"first_name,omitempty"`
	IsNewUser                   bool         `json:"is_new_user"`
	IsOnboardingComplete        bool         `json:"is_onboarding_complete"`
	XiApiKey                    string       `json:"xi_api_key"`
	CanUseDelayedPaymentMethods bool         `json:"can_use_delayed_payment_methods"`
}

type AddEditVoiceRequest struct {
	Name        string
	FilePaths   []string
	Description string
	Labels      map[string]string
}

func (r *AddEditVoiceRequest) buildRequestBody() (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	buildFailed := func(err error) (*bytes.Buffer, string, error) {
		return nil, "", fmt.Errorf("failed to build request body: %w", err)
	}

	if err := w.WriteField("name", r.Name); err != nil {
		return buildFailed(err)
	}
	if r.Description != "" {
		if err := w.WriteField("description", r.Description); err != nil {
			return buildFailed(err)
		}
	}
	if len(r.Labels) > 0 {
		labelsJson, err := json.Marshal(r.Labels)
		if err != nil {
			return buildFailed(err)
		}
		if err := w.WriteField("labels", string(labelsJson)); err != nil {
			return buildFailed(err)
		}
	}

	for _, file := range r.FilePaths {
		f, err := os.Open(file)
		if err != nil {
			return buildFailed(err)
		}
		defer f.Close()

		fw, err := w.CreateFormFile("files", filepath.Base(file))
		if err != nil {
			return buildFailed(err)
		}
		if _, err = io.Copy(fw, f); err != nil {
			return buildFailed(err)
		}
	}

	err := w.Close()
	if err != nil {
		return buildFailed(err)
	}

	return &b, w.FormDataContentType(), nil
}

// SpeechToTextRequest represents the request parameters for speech-to-text conversion
type SpeechToTextRequest struct {
	ModelID               string               `json:"model_id"`
	File                  io.Reader            `json:"-"` // File content, handled separately in multipart
	FileName              string               `json:"-"` // Original filename for multipart
	LanguageCode          *string              `json:"language_code,omitempty"`
	TagAudioEvents        *bool                `json:"tag_audio_events,omitempty"`
	NumSpeakers           *int                 `json:"num_speakers,omitempty"`
	TimestampsGranularity *string              `json:"timestamps_granularity,omitempty"`
	Diarize               *bool                `json:"diarize,omitempty"`
	DiarizationThreshold  *float64             `json:"diarization_threshold,omitempty"`
	AdditionalFormats     []SpeechToTextFormat `json:"additional_formats,omitempty"`
	FileFormat            *string              `json:"file_format,omitempty"`
	CloudStorageURL       *string              `json:"cloud_storage_url,omitempty"`
	Webhook               *bool                `json:"webhook,omitempty"`
	WebhookID             *string              `json:"webhook_id,omitempty"`
	Temperature           *float64             `json:"temperature,omitempty"`
	Seed                  *int                 `json:"seed,omitempty"`
	UseMultiChannel       *bool                `json:"use_multi_channel,omitempty"`
	WebhookMetadata       interface{}          `json:"webhook_metadata,omitempty"`
}

// SpeechToTextFormat represents additional export formats for speech-to-text
type SpeechToTextFormat struct {
	Type                  string  `json:"type"`
	TimestampsGranularity *string `json:"timestamps_granularity,omitempty"`
	Diarize               *bool   `json:"diarize,omitempty"`
	MaxSpeakers           *int    `json:"max_speakers,omitempty"`
	MaxWords              *int    `json:"max_words,omitempty"`
	PunctuateOnly         *bool   `json:"punctuate_only,omitempty"`
}

// SpeechToTextWord represents a word in the transcription with timing and speaker information
type SpeechToTextWord struct {
	Text         string   `json:"text"`
	Start        float64  `json:"start"`
	End          float64  `json:"end"`
	Type         string   `json:"type"`
	SpeakerID    *string  `json:"speaker_id,omitempty"`
	ChannelIndex *int     `json:"channel_index,omitempty"`
	LogProb      *float64 `json:"logprob,omitempty"`
}

// SpeechToTextResponse represents the response from speech-to-text conversion
type SpeechToTextResponse struct {
	LanguageCode        string             `json:"language_code"`
	LanguageProbability float64            `json:"language_probability"`
	Text                string             `json:"text"`
	Words               []SpeechToTextWord `json:"words"`
}

// MultichannelSpeechToTextResponse represents the response for multi-channel audio
type MultichannelSpeechToTextResponse struct {
	Transcripts []SpeechToTextResponse `json:"transcripts"`
}

// SpeechToTextWebhookResponse represents the response when webhook is enabled
type SpeechToTextWebhookResponse struct {
	RequestID string `json:"request_id"`
	Message   string `json:"message"`
}

// buildRequestBody creates the multipart form request body for speech-to-text
func (r *SpeechToTextRequest) buildRequestBody() (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	buildFailed := func(err error) (*bytes.Buffer, string, error) {
		return nil, "", fmt.Errorf("failed to build speech-to-text request body: %w", err)
	}

	// Add model_id
	if err := w.WriteField("model_id", r.ModelID); err != nil {
		return buildFailed(err)
	}

	// Add file if provided
	if r.File != nil && r.FileName != "" {
		fw, err := w.CreateFormFile("file", r.FileName)
		if err != nil {
			return buildFailed(err)
		}
		if _, err = io.Copy(fw, r.File); err != nil {
			return buildFailed(err)
		}
	}

	// Add optional fields
	if r.LanguageCode != nil {
		if err := w.WriteField("language_code", *r.LanguageCode); err != nil {
			return buildFailed(err)
		}
	}

	if r.TagAudioEvents != nil {
		if err := w.WriteField("tag_audio_events", fmt.Sprintf("%t", *r.TagAudioEvents)); err != nil {
			return buildFailed(err)
		}
	}

	if r.NumSpeakers != nil {
		if err := w.WriteField("num_speakers", fmt.Sprintf("%d", *r.NumSpeakers)); err != nil {
			return buildFailed(err)
		}
	}

	if r.TimestampsGranularity != nil {
		if err := w.WriteField("timestamps_granularity", *r.TimestampsGranularity); err != nil {
			return buildFailed(err)
		}
	}

	if r.Diarize != nil {
		if err := w.WriteField("diarize", fmt.Sprintf("%t", *r.Diarize)); err != nil {
			return buildFailed(err)
		}
	}

	if r.DiarizationThreshold != nil {
		if err := w.WriteField("diarization_threshold", fmt.Sprintf("%f", *r.DiarizationThreshold)); err != nil {
			return buildFailed(err)
		}
	}

	if len(r.AdditionalFormats) > 0 {
		formatsJSON, err := json.Marshal(r.AdditionalFormats)
		if err != nil {
			return buildFailed(err)
		}
		if err := w.WriteField("additional_formats", string(formatsJSON)); err != nil {
			return buildFailed(err)
		}
	}

	if r.FileFormat != nil {
		if err := w.WriteField("file_format", *r.FileFormat); err != nil {
			return buildFailed(err)
		}
	}

	if r.CloudStorageURL != nil {
		if err := w.WriteField("cloud_storage_url", *r.CloudStorageURL); err != nil {
			return buildFailed(err)
		}
	}

	if r.Webhook != nil {
		if err := w.WriteField("webhook", fmt.Sprintf("%t", *r.Webhook)); err != nil {
			return buildFailed(err)
		}
	}

	if r.WebhookID != nil {
		if err := w.WriteField("webhook_id", *r.WebhookID); err != nil {
			return buildFailed(err)
		}
	}

	if r.Temperature != nil {
		if err := w.WriteField("temperature", fmt.Sprintf("%f", *r.Temperature)); err != nil {
			return buildFailed(err)
		}
	}

	if r.Seed != nil {
		if err := w.WriteField("seed", fmt.Sprintf("%d", *r.Seed)); err != nil {
			return buildFailed(err)
		}
	}

	if r.UseMultiChannel != nil {
		if err := w.WriteField("use_multi_channel", fmt.Sprintf("%t", *r.UseMultiChannel)); err != nil {
			return buildFailed(err)
		}
	}

	if r.WebhookMetadata != nil {
		metadataJSON, err := json.Marshal(r.WebhookMetadata)
		if err != nil {
			return buildFailed(err)
		}
		if err := w.WriteField("webhook_metadata", string(metadataJSON)); err != nil {
			return buildFailed(err)
		}
	}

	err := w.Close()
	if err != nil {
		return buildFailed(err)
	}

	return &b, w.FormDataContentType(), nil
}
