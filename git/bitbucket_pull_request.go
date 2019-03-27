package git

import "time"

type BitbucketPullRequest struct {
	Approval struct {
		Date time.Time `json:"date"`
		User struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			AccountID   string `json:"account_id"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname string `json:"nickname"`
			Type     string `json:"type"`
			UUID     string `json:"uuid"`
		} `json:"user"`
	} `json:"approval"`
	PullRequest struct {
		Rendered struct {
			Description struct {
				Raw    string `json:"raw"`
				Markup string `json:"markup"`
				HTML   string `json:"html"`
				Type   string `json:"type"`
			} `json:"description"`
			Title struct {
				Raw    string `json:"raw"`
				Markup string `json:"markup"`
				HTML   string `json:"html"`
				Type   string `json:"type"`
			} `json:"title"`
		} `json:"rendered"`
		Type        string `json:"type"`
		Description string `json:"description"`
		Links       struct {
			Decline struct {
				Href string `json:"href"`
			} `json:"decline"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			Merge struct {
				Href string `json:"href"`
			} `json:"merge"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Activity struct {
				Href string `json:"href"`
			} `json:"activity"`
			Diff struct {
				Href string `json:"href"`
			} `json:"diff"`
			Approve struct {
				Href string `json:"href"`
			} `json:"approve"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"links"`
		Title             string        `json:"title"`
		CloseSourceBranch bool          `json:"close_source_branch"`
		Reviewers         []interface{} `json:"reviewers"`
		ID                int           `json:"id"`
		Destination       struct {
			Commit struct {
				Hash  string `json:"hash"`
				Type  string `json:"type"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"commit"`
			Repository struct {
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Type     string `json:"type"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				UUID     string `json:"uuid"`
			} `json:"repository"`
			Branch struct {
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"destination"`
		CreatedOn time.Time `json:"created_on"`
		Summary   struct {
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			HTML   string `json:"html"`
			Type   string `json:"type"`
		} `json:"summary"`
		Source struct {
			Commit struct {
				Hash  string `json:"hash"`
				Type  string `json:"type"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"commit"`
			Repository struct {
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Type     string `json:"type"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				UUID     string `json:"uuid"`
			} `json:"repository"`
			Branch struct {
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"source"`
		CommentCount int    `json:"comment_count"`
		State        string `json:"state"`
		TaskCount    int    `json:"task_count"`
		Participants []struct {
			Role           string    `json:"role"`
			ParticipatedOn time.Time `json:"participated_on"`
			Type           string    `json:"type"`
			User           struct {
				Username    string `json:"username"`
				DisplayName string `json:"display_name"`
				AccountID   string `json:"account_id"`
				Links       struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Nickname string `json:"nickname"`
				Type     string `json:"type"`
				UUID     string `json:"uuid"`
			} `json:"user"`
			Approved bool `json:"approved"`
		} `json:"participants"`
		Reason    string    `json:"reason"`
		UpdatedOn time.Time `json:"updated_on"`
		Author    struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			AccountID   string `json:"account_id"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname string `json:"nickname"`
			Type     string `json:"type"`
			UUID     string `json:"uuid"`
		} `json:"author"`
		MergeCommit interface{} `json:"merge_commit"`
		ClosedBy    interface{} `json:"closed_by"`
	} `json:"pullrequest"`
	Repository struct {
		Scm     string `json:"scm"`
		Website string `json:"website"`
		Name    string `json:"name"`
		Links   struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Project struct {
			Key   string `json:"key"`
			Type  string `json:"type"`
			UUID  string `json:"uuid"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Name string `json:"name"`
		} `json:"project"`
		FullName string `json:"full_name"`
		Owner    struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			Type        string `json:"type"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
		} `json:"owner"`
		Type      string `json:"type"`
		IsPrivate bool   `json:"is_private"`
		UUID      string `json:"uuid"`
	} `json:"repository"`
	Actor struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		AccountID   string `json:"account_id"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Nickname string `json:"nickname"`
		Type     string `json:"type"`
		UUID     string `json:"uuid"`
	} `json:"actor"`
}
