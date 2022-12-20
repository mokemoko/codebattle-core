package commands

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"os"
	"strings"
)

func genRepoHash(entry *models.Entry) string {
	h := sha1.New()
	h.Write([]byte(entry.Repository))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getBranchName(entry *models.Entry) string {
	s := strings.SplitN(entry.Repository, "#", 2)
	if len(s) == 2 {
		return s[1]
	} else {
		return "main"
	}

}

func getEntry() (*models.Entry, error) {
	entry, err := models.Entries(
		models.EntryWhere.Status.EQ(models.EntryStatusAccepted.Code),
		Load(models.EntryRels.User),
	).OneG(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	// 取得時点でステータスは処理中に
	entry.Status = models.EntryStatusProcessing.Code
	err = updateEntry(entry)
	return entry, err
}

func pullRepo(entry *models.Entry) error {
	dir := fmt.Sprintf("work/%s", genRepoHash(entry))
	_ = os.RemoveAll(dir)
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           entry.Repository,
		ReferenceName: plumbing.NewBranchReferenceName(getBranchName(entry)),
		Auth: &http.BasicAuth{
			Username: entry.R.User.Name,
			Password: entry.R.User.Token.String,
		},
		SingleBranch: true,
		Depth:        1,
	})
	if err != nil {
		return err
	}
	return nil
}

func buildImage(entry *models.Entry) error {
	hash := genRepoHash(entry)
	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	buildCtx, err := archive.Tar(fmt.Sprintf("work/%s", hash), 0)
	if err != nil {
		return err
	}
	_, err = c.ImageBuild(context.Background(), buildCtx, types.ImageBuildOptions{
		Tags:           []string{hash},
		SuppressOutput: true,
		Remove:         true,
		ForceRemove:    true,
		Dockerfile:     "Dockerfile",
	})
	if err != nil {
		return err
	}
	return nil
}

func updateEntry(entry *models.Entry) error {
	_, err := entry.UpdateG(context.Background(), boil.Infer())
	return err
}

func applyError(entry *models.Entry, err error) error {
	entry.Status = models.EntryStatusError.Code
	entry.Error.SetValid(err.Error())
	return updateEntry(entry)
}

func RunEntry() {
	count := 0
	for {
		entry, err := getEntry()
		if err != nil {
			log.Fatal(err)
		}
		if entry == nil {
			break
		}
		count += 1
		err = pullRepo(entry)
		if err != nil {
			log.Print(err)
			err = applyError(entry, err)
			if err != nil {
				log.Fatal(err)
			} else {
				continue
			}
		}
		err = buildImage(entry)
		if err != nil {
			log.Print(err)
			err = applyError(entry, err)
			if err != nil {
				log.Fatal(err)
			} else {
				continue
			}
		}
		// success
		entry.Status = models.EntryStatusRegistered.Code
		entry.Error = null.String{}
		err = updateEntry(entry)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Execute %d entries.", count)
}
