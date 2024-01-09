package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/platforms"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

func GetSavedPosts(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }

    id := uuid.MustParse(r.Header.Get("id"))

    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `SELECT
    p.id AS post_id,
    p.created_at AS post_created_at,
    p.caption AS post_caption,
    p.type AS post_type,
    p.type_specific_id AS post_specific_id,
    u.username AS post_username,
    COUNT(lp.id) AS likes_count,
    CASE WHEN lp.user_id IS NOT NULL THEN true ELSE false END AS user_liked,
    CASE WHEN sp.user_id IS NOT NULL THEN true ELSE false END AS user_saved
    FROM
        saved_posts sp
    JOIN
        posts p ON sp.post_id = p.id
    LEFT JOIN
        liked_posts lp ON p.id = lp.post_id
    JOIN
        users u ON p.user_id = u.id
    WHERE
        sp.user_id = $1
    GROUP BY
        sp.id,
        p.id,
        lp.user_id,
        u.id
    LIMIT $2
    OFFSET $3;`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        err = rows.Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username, &post.LikeCount, &post.HasLiked, &post.HasSaved)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        content, err := repositories.GetPostContent(post.Type, typeSpecificId)

        if err != nil {
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

func PostSavedPost(w http.ResponseWriter, r *http.Request) {
    id := r.Header.Get("id")
    postId := r.URL.Query().Get("id")

    user, err := auth.GetUserFromSession(uuid.MustParse(id))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    var savedPostId uuid.UUID
    sqlStatement := `INSERT INTO saved_posts (id, user_id, post_id)
                     VALUES (uuid_generate_v4(), $1, $2)
                     ON CONFLICT (user_id, post_id) DO UPDATE
                     SET id = saved_posts.id
                     RETURNING id;`
    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, user.ID, postId).Scan(&savedPostId)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error saving post")
        return
    }

    platformItems, err := repositories.GetPostPlatform(uuid.MustParse(postId))

    if err != nil {
        fmt.Println(err.Error())
        models.Error(w, http.StatusInternalServerError, "Error saving post")
    }

    if len(platformItems) == 0 {
        fmt.Println("No platforms")
        models.Error(w, http.StatusInternalServerError, "Error saving post")
        return
    }

    platform := platforms.GetPlatform(platformItems[0].Platform, user.ID)

    result, err := platform.Save(platformItems[0].Type, platformItems[0].ID)

    if err != nil {
        fmt.Println(err.Error())
        models.Error(w, http.StatusInternalServerError, "Error saving post")
        return
    }

    if !result {
        fmt.Println("Failed to save")
        models.Error(w, http.StatusInternalServerError, "Error saving post")
        return
    }

    models.Result(w, savedPostId)
}

func DeleteSavedPost(w http.ResponseWriter, r *http.Request) {
    id := r.Header.Get("id")
    postId := r.URL.Query().Get("id")

    user, err := auth.GetUserFromSession(uuid.MustParse(id))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `DELETE FROM saved_posts
                     WHERE user_id = $1
                     AND post_id = $2;`
    tag, err := repositories.Pool.Exec(context.Background(), sqlStatement, user.ID, postId)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error saving post")
        return
    }

    if tag.RowsAffected() != 1 {
        models.Error(w, http.StatusInternalServerError, "Failed to delete post")
        return
    }

    models.Result(w, "Deleted")
}
