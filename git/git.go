package git

import (
	"context"
	"github.com/ebauman/hf-git/apis/v1alpha1"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"time"
)

type gitworker struct {
	gr v1alpha1.GitRepo
	r *git.Repository
	messages chan<- Message
}

func WatchRepo(repo v1alpha1.GitRepo, ctx context.Context, messages chan<- Message) {
	gw := &gitworker{gr: repo}
	messages <- Message{repo.Name, "watching repository"}
	err := gw.clone()
	if err != nil {
		gw.messagef("error cloning repo: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			rev, err := gw.poll()
			if err != nil {
				gw.messagef("error polling repo: %v", err)
			}
			if rev != gw.gr.Status.Revision {
				// clone again and rebuild
			}
		}
		time.Sleep(time.Second * 60)
	}
}

func (gw *gitworker) clone() error {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: gw.gr.Spec.Repo,
	})
	if err != nil {
		return err
	}

	gw.r = r

	return nil
}

func (gw *gitworker) poll() (string, error) {
	w, err := gw.r.Worktree()
	if err != nil {
		return "", err
	}

	err = w.Pull(&git.PullOptions{})
	if err != nil {
		return "", err
	}

	if rev, err := gw.r.Head(); err != nil {
		 return "", err
	} else {
		return rev.Hash().String(), nil
	}
}