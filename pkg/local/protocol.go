package local

import (
	"github.com/pkg/errors"

	"github.com/havoc-io/mutagen/pkg/session"
	urlpkg "github.com/havoc-io/mutagen/pkg/url"
)

// localProtocolHandler is a protocol handler for connecting to local endpoints.
type localProtocolHandler struct{}

// Dial connects to a local endpoint.
func (h *localProtocolHandler) Dial(
	url *urlpkg.URL,
	session string,
	version session.Version,
	configuration *session.Configuration,
	alpha bool,
	prompter string,
) (session.Endpoint, error) {
	// Verify that the URL is of the correct protocol.
	if url.Protocol != urlpkg.Protocol_Local {
		panic("non-local URL dispatched to local URL handler")
	}

	// Create a local endpoint.
	endpoint, err := NewEndpoint(session, version, url.Path, configuration, alpha)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create local endpoint")
	}

	// Success.
	return endpoint, nil
}

func init() {
	// Register the local protocol handler with the session package.
	session.ProtocolHandlers[urlpkg.Protocol_Local] = &localProtocolHandler{}
}