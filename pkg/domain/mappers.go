package domain

import (
	"fmt"
	"icl-broker/pkg/model"
	"net/url"
)

func ListPostParamsToUrlValues(dto *PostsFiltersDTO) *url.Values {
	queryParams := url.Values{}
	queryParams.Add("limit", fmt.Sprint(dto.Limit))
	queryParams.Add("offset", fmt.Sprint(dto.Offset))
	return &queryParams
}

func ServicePostsToCompactPosts(postsFromService []ServicePost) []*PostCompact {
	posts := []*PostCompact{}
	for _, srvsPost := range postsFromService {
		post := PostCompact{
			ID:          srvsPost.ID,
			UserId:      srvsPost.UserId,
			Description: srvsPost.Description,
			CreatedAt:   srvsPost.CreatedAt,
			UpdatedAt:   srvsPost.UpdatedAt,
			Images:      srvsPost.Images,
			LikesCount:  len(srvsPost.Likers),
		}

		posts = append(posts, &post)
	}

	return posts
}

func ServicePostToDetailed(srvsPost ServicePost) *PostDetailed {
	post := PostDetailed{
		ID:          srvsPost.ID,
		UserId:      srvsPost.UserId,
		Description: srvsPost.Description,
		CreatedAt:   srvsPost.CreatedAt,
		UpdatedAt:   srvsPost.UpdatedAt,
		Images:      srvsPost.Images,
	}
	return &post
}

func DomainToService(post *model.Post) *ServicePost {
	return &ServicePost{
		ID:          post.ID,
		UserId:      post.UserId,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Images:      post.Images,
	}
}
