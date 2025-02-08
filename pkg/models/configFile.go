package models

type AwsProfile struct {
	Profiles map[string]ProfileDetails `yaml:"profiles"`
}

type ProfileDetails struct {
	SSOEnabled  bool   `yaml:"ssoEnabled"`
	Config      string `yaml:"config"`
	Region      string `yaml:"region"`
	Credentials string `yaml:"credentials"`
}
