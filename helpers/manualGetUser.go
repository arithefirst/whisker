package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// DiscordUser represents a Discord user from the Discord API
type DiscordUser struct {
	ID                   string                `json:"id"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	GlobalName           *string               `json:"global_name,omitempty"`
	Avatar               *string               `json:"avatar,omitempty"`
	Bot                  bool                  `json:"bot,omitempty"`
	System               bool                  `json:"system,omitempty"`
	MFAEnabled           bool                  `json:"mfa_enabled,omitempty"`
	Banner               *string               `json:"banner,omitempty"`
	AccentColor          *int                  `json:"accent_color,omitempty"`
	Locale               *string               `json:"locale,omitempty"`
	Verified             bool                  `json:"verified,omitempty"`
	Email                *string               `json:"email,omitempty"`
	Flags                *int                  `json:"flags,omitempty"`
	PremiumType          *int                  `json:"premium_type,omitempty"`
	PublicFlags          *int                  `json:"public_flags,omitempty"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	PrimaryGuild         *UserPrimaryGuild     `json:"primary_guild,omitempty"`
}

// AvatarDecorationData represents avatar decoration for a Discord user
type AvatarDecorationData struct {
	// Fields would be added here based on Discord's documentation
	// This is a placeholder as the specific structure isn't detailed in the table
	Asset *string `json:"asset,omitempty"`
	SkuID *string `json:"sku_id,omitempty"`
}

// UserPrimaryGuild represents a user's primary guild information
type UserPrimaryGuild struct {
	IdentityGuildID *string `json:"identity_guild_id,omitempty"`
	IdentityEnabled bool    `json:"identity_enabled,omitempty"`
	Tag             *string `json:"tag,omitempty"`
	Badge           *string `json:"badge,omitempty"`
}

func ManualGetUser(userId string, s *discordgo.Session) (*DiscordUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://discord.com/api/v10/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", s.Token)
	req.Header.Add("User-Agent", "Whisker Discord Bot (https://github.com/arithefirst/whisker, v1.0)")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Discord API returned status code %d", resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user DiscordUser
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error unmarshaling user data:", err)
		return nil, err
	}

	return &user, nil
}
