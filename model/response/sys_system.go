package response

import "innovation/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
