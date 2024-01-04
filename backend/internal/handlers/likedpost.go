package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"

    "github.com/google/uuid"
)

func GetLikedPosts(w http.ResponseWriter, r *http.Request) {
    limit, offset, user, err := GetLimitOffsetSession(r)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `SELECT
    lp.id AS liked_post_id,
    p.id AS post_id,
    p.user_id AS post_owner_id,
    u.username AS post_username,
    p.created_at AS post_created_at,
    p.caption AS post_caption,
    p.type AS post_type,
    p.type_specific_id AS post_specific_id,
    COUNT(lp.id) AS likes_count,
    CASE WHEN lp.user_id IS NOT NULL THEN true ELSE false END AS user_liked,
    CASE WHEN sp.user_id IS NOT NULL THEN true ELSE false END AS user_saved
    FROM
        liked_posts lp
    JOIN
        posts p ON lp.post_id = p.id
    LEFT JOIN
        saved_posts sp ON p.id = sp.post_id
    JOIN
        users u ON p.user_id = u.id
    WHERE
        lp.user_id = $1
    GROUP BY
    lp.id,
    p.id,
    sp.user_id,
    u.id
    LIMIT $2 OFFSET $3;
    `
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)
    if err != nil {
        fmt.Println(err.Error())
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        err = rows.Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username, &post.LikeCount, &post.HasLiked, &post.HasLiked)
        if err != nil {
            fmt.Println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        content,err := repositories.GetPostContent(post.Type, typeSpecificId)

        if err != nil {
            println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Content = content

        processedComments, err := repositories.GetPostComments(post.ID)

        if err != nil {
            println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Comments = processedComments

        posts = append(posts, post)
    }
    models.Result(w, posts)
}

func PostLikedPost(w http.ResponseWriter, r *http.Request) {
    id := r.Header.Get("id")
    postId := r.URL.Query().Get("id")

    user, err := auth.GetUserFromSession(uuid.MustParse(id))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    var savedPostId uuid.UUID
    sqlStatement := `INSERT INTO liked_posts (id, user_id, post_id)
                     VALUES (uuid_generate_v4(), $1, $2)
                     ON CONFLICT (user_id, post_id) DO UPDATE
                     SET id = liked_posts.id
                     RETURNING id;`
    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, user.ID, postId).Scan(&savedPostId)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error liking post")
        return
    }

    models.Result(w, savedPostId)
}

func DeleteLikedPost(w http.ResponseWriter, r *http.Request) {
    id := r.Header.Get("id")
    postId := r.URL.Query().Get("id")

    user, err := auth.GetUserFromSession(uuid.MustParse(id))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `DELETE FROM liked_posts
                     WHERE user_id = $1
                     AND post_id = $2;`
    tag, err := repositories.Pool.Exec(context.Background(), sqlStatement, user.ID, postId)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error liking post")
        return
    }

    if tag.RowsAffected() != 1 {
        models.Error(w, http.StatusInternalServerError, "Failed to delete post")
        return
    }

    models.Result(w, "Deleted")
}
