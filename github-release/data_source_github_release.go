package github_release

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubRelease() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubReleaseRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retrieve_by": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"latest",
					"id",
					"tag",
				}, false),
			},
			"release_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//"target_commitish": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//"body": {
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			"is_draft": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_prerelease": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"published_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"author_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_asset_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mention_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourceGithubReleaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	repository := d.Get("repository").(string)
	owner := d.Get("owner").(string)
	getBy := d.Get("retrieve_by").(string)
	//_ := d.Get("release_id").(string)
	tag := d.Get("release_tag").(string)

	client := meta.(*Owner).client
	//_ = meta.(*Owner).Context //context.Background()

	var diags diag.Diagnostics

	var query struct {
		Repository struct {
			Release  Release `graphql:"release(tagName: $tagName)"`
			Releases struct {
				Nodes []Release
			} `graphql:"releases(first: 1, orderBy: { field: CREATED_AT, direction: DESC })"`
		} `graphql:"repository(owner: $login, name: $repository)"`
	}

	var variables = map[string]interface{}{
		"login":      githubv4.String(owner),
		"repository": githubv4.String(repository),
		"tagName":    githubv4.String(tag),
	}

	var err = client.Query(ctx, &query, variables)
	if err != nil {
		return diag.FromErr(err)
	}

	var release Release

	if getBy == "latest" {
		if query.Repository.Releases.Nodes == nil {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Warning,
				Summary:       "No releases found",
				AttributePath: nil,
			})
			return diags
		}
		release = query.Repository.Releases.Nodes[0]
	} else {
		if query.Repository.Release.TagName == "" {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Warning,
				Summary:       "No release found with this TagName",
				AttributePath: nil,
			})
			return diags
		}
	}

	d.SetId(fmt.Sprint(release.DatabaseId))
	d.Set("release_id", release.Id)
	d.Set("release_tag", release.TagName)
	d.Set("owner", owner)
	//d.Set("target_commitish", release.GetTargetCommitish())
	d.Set("name", release.Name)
	d.Set("description", release.Description)
	d.Set("retrieve_by", getBy)

	d.Set("is_draft", release.IsDraft)
	d.Set("is_prerelease", release.IsPrerelease)
	d.Set("is_latest", release.IsLatest)
	d.Set("created_at", release.CreatedAt)
	d.Set("updated_at", release.UpdatedAt)
	d.Set("published_at", release.PublishedAt)
	d.Set("url", release.Url)
	d.Set("author_id", release.Author.Id)
	d.Set("release_asset_count", release.ReleaseAssets.TotalCount)
	d.Set("mention_count", release.Mentions.TotalCount)
	//d.Set("asserts_url", release.GetAssetsURL())
	//d.Set("upload_url", release.GetUploadURL())
	//d.Set("zipball_url", release.GetZipballURL())
	//d.Set("tarball_url", release.GetTarballURL())

	//var err error
	//var release *github.RepositoryRelease
	//
	//switch retrieveBy := strings.ToLower(d.Get("retrieve_by").(string)); retrieveBy {
	//case "latest":
	//	release, _, err = client.Repositories.GetLatestRelease(ctx, owner, repository)
	//case "id":
	//	releaseID := int64(d.Get("release_id").(int))
	//	if releaseID == 0 {
	//		return fmt.Errorf("`release_id` must be set when `retrieve_by` = `id`")
	//	}
	//
	//	release, _, err = client.Repositories.GetRelease(ctx, owner, repository, releaseID)
	//case "tag":
	//	tag := d.Get("release_tag").(string)
	//	if tag == "" {
	//		return fmt.Errorf("`release_tag` must be set when `retrieve_by` = `tag`")
	//	}
	//
	//	release, _, err = client.Repositories.GetReleaseByTag(ctx, owner, repository, tag)
	//default:
	//	return fmt.Errorf("one of: `latest`, `id`, `tag` must be set for `retrieve_by`")
	//}
	//
	//if err != nil {
	//	return err
	//}
	//
	//d.SetId(strconv.FormatInt(release.GetID(), 10))
	//d.Set("release_tag", release.GetTagName())
	//d.Set("target_commitish", release.GetTargetCommitish())
	//d.Set("name", release.GetName())
	//d.Set("body", release.GetBody())
	//d.Set("draft", release.GetDraft())
	//d.Set("prerelease", release.GetPrerelease())
	//d.Set("created_at", release.GetCreatedAt())
	//d.Set("published_at", release.GetPublishedAt())
	//d.Set("url", release.GetURL())
	//d.Set("html_url", release.GetHTMLURL())
	//d.Set("asserts_url", release.GetAssetsURL())
	//d.Set("upload_url", release.GetUploadURL())
	//d.Set("zipball_url", release.GetZipballURL())
	//d.Set("tarball_url", release.GetTarballURL())

	return diags
}
