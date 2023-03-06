package ezhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Header(key, val string) ConfigPiece {
	return After(func(conf *Config) {
		conf.Request.Header.Set(key, val)
	})
}

func Cookie(cookie http.Cookie) ConfigPiece {
	return After(func(conf *Config) {
		conf.Request.AddCookie(&cookie)
	})
}

func AuthBasic(username, password string) ConfigPiece {
	return After(func(conf *Config) {
		conf.Request.SetBasicAuth(username, password)
	})
}

func AuthBearer(token string) ConfigPiece { return Header("Authorization", "Bearer "+token) }

func SendJSON(ref interface{}) ConfigPiece {
	return ConfigPiece{
		BeforeInit: func(conf *Config) {
			jsonBytes, err := json.Marshal(ref)
			if err != nil {
				conf.Abort = err
				return
			}
			conf.RequestBody = bytes.NewBuffer(jsonBytes)
		},
		AfterInit: func(conf *Config) {
			conf.Request.Header.Set("Content-Type", jsonContentType)
		},
	}
}

func SendBody(body io.Reader, contentType string) ConfigPiece {
	return ConfigPiece{
		BeforeInit: func(conf *Config) {
			conf.RequestBody = body
		},
		AfterInit: func(conf *Config) {
			conf.Request.Header.Set("Content-Type", contentType)
		},
	}
}

// use this when you want to be forward compatible, i.e. server is allowed to add new fields to
// JSON structure.
func RespondsJSONAllowUnknownFields(obj interface{}) ConfigPiece {
	return respondsJSON(obj, true)
}

// use this when you DON'T want to be forward compatible, i.e. server is not allowed to add new
// fields to JSON structure without yielding an error.
func RespondsJSONDisallowUnknownFields(obj interface{}) ConfigPiece {
	return respondsJSON(obj, false)
}

func respondsJSON(ref interface{}, allowUnknownFields bool) ConfigPiece {
	return After(func(conf *Config) {
		if conf.Request.Header.Get("Accept") == "" { // don't override if we have explicit `Accept` header
			conf.Request.Header.Set("Accept", jsonContentType)
		}

		conf.OutputsJson = true
		conf.OutputsJsonRef = ref
		conf.OutputsJsonAllowUnknownFields = allowUnknownFields
	})
}

func Client(client *http.Client) ConfigPiece {
	return After(func(conf *Config) {
		conf.Client = client
	})
}

var TolerateNon2xxResponse = After(func(conf *Config) {
	conf.TolerateNon2xxResponse = true
})

// backwards compat

// Deprecated: use explicit allow/disallow instead
func RespondsJson(ref interface{}, allowUnknownFields bool) ConfigPiece {
	return respondsJSON(ref, allowUnknownFields)
}

// Deprecated: use fn with JSON uppercased
func SendJson(ref interface{}) ConfigPiece {
	return SendJSON(ref)
}

// Deprecated: use fn with JSON uppercased
func RespondsJsonAllowUnknownFields(obj interface{}) ConfigPiece {
	return RespondsJSONAllowUnknownFields(obj)
}

// Deprecated: use fn with JSON uppercased
func RespondsJsonDisallowUnknownFields(obj interface{}) ConfigPiece {
	return RespondsJSONDisallowUnknownFields(obj)
}
