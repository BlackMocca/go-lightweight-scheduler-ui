package constants

import (
	"strings"

	"github.com/lnquy/cron"
)

type LocalStorageKey string

const (
	STORAGE_CONNECTION_LIST LocalStorageKey = "connection-list"
)

type AssetPath string

const (
	LOGO_NO_BACKGROUND          AssetPath = "/web/resources/assets/logo/logo-no-background.svg"
	ICON_FAVOURITE              AssetPath = "/web/resources/assets/icon/favourite.png"
	ICON_ADD_PRIMARY            AssetPath = "/web/resources/assets/icon/add_primary_color.png"
	ICON_ADD_SECONDARY          AssetPath = "/web/resources/assets/icon/add_secondary_color.svg"
	ICON_DELETE_PRIMARY         AssetPath = "/web/resources/assets/icon/bin_primary_color.svg"
	ICON_DELETE_SECONDARY       AssetPath = "/web/resources/assets/icon/bin_secondary_color.svg"
	ICON_SETTING                AssetPath = "/web/resources/assets/icon/setting.svg"
	ICON_SIGN_OUT               AssetPath = "/web/resources/assets/icon/signout.svg"
	ICON_PLAY                   AssetPath = "/web/resources/assets/icon/play.svg"
	ICON_PAGINATION_LEFT_ARROW  AssetPath = "/web/resources/assets/icon/pagination_left_arrow.svg"
	ICON_PAGINATION_RIGHT_ARROW AssetPath = "/web/resources/assets/icon/pagination_right_arrow.svg"
)

const (
	DATE_LAYOUT      = "2006-01-02"
	TIMESTAMP_LAYOUT = "2006-01-02 15:04:05"
)

var (
	CRONJOB_READABLE = func(cronExpression string) (string, error) {
		if cronExpression == "" {
			return "-", nil
		}
		exprDesc, _ := cron.NewDescriptor(
			cron.Use24HourTimeFormat(true),
			cron.DayOfWeekStartsAtOne(false),
		)

		desc, err := exprDesc.ToDescription(cronExpression, cron.Locale_en)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(desc), nil
	}
)
