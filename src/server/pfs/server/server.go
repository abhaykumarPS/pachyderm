package server

import (
	"fmt"

	pfsclient "github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/server/pkg/obj"
)

// Valid object storage backends
const (
	MinioBackendEnvVar     = "MINIO"
	AmazonBackendEnvVar    = "AMAZON"
	GoogleBackendEnvVar    = "GOOGLE"
	MicrosoftBackendEnvVar = "MICROSOFT"
	LocalBackendEnvVar     = "LOCAL"
)

var (
	// The size of file chunks when resumable upload is used
	ChunkSize = int64(16 * 1024 * 1024) // 16 MB
)

// APIServer represents and api server.
type APIServer interface {
	pfsclient.APIServer
}

// BlockAPIServer combines BlockAPIServer and ObjectAPIServer.
type BlockAPIServer interface {
	pfsclient.ObjectAPIServer
}

// NewAPIServer creates an APIServer.
// cacheSize is the number of commit trees which will be cached in the server.
func NewAPIServer(address string, etcdAddresses []string, etcdPrefix string, cacheSize int64) (APIServer, error) {
	return newAPIServer(address, etcdAddresses, etcdPrefix, cacheSize)
}

// NewHTTPServer creates an APIServer.
// cacheSize is the number of commit trees which will be cached in the server.
func NewHTTPServer(address string, etcdAddresses []string, etcdPrefix string, cacheSize int64) (*HTTPServer, error) {
	return newHTTPServer(address, etcdAddresses, etcdPrefix, cacheSize)
}

// NewObjBlockAPIServer create a BlockAPIServer from an obj.Client.
func NewObjBlockAPIServer(dir string, cacheBytes int64, etcdAddress string, objClient obj.Client) (BlockAPIServer, error) {
	return newObjBlockAPIServer(dir, cacheBytes, etcdAddress, objClient)
}

// NewBlockAPIServer creates a BlockAPIServer using the credentials it finds in
// the environment
func NewBlockAPIServer(dir string, cacheBytes int64, backend string, etcdAddress string) (BlockAPIServer, error) {
	switch backend {
	case MinioBackendEnvVar:
		// S3 compatible doesn't like leading slashes
		if len(dir) > 0 && dir[0] == '/' {
			dir = dir[1:]
		}
		blockAPIServer, err := newMinioBlockAPIServer(dir, cacheBytes, etcdAddress)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case AmazonBackendEnvVar:
		// amazon doesn't like leading slashes
		if len(dir) > 0 && dir[0] == '/' {
			dir = dir[1:]
		}
		blockAPIServer, err := newAmazonBlockAPIServer(dir, cacheBytes, etcdAddress)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case GoogleBackendEnvVar:
		// TODO figure out if google likes leading slashses
		blockAPIServer, err := newGoogleBlockAPIServer(dir, cacheBytes, etcdAddress)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case MicrosoftBackendEnvVar:
		blockAPIServer, err := newMicrosoftBlockAPIServer(dir, cacheBytes, etcdAddress)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case LocalBackendEnvVar:
		blockAPIServer, err := newLocalBlockAPIServer(dir, cacheBytes, etcdAddress)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	default:
		return nil, fmt.Errorf("unrecognized storage backend")
	}
}
