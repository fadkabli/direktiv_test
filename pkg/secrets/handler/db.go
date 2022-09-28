package handler

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/direktiv/direktiv/ent/predicate"
	"github.com/direktiv/direktiv/pkg/dlog"
	"github.com/direktiv/direktiv/pkg/secrets/ent"
	entc "github.com/direktiv/direktiv/pkg/secrets/ent"
	"github.com/direktiv/direktiv/pkg/secrets/ent/namespacesecret"
	_ "github.com/lib/pq" // lib needs to imported for postgres
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* #nosec */
const (
	secretsConn = "DIREKTIV_SECRETS_DB"
	secretsKey  = "DIREKTIV_SECRETS_KEY"
)

type dbHandler struct {
	db  *entc.Client
	key string
}

var logger *zap.SugaredLogger

func setupDB() (SecretsHandler, error) {

	var err error

	dbEnv := os.Getenv(secretsConn)
	keyEnv := os.Getenv(secretsKey)

	logger, err = dlog.ApplicationLogger("secrets-db")
	if err != nil {
		return nil, err
	}

	if keyEnv == "" || dbEnv == "" {
		return nil, fmt.Errorf("DB and Key have to be set")
	}

	if len(keyEnv) != 32 {
		return nil, fmt.Errorf("key needs to be 32 characters")
	}

	db, err := ent.Open("postgres", dbEnv)
	if err != nil {
		logger.Errorf("can not connect to secrets db: %v", err)
		return nil, err
	}

	if err := db.Schema.Create(context.Background()); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}

	return &dbHandler{
		db:  db,
		key: keyEnv,
	}, err
}

func (db *dbHandler) AddSecret(namespace, name string, secret []byte) error {

	logger.Infof("adding secret %s", name)

	bs, _ := db.db.NamespaceSecret.
		Query().
		Where(
			namespacesecret.And(
				namespacesecret.NsEQ(namespace),
				namespacesecret.NameEQ(name),
			)).
		Only(context.Background())

	if bs != nil {
		return fmt.Errorf("secret already exists")
	}

	d, err := encryptData([]byte(db.key), secret)
	if bs != nil {
		return fmt.Errorf("error encrypting data: %v", err)
	}

	_, err = db.db.NamespaceSecret.
		Create().
		SetName(name).
		SetSecret(d).
		SetNs(namespace).
		Save(context.Background())

	return err

}

func (db *dbHandler) GetSecret(namespace, name string) ([]byte, error) {

	bs, err := db.db.NamespaceSecret.
		Query().
		Where(
			namespacesecret.And(
				namespacesecret.NsEQ(namespace),
				namespacesecret.NameEQ(name),
			)).
		Only(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "secret '%s' not found", name)
		}
		return nil, err
	}

	return decryptData([]byte(db.key), bs.Secret)

}

func (db *dbHandler) GetSecrets(namespace string) ([]string, error) {

	var names []string
	dbs := db
	var err error
	prefix := ""
	prefix = strings.TrimPrefix(prefix, "/")

	test := []predicate.NamespaceSecret{
		namespacesecret.NameHasPrefix(prefix),
	}

	if prefix[len(prefix)-1:] != "/" { //FILE  BETTER PROVE IF PREFIX == ""
		dbs, err = db.db.NamespaceSecret.
			Query().
			Where(
				namespacesecret.And(
					namespacesecret.NsEQ(namespace),
				)).
			All(context.Background())
	} else { // FOLDER
		dbs, err := db.db.NamespaceSecret.
			Query().
			Where(
				namespacesecret.And(
					namespacesecret.NsEQ(namespace),
				)).
			All(context.Background())
	}

	if err != nil {
		return nil, err
	}

	for _, s := range dbs {
		if prefix[len(prefix)-1:] != "/" {
			names = append(names, s.Name) //IGNORE FOLDERS, ONLY GET THE SECRETS
		}

	}

	return names, nil

}

//var := []whatevr{
//	namespacesecret.NsEQ(namespace)
//}

//if true {
//	var 0 append(var, namespacesecret.NameHasPrefix(prefix),)
//}

//dbs, err := db.db.NamespaceSecret.
//	Query().
//	Where(
//		namespacesecret.And(
//			var...
//					namespacesecret.NsEQ(namespace),
//					namespacesecret.NameHasPrefix(prefix),
//		)).
//	All(context.Background())

//}

func (db *dbHandler) RemoveSecret(namespace, name string) error {

	if name[len(name)-1:] != "/" { //FILE

		_, err := db.db.NamespaceSecret.
			Delete().
			Where(
				namespacesecret.And(
					namespacesecret.NsEQ(namespace),
					namespacesecret.NameEQ(name),
				)).
			Exec(context.Background())

		return err

	} else { //FOLDER
		_, err := db.db.NamespaceSecret.
			Delete().
			Where(
				namespacesecret.And(
					namespacesecret.NsEQ(namespace),
					namespacesecret.NameHasPrefix(name),
				)).
			Exec(context.Background())
		return err

	}

}

//TODO rename to RemoveNamespaceSecrets
func (db *dbHandler) RemoveSecrets(namespace string) error {

	_, err := db.db.NamespaceSecret.
		Delete().
		Where(
			namespacesecret.And(
				namespacesecret.NsEQ(namespace),
			)).
		Exec(context.Background())

	return err

}

func encryptData(key, data []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil

}

func decryptData(key, data []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := data[:gcm.NonceSize()]
	data = data[gcm.NonceSize():]
	return gcm.Open(nil, nonce, data, nil)

}
