package enum

const (
	ConfigFile = "./config/config.yaml"

	JwtTokenName = "Token"
	JwtIssuer    = "jwt-issuer"
	JwtTtl       = 24 * 7 // hour
	JwtSecret    = "tiktok"

	ContextUserId = "user"
)
