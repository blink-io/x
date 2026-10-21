package dotenv

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"

	EnvSuffix            = ".env"
	EnvSuffixFormat      = ".env.%s"
	EnvLocalSuffix       = ".env.local"
	EnvLocalSuffixFormat = ".env.%s.local"
)

// Flow load multiple files based on the following orders (from lower to higher)
// .env
// .env.production
// .env.test
// .env.development
// .env.local
// .env.production.local
// .env.test.local
// .env.development.local
// supported env: development|test|production
// See https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func Flow(paths ...string) error {
	if len(paths) == 0 {
		return nil
	}
	envs := []string{EnvProduction, EnvTest, EnvDevelopment}

	// File names do not depend on the path; build them once instead of
	// re-formatting format strings inside the nested per-path loop.
	files := make([]string, 0, len(envs)*2)
	for _, env := range envs {
		files = append(files, ".env."+env, ".env."+env+".local")
	}
	envFiles := files[:len(envs)]
	localFiles := files[len(envs):]

	for _, path := range paths {
		// Load .env first
		_ = godotenv.Overload(filepath.Join(path, EnvSuffix))
		// Load .env.(production|test|development)
		for _, name := range envFiles {
			_ = godotenv.Overload(filepath.Join(path, name))
		}
		// Load .env.local
		_ = godotenv.Overload(filepath.Join(path, EnvLocalSuffix))
		// Load .env.local.(production|test|development)
		for _, name := range localFiles {
			_ = godotenv.Overload(filepath.Join(path, name))
		}
	}
	return nil
}
