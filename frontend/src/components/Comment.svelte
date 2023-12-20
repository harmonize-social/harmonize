<script lang="ts">
	import type CommentModel from "../models/comment";

    export let content: CommentModel;
</script>
<style></style>

<p>{content.content_comment}</p>