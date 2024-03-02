package constants

type LocalStorageKey string

const (
	STORAGE_CONNECTION_LIST LocalStorageKey = "connection-list"
)

type AssetPath string

const (
	LOGO_NO_BACKGROUND    AssetPath = "/web/resources/assets/logo/logo-no-background.svg"
	ICON_FAVOURITE        AssetPath = "/web/resources/assets/icon/favourite.png"
	ICON_ADD_PRIMARY      AssetPath = "/web/resources/assets/icon/add_primary_color.png"
	ICON_ADD_SECONDARY    AssetPath = "/web/resources/assets/icon/add_secondary_color.svg"
	ICON_DELETE_PRIMARY   AssetPath = "/web/resources/assets/icon/bin_primary_color.svg"
	ICON_DELETE_SECONDARY AssetPath = "/web/resources/assets/icon/bin_secondary_color.svg"
	ICON_SIGN_OUT         AssetPath = "/web/resources/assets/icon/signout.svg"
)

var (
// SVG_RING_WEDDING_STRING   = GetSVGString("assets/icon/rings-wedding.svg")
)
